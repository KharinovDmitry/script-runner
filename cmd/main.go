package main

import (
	"TestTask-PGPro/internal/app"
	"TestTask-PGPro/internal/config"
	"flag"
	"log"
)

func main() {
	var configPath string
	flag.StringVar(&configPath, "path", "", "config path")
	flag.Parse()
	if configPath == "" {
		log.Fatal("Необходимо указать путь до файлов конфигурации")
	}
	cfg := config.MustLoad(configPath)

	app.MustRun(cfg)
}
