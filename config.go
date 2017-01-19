package mcsm

import "github.com/BurntSushi/toml"

// Config ... This struct maekes from toml file.
type Config struct {
	APIKey string `toml:"api_key"`
	Rule   []Rule `toml:"rule"`
}

// Rule ... This struct makes [[rule]] from toml file.
type Rule struct {
	Cmd         string `toml:"cmd"`
	ServiceName string `toml:"service_name"`
	MetricName  string `toml:"metric_name"`
}

// LoadConfig ...Decoding from toml file.
func LoadConfig(path string, config *Config) error {
	if _, err := toml.DecodeFile(path, &config); err != nil {
		return err
	}

	return nil
}
