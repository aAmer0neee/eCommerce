package config

import (
	"flag"
	"log"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

var (
	configPath = flag.String("config", "", "specify path to configure file")
)

func MustLoad(cfg any) {
	flag.Parse()

	if *configPath == "" {
		log.Fatalln("empty value of required argument: conigure path")
	}

	if _, err := os.Stat(*configPath); err == os.ErrNotExist {
		log.Fatalf("file %s not exist", *configPath)
	}

	if err := cleanenv.ReadConfig(*configPath, cfg); err != nil {
		log.Fatalf("error read config message: %s", err.Error())
	}
}
