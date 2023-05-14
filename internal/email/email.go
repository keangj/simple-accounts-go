package email

import (
	"log"

	gomail "gopkg.in/gomail.v2"
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

	d := gomail.NewDialer("smtp.qq.com", 465, "1187688808@qq.com", "")
	// d := gomail.NewDialer("smtp.gmail.com", 587, "keangjay@gmail.com", "")

	// Send the email to Bob, Cora and Dan.
	if err := d.DialAndSend(m); err != nil {
		log.Println(err)
		panic(err)
	}
}
