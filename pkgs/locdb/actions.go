package locdb

import (
	"fmt"
	"time"
)

func (db *DBS) SaveRequest(data NotificationRequest) error {

	err := db.DB.Create(&data)

	if err.Error != nil {

		fmt.Println("failed to save msg")
		return err.Error

	}
	fmt.Println("saving in data")
	fmt.Println(data)
	return err.Error

}

func (db *DBS) GetAllResources() []ResourceID {

	out := make([]ResourceID, 0)

	db.DB.Where(&ResourceID{Alive: true}).Find(&out)

	return out
}

func (db *DBS) SaveReSourceID(name string) error {

	res := NewResourceId(name)

	err := db.DB.Create(&res)

	if err.Error != nil {
		fmt.Println("Failed to save resource id")
		return err.Error

	}
	fmt.Println("Save resource id sucess")

	return err.Error

}

func (db *DBS) GetAllOpen(startime, endtime time.Time, resourceid string) []NotificationRequest {

	out := make([]NotificationRequest, 0)

	db.DB.Where("created_time BETWEEN ? And ?", startime, endtime).Where(&NotificationRequest{MailingId: resourceid, TrackingType: "1"}).Find(&out)
	//db.DB.Where("created_time BETWEEN ? And ?", startime, endtime).Find(&out)

	//db.DB.Find(&out)
	return out

}

func (db *DBS) GetAllClick(startime, endtime time.Time, resourceid string) []NotificationRequest {

	out := make([]NotificationRequest, 0)

	db.DB.Where("created_time BETWEEN ? And ?", startime, endtime).Where(&NotificationRequest{MailingId: resourceid, TrackingType: "0"}).Find(&out)
	//db.DB.Where("created_time BETWEEN ? And ?", startime, endtime).Find(&out)

	//db.DB.Find(&out)
	return out
}

func (db *DBS) GetClientResource(emails []string) []ClientResource {

	out := make([]ClientResource, 0)

	return out
}

func (db *DBS) NotificationRequestsToClientResource(notifications []NotificationRequest) []ClientResource {

	emails := make([]string, 0)

	for _, v := range notifications {
		emails = append(emails, v.Address)
	}

	res := db.GetClientResource(emails)

	return res

}
