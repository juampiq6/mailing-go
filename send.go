package main

import (
	"bytes"
	html "html/template"
	"log"

	gomail "gopkg.in/gomail.v2"
)

type callBody struct {
	Direccion []string `binding:"required" json:"direccion"`
	IDUsuario []int    `json:"id_usuario"`
	Subject   string   `binding:"required" json:"subject"`
}

func parseTemplate(templateFileName string, data interface{}) (string, error) {
	templateData := struct {
		Name string
		URL  string
	}{
		Name: "Dhanush",
		URL:  "http://geektrust.in",
	}
	parseo, err := html.ParseFiles(templateFileName)
	if err != nil {
		return "", err
	}
	buf := new(bytes.Buffer)
	if err = parseo.Execute(buf, templateData); err != nil {
		return "", err
	}
	result := buf.String()
	return result, nil
}

func (b callBody) sendSpecific(tempname string) error {
	m := gomail.NewMessage()
	m.SetHeader("From", "microservicios.mailing@gmail.com")
	m.SetHeader("To", b.Direccion...)
	// m.SetAddressHeader("Cc", "dan@example.com", "Dan")
	m.SetHeader("Subject", b.Subject)
	templateData := struct {
		Name string
		URL  string
	}{
		Name: "Pablo",
		URL:  "http://facebook.com",
	}
	res, err := parseTemplate("./templateee.html", templateData)
	if err != nil {
		log.Print(err)
		return err
	}
	log.Print(res)
	m.SetBody("text/html", res)

	d := gomail.NewDialer("smtp.gmail.com", 587, "microservicios.mailing@gmail.com", "lkssoxmjqoyywgnb")

	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}
	return nil
}

// func getListaBroadcast() []string {
// 	client, ctx, err := getFireStore()
// 	if err != nil {
// 		panic(err)
// 	}
// 	var direcciones []string

// 	return direcciones
// }
