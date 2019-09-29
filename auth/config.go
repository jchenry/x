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
	CallbackFunc CallbackFunc
}

func FromEnv(c CallbackFunc) Config {
	return Config{
		Domain:       os.Getenv("AUTH_DOMAIN"),
		ClientID:     os.Getenv("AUTH_CLIENT_ID"),
		ClientSecret: os.Getenv("AUTH_CLIENT_SECRET"),
		CallbackURL:  os.Getenv("AUTH_CALLBACK_URL"),
		CallbackFunc: c,
	}
}

func PrintConfig() {
	fmt.Printf("%#v\n", FromEnv(nil))
}

type CallbackFunc func(c Config, u User) error
