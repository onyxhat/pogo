package main

import (
	log "github.com/Sirupsen/logrus"
	"github.com/gorilla/mux"
	"github.com/kardianos/osext"
	"github.com/kardianos/service"
	config "github.com/spf13/viper"
	"net/http"
)

type program struct{}

//Runtime
func (p *program) Start(s service.Service) error {
	exePath, err := osext.ExecutableFolder()
	if err != nil {
		exePath = ".\\"
	}

	config.AddConfigPath(exePath)
	config.SetConfigName("config")
	config.ReadInConfig()

	config.SetDefault("Binding", "0.0.0.0:8080")
	config.SetDefault("ScriptFolder", ".\\scripts\\")
	config.SetDefault("CommandsEnabled", true)

	go p.run()
	return nil
}

func (p *program) run() {
	log.Info("Listening at http://" + config.GetString("Binding"))

	mx := mux.NewRouter()
	mx.HandleFunc("/", IndexHandler)

	if config.GetBool("ScriptsEnabled") {
		mx.HandleFunc("/scripts/{name:\\S+}", RunScript)
		log.Info("Script router listening at http://" + config.GetString("Binding") + "/scripts/")
	}

	if config.GetBool("CommandsEnabled") {
		mx.HandleFunc("/command/{name:\\S+}", RunCommand)
		log.Info("Command router listening at http://" + config.GetString("Binding") + "/command/")
	}

	http.ListenAndServe(config.GetString("Binding"), mx)
}

func (p *program) Stop(s service.Service) error {
	return nil
}

func main() {
	svcConfig := &service.Config{
		Name:        "pogo",
		DisplayName: "PoGo Service",
		Description: "PoGo API Service.",
	}

	prg := &program{}
	s, err := service.New(prg, svcConfig)
	if err != nil {
		log.Fatal(err)
	}

	logger, err := s.Logger(nil)
	if err != nil {
		log.Fatal(err)
	}

	err = s.Run()
	if err != nil {
		logger.Error(err)
	}
}
