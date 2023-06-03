package ifirebase

import (
	"context"
	"log"

	"firebase.google.com/go/v4"
	"firebase.google.com/go/v4/auth"
	"google.golang.org/api/option"

	"backend/utils/text"
)

func Init() (*firebase.App, *auth.Client) {
	opt := option.WithCredentialsFile(text.RelativePath("firebase-adminsdk.json"))
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		log.Fatalf("error initializing app: %v\n", err)
	}

	authClient, err := app.Auth(context.Background())
	if err != nil {
		log.Fatalf("error getting Auth client: %v\n", err)
	}

	return app, authClient
}
