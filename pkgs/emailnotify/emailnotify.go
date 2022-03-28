package emailnotify

import (
	"crypto/tls"
	"fmt"
	"strconv"
	"time"

	"github.com/qobbysam/socketnotify/pkgs/config"
	gomail "gopkg.in/mail.v2"
)

type SmtpAuth struct {
	From     string
	Password string
	Host     string
	Port     string
}

type EmailNotifyExecutor struct {
	SmtpAuth *SmtpAuth
	Notify   []string
}

type Message struct {
	//Action string
	Subject string
	Msg     string
}

func NewEmailNotifyExecutor(cfg *config.EmailConfig) *EmailNotifyExecutor {

	out := EmailNotifyExecutor{}

	smtpauth := SmtpAuth{
		From:     cfg.Smtpfrom,
		Password: cfg.SmtpPassword,
		Host:     cfg.SmtpHost,
		Port:     cfg.SmtpPort,
	}

	out.Notify = cfg.Notify
	out.SmtpAuth = &smtpauth

	return &out

}

func (em *EmailNotifyExecutor) TestExecutorServer() error {

	msg := em.BuildTestMsg()

	err := em.SendOneMessage(*msg)

	return err
}

func (em *EmailNotifyExecutor) SendOneMessage(msg Message) error {
	// Sender data.
	// from := em.SmtpAuth.From
	// password := em.SmtpAuth.Password

	// // Receiver email address.
	// to := em.Notify

	// // smtp server configuration.
	// smtpHost := em.SmtpAuth.Host
	// smtpPort := em.SmtpAuth.Password

	// // Message.
	// //message := []byte("This is a test email message.")

	// message := msg
	// // Authentication.
	// auth := smtp.PlainAuth("", from, password, smtpHost)

	// // Sending email.
	// err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, message)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return err
	// }
	// fmt.Println("Email Sent Successfully!")
	m := gomail.NewMessage()

	// Set E-Mail sender
	m.SetHeader("From", em.SmtpAuth.From)

	// Set E-Mail receivers
	m.SetHeader("To", em.Notify...)

	// Set E-Mail subject
	m.SetHeader("Subject", msg.Subject)

	// Set E-Mail body. You can set plain text or html with text/html
	m.SetBody("text/plain", msg.Msg)

	// Settings for SMTP server
	port, _ := strconv.Atoi(em.SmtpAuth.Port)
	d := gomail.NewDialer(em.SmtpAuth.Host, port, em.SmtpAuth.From, em.SmtpAuth.Password)

	// This is only needed when SSL/TLS certificate is not valid on server.
	// In production this should be set to false.
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	// Now send E-Mail
	if err := d.DialAndSend(m); err != nil {
		fmt.Println(err)
		//panic(err)
	}

	return nil
}

func (em *EmailNotifyExecutor) BuildTestMsg() *Message {

	msg := time.Now().Format("02-02-2002 Mon 19:00:00")

	return &Message{
		Subject: "Server Is on ",

		Msg: "Server is on now   " + msg,
	}
}

func (em *EmailNotifyExecutor) BuildClickMsg(email string) *Message {

	subject := "New Click MSG"

	nowmsg := time.Now().Format("02-02-2002 Mon 19:00:00")

	msgout := fmt.Sprint(email, " \n", " click ", nowmsg)

	return &Message{Subject: subject, Msg: msgout}
}

func (em *EmailNotifyExecutor) BuildOpenMsg(email string) *Message {

	subject := "New Open MSG"

	nowmsg := time.Now().Format("02-02-2002 Mon 19:00:00")

	msgout := fmt.Sprint(email, " \n", " open  ", nowmsg)

	return &Message{Subject: subject, Msg: msgout}
}
