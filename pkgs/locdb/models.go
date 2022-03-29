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
	ID          string `gorm:"primaryKey"`
	ResourceID  string
	Email       string
	FirstName   string
	LastName    string
	PhoneNumber string
	DotNumber   string
	McNumber    string
	City        string
	State       string
	Address     string
}

func NewClientResource() *ClientResource {
	return &ClientResource{
		ID: shortuuid.New(),
	}
}
