package config

import (
	"os"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
)

type Config struct {
	Env        string `yaml:"env" env-required:"true" env:"Env"`
	HTTPServer `yaml:"http_server"`
	Storage    `yaml:"storage"`
}

type HTTPServer struct {
	Address     string        `yaml:"address" env-required:"true"`
	Timeout     time.Duration `yaml:"timeout" env-default:"4s"`
	IdleTimeout time.Duration `yaml:"idle_timeout" env-default:"60s"`
}

type Storage struct {
	Host     string `yaml:"db_host" env-required:"true"`
	Port     int    `yaml:"db_port" env-default:"5432"`
	Name     string `yaml:"db_name" env_required:"true"`
	Username string `yaml:"db_username" env-required:"true"`
	Password string `yaml:"db_password" env-required:"true"`
}

func Init() *Config {
	var cfg Config

	err := godotenv.Load()

	if err != nil {
		panic("Couldn't load .env file!")
	}

	path, ok := os.LookupEnv("CONFIG_PATH")

	if !ok {
		panic("Couldn't find env variable for config path!")
	}

	err = cleanenv.ReadConfig(path, &cfg)

	if err != nil {
		panic("Couldn't read config!")
	}

	return &cfg
}
