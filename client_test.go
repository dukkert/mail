package mail_test

import (
	"fmt"
	"log"
	"testing"

	"github.com/dukkert/mail"
)

func TestMain(m *testing.M) {
	mailClient := &mail.MailClient{
		Key:                       "",
		VerifedReplySender:        "",
		VerifedReplySenderAddress: "",
	}

	client := mail.NewMailClient(mailClient)

	mail := &mail.Mail{
		From:        "",
		FromAddress: "",
		To:          "",
		ToAddress:   "",
		Subject:     "",
		Message:     "",
	}

	res, err := client.NewMail(mail)
	if err != nil {
		log.Println(err)
	}
	fmt.Print(string(res.ClientResponse.Body))

	replyRes, err := res.Reply("Hi again", "Just replying back")
	if err != nil {
		log.Println(err)
	}
	fmt.Print(string(replyRes.Body))
}
