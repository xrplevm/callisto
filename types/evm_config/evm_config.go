package evm

import (
	"fmt"
	"os"
	"path"

	"github.com/forbole/juno/v6/types/config"
	"gopkg.in/yaml.v3"
)

// Config holds EVM-specific configuration for the explorer.
type Config struct {
	ChainID uint64 `yaml:"chain_id"`
}

func DefaultConfig() Config {
	return Config{ChainID: 262144}
}

// ParseConfig parses the EVM configuration from the given YAML bytes.
func ParseConfig(bz []byte) Config {
	type T struct {
		EVM *Config `yaml:"evm"`
	}
	var cfg T
	if err := yaml.Unmarshal(bz, &cfg); err != nil || cfg.EVM == nil {
		return DefaultConfig()
	}
	return *cfg.EVM
}

// Cfg is the global EVM configuration used during execution.
var Cfg = DefaultConfig()

// GetConfig returns the configuration reading it from the config.yaml file present inside the home directory
func GetConfig() (Config, error) {
	file := path.Join(config.HomePath, "config.yaml")

	// Make sure the path exists
	if _, err := os.Stat(file); os.IsNotExist(err) {
		return Config{}, fmt.Errorf("config file does not exist")
	}

	bz, err := os.ReadFile(file)
	if err != nil {
		return Config{}, fmt.Errorf("error while reading config file: %s", err)
	}

	return ParseConfig(bz), nil
}

// ReadConfigFromFile reads the EVM configuration from the config.yaml file.
// Returns the default config if the file does not exist or cannot be parsed.
func ReadConfigFromFile() Config {
	cfg, err := GetConfig()
	if err != nil {
		return DefaultConfig()
	}
	return cfg
}
