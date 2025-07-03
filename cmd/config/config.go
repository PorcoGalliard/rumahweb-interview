package config

import "github.com/PorcoGalliard/rumahweb-interview/pkg/config"

type UserConfig struct {
	App config.AppConfig
	Database config.PostgreConfig
	Secret config.SecretConfig
}