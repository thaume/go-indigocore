// Copyright 2016 Stratumn SAS. All rights reserved.
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

// Package types defines common types.
package types

import (
	"encoding/hex"
	"encoding/json"
)

// Bytes20Size is the size of a 20-byte long byrte array...
const Bytes20Size = 20

// Bytes20 is the 20-byte long byte array.
type Bytes20 [Bytes20Size]byte

// NewBytes20FromString creates a Bytes20 from a hex encoded string.
func NewBytes20FromString(src string) (*Bytes20, error) {
	var b Bytes20
	if err := b.Unstring(src); err != nil {
		return nil, err
	}
	return &b, nil
}

// String returns a hex encoded string.
func (b *Bytes20) String() string {
	return hex.EncodeToString(b[:])
}

// Unstring sets the value from a hex encoded string.
func (b *Bytes20) Unstring(src string) error {
	_, err := hex.Decode(b[:], []byte(src))
	return err
}

// MarshalJSON implements encoding/json.Marshaler.MarshalJSON.
func (b *Bytes20) MarshalJSON() ([]byte, error) {
	return json.Marshal(b.String())
}

// UnmarshalJSON implements encoding/json.Unmarshaler.UnmarshalJSON.
func (b *Bytes20) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}

	buf, err := hex.DecodeString(s)
	if err != nil {
		return err
	}

	copy(b[:], buf)
	return nil
}

// Reverse reverse the bytes order.
func (b *Bytes20) Reverse(rb *ReversedBytes20) {
	for i, v := range b {
		rb[Bytes20Size-i-1] = v
	}
}

// ReversedBytes20 is the 20-byte long byte reversed array.
type ReversedBytes20 [Bytes20Size]byte

// NewReversedBytes20FromString creates a ReversedBytes20 from a hex encoded string.
func NewReversedBytes20FromString(src string) (*ReversedBytes20, error) {
	b, err := NewBytes20FromString(src)
	if err != nil {
		return nil, err
	}
	var rb ReversedBytes20
	b.Reverse(&rb)
	return &rb, nil
}

// String returns a hex encoded string.
func (rb *ReversedBytes20) String() string {
	var b Bytes20
	rb.Reverse(&b)
	return b.String()
}

// Unstring sets the value from a hex encoded string.
func (rb *ReversedBytes20) Unstring(src string) error {
	b, err := NewBytes20FromString(src)
	if err != nil {
		return err
	}
	b.Reverse(rb)
	return nil
}

// MarshalJSON implements encoding/json.Marshaler.MarshalJSON.
func (rb *ReversedBytes20) MarshalJSON() ([]byte, error) {
	var b Bytes20
	rb.Reverse(&b)
	return b.MarshalJSON()
}

// UnmarshalJSON implements encoding/json.Unmarshaler.UnmarshalJSON.
func (rb *ReversedBytes20) UnmarshalJSON(data []byte) error {
	var b Bytes20
	if err := json.Unmarshal(data, &b); err != nil {
		return err
	}
	b.Reverse(rb)
	return nil
}

// Reverse reverse the bytes order.
func (rb *ReversedBytes20) Reverse(b *Bytes20) {
	for i, v := range rb {
		b[Bytes20Size-i-1] = v
	}
}

// Bytes32Size is the size of a 32-byte long byrte array...
const Bytes32Size = 32

// Bytes32 is the 32-byte long byte array.
type Bytes32 [Bytes32Size]byte

// NewBytes32FromString creates a Bytes32 from a hex encoded string.
func NewBytes32FromString(src string) (*Bytes32, error) {
	var b Bytes32
	if err := b.Unstring(src); err != nil {
		return nil, err
	}
	return &b, nil
}

// String returns a hex encoded string.
func (b *Bytes32) String() string {
	return hex.EncodeToString(b[:])
}

// Unstring sets the value from a hex encoded string.
func (b *Bytes32) Unstring(src string) error {
	_, err := hex.Decode(b[:], []byte(src))
	return err
}

// MarshalJSON implements encoding/json.Marshaler.MarshalJSON.
func (b *Bytes32) MarshalJSON() ([]byte, error) {
	return json.Marshal(b.String())
}

// UnmarshalJSON implements encoding/json.Unmarshaler.UnmarshalJSON.
func (b *Bytes32) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}

	buf, err := hex.DecodeString(s)
	if err != nil {
		return err
	}

	copy(b[:], buf)
	return nil
}

// Reverse reverse the bytes order.
func (b *Bytes32) Reverse(rb *ReversedBytes32) {
	for i, v := range b {
		rb[Bytes32Size-i-1] = v
	}
}

// ReversedBytes32 is the 32-byte long byte reversed array.
type ReversedBytes32 [Bytes32Size]byte

// NewReversedBytes32FromString creates a ReversedBytes32 from a hex encoded string.
func NewReversedBytes32FromString(src string) (*ReversedBytes32, error) {
	b, err := NewBytes32FromString(src)
	if err != nil {
		return nil, err
	}
	var rb ReversedBytes32
	b.Reverse(&rb)
	return &rb, nil
}

// String returns a hex encoded string.
func (rb *ReversedBytes32) String() string {
	var b Bytes32
	rb.Reverse(&b)
	return b.String()
}

// Unstring sets the value from a hex encoded string.
func (rb *ReversedBytes32) Unstring(src string) error {
	b, err := NewBytes32FromString(src)
	if err != nil {
		return err
	}
	b.Reverse(rb)
	return nil
}

// MarshalJSON implements encoding/json.Marshaler.MarshalJSON.
func (rb *ReversedBytes32) MarshalJSON() ([]byte, error) {
	var b Bytes32
	rb.Reverse(&b)
	return b.MarshalJSON()
}

// UnmarshalJSON implements encoding/json.Unmarshaler.UnmarshalJSON.
func (rb *ReversedBytes32) UnmarshalJSON(data []byte) error {
	var b Bytes32
	if err := json.Unmarshal(data, &b); err != nil {
		return err
	}
	b.Reverse(rb)
	return nil
}

// Reverse reverse the bytes order.
func (rb *ReversedBytes32) Reverse(b *Bytes32) {
	for i, v := range rb {
		b[Bytes32Size-i-1] = v
	}
}
