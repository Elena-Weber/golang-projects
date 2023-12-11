package config

import (
	"html/template"
	"log"
	"github.com/alexedwards/scs/v2"
)

// AppConfig is config available all throughout app right away
type AppConfig struct {
	UseCache bool
	TemplateCache map[string]*template.Template
	InfoLog *log.Logger
	InProduction bool
	Session *scs.SessionManager
}