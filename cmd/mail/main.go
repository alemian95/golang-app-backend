package main

import (
	"bytes"
	"golang-app/app/utils/mail"
	"golang-app/cmd"
	"html/template"
)

func main() {

	cmd.InitCmdApplication()

	t, err := template.ParseFiles("./cmd/mail/template.html")

	if err != nil {
		panic(err.Error())
	}

	var b bytes.Buffer

	t.Execute(&b, struct {
		Body string
	}{
		Body: "Test Body",
	})

	d, m := mail.Prepare("admin@example.com", "user@example.com", "Test Subject", b.String())
	mail.Send(d, m)
}
