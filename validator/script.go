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

package validator

import (
	"context"
	"crypto/sha256"
	"path"
	"plugin"
	"strings"

	cj "github.com/gibson042/canonicaljson-go"
	"github.com/pkg/errors"
	"github.com/stratumn/go-indigocore/cs"
	"github.com/stratumn/go-indigocore/store"
	"github.com/stratumn/go-indigocore/types"
)

const baseValidatorPath = ""

var validScriptTypes = []string{"go"}

// TransitionValidator is the function called when validating a transition
type TransitionValidator = func(l *cs.Link) error

type scriptValidator struct {
	Script TransitionValidator
	Config *validatorBaseConfig
}

func checkScriptType(cfg *scriptConfig) error {
	switch cfg.Type {
	case "go":
		return nil
	default:
		return errors.Errorf("Validation engine does not handle script of type %s, valid types are %v", cfg.Type, validScriptTypes)
	}
}

func newScriptValidator(baseConfig *validatorBaseConfig, scriptCfg *scriptConfig) (Validator, error) {
	if err := checkScriptType(scriptCfg); err != nil {
		return nil, err
	}

	p, err := plugin.Open(path.Join(baseValidatorPath, scriptCfg.File))
	if err != nil {
		return nil, errors.Wrap(err, "Could not load validation plugin")
	}

	symbol, err := p.Lookup(strings.Title(baseConfig.LinkType))
	if err != nil {
		return nil, errors.Wrapf(err, "Error while loading validation script for process %s and type %s", baseConfig.Process, baseConfig.LinkType)
	}

	customValidator, ok := symbol.(TransitionValidator)
	if !ok {
		return nil, errors.Errorf("Could not load validation script for process %s and linkType %s: script does not implement the CustomValidator interface", baseConfig.Process, baseConfig.LinkType)
	}

	return &scriptValidator{
		Config: baseConfig,
		Script: customValidator,
	}, nil
}

func (sv scriptValidator) Hash() (*types.Bytes32, error) {
	b, err := cj.Marshal(sv)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	validationsHash := types.Bytes32(sha256.Sum256(b))
	return &validationsHash, nil
}

func (sv scriptValidator) ShouldValidate(link *cs.Link) bool {
	return sv.Config.ShouldValidate(link)
}

// Validate checks that the provided signatures match the required ones.
// a requirement can either be: a public key, a name defined in PKI, a role defined in PKI.
func (sv scriptValidator) Validate(_ context.Context, storeReader store.SegmentReader, link *cs.Link) error {
	err := sv.Script(link)
	return err
}
