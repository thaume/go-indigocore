// Copyright 2017 Stratumn SAS. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package cs defines types to work with Chainscripts.
package cs

import (
	"encoding/json"
	"fmt"

	"github.com/pkg/errors"
)

// DeserializeMethods maps a proof backend (like "TMPop") to a deserializer function returning a specific proof
var DeserializeMethods = make(map[string]func(json.RawMessage) (Proof, error))

// Evidences encapsulates a list of evidences contained in Segment.Meta
type Evidences []*Evidence

// AddEvidence sets the segment evidence
func (e *Evidences) AddEvidence(evidence Evidence) error {
	if e.GetEvidence(evidence.Provider) != nil {
		return fmt.Errorf("evidence already exist for provider %s", evidence.Provider)
	}
	*e = append(*e, &evidence)
	return nil
}

// GetEvidence gets an evidence from a provider
func (e *Evidences) GetEvidence(provider string) *Evidence {
	for _, evidence := range *e {
		if evidence.Provider == provider {
			return evidence
		}
	}
	return nil
}

// FindEvidences find all evidences generated by a specified backend ("TMPop", "bcbatchfossilizer"...)
func (e *Evidences) FindEvidences(backend string) (res Evidences) {
	for _, evidence := range *e {
		if evidence.Backend == backend {
			res = append(res, evidence)
		}
	}
	return res
}

// Evidence contains data that can be used to externally verify a segment's proof of existence
type Evidence struct {
	Backend  string `json:"backend"`  // can either be "TMPop", "bitcoin", "dummy"...
	Provider string `json:"provider"` // can either be a chainId (in case of a blockchain) or an identifier for a trusted third-party (timestamping authority, regulator...)
	Proof    Proof  `json:"proof"`
}

// UnmarshalJSON serializes bytes into an Evidence
func (e *Evidence) UnmarshalJSON(data []byte) error {
	serialized := struct {
		Backend  string          `json:"backend"`
		Provider string          `json:"provider"`
		Proof    json.RawMessage `json:"proof"`
	}{}

	err := json.Unmarshal(data, &serialized)
	if err != nil {
		return errors.WithStack(err)
	}

	deserializer, exists := DeserializeMethods[serialized.Backend]
	if !exists {
		return errors.New("Evidence type does not exist")
	}

	proof, err := deserializer(serialized.Proof)
	if err != nil {
		return err
	}

	*e = Evidence{
		Backend:  serialized.Backend,
		Provider: serialized.Provider,
		Proof:    proof,
	}
	return nil
}

// UnmarshalRQL serializes an interface{} into an Evidence
func (e *Evidence) UnmarshalRQL(data interface{}) error {
	return e.UnmarshalJSON(data.([]byte))
}

// MarshalRQL serializes an Evidence into an JSON-formatted byte array
func (e *Evidence) MarshalRQL() (interface{}, error) {
	return json.Marshal(e)
}

// Proof is the generic interface which a custom proof type has to implement.
// Each package defining its own implementation of the cs.Proof interface
// needs to provide a way to serialize/deserialize such proof.
// The init() function of each such package should register the deserialize
// method to the DeserializeMethods map.
type Proof interface {
	Time() uint64            // returns the timestamp (UNIX format) contained in the proof
	FullProof() []byte       // returns data to independently validate the proof
	Verify(interface{}) bool // Checks the validity of the proof
}

// GenericProof implements the Proof interface
type GenericProof struct {
	Timestamp uint64      `json:"timestamp"`
	Data      interface{} `json:"data"`
	Pubkey    []byte      `json:"pubkey"`
	Signature []byte      `json:"signature"` //sign(hash(time+data))
}

// Time returns the timestamp from the block header
func (p *GenericProof) Time() uint64 {
	return p.Timestamp
}

// FullProof returns a JSON formatted proof
func (p *GenericProof) FullProof() []byte {
	bytes, err := json.MarshalIndent(p, "", "   ")
	if err != nil {
		return nil
	}
	return bytes
}

// Verify returns true if the proof of a given linkHash is correct
func (p *GenericProof) Verify(_ interface{}) bool {
	return true
}

// init needs to define a way to deserialize a DummyProof
func init() {
	DeserializeMethods["generic"] = func(rawProof json.RawMessage) (Proof, error) {
		p := GenericProof{}
		if err := json.Unmarshal(rawProof, &p); err != nil {
			return nil, err
		}
		return &p, nil
	}
}
