package locdb

import (
	"fmt"

	"github.com/qobbysam/socketnotify/pkgs/config"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type DBS struct {
	DB *gorm.DB
}

func NewDBS(cfg *config.BigConfig) (*DBS, error) {

	db, err := gorm.Open(sqlite.Open(cfg.DB.Name))

	if err != nil {
		fmt.Println("failed to connect to the server")
		return nil, err
	}

	return &DBS{DB: db}, nil
}

func (dbs *DBS) RunMigrations() {

	dbs.DB.AutoMigrate(&NotificationRequest{}, &ClientResource{}, &ResourceID{})
}
