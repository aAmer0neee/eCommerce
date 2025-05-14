package domain

import "time"

type User struct {
	Email        string
	Name         string
	PasswordHash string
	Access       string
	Refresh      string
}

type Cfg struct {
	Server struct {
		Port string `yaml:"port" env:"PORT" env-default:"8080"`
		Host string `yaml:"host" env:"HOST" env-default:"localhost"`
	} `yaml:"server" env-required:"true"`
	Logger struct {
		Level string `yaml:"level" env-default:"info"`
	} `yaml:"logger"`

	Services struct {
		Timeout time.Duration `yaml:"timeout" env-dafault:"5s"`
		User    string        `yaml:"user"`
	} `yaml:"services"`
}
