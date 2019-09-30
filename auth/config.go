package auth

import (
	"fmt"
	"os"
)

type Config struct {
	Domain       string
	ClientID     string
	ClientSecret string
	CallbackURL  string
	RedirectURL  string
}

func FromEnv() Config {
	return Config{
		Domain:       os.Getenv("AUTH_DOMAIN"),
		ClientID:     os.Getenv("AUTH_CLIENT_ID"),
		ClientSecret: os.Getenv("AUTH_CLIENT_SECRET"),
		CallbackURL:  os.Getenv("AUTH_CALLBACK_URL"),
		RedirectURL:  "/user",
	}
}

func PrintConfig() {
	fmt.Printf("%#v\n", FromEnv())
}

type CallbackFunc func(c Config, u User) error
