package locdb

import (
	"fmt"
	"log"
	"strings"
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

func (db *DBS) GetEmails(nots []NotificationRequest) []string {

	out := make([]string, 0)

	for _, v := range nots {
		upper := strings.ToUpper(v.Address)
		out = append(out, upper)
	}

	return out

}

func (db *DBS) GetOneResource(name string) *ResourceID {
	res := ResourceID{}
	err := db.DB.Where(&ResourceID{Name: name}).Find(&res)

	if err.Error != nil {
		return nil
	} else if res.ID == "" {

		return nil

	}

	return &res

}

func (db *DBS) GetAllResources() []ResourceID {

	out := make([]ResourceID, 0)

	db.DB.Where(&ResourceID{Alive: true}).Find(&out)

	return out
}

func (db *DBS) SaveReSourceID(name string, alive bool) error {

	knownname := ResourceID{}

	get := db.DB.Where(&ResourceID{Name: name}).Find(&knownname)

	if get.Error != nil {
		log.Println("failed to find resource name")
	}

	if knownname.ID == "" {
		res := NewResourceId(name)
		res.Alive = alive
		err := db.DB.Create(&res)

		if err.Error != nil {
			fmt.Println("Failed to save resource id")
			return err.Error

		}
		fmt.Println("Save resource id sucess")
	} else {

		knownname.Alive = alive

		up := db.DB.Save(&knownname)

		if up.Error != nil {
			log.Println("failed to update name")
			return up.Error
		}
		log.Println("updated status of ,  ", name)
	}

	return nil

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
