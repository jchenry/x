package payments

import (
	"fmt"
	"os"
)

type Config struct {
	StripeKey       string
	StripeProductID string
	RedirectURL     string
	TenantSetup     func(subscriptionID, customerID string) (tenantID string)
}

func FromEnv() Config {
	return Config{
		StripeKey:       os.Getenv("STRIPE_KEY"),
		StripeProductID: os.Getenv("STRIPE_PRODUCT_ID"),
		RedirectURL:     "/",
	}
}

func PrintConfig() {
	fmt.Printf("%#v\n", FromEnv())
}
