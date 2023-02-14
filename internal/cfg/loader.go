package cfg

import (
	"io/fs"

	"github.com/caarlos0/env/v6"
	envLoader "github.com/joho/godotenv"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
)

func LoadConfig() (*Config, error) {
	err := envLoader.Load()
	if err == nil {
		log.Info().Msg("Loaded configuration from local .env file")
	} else if !errors.Is(err, fs.ErrNotExist) {
		return nil, errors.Wrap(err, "failed parsing .env file")
	}

	config := Config{}
	if err := env.Parse(&config); err != nil {
		return nil, err
	}

	return &config, nil
}
