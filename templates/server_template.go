package templates

// Gomod ....
func Gomod() (string, string) {
	filename := `go.mod`
	body := `module {{.}}

go 1.12

require (
	github.com/BurntSushi/toml v0.3.1
	github.com/gin-contrib/sse v0.1.0
	github.com/gin-gonic/gin v1.4.0
	github.com/google/uuid v1.1.1
	github.com/jinzhu/gorm v1.9.10
	github.com/sirupsen/logrus v1.4.2
)`
	return filename, body
}

// ServerGO ...
func ServerGO() (string, string) {
	filename := "server/server.go"
	body := `package server

import (
	"log"
	"{{.}}/api"
	"{{.}}/database"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

// Server ...
type Server struct {
	config *Config
	db     *gorm.DB
	router *gin.Engine
}

// NewServer ...
func NewServer(c *Config) *Server {
	return &Server{
		config: c,
		router: gin.Default(),
	}
}

// Start ...
func (s *Server) Start() error {
	db, err := database.Init(
		s.config.DbUser,
		s.config.DbPassword,
		s.config.DbName,
		s.config.DbHost,
		s.config.DbPort,
	)
	if err != nil {
		log.Fatal(err)
	}
	s.db = db
	api.Health(s.router)
	s.router.Run()
	return nil
}
`
	return filename, body
}

// ConfigGO ...
func ConfigGO() (string, string) {
	filename := "server/config.go"
	body := `package server

// Config ...
type Config struct {
	BindAddr   string 'toml:"bind_addr"'
	DbUser     string 'toml:"db_user"'
	DbPassword string 'toml:"db_password"'
	DbHost     string 'toml:"db_host"'
	DbName     string 'toml:"db_name"'
	DbPort     string 'toml:"db_port"'
}

// NewConfig ...
func NewConfig() *Config {
	return &Config{
		BindAddr:   ":8080",
		DbUser:     "",
		DbPassword: "",
		DbHost:     "localhost",
		DbName:     "",
		DbPort:     "5432",
	}
}
`
	return filename, body
}

// RouterGO ...
func RouterGO() (string, string) {
	filename := "server/router.go"
	body := `package server
`
	return filename, body
}
