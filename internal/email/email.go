package email

import (
	"log"

	"github.com/spf13/viper"
	gomail "gopkg.in/gomail.v2"
)

func Send() {
	m := gomail.NewMessage()
	m.SetHeader("From", "1187688808@qq.com")
	// m.SetHeader("From", "keangjay@gmail.com")
	m.SetHeader("To", "keangj@outlook.com")
	// m.SetAddressHeader("Cc", "dan@example.com", "Dan")
	m.SetHeader("Subject", "Hello!")
	m.SetBody("text/html", "Hello <i>jay</i>!")
	// m.Attach("/home/Alex/lolcat.jpg")

	d := gomail.NewDialer(
		viper.GetString("email.smtp.host"),
		viper.GetInt("email.smtp.port"),
		viper.GetString("email.smtp.user"),
		viper.GetString("email.smtp.password"),
	)
	// d := gomail.NewDialer("smtp.gmail.com", 587, "keangjay@gmail.com", "")
	// Send the email to Bob, Cora and Dan.
	if err := d.DialAndSend(m); err != nil {
		log.Println(err)
		panic(err)
	}
}
