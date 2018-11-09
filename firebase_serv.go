package main

import (
	"context"

	"cloud.google.com/go/firestore"
	"firebase.google.com/go"
	"google.golang.org/api/option"
)

func getFireStore() (*firestore.Client, context.Context, error) {

	ctx := context.Background()
	accountkey := option.WithCredentialsFile("./mailing-e-commerce-firebase-adminsdk-2jetw-1edd40aaf3.json")
	app, err := firebase.NewApp(ctx, nil, accountkey)
	if err != nil {
		return nil, nil, err
	}
	client, err := app.Firestore(ctx)
	if err != nil {
		return nil, nil, err
	}
	return client, ctx, nil
}
