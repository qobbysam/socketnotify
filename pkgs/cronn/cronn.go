package cronn

import (
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/qobbysam/socketnotify/pkgs/config"
	"github.com/qobbysam/socketnotify/pkgs/emailnotify"

	//"github.com/qobbysam/socketnotify/pkgs/emailnotify"

	//"github.com/qobbysam/socketnotify/pkgs/emailnotify"

	//	"github.com/qobbysam/socketnotify/pkgs/emailnotify"

	//"github.com/qobbysam/socketnotify/pkgs/emailnotify"
	"github.com/qobbysam/socketnotify/pkgs/locdb"
)

type State struct {
	CanTick           bool
	CanUpdate         bool
	NewMsgReceived    bool
	LastMsgReportTime time.Time
	SecondInterval    int
	DB                *locdb.DBS
	//	EmailNotify       *emailnotify.EmailNotifyExecutor
	Superwaitgroup sync.WaitGroup
	Bufferwait     []locdb.NotificationRequest
}

func NewState(cfg *config.BigConfig, db *locdb.DBS, not *emailnotify.EmailNotifyExecutor) *State {
	return &State{
		CanUpdate:         true,
		NewMsgReceived:    false,
		LastMsgReportTime: time.Now(),
		SecondInterval:    cfg.Cron.SecondInterval,
		DB:                db,
		Bufferwait:        make([]locdb.NotificationRequest, 0),
		//	EmailNotify:       not,
		CanTick: cfg.Cron.CanTick,
	}
}

func (st *State) StartWatching(wg sync.WaitGroup, errchan chan error) {

	fmt.Println("starting watcher process")
	st.Superwaitgroup = wg
	ticker := time.NewTicker(time.Duration(st.SecondInterval) * time.Second)
	fmt.Println("ticker started")
	for t := range ticker.C {
		fmt.Println("ticker run")

		// if !st.CanTickOn() {

		// 	return
		// }

		if st.CheckToReport() && st.CanTickOn() {

			st.LockUpdate()

			errchan := make(chan error, 1)
			st.Superwaitgroup.Add(1)
			go st.HandleReport(errchan, t)
			st.Superwaitgroup.Add(1)
			go st.WatchChannel(errchan, t)
			// if err != nil {
			// 	fmt.Println("faile to build report,  ", t)
			// }

			//st.UnlockUpdate()

		} else {
			fmt.Println("Nothing to report")
		}

	}

}

func (st *State) CheckToReport() bool {
	if st.NewMsgReceived {
		return st.NewMsgReceived
	}

	return false
}

func (st *State) HandleReport(errchan chan error, t time.Time) {
	fmt.Println("starting handle report")
	all_resource_id := st.DB.GetAllResources()

	fmt.Println(all_resource_id)

	open_reports, _ := st.BuildOpenReport(all_resource_id)

	click_reports, _ := st.BuildClickReport(all_resource_id)

	//one_report := st.BuildOneReport(open_reports, click_reports)
	sub := "Open Reports At time " + time.Now().Format("2006-01-02 15:04:05")
	err := st.SendReports(sub, open_reports)

	// err := st.NotifyInterests(one_report)

	if err != nil {

		fmt.Println("failed to send open reports")
		errchan <- err
	}

	// err = st.SendReports(sub, open_resources)

	// if err != nil {

	// 	fmt.Println("failed to send open  resource reports")
	// 	errchan <- err
	// }

	sub2 := "Click Reports At time " + time.Now().Format("2006-01-02 15:04:05")

	err = st.SendReports(sub2, click_reports)

	if err != nil {

		fmt.Println("failed to send click reports")
		errchan <- err
	}

	// err = st.SendReports(sub2, click_resources)

	// if err != nil {
	// 	fmt.Println("Failed to send click report")
	// 	errchan <- err
	// }

	//time.Sleep(20 * time.Second)
	errchan <- nil

}

func (st *State) SendReports(sub string, reports []Report) error {
	fmt.Println("sending report , ", sub)

	for _, v := range reports {

		err := st.NotifyInterests(v)

		if err != nil {
			return err
		}

	}
	return nil
}

func (st *State) WatchChannel(errchan chan error, t time.Time) {

	for err := range errchan {

		if err != nil {
			fmt.Println("err received on this channel ")
			fmt.Println(err)
		} else {
			st.CleanBuffer()
			st.LastMsgReportTime = time.Now()
			st.UnlockUpdate()
			st.LockNewMsg()
			st.Superwaitgroup.Done()
			st.Superwaitgroup.Done()
			fmt.Println("sucess handle watch at : ", t)

		}

	}
}

func (st *State) NotifyInterests(msg Report) error {

	msgg := emailnotify.Message{}

	msgg.Subject = msg.Subject
	msgg.Msg = msg.Txt

	//	err := st.EmailNotify.SendOneMessage(msgg)

	// if err != nil {
	// 	fmt.Println("failed to send msg")
	// 	return err
	// }

	fmt.Println("msg sending success")

	return nil
}

func (st *State) BuildSpecialReport(in []locdb.NotificationRequest, action string) Report {

	outreport := Report{}

	for k, v := range in {

		top := fmt.Sprint("number :  ", k, "mailingid ", v.MailingId, "\n")
		middle := fmt.Sprint("email : ", action, "  on "+"\n")

		outreport.Txt = outreport.Txt + top + middle
	}

	outreport.Subject = fmt.Sprint(action, " report ", len(in))

	return outreport

}

func (st *State) BuildOpenReport(allresourceid []locdb.ResourceID) ([]Report, []Report) {

	//all_open := st.DB.GetAllOpen()

	outNotification := make([]Report, 0)

	outResource := make([]Report, 0)

	for _, v := range allresourceid {

		allopens := st.DB.GetAllOpen(st.LastMsgReportTime, time.Now(), v.Name)
		fmt.Println(allopens)
		sub := "msgid: " + v.Name
		report := st.NotificationRequestsToReport(sub, "open", allopens)

		outNotification = append(outNotification, report)

		clientResource := st.DB.NotificationRequestsToClientResource(allopens)

		report2 := st.ClientResourceToReport(sub, clientResource)

		outResource = append(outResource, report2)
	}

	return outNotification, outResource

}

func (st *State) BuildClickReport(allresourceid []locdb.ResourceID) ([]Report, []Report) {
	outNotification := make([]Report, 0)
	outResource := make([]Report, 0)

	for _, v := range allresourceid {
		sub := "msgid: " + v.Name

		allopens := st.DB.GetAllClick(st.LastMsgReportTime, time.Now(), v.Name)

		report := st.NotificationRequestsToReport(sub, "click", allopens)

		outNotification = append(outNotification, report)

		cresources := st.DB.NotificationRequestsToClientResource(allopens)

		report2 := st.ClientResourceToReport(sub, cresources)

		outResource = append(outResource, report2)

	}

	return outNotification, outResource
}

func (st *State) ClientResourceToReport(sub string, reslist []locdb.ClientResource) Report {
	report := Report{
		Subject: sub,
		Txt:     "",
	}

	for k, v := range reslist {

		middle := fmt.Sprint(v, " ", v)
		top := fmt.Sprint("Number: ", k)
		report.Txt = report.Txt + "\n" + top + "\n" + middle + "\n"

	}

	return report
}

func (st *State) NotificationRequestsToReport(sub, action string, reqlist []locdb.NotificationRequest) Report {

	report := Report{
		Subject: sub,
		Txt:     "",
	}

	for k, v := range reqlist {

		middle := fmt.Sprint(v.Address, " ", action, " ", v.MailingId)
		top := fmt.Sprint("Number: ", k)
		report.Txt = report.Txt + "\n" + top + "\n" + middle + "\n"

	}
	fmt.Println("printing report,  ", action)
	fmt.Println(report)

	return report
}

func (st *State) BuildResourceReport(emails []string, action string) Report {

	known, err := st.DB.GetClientResourceList(emails)

	if err != nil {
		log.Println("failed to get from db ,  ", err)
	}
	action_ := ""

	if action == "1" {
		action_ = "open"
	} else if action == "0" {
		action_ = "click"
	} else {
		action_ = "unkown"
	}

	now := time.Now()

	sub := fmt.Sprint("Resource Interest  ", action_, " ", now.Month(), " : ", now.Hour(), " ", now.Minute(), "\n")

	report := st.ClientResourceToReport(sub, known)

	return report
}

// func (st *State) BuildOneReport(sub string) Report {

// }

type Report struct {
	Subject string
	Txt     string
}
