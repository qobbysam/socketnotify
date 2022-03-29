package app

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/qobbysam/socketnotify/pkgs/config"
	"github.com/qobbysam/socketnotify/pkgs/cronn"
	"github.com/qobbysam/socketnotify/pkgs/emailnotify"
	"github.com/qobbysam/socketnotify/pkgs/locdb"
)

type EmailNotifyApp struct {
	//RestServer *server.RestServer
	NotifyEx *emailnotify.EmailNotifyExecutor
	Router   chi.Router
	Addr     string
	DB       *locdb.DBS
	State    cronn.State
}

func NewEmailNotifyApp(cfg *config.BigConfig, db *locdb.DBS) *EmailNotifyApp {

	//restserver := server.NewRestServer(cfg.Rest)

	notifyapp := emailnotify.NewEmailNotifyExecutor(&cfg.Email)

	state := cronn.NewState(cfg, db, notifyapp)
	out := EmailNotifyApp{
		//RestServer: restserver,
		NotifyEx: notifyapp,
		Addr:     cfg.Rest.Address,
		DB:       db,
		State:    *state,
	}
	return &out
}

func (ema *EmailNotifyApp) Init() error {

	ema.DB.RunMigrations()

	rou := chi.NewRouter()

	rou.Post("/receive", ema.ReceiveHandler)
	rou.Post("/val", ema.ValHandler)
	rou.Post("/sendtest", ema.SendTestHandler)

	rou.Post("/turnsend", ema.TurnSendHandler)

	ema.Router = rou

	return nil
}

func (ema *EmailNotifyApp) RequestToDB(data EngageMentNotificationPostRequest, origin string) locdb.NotificationRequest {

	dbs := locdb.NewRestNotificationRequest(origin)

	dbs.Address = data.Address
	dbs.ClientIp = data.ClientIp
	dbs.DateTime = data.DateTime
	dbs.MailingId = data.MailingId
	dbs.MessageId = data.MessageId
	//dbs.Origin = origin
	dbs.SecretKey = data.SecretKey
	dbs.ServerId = data.ServerId
	dbs.TrackingType = data.TrackingType
	dbs.Type = data.Type
	dbs.Url = data.Url
	dbs.UserAgent = data.UserAgent

	return *dbs
}

func (ema *EmailNotifyApp) HandleReceiveData(data EngageMentNotificationPostRequest) {

	data_to_save := ema.RequestToDB(data, "socketlab")

	if ema.State.CanSave() {
		err := ema.DB.SaveRequest(data_to_save)

		if err != nil {
			fmt.Println("failed to save data")
			fmt.Println(err)
		}

		log.Println("save success full")

		ema.State.UnlockNewMsg()

	} else {

		ema.State.AddToBuffer(data_to_save)
	}

	//return nil

	// switch data.TrackingType {
	// case "0":
	// 	//click event
	// 	err := ema.HandleClick(*data)

	// 	// if err != nil {
	// 	// 	render.Render(rw, r, ErrInvalidRequest(err))
	// 	// 	return
	// 	// }

	// case "1":
	// 	//open event

	// 	msg := ema.NotifyEx.BuildOpenMsg(data.Address)

	// 	err := ema.NotifyEx.SendOneMessage(*msg)

	// 	if err != nil {
	// 		return
	// 	}
	// 	return
	// default:
	// 	return

	// }
}

func (ema *EmailNotifyApp) TurnSendHandler(rw http.ResponseWriter, r *http.Request) {

	data := &TurnRequest{}

	if err := render.Bind(r, data); err != nil {

		ema.WriteError(rw, r)
		return

	}

	if data.Action == "1" {

		ema.NotifyEx.LockCanSend()

		ema.WriteSuccess(rw, r)
		return
	} else if data.Action == "0" {
		ema.NotifyEx.UnlockCanSend()
		ema.WriteSuccess(rw, r)
		return
	} else {
		ema.WriteError(rw, r)
		return
	}

}

func (ema *EmailNotifyApp) WriteSuccess(rw http.ResponseWriter, r *http.Request) {
	rw.WriteHeader(http.StatusOK)
	rw.Header().Set("Content-Type", "application/json")

	resp := make(map[string]string)

	resp["message"] = "Status OK"

	jsonResp, err := json.Marshal(resp)

	if err != nil {
		log.Println("Error happend in json marshal ", err)
	}

	rw.Write(jsonResp)
}

func (ema *EmailNotifyApp) WriteError(rw http.ResponseWriter, r *http.Request) {
	rw.WriteHeader(http.StatusUnauthorized)
	rw.Header().Set("Content-Type", "application/json")

	resp := make(map[string]string)

	resp["message"] = "Status Unauthorized"

	jsonResp, err := json.Marshal(resp)

	if err != nil {
		log.Println("Error happend in json marshal ", err)
	}

	rw.Write(jsonResp)
}

func (ema *EmailNotifyApp) ReceiveHandler(rw http.ResponseWriter, r *http.Request) {
	data := &EngageMentNotificationPostRequest{}

	if err := render.Bind(r, data); err != nil {
		//render.Render(rw, r, ErrInvalidRequest(err))
		rw.WriteHeader(http.StatusUnauthorized)
		rw.Header().Set("Content-Type", "application/json")

		resp := make(map[string]string)

		resp["message"] = "Status Unauthorized"

		jsonResp, err := json.Marshal(resp)

		if err != nil {
			log.Println("Error happend in json marshal ", err)
		}

		rw.Write(jsonResp)

		return
	}

	log.Println(data)

	ema.HandleReceiveData(*data)

	rw.WriteHeader(http.StatusOK)
	rw.Header().Set("Content-Type", "application/json")

	resp := make(map[string]string)

	resp["message"] = "Status OK"

	jsonResp, err := json.Marshal(resp)

	if err != nil {
		log.Println("Error happend in json marshal ", err)
	}

	rw.Write(jsonResp)

	return

	//render.Status(rw, http.StatusOK)

}
func (ema *EmailNotifyApp) ValHandler(rw http.ResponseWriter, r *http.Request) {
	data := &EngageMentNotificationPostRequest{}

	if err := render.Bind(r, data); err != nil {
		render.Render(rw, r, ErrInvalidRequest(err))
		return
	}

	log.Println(data)
	fmt.Println("succes colled on val")
	render.Status(r, http.StatusOK)
	return
}

func (ema *EmailNotifyApp) SendTestHandler(rw http.ResponseWriter, r *http.Request) {
	fmt.Println("building test msg")
	msg := ema.NotifyEx.BuildTestMsg()

	err := ema.NotifyEx.SendOneMessage(*msg)

	if err != nil {
		fmt.Println("failed to send error msg")

		fmt.Println(err)
		render.Status(r, http.StatusOK)

		return
	}

	fmt.Println("send success")
	render.Status(r, http.StatusOK)

}

func (ema *EmailNotifyApp) SendTest() error {

	return nil
}

func (ema *EmailNotifyApp) StartApplicationServer() error {

	err := ema.Init()

	if err != nil {

		fmt.Println("failed to init app server")
	}
	fmt.Println("starting http server on , ", ema.Addr)
	errchan := make(chan error, 1)
	donechan := make(chan struct{}, 1)
	var wg sync.WaitGroup
	wg.Add(1)
	go ema.State.StartWatching(wg, errchan)

	wg.Add(1)
	go ema.WatchChannel(errchan, donechan)

	if err != nil {

		fmt.Println("failed to start watching server")
	}

	fmt.Println("hitting http listen and serve")
	err = http.ListenAndServe(ema.Addr, ema.Router)
	wg.Wait()
	return err
}

func (ema *EmailNotifyApp) WatchChannel(errchan chan error, donechan chan struct{}) {
	fmt.Println("starting app watch chanel")
	for {
		select {
		case <-donechan:
			fmt.Println("shutting down now")
		case err := <-errchan:

			if err != nil {
				fmt.Println("error happend while starting")

				fmt.Println(err)

			}
		}
	}
}

func (ema *EmailNotifyApp) HandleClick(data EngageMentNotificationPostRequest) error {
	msg := ema.NotifyEx.BuildClickMsg(data.Address)

	err := ema.NotifyEx.SendOneMessage(*msg)

	return err

}
