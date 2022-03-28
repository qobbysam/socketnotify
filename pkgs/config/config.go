package config

//config for the rest server and modes to notify

type BigConfig struct {
	Rest  RestConfig  `json:"rest"`
	Email EmailConfig `json:"email"`
	Phone PhoneConfig `json:"phone"`
}

type RestConfig struct {
	Address   string `json:"address"`
	SendOpen  bool   `json:"sendopen"`
	SendClick bool   `json:"sendclick"`
}

type EmailConfig struct {
	Smtpfrom     string   `json:"smtpfrom"`
	SmtpPassword string   `json:"smtpPassword"`
	SmtpHost     string   `json:"smtpHost"`
	SmtpPort     string   `json:"smtpPort"`
	Notify       []string `json:"notify"`
}

type PhoneConfig struct{}

type NotificationsConfig struct{}
