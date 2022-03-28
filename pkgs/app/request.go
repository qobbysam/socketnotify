package app

import "net/http"

type EngageMentNotificationPostRequest struct {
	Type string
	//	Address   string
	DateTime  string
	MailingId string
	MessageId string
	Address   string
	ServerId  string
	SecretKey string
	Response  string
	LocalIP   string
	RemoteMta string
}

func (e *EngageMentNotificationPostRequest) Bind(r *http.Request) error {

	return nil

}

type ValRequest struct {
	Type      string
	ServerId  string
	SecretKey string
}

func (e *ValRequest) Bind(r *http.Request) error {

	return nil

}
