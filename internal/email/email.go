package email

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
	gomail "gopkg.in/gomail.v2"
)

func newMessage(to, subject, body string) *gomail.Message {
	m := gomail.NewMessage()
	m.SetHeader("From", "1187688808@qq.com")
	// m.SetHeader("From", "keangjay@gmail.com")
	m.SetHeader("To", to)
	// m.SetAddressHeader("Cc", "dan@example.com", "Dan")
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", body)
	// m.Attach("/home/Alex/lolcat.jpg")
	return m
}
func newDialer() *gomail.Dialer {
	d := gomail.NewDialer(
		viper.GetString("email.smtp.host"),
		viper.GetInt("email.smtp.port"),
		viper.GetString("email.smtp.user"),
		viper.GetString("email.smtp.password"),
	)
	return d
}

func Send() {
	m := newMessage("keangj@outlook.com", "Hello!", "Hello <i>jay</i>!")
	d := newDialer()
	if err := d.DialAndSend(m); err != nil {
		log.Println(err)
		panic(err)
	}
}

func SendValidateCode(email, code string) error {
	m := newMessage(
		email,
		fmt.Sprintf("[%s] 验证码", code),
		fmt.Sprintf("您的验证码为: %s", code),
	)
	d := newDialer()
	return d.DialAndSend(m)
}
