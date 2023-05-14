package email

import (
	"log"
	"os"
	"strconv"

	gomail "gopkg.in/gomail.v2"
)

var (
	EMAIL_SMTP_HOST = os.Getenv("EMAIL_SMTP_HOST")
	EMAIL_SMTP_PORT = os.Getenv("EMAIL_SMTP_PORT")
	EMAIL_USER      = os.Getenv("EMAIL_SMTP_USER")
	EMAIL_PWD       = os.Getenv("EMAIL_SMTP_PWD")
)

func Send() {
	m := gomail.NewMessage()
	m.SetHeader("From", "1187688808@qq.com")
	// m.SetHeader("From", "keangjay@gmail.com")
	m.SetHeader("To", "keangj@outlook.com")
	// m.SetAddressHeader("Cc", "dan@example.com", "Dan")
	m.SetHeader("Subject", "Hello!")
	m.SetBody("text/html", "Hello <b>test</b> and <i>jay</i>!")
	// m.Attach("/home/Alex/lolcat.jpg")
	var d *gomail.Dialer
	if port, err := strconv.Atoi(EMAIL_SMTP_PORT); err != nil {
		log.Fatalln(err)
	} else {
		d = gomail.NewDialer(EMAIL_SMTP_HOST, port, EMAIL_USER, EMAIL_PWD)
		// d := gomail.NewDialer("smtp.gmail.com", 587, "keangjay@gmail.com", "")
	}
	// Send the email to Bob, Cora and Dan.
	if err := d.DialAndSend(m); err != nil {
		log.Println(err)
		panic(err)
	}
}
