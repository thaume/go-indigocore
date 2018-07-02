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

package utils_test

import (
	"testing"

	"github.com/stratumn/go-indigocore/utils"
	"github.com/stretchr/testify/assert"
)

func TestStructurize(t *testing.T) {

	type testStruct struct {
		Test string `json:"test"`
	}

	t.Run("transforms into custom type", func(t *testing.T) {
		src := map[string]interface{}{
			"test": "jean-pierre",
		}
		dest := testStruct{}
		err := utils.Structurize(src, &dest)
		assert.NoError(t, err)
		assert.Equal(t, "jean-pierre", dest.Test)
	})

	t.Run("fails when type does not match", func(t *testing.T) {
		src := map[string]interface{}{
			"test": true,
		}
		err := utils.Structurize(src, &testStruct{})
		assert.Error(t, err)
	})
}
