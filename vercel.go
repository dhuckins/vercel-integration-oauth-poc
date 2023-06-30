package main

import (
	"context"
	"os"

	"golang.org/x/oauth2"
)

func getAccessToken(code string) (*oauth2.Token, error) {
	config := &oauth2.Config{
		ClientID:     os.Getenv("VERCEL_CLIENT_ID"),           // vercel provided
		ClientSecret: os.Getenv("VERCEL_CLIENT_SECRET"),       // vercel provided
		RedirectURL:  "http://localhost:8000/vercel/callback", // MUST be what we put in the vercel integration config
		Endpoint: oauth2.Endpoint{
			TokenURL:  "https://api.vercel.com/v2/oauth/access_token",
			AuthURL:   "",                       // not provided/not needed
			AuthStyle: oauth2.AuthStyleInParams, // per documentation
		},
		Scopes: []string{"project-env-vars"}, // not sure if needed
	}

	token, err := config.Exchange(context.Background(), code)
	if err != nil {
		return nil, err
	}
	return token, nil
}
