package internal

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/qobbysam/socketnotify/pkgs/app"
	"github.com/qobbysam/socketnotify/pkgs/config"
)

type InternalStruct struct {
	App *app.EmailNotifyApp
}

func (in *InternalStruct) GetConfig(input string) string {

	if input != "" {
		return input
	} else {

		dirpath, _ := os.Executable()

		dir := filepath.Dir(dirpath)

		cfg_path := filepath.Join(dir, "config", "conf.json")

		return cfg_path
	}

}

func (in *InternalStruct) Init(input string) error {
	path := in.GetConfig(input)

	cfg, err := in.BuildConfig(path)

	if err != nil {
		return err
	}

	app := app.NewEmailNotifyApp(cfg)

	in.App = app

	return nil
}

func (in *InternalStruct) StartServer() {

	err := in.App.StartApplicationServer()

	if err != nil {
		fmt.Println("failed to start server, ", err)
	}
	fmt.Println("started rest server successfully")
}

func (in *InternalStruct) EmailTest() {

	err := in.App.SendTest()

	if err != nil {
		fmt.Println("failed to send test msg,  ", err)
	}

	fmt.Println("sent test message sucessfully")
}

func (in *InternalStruct) BuildConfig(path string) (*config.BigConfig, error) {

	file, err := os.Open(path)

	if err != nil {
		fmt.Println("failed to open file  ,  ", err)
		return nil, err
	}

	by, err := ioutil.ReadAll(file)

	if err != nil {
		fmt.Println("failed read file ,  ", err)
		return nil, err
	}

	var bigConfig config.BigConfig

	err = json.Unmarshal(by, &bigConfig)

	if err != nil {
		fmt.Println("failed to marshal json check order,  ", err)
		return nil, err
	}

	return &bigConfig, nil
}

func (in *InternalStruct) StartApplication(action, input string) {

	err := in.Init(input)

	if err != nil {

		fmt.Println("failed to init,  ", err)

		panic("failed to init")
	}

	switch action {
	case "server":
		fmt.Println("starting server")
		in.StartServer()
	case "emailtest":
		in.EmailTest()
	default:
		fmt.Println("not a valid action received")
	}

}
