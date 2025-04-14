package config

import (
	"log"
)

// AppConfig provides a way to share the configuration of the app globally
type AppConfig struct {
	InfoLog *log.Logger
}
