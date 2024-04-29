package platforms

import (
	"os"

	toml "github.com/pelletier/go-toml/v2"
)

const platformsConfigPath = "/etc/containerd/platform-config.toml"

var cfg PlatformsConfig

type PlatformsConfig struct {
	initialized bool

	Features        []string          `toml:"features"`
	Compatibilities map[string]string `toml:"compatibilities"`
}

func (p *PlatformsConfig) IsInitialized() bool {
	return p.initialized
}

func NewPlatformsConfig() (*PlatformsConfig, error) {
	if cfg.IsInitialized() {
		return &cfg, nil
	}

	b, err := os.ReadFile(platformsConfigPath)
	if err != nil {
		return nil, err
	}

	err = toml.Unmarshal(b, &cfg)
	if err != nil {
		return nil, err
	}
	cfg.initialized = true

	return &cfg, nil
}
