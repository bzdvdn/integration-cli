package templates

// MainGO ...
func MainGO() (string, string) {
	filename := "main.go"
	body := `package main

import (
	"flag"
	"log"
	"{{.}}/server"

	"github.com/BurntSushi/toml"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "config-path", "config/dev-config.toml", "path to config file")
}

func main() {
	flag.Parse()

	config := server.NewConfig()
	_, err := toml.DecodeFile(configPath, config)

	if err != nil {
		log.Fatal(err)
	}

	server := server.NewServer(config)

	if err := server.Start(); err != nil {
		log.Fatal(err)
	}
}

`
	return filename, body
}
