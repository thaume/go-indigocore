package validator

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"testing"

	"github.com/stratumn/go-indigocore/cs/cstesting"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type testCase struct {
	name      string
	baseCfg   *validatorBaseConfig
	scriptCfg *scriptConfig
	valid     bool
	err       string
}

func TestNewScriptValidator(t *testing.T) {
	testLink := cstesting.RandomLink()
	testLink.Meta.Type = "valid"

	sourceFile := filepath.Join("testdata", "custom_validator.go")
	pluginFile := filepath.Join("testdata", "custom_validator.so")
	defer os.Remove(pluginFile)

	fmt.Println("Compiling go plugin...")
	cmd := exec.Command("go", "build", "-o", pluginFile, "-buildmode=plugin", sourceFile)
	require.NoError(t, cmd.Run())
	fmt.Println("Done!")

	testCases := []testCase{
		{
			name: "valid-config",
			baseCfg: &validatorBaseConfig{
				Process:  "test",
				LinkType: "valid",
			},
			scriptCfg: &scriptConfig{
				File: pluginFile,
				Type: "go",
			},
			valid: true,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			s, err := newScriptValidator(tt.baseCfg, tt.scriptCfg)
			if tt.valid {
				assert.NoError(t, err)
				assert.Nil(t, s.Validate(context.Background(), nil, testLink))
			} else {
				assert.EqualError(t, err, tt.err)
			}
		})
	}

	// t.Run("valid script", func(t *testing.T) {
	// 	baseConfig := &validatorBaseConfig{
	// 		Process:  "test",
	// 		LinkType: "valid",
	// 	}

	// 	validScriptCfg := &scriptConfig{
	// 		File: pluginFile,
	// 		Type: "go",
	// 	}
	// 	s, err := newScriptValidator(baseConfig, validScriptCfg)
	// 	require.NoError(t, err)
	// 	assert.Nil(t, s.Validate(context.Background(), nil, testLink))
	// })
}
