package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gopkg.in/gomail.v2"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	var file string
	fmt.Print("give me a full file path: ")
	fmt.Scan(&file)

	var address string
	fmt.Print("give me a sending address: ")
	fmt.Scan(&address)

	m := gomail.NewMessage()
	m.SetHeader("From", "MS_gxA32N@trial-3vz9dlenp9nlkj50.mlsender.net")
	m.SetHeader("To", address)
	m.SetHeader("Subject", "Bijlagen")
	m.SetBody("text/html", "Hoi, zie bijlagen!")
	m.Attach(file) //pas aan

	d := gomail.NewDialer("smtp.mailersend.net", 587, os.Getenv("NAME"), os.Getenv("PASSWORD")) //change name and password

	// Send the email
	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}
}
