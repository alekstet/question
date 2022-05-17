package conf

import (
	"database/sql"
	"io/ioutil"

	"github.com/gorilla/sessions"
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/julienschmidt/httprouter"
	"github.com/olebedev/config"
	"github.com/sirupsen/logrus"
)

type Config struct {
	Ip          string
	Port        string
	DbPath      string
	SQLInitPath string
	SessionsKey string
}

type ConfigDatabase struct {
	Host        string `yaml:"host" env:"HOST" env-default:"127.0.0.1"`
	Port        string `yaml:"port" env:"PORT" env-default:"8080"`
	DbPath      string `yaml:"db" env:"PORT"`
	SQLInitPath string `yaml:"sql_init_path" env:"PORT"`
	SessionsKey string `yaml:"sessions_key" env:"PORT"`
}

var cfg ConfigDatabase

type Store struct {
	Db      *sql.DB
	Log     *logrus.Logger
	Routes  *httprouter.Router
	Session *sessions.CookieStore
}

func New(key string) *Store {
	s := &Store{
		Log:     logrus.New(),
		Routes:  httprouter.New(),
		Session: sessions.NewCookieStore([]byte(key)),
	}
	return s
}

func (s *Store) InitDB(c *Config) error {
	db, err := sql.Open("sqlite3", c.DbPath)
	if err != nil {
		return err
	}

	sql, err := ioutil.ReadFile(c.SQLInitPath)
	if err != nil {
		panic(err)
	}

	_, err = db.Query(string(sql))
	if err != nil {
		return err
	}

	if err = db.Ping(); err != nil {
		return err
	}

	s.Db = db
	return nil
}

func Cnf() (*Config, error) {
	file, err := ioutil.ReadFile("conf/cnf.yml")
	if err != nil {
		return nil, err
	}
	yamlString := string(file)

	err = cleanenv.ReadConfig("conf/cnf.yml", &cfg)
	if err != nil {
		return nil, err
	}

	cfg, err := config.ParseYaml(yamlString)
	if err != nil {
		return nil, err
	}
	SQLInitPath, err := cfg.String("sql_init_path")
	if err != nil {
		return nil, err
	}
	Ip, err := cfg.String("ip")
	if err != nil {
		return nil, err
	}
	Port, err := cfg.String("port")
	if err != nil {
		return nil, err
	}
	Db, err := cfg.String("db")
	if err != nil {
		return nil, err
	}
	SessionsKey, err := cfg.String("sessions_key")
	if err != nil {
		return nil, err
	}
	return &Config{
		Ip:          Ip,
		Port:        Port,
		DbPath:      Db,
		SQLInitPath: SQLInitPath,
		SessionsKey: SessionsKey,
	}, nil
}
