package emailnotify

import (
	"crypto/tls"
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/qobbysam/socketnotify/pkgs/config"
	//"github.com/qobbysam/socketnotify/pkgs/cronn"
	"github.com/qobbysam/socketnotify/pkgs/locdb"
	gomail "gopkg.in/mail.v2"
)

type SmtpAuth struct {
	From     string
	Password string
	Host     string
	Port     string
	//CanSend  bool
}

type EmailNotifyExecutor struct {
	SmtpAuth *SmtpAuth
	Notify   []string
	CanSend  bool
	DB       *locdb.DBS
	AuthKey  string
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
	out.CanSend = cfg.CanSend
	out.AuthKey = cfg.AuthKey

	return &out

}

func (em *EmailNotifyExecutor) TestExecutorServer() error {

	msg := em.BuildTestMsg()

	err := em.SendOneMessage(*msg)

	return err
}

func (em *EmailNotifyExecutor) CanSendOn() bool {

	if em.CanSend {
		log.Println("Can send is on")
		//return em.CanSend
	} else {
		log.Println("Can send if off")

	}

	return em.CanSend

}

func (em *EmailNotifyExecutor) LockCanSend() {

	em.CanSend = false

	log.Println("locking can send")
}

func (em *EmailNotifyExecutor) UnlockCanSend() {
	em.CanSend = true
	log.Println("unlocking can send")
}

func (em *EmailNotifyExecutor) SendGenMsg(msg Message, notify []string) error {
	if !em.CanSendOn() {
		log.Println("call was success but we can't send yet")

		return nil

	}
	m := gomail.NewMessage()

	// Set E-Mail sender
	m.SetHeader("From", em.SmtpAuth.From)

	// Set E-Mail receivers
	m.SetHeader("To", notify...)

	// Set E-Mail subject
	m.SetHeader("Subject", msg.Subject)

	// Set E-Mail body. You can set plain text or html with text/html
	m.SetBody("text/plain", msg.Msg)

	// Settings for SMTP server
	port, _ := strconv.Atoi(em.SmtpAuth.Port)
	d := gomail.NewDialer(em.SmtpAuth.Host, port, em.SmtpAuth.From, em.SmtpAuth.Password)
	//port, _ := strconv.Atoi(em.SmtpAuth.Port)
	//d := gomail.NewDialer(em.SmtpAuth.Host, port, em.SmtpAuth.From, em.SmtpAuth.Password)

	// This is only needed when SSL/TLS certificate is not valid on server.
	// In production this should be set to false.
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	// Now send E-Mail
	if err := d.DialAndSend(m); err != nil {
		fmt.Println(err)
		//panic(err)
	}

	return nil
	//return nil
}

func (em *EmailNotifyExecutor) SendOneMessage(msg Message) error {

	if !em.CanSendOn() {
		log.Println("call was success but we can't send yet")

		return nil

	}

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
