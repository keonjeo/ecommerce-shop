package config

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

// AppSettings ...
type AppSettings struct {
	Host string `envconfig:"HOST"`
	Port int    `envconfig:"PORT"`
	ENV  string `envconfig:"ENV"`
}

// DatabaseSettings ...
type DatabaseSettings struct {
	PostgresHost string `envconfig:"POSTGRES_HOST"`
	PostgresURI  string `envconfig:"POSTGRES_URI"`
	PostgresDB   string `envconfig:"POSTGRES_DB"`
	PostgresUser string `envconfig:"POSTGRES_USER"`
	PostgresPass string `envconfig:"POSTGRES_PASSWORD"`
}

// AuthSettings ...
type AuthSettings struct {
	VerificationRequired  bool   `envconfig:"VERIFICATION_REQUIRED"`
	ResetPasswordValidFor int    `envconfig:"RESET_PASSWORD_VALID_FOR"`
	AccessTokenSecret     string `envconfig:"ACCESS_TOKEN_SECRET"`
	RefreshTokenSecret    string `envconfig:"REFRESH_TOKEN_SECRET"`
}

// EmailSettings ...
type EmailSettings struct {
	Enabled   bool   `envconfig:"EMAIL_ENABLED"`
	Transport string `envconfig:"EMAIL_TRANSPORT"`
	From      string `envconfig:"EMAIL_FROM"`
	Host      string `envconfig:"EMAIL_HOST"`
	Port      int    `envconfig:"EMAIL_PORT"`
	User      string `envconfig:"EMAIL_USER"`
	Pass      string `envconfig:"EMAIL_PASS"`
}

// CookieSettings ...
type CookieSettings struct {
	Name     string `envconfig:"COOKIE_NAME"`
	Path     string `envconfig:"COOKIE_PATH"`
	Secret   string `envconfig:"COOKIE_SECRET"`
	HTTPOnly bool   `envconfig:"COOKIE_HTTP_ONLY"`
	Secure   bool   `envconfig:"COOKIE_SECURE"`
	MaxAge   int    `envconfig:"COOKIE_MAX_AGE"`
}

// PasswordSettings ...
type PasswordSettings struct {
	MinLength int
	MaxLength int
	Lowercase bool
	Uppercase bool
	Number    bool
	Symbol    bool
}

// Config represents the app config
type Config struct {
	AppSettings
	DatabaseSettings DatabaseSettings
	AuthSettings     AuthSettings
	EmailSettings    EmailSettings
	CookieSettings   CookieSettings
	PasswordSettings PasswordSettings
}

func loadEnvironment() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("error loading .env file: %v", err)
	}
}

// ApplyDefaults sets all config default values
func (c *Config) ApplyDefaults() {
	c.AppSettings.SetDefaults()
	c.DatabaseSettings.SetDefaults()
	c.AuthSettings.SetDefaults()
	c.EmailSettings.SetDefaults()
	c.CookieSettings.SetDefaults()
	c.PasswordSettings.SetDefaults()
}

// New creates the new config
func New() *Config {
	loadEnvironment()

	cfg := &Config{}
	if err := envconfig.Process("", cfg); err != nil {
		panic(err)
	}

	cfg.ApplyDefaults()

	return cfg
}
