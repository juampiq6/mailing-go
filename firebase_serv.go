package main

import (
	"context"
	"fmt"
	"log"

	"firebase.google.com/go"
	"google.golang.org/api/option"
)

type Template struct {
	template_id   int
	template_name string
	template_body string
}

func (t Template) postTemplateFirebase() int {

	//configuracion firebase
	ctx := context.Background()
	accountkey := option.WithCredentialsFile("./mailing-e-commerce-firebase-adminsdk-2jetw-1edd40aaf3.json")
	app, err := firebase.NewApp(ctx, nil, accountkey)
	if err != nil {
		log.Fatal(err)
	}
	client, err := app.Firestore(ctx)
	if err != nil {
		log.Fatal(err)
	}

	//nuevoTemplate := Template{template_id: templ_id, template_name: templ_name, template_body: templ_body}
	if t.template_id == 0 {
		result, err := client.Collection("template").Doc("eZOdtrDhI50fVVKCfLLV").Set(ctx, t)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Print(result)

	}
	return 200

	// 	log.Print(result)
	// 	getter, err := client.Collection("template").Doc("nombre").Get(ctx)
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	log.Print(dsnap.Data())
}
