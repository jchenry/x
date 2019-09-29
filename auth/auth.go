package auth

import (
	"context"
	"log"

	"golang.org/x/oauth2"

	oidc "github.com/coreos/go-oidc"
)

type Authenticator struct {
	Provider *oidc.Provider
	Config   oauth2.Config
	Ctx      context.Context
}

func NewAuthenticator(domain, clientID, clientSecret, callback string) (*Authenticator, error) {
	ctx := context.Background()

	provider, err := oidc.NewProvider(ctx, domain)
	if err != nil {
		log.Printf("failed to get provider: %v", err)
		return nil, err
	}

	conf := oauth2.Config{
		ClientID:     "ae1e02bTwXA35O3r3Xxk4kbRf31j5ge9",
		ClientSecret: "NFC5KYeM9GA2z0ptvzKPo9jmkQDRjx_WcsWyK0hzOJmr1CykS9cEmTcNh0-hKiMd",
		RedirectURL:  "http://localhost:3000/callback",
		Endpoint:     provider.Endpoint(),
		Scopes:       []string{oidc.ScopeOpenID, "profile"},
	}

	return &Authenticator{
		Provider: provider,
		Config:   conf,
		Ctx:      ctx,
	}, nil
}
