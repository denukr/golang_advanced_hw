package verify

import (
	"fmt"
	"go/email-verify/config"
	"net/smtp"

	"github.com/jordan-wright/email"
)

func SendEmail(hash string, cfg *config.VerifyConfig) {
	e := email.NewEmail()
	e.From = "Jordan Wright <noreplay@gmail.com>"
	e.To = []string{"denisukrainskii123@gmail.com"}
	e.Subject = "Awesome Subject"
	e.Text = []byte("Text Body is, of course, supported!")
	text := fmt.Sprintf(`<a href="http://localhost:8081/verify/%s">Нажми сюда</a>`, hash)
	e.HTML = []byte(text)
	auth := smtp.PlainAuth("", "denisukrainskii123@gmail.com", cfg.Password, cfg.Address) // Пароль без пробелов
	err := e.Send(fmt.Sprintf("%s:587", cfg.Address), auth)
	if err != nil {
		fmt.Println("Ошибка отправки:", err)
	}
}
