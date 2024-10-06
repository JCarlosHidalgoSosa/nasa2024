package main

import (
	"bytes"
	"fmt"
	"net/smtp"
	"html/template"
	"github.com/gofiber/fiber/v2"
    "github.com/gofiber/fiber/v2/middleware/cors"
	"log"
	"time"
)

func sendMailSimple(subject string, body string, to string) {
	auth := smtp.PlainAuth (
		"",
		"casiperfectos573@gmail.com",
		"rotvvtimfiyonojg",
		"smtp.gmail.com",
	)

	msg := "Subject: " + subject + "\n" + body
	err := smtp.SendMail(
		"smtp.gmail.com:587",
		auth,
		"casiperfectos573@gmail.com",
		[]string{to},
		[]byte(msg),
	)

	if err != nil {
		fmt.Println(err)
	}
}

func sendMailSimpleHTML(subject string, to string) {
	var body bytes.Buffer

	tmpl, err := template.ParseFiles("index.html")
	
	tmpl.Execute(&body, struct{ Name string }{Name:"casi"})
	auth := smtp.PlainAuth (
		"",
		"casiperfectos573@gmail.com",
		"rotvvtimfiyonojg",
		"smtp.gmail.com",
	)

	headers :=  "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";"

	msg := "Subject: " + subject + "\n" + headers + "\n\n" + body.String()
	
	err = smtp.SendMail(
		"smtp.gmail.com:587",
		auth,
		"casiperfectos573@gmail.com",
		[]string{to},
		[]byte(msg),
	)

	if err != nil {
		fmt.Println(err)
	}
}

func main() {
	app:= fiber.New()

	app.Use(cors.New())

	app.Post("/:email", func(c *fiber.Ctx) error {
		file, err := c.FormFile("body")

		if err != nil {
			c.WriteString("error")
		}

		c.SaveFile(file, fmt.Sprintf("./%s", "index.html"))
		
		actualTime := time.Now().Format(time.RFC850)

		sendMailSimpleHTML(
			"REPORT: " + actualTime,
			c.Params("email","casiperfectos573@gmail.com"),
		)


		// Save file to root directory:
		return c.JSON(fiber.Map{
			"time" : actualTime,
		})
	})
	
	log.Fatal(app.Listen(":8000"))
}