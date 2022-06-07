package conf

import (
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

type ConfigDatabase struct {
	Host           string        `yaml:"host" env:"HOST" env-default:"127.0.0.1"`
	Port           string        `yaml:"port" env:"PORT" env-default:"8080"`
	DbPath         string        `yaml:"db_path" env:"DbPath"`
	SQLInitPath    string        `yaml:"sql_init_path" env:"SQLInitPath"`
	SymmetricKey   string        `yaml:"symmetric_key" env:"SymmetricKey"`
	AccessDuration time.Duration `yaml:"access_duration" env:"AccessDuration"`
}

var cfg ConfigDatabase

func Cnf() (*ConfigDatabase, error) {
	err := cleanenv.ReadConfig("conf/cnf.yml", &cfg)
	if err != nil {
		return nil, err
	}

	return &cfg, nil
}
