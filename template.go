package main

import (
	"fmt"
	"log"

	"google.golang.org/api/iterator"
)

type Template struct {
	Template_name string `json:"template_name" firestore:"-"`
	Template_body string `binding:"required" json:"template_body" firestore:"template_body,omitempty"`
}

func (t Template) postTemplateFirebase() error {

	client, ctx, err := getFireStore()
	if err != nil {
		panic(err)
	}
	result, err := client.Collection("template").Doc(t.Template_name).Set(ctx, t)
	if err != nil {
		return err
	}
	fmt.Print(result)
	return nil
}

func getTemplateFirebase(temp_name string) ([]Template, error) {

	client, ctx, err := getFireStore()
	if err != nil {
		panic(err)
	}

	//obtener los template por nombre
	if temp_name == "" {
		sliceTemp := make([]Template, 0)
		array := client.Collection("template").Documents(ctx)
		i := 0
		for {
			var tempElem Template
			doc, err := array.Next()

			if err == iterator.Done {
				break
			}
			if err != nil {
				return nil, err
			}
			tempElem.Template_name = doc.Ref.ID
			doc.DataTo(&tempElem)
			sliceTemp = append(sliceTemp, tempElem)
			i++
		}
		return sliceTemp, nil
	} else {
		sliceTemp := make([]Template, 0)
		var tempzero Template
		doc, err := client.Collection("template").Doc(temp_name).Get(ctx)
		if err != nil {
			return nil, err
		}
		tempzero.Template_name = doc.Ref.ID
		err = doc.DataTo(&tempzero)
		if err != nil {
			log.Print("error datato ", err)
		}
		sliceTemp = append(sliceTemp, tempzero)
		return sliceTemp, nil
	}
}
