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
	ENV  string `env:"ENV" envDefault:development`
}

// DatabaseConfig ...
type DatabaseConfig struct {
	DBHost string `env:"DB_HOST"`
	DBURI  string `env:"DB_URI"`
	DBName string `env:"DB_NAME"`
	DBUser string `env:"DB_USER"`
	DBPass string `env:"DB_PASSWORD"`
}

// AuthConfig ...
type AuthConfig struct {
	VerificationRequired  bool   `env:"VERIFICATION_REQUIRED" envDefault:false`
	ResetPasswordValidFor int    `env:"RESET_PASSWORD_VALID_FOR" envDefault:9999`
	AccessTokenSecret     string `env:"ACCESS_TOKEN_SECRET" envDefault:xxxxx`
	RefreshTokenSecret    string `env:"REFRESH_TOKEN_SECRET" envDefault:xxxxx`
}

// EmailConfig ...
type EmailConfig struct {
	Enabled   bool   `env:"EMAIL_ENABLED" envDefault:false`
	Transport string `env:"EMAIL_TRANSPORT" envDefault:"sendgrid"`
	From      string `env:"EMAIL_FROM" envDefault:"dp24031995@gmail.com"`
	Host      string `env:"EMAIL_HOST"`
	Port      int    `env:"EMAIL_PORT"`
	User      string `env:"EMAIL_USER"`
	Pass      string `env:"EMAIL_PASS"`
}

// CookieConfig ...
type CookieConfig struct {
	Name     string    `env:"COOKIE_NAME", default:"cookie"`
	Path     string    `env:"COOKIE_PATH" default:"/"`
	Secret   string    `env:"COOKIE_SECRET" default:"xxxxx"`
	HTTPOnly bool      `env:"COOKIE_HTTP_ONLY" default:false`
	Secure   bool      `env:"COOKIE_SECURE" default:false`
	MaxAge   time.Time `env:"COOKIE_MAX_AGE" default:0`
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
	if c.App.ENV == "development" {
		return true
	}
	return false
}
func (c Config) isProd() bool {
	if c.App.ENV == "production" {
		return true
	}
	return false
}
func (c Config) isTest() bool {
	if c.App.ENV == "test" {
		return true
	}
	return false
}

func loadEnvironment() {
	err := godotenv.Load("./env/.env")
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
