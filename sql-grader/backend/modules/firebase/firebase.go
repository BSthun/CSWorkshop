package ifirebase

import (
	"context"
	"log"

	"firebase.google.com/go/v4"
	"google.golang.org/api/option"

	"backend/modules"
)

func Init(b *modules.Base) {
	opt := option.WithCredentialsFile("firebase-adminsdk.json")
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		log.Fatalf("error initializing app: %v\n", err)
	}

	authClient, err := app.Auth(context.Background())
	if err != nil {
		log.Fatalf("error getting Auth client: %v\n", err)
	}

	b.FirebaseApp = app
	b.FirebaseAuth = authClient
}
