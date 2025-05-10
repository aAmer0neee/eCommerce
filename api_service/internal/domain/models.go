package domain

type ApiCfg struct {
	Server struct {
		Port string `yaml:"port" env:"PORT" env-default:"8080"`
		Host string `yaml:"host" env:"HOST" env-default:"localhost"`
	} `yaml:"server" env-required:"true"`
	Logger struct {
		Level string `yaml:"level" env-default:"info"`
	} `yaml:"logger"`
}
