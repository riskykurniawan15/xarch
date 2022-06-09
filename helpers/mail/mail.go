package mail

import (
	mail_server "github.com/xhit/go-simple-mail/v2"

	"github.com/riskykurniawan15/xarch/config"
)

func Send(cfg config.Config, to, subject, body string) error {
	server := mail_server.NewSMTPClient()
	server.Host = cfg.Email.EMAIL_HOST
	server.Port = cfg.Email.EMAIL_PORT
	server.Username = cfg.Email.EMAIL_EMAIL
	server.Password = cfg.Email.EMAIL_PASSWORD
	server.Encryption = mail_server.EncryptionSSLTLS

	smtpClient, err := server.Connect()
	if err != nil {
		return err
	}

	email := mail_server.NewMSG()
	email.SetFrom(cfg.Email.EMAIL_NAME + " <" + cfg.Email.EMAIL_EMAIL + ">")
	email.AddTo(to)
	email.SetSubject(subject)

	email.SetBody(mail_server.TextHTML, body)

	// Send email
	err = email.Send(smtpClient)
	if err != nil {
		return err
	}

	return nil
}
