package config

import (
	"html/template"
	"log"

	"github.com/alexedwards/scs/v2"
)

// AppConfig holds the application config
type AppConfig struct {
	UseCache      bool                          `json:"use_cache"`
	TemplateCache map[string]*template.Template `json:"template_cache"`
	InfoLog       *log.Logger                   `json:"info_log"`
	InProduction  bool                          `json:"in_production"`
	Session       *scs.SessionManager           `json:"session"`
}
