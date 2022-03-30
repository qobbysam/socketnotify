package locdb

import (
	"time"

	"github.com/lithammer/shortuuid"
)

type NotificationRequest struct {
	ID           string `gorm:"primaryKey"`
	CreatedTime  time.Time
	UpdatedAt    time.Time
	New          bool
	Origin       string
	Type         string
	DateTime     string
	MailingId    string
	MessageId    string
	Address      string
	ServerId     int
	SecretKey    string
	ClientIp     string
	TrackingType string
	Url          string
	UserAgent    string
}

func NewRestNotificationRequest(origin string) *NotificationRequest {

	return &NotificationRequest{
		Origin:      origin,
		ID:          shortuuid.New(),
		New:         true,
		CreatedTime: time.Now(),
	}
}

type ResourceID struct {
	ID    string `gorm:"primaryKey"`
	Name  string
	Alive bool
}

func NewResourceId(name string) *ResourceID {

	return &ResourceID{
		ID:    shortuuid.New(),
		Alive: true,
		Name:  name,
	}
}

type ClientResource struct {
	ID string `gorm:"primaryKey"`
	//ResourceID  string
	Name   string
	Phone  string
	Email  string
	Mc     string
	Dot    string
	Street string
	City   string
	State  string
	//Address      string
	DotApplyDate string
	McAppyDate   string
	McGrantDate  string
	PowerUnit    string
	DriverTotal  string
	McsFileDate  string

	FirstName   string
	LastName    string
	CreatedTime time.Time
	UpdatedTime time.Time
}

func NewClientResource() *ClientResource {
	return &ClientResource{
		ID: shortuuid.New(),
	}
}
