package config

import (
	"html/template"
	"log"
)

// AppConfig is app config
type AppConfig struct {
	UseCache bool
	TemplateCache map[string]*template.Template
	InfoLog *log.Logger
}