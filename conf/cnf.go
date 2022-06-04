package conf

import (
	"github.com/ilyakaznacheev/cleanenv"
)

type ConfigDatabase struct {
	Host        string `yaml:"host" env:"HOST" env-default:"127.0.0.1"`
	Port        string `yaml:"port" env:"PORT" env-default:"8080"`
	DbPath      string `yaml:"db" env:"PORT"`
	SQLInitPath string `yaml:"sql_init_path" env:"PORT"`
	SessionsKey string `yaml:"sessions_key" env:"PORT"`
}

var cfg ConfigDatabase

func Cnf() (*ConfigDatabase, error) {
	err := cleanenv.ReadConfig("conf/cnf.yml", &cfg)
	if err != nil {
		return nil, err
	}

	return &cfg, nil
}
