package mail

import (
	"errors"
	"fmt"

	"github.com/sendgrid/rest"
	sg "github.com/sendgrid/sendgrid-go"
	sgMail "github.com/sendgrid/sendgrid-go/helpers/mail"
)

var sgClient *sg.Client

type MailClient struct {
	Key                       string
	client                    *sg.Client
	VerifedReplySender        string
	VerifedReplySenderAddress string
}

type Mail struct {
	From        string
	FromAddress string
	To          string
	ToAddress   string
	Subject     string
	Message     string
}

type MailReply struct {
	From *sgMail.Email
	To   *sgMail.Email
}

type MailResponse struct {
	ClientResponse *rest.Response
	Mail           *MailReply
}

func NewMailClient(client *MailClient) *MailClient {
	sgClient = sg.NewSendClient(client.Key)
	return &MailClient{
		Key:    client.Key,
		client: sgClient,
	}
}

func (mc *MailClient) NewMail(mail *Mail) (*MailResponse, error) {
	from := sgMail.NewEmail(mail.From, mail.FromAddress)
	to := sgMail.NewEmail(mail.To, mail.ToAddress)

	message := sgMail.NewSingleEmail(from, mail.Subject, to, mail.Message, "")

	resp, err := mc.client.Send(message)

	fmt.Print("Sent the message")

	if err != nil {
		return nil, errors.New("failed to send message")
	}

	verifiedReplyEmail := sgMail.NewEmail(mc.VerifedReplySender, mc.VerifedReplySenderAddress)

	response := &MailResponse{
		ClientResponse: resp,
		Mail: &MailReply{
			From: verifiedReplyEmail,
			To:   from,
		},
	}

	return response, nil
}

func (mr *MailResponse) Reply(subject string, replyMessage string) (*rest.Response, error) {
	message := sgMail.NewSingleEmail(mr.Mail.From, subject, mr.Mail.To, replyMessage, "")

	res, err := sgClient.Send(message)
	if err != nil {
		return nil, errors.New("failed to reply")
	}

	return res, nil
}
