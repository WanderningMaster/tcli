package config

import (
	"context"
	"errors"
	"fmt"
	"os"
	"path"

	"github.com/WanderningMaster/tcli/internal"
	"github.com/WanderningMaster/tcli/internal/encoding"
	"github.com/WanderningMaster/tcli/internal/logger"
)

var (
	ConfigNotFound = errors.New("Config not found in desired location")
)

type Config struct {
	StoragePath string
}

func LoadConfig(ctx context.Context, p encoding.Parser, cfg *Config) error {
	filePath := path.Join(internal.GetConfigDir(ctx), "tcli", fmt.Sprintf("config.%s", p.Extension()))

	b, err := os.ReadFile(filePath)
	if err != nil {
		return ConfigNotFound
	}

	err = p.Unmarshal(ctx, b, cfg)
	if err != nil {
		return err
	}

	return nil
}

func LoadDefaultConfig(ctx context.Context, p encoding.Parser, cfg *Config) error {
	configDir := path.Join(internal.GetConfigDir(ctx), "tcli")

	cfg.StoragePath = internal.GetHomeDir(ctx)
	b, err := p.Marshal(ctx, &cfg)

	if err != nil {
		panic(err)
	}

	if _, err := os.Stat(configDir); err != nil {
		err = os.MkdirAll(configDir, 0700)
		if err != nil {
			return err
		}
	}

	filePath := path.Join(configDir, fmt.Sprintf("config.%s", p.Extension()))

	fd, err := os.Create(filePath)
	defer fd.Close()

	if err != nil {
		return nil
	}

	_, err = fd.Write(b)
	if err != nil {
		return err
	}

	return nil
}

func NewConfig(ctx context.Context, p encoding.Parser) *Config {
	logger := logger.FromContext(ctx)
	var cfg Config

	err := LoadConfig(ctx, p, &cfg)
	if errors.Is(err, ConfigNotFound) {
		logger.Warn("Failed to load config. Creating default one...")
		err = LoadDefaultConfig(ctx, p, &cfg)

		if err != nil {
			logger.Error(err.Error())
			return nil
		}
	}

	return &cfg
}
