package main

import (
	"bytes"
	"fmt"
	"net/smtp"
	"text/template"
)

// to send email
func sendSimpleMail(subject, body string, to []string) {
	auth := smtp.PlainAuth(
		"",
		"fazalbkabeer@gmail.com",
		"ldsc mwya jqur vrey",
		"smtp.gmail.com",
	)

	msg := "Subject: " + subject + "\n" + body

	err := smtp.SendMail(
		"smtp.gmail.com:587",
		auth,
		"fazalbkabeer@gmail.com",
		to,
		[]byte(msg),
	)

	if err != nil {
		fmt.Println(err)
	}
}

// to send html
func sendSimpleMailHTML(subject, html string, to []string) {
	auth := smtp.PlainAuth(
		"",
		"fazalbkabeer@gmail.com",
		"ldsc mwya jqur vrey",
		"smtp.gmail.com",
	)

	Headers := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";"

	msg := "Subject: " + subject + "\n" + Headers + "\n\n" + html

	err := smtp.SendMail(
		"smtp.gmail.com:587",
		auth,
		"fazalbkabeer@gmail.com",
		to,
		[]byte(msg),
	)

	if err != nil {
		fmt.Println(err)
	}
}

// to send html template
func sendSimplHTMLTemplate(subject, templatePath string, to []string) {

	// Get html
	var body bytes.Buffer
	t, err := template.ParseFiles(templatePath)

	if err != nil {
		fmt.Println(err)
		return
	}

	err = t.Execute(&body, struct{ Name string }{Name: "fazal"})
	if err != nil {
		fmt.Printf("Error executing template: %v\n", err)
		return
	}

	auth := smtp.PlainAuth(
		"",
		"fazalbkabeer@gmail.com",
		"ldsc mwya jqur vrey",
		"smtp.gmail.com",
	)

	Headers := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";"

	msg := "Subject: " + subject + "\n" + Headers + "\n\n" + body.String()

	err = smtp.SendMail(
		"smtp.gmail.com:587",
		auth,
		"fazalbkabeer@gmail.com",
		to,
		[]byte(msg),
	)

	if err != nil {
		fmt.Println(err)
	}
}

func main() {
	// sendSimpleMail("Second Mail",
	// 	"Second Body",
	// 	[]string{"fazalkdy306@gmail.com"})

	// sendSimpleMailHTML("HTML Mail",
	// 	"<h1>Body Heading<h/1><p>Body Paragraph...</p>",
	// 	[]string{"fazalkdy306@gmail.com"})/

	sendSimplHTMLTemplate("Template Mail",
		"test.html",
		[]string{"fazalkdy306@gmail.com"})

}
