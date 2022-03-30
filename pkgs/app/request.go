package app

import (
	"errors"
	"net/http"
)

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

type GenSpecRequest struct {
	MailingId      string   `json:"mailingid"`
	Action         string   `json:"action"`
	Last           int      `json:"last"`
	Auth           string   `json:"auth"`
	NotifyInterest []string `json:"notify"`
	GenResource    bool     `json:"genresource"`
}

func (e *GenSpecRequest) Bind(r *http.Request) error {

	if e.Auth == "" {

		return errors.New("auth cannot be empty")
	}

	if e.Action == "" {
		return errors.New("action cannot be empty")
	}
	return nil
}

type TurnRequest struct {
	Action string `json:"action"`
	Auth   string `json:"auth"`
}

func (e *TurnRequest) Bind(r *http.Request) error {

	if e.Auth == "" {

		return errors.New("auth cannot be empty")
	}

	if e.Action == "" {
		return errors.New("action cannot be empty")
	}
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
