package main

import gomail "gopkg.in/gomail.v2"

type CallBody struct {
	Direccion  []string `binding:"required" json:"direccion"`
	Id_usuario []int    `json:"id_usuario"`
	Subject    string   `binding:"required" json:"subject"`
}

func (b CallBody) sendSpecific(tempname string) error {
	m := gomail.NewMessage()
	m.SetHeader("From", "microservicios.mailing@gmail.com")
	m.SetHeader("To", b.Direccion...)
	// m.SetAddressHeader("Cc", "dan@example.com", "Dan")
	m.SetHeader("Subject", b.Subject)
	m.SetBody("text/html", "<b> Hello template name recibido: </b> <i>"+tempname+"</i>!")

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
