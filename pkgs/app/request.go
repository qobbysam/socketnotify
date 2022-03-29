package app

import "net/http"

type EngageMentNotificationPostRequest struct {
	Type         string `json:"Type"`
	DateTime     string `json:"DateTime"`
	MailingId    string `json:"MailingId"`
	MessageId    string `json:"MessageId"`
	Address      string `json:"Address"`
	ServerId     int    `json:"ServerId"`
	SecretKey    string `json:"SecretKey"`
	ClientIp     string `json:"ClientIp"`
	TrackingType string `json:"TrackingType"`
	Url          string `json:"Url"`
	UserAgent    string `json:"UserAgent"`
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
