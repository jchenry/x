package payments

import (
	"fmt"
	"os"
)

type Config struct {
	StripeKey       string
	StripeProductID string
}

func FromEnv() Config {
	return Config{
		StripeKey:       os.Getenv("STRIPE_KEY"),
		StripeProductID: os.Getenv("STRIPE_PRODUCT_ID"),
	}
}

func PrintConfig() {
	fmt.Printf("%#v\n", FromEnv())
}
