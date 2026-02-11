package evm

import (
	"os"
	"path"
	"testing"

	"github.com/forbole/juno/v6/types/config"
	"github.com/stretchr/testify/require"
)

func TestDefaultConfig(t *testing.T) {
	cfg := DefaultConfig()
	require.Equal(t, uint64(262144), cfg.ChainID)
}

func TestParseConfig(t *testing.T) {
	validYaml := []byte(`
evm:
  chain_id: 12345
`)
	invalidYaml := []byte(`invalid`)

	t.Run("Valid YAML", func(t *testing.T) {
		cfg := ParseConfig(validYaml)
		require.Equal(t, uint64(12345), cfg.ChainID)
	})

	t.Run("Invalid YAML", func(t *testing.T) {
		cfg := ParseConfig(invalidYaml)
		require.Equal(t, DefaultConfig(), cfg)
	})

	t.Run("Missing EVM Config", func(t *testing.T) {
		missingEvmYaml := []byte(`other: value`)
		cfg := ParseConfig(missingEvmYaml)
		require.Equal(t, DefaultConfig(), cfg)
	})
}

func TestGetConfig(t *testing.T) {
	t.Run("Valid Config File", func(t *testing.T) {
		tempDir, err := os.MkdirTemp("", "callisto-test-*")
		require.NoError(t, err)
		defer os.RemoveAll(tempDir)

		validYaml := []byte(`
evm:
  chain_id: 67890
`)
		err = os.WriteFile(path.Join(tempDir, "config.yaml"), validYaml, 0644)
		require.NoError(t, err)

		originalHomePath := config.HomePath
		config.HomePath = tempDir
		defer func() { config.HomePath = originalHomePath }()

		cfg, err := GetConfig()
		require.NoError(t, err)
		require.Equal(t, uint64(67890), cfg.ChainID)
	})

	t.Run("Missing Config File", func(t *testing.T) {
		tempDir, err := os.MkdirTemp("", "callisto-test-*")
		require.NoError(t, err)
		defer os.RemoveAll(tempDir)
		tempFile, err := os.CreateTemp("", "config-*.yaml")
		require.NoError(t, err)
		defer os.Remove(tempFile.Name())

		invalidYaml := []byte(`invalid: yaml: content`)
		_, err = tempFile.Write(invalidYaml)
		require.NoError(t, err)
		require.NoError(t, tempFile.Close())

		t.Setenv("JUNO_CONFIG", tempFile.Name())

		cfg := ReadConfigFromFile()
		require.Equal(t, DefaultConfig(), cfg)
	})
}
