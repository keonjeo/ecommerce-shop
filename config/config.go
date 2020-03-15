package config

import (
	"fmt"
	"log"
	"time"

	"github.com/caarlos0/env/v6"
	"github.com/joho/godotenv"
)

// AppConfig ...
type AppConfig struct {
	Host string `env:"HOST" envDefault:"127.0.0.1"`
	Port int    `env:"PORT" envDefault:"3001"`
	ENV  string `env:"ENV" envDefault:"development"`
}

// DatabaseConfig ...
type DatabaseConfig struct {
	PostgresHost string `env:"POSTGRES_HOST"`
	PostgresURI  string `env:"POSTGRES_URI"`
	PostgresDB   string `env:"POSTGRES_DB"`
	PostgresUser string `env:"POSTGRES_USER"`
	PostgresPass string `env:"POSTGRES_PASSWORD"`
}

// AuthConfig ...
type AuthConfig struct {
	VerificationRequired  bool   `env:"VERIFICATION_REQUIRED" envDefault:"false"`
	ResetPasswordValidFor int    `env:"RESET_PASSWORD_VALID_FOR" envDefault:"9999"`
	AccessTokenSecret     string `env:"ACCESS_TOKEN_SECRET" envDefault:"xxxxx"`
	RefreshTokenSecret    string `env:"REFRESH_TOKEN_SECRET" envDefault:"xxxxx"`
}

// EmailConfig ...
type EmailConfig struct {
	Enabled   bool   `env:"EMAIL_ENABLED" envDefault:"false"`
	Transport string `env:"EMAIL_TRANSPORT" envDefault:"sendgrid"`
	From      string `env:"EMAIL_FROM" envDefault:"dp24031995@gmail.com"`
	Host      string `env:"EMAIL_HOST"`
	Port      int    `env:"EMAIL_PORT"`
	User      string `env:"EMAIL_USER"`
	Pass      string `env:"EMAIL_PASS"`
}

// CookieConfig ...
type CookieConfig struct {
	Name     string    `env:"COOKIE_NAME" envDefault:"cookie"`
	Path     string    `env:"COOKIE_PATH" envDefault:"/"`
	Secret   string    `env:"COOKIE_SECRET" envDefault:"xxxxx"`
	HTTPOnly bool      `env:"COOKIE_HTTP_ONLY" envDefault:"false"`
	Secure   bool      `env:"COOKIE_SECURE" envDefault:"false"`
	MaxAge   time.Time `env:"COOKIE_MAX_AGE" envDefault:"0"`
}

// Config represents the app config
type Config struct {
	App    AppConfig
	DB     DatabaseConfig
	Auth   AuthConfig
	Email  EmailConfig
	Cookie CookieConfig
}

func (c Config) isDev() bool {
	return c.App.ENV == "development"
}
func (c Config) isProd() bool {
	return c.App.ENV == "production"
}
func (c Config) isTest() bool {
	return c.App.ENV == "test"
}

func loadEnvironment() {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal("error loading .env file")
	}
}

func init() {
	loadEnvironment()
	cfg := Config{}
	if err := env.Parse(&cfg); err != nil {
		panic("config error")
	}
	fmt.Printf("%+v\n", cfg)
}
