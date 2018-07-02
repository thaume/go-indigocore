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

package utils

import (
	"encoding/json"

	"github.com/pkg/errors"
)

var (
	// ErrBadTypeStructure is returned when trying to (de)serialize
	// an interface from/to an incompatible type.
	ErrBadTypeStructure = errors.New("bad type structure")
)

// Structurize transforms the state into a custom type.
// The provided 'dest' argument should be a pointer to struct (passing a literal type will fail).
// On success, 'stateType' will be overriden with the state's data matching its JSON structure.
func Structurize(src interface{}, dest interface{}) error {
	srcBytes, err := json.Marshal(src)
	if err != nil {
		return errors.Wrap(err, ErrBadTypeStructure.Error())
	}
	if err := json.Unmarshal(srcBytes, dest); err != nil {
		return errors.Wrap(err, ErrBadTypeStructure.Error())
	}
	return nil
}
