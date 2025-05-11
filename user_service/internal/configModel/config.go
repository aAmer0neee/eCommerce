package configModel

import "time"

type Cfg struct {
	GRPC struct {
		Port    string        `yaml:"port" env:"PORT" env-required:"true"`
		Timeout time.Duration `yaml:"timeout" env-default:"5s"`
	} `yaml:"grpc" env-required:"true"`
	Logger struct {
		Level string `yaml:"level" env-default:"info"`
	} `yaml:"logger"`
	Postgres struct {
		Port     string `yaml:"port" env-required:"true"`
		Host     string `yaml:"host" env-required:"true"`
		Name     string `yaml:"name" env-required:"true"`
		Password string `yaml:"password" env-required:"true"`
		User     string `yaml:"user" env-required:"true"`
		Migrate  bool   `yaml:"migrate" env-default:"false"`
		Sslmode  string `yaml:"sslmode" env-default:"disable"`
	} `yaml:"postgres" env-required:"true"`
}