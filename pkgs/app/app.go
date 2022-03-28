package app

import (
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/qobbysam/socketnotify/pkgs/config"
	"github.com/qobbysam/socketnotify/pkgs/emailnotify"
)

type EmailNotifyApp struct {
	//RestServer *server.RestServer
	NotifyEx *emailnotify.EmailNotifyExecutor
	Router   chi.Router
	Addr     string
}

func NewEmailNotifyApp(cfg *config.BigConfig) *EmailNotifyApp {

	//restserver := server.NewRestServer(cfg.Rest)

	notifyapp := emailnotify.NewEmailNotifyExecutor(&cfg.Email)

	out := EmailNotifyApp{
		//RestServer: restserver,
		NotifyEx: notifyapp,
		Addr:     cfg.Rest.Address,
	}
	return &out
}

func (ema *EmailNotifyApp) Init() error {

	rou := chi.NewRouter()

	rou.Post("/receive", ema.ReceiveHandler)
	rou.Post("/val", ema.ValHandler)
	rou.Post("/sendtest", ema.SendTestHandler)

	ema.Router = rou

	return nil
}

func (ema *EmailNotifyApp) ReceiveHandler(rw http.ResponseWriter, r *http.Request) {
	data := &EngageMentNotificationPostRequest{}

	if err := render.Bind(r, data); err != nil {
		render.Render(rw, r, ErrInvalidRequest(err))
		return
	}

	log.Println(data)

	switch data.Type {
	case "Click":
		err := ema.HandleClick(*data)

		if err != nil {
			render.Render(rw, r, ErrInvalidRequest(err))
			return
		}

		render.Status(r, http.StatusOK)
		return
	case "Open":
		msg := ema.NotifyEx.BuildOpenMsg(data.Address)

		err := ema.NotifyEx.SendOneMessage(*msg)

		if err != nil {
			render.Render(rw, r, ErrInvalidRequest(err))
			return
		}
		render.Status(r, http.StatusOK)
		return
	default:
		render.Render(rw, r, ErrInvalidRequest(errors.New("Not open or click")))
		return

	}

}
func (ema *EmailNotifyApp) ValHandler(rw http.ResponseWriter, r *http.Request) {
	data := &ValRequest{}

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
	err = http.ListenAndServe(ema.Addr, ema.Router)
	return err
}

func (ema *EmailNotifyApp) HandleClick(data EngageMentNotificationPostRequest) error {
	msg := ema.NotifyEx.BuildClickMsg(data.Address)

	err := ema.NotifyEx.SendOneMessage(*msg)

	return err

}
