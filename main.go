package main

import (
	"bytes"
	"fmt"
	log "github.com/Sirupsen/logrus"
	"github.com/gorilla/mux"
	"github.com/kardianos/osext"
	config "github.com/spf13/viper"
	"net/http"
	"net/url"
	"os/exec"
	"path/filepath"
	"strings"
    "github.com/kardianos/service"
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

//Functions
func ParseArgs(r *http.Request) string {
	var argsBuffer bytes.Buffer

	pUrl, err := url.Parse(r.RequestURI)
	if err != nil {
		log.Error("Error Parsing URL")
	}

	for key, value := range pUrl.Query() {
		argsBuffer.WriteString(" " + key)
		argsBuffer.WriteString(" " + value[0])
	}

	return argsBuffer.String()
}

func exec_script(sc string) string {
	sc = string(strings.TrimSpace(sc))

	cmd := exec.Command("powershell.exe", "-NoLogo", "-NonInteractive", "-Command", "&{", sc, "}")
	out, err := cmd.Output()

	if err != nil {
		log.WithFields(log.Fields{
			"command": string(sc),
		}).Error(err.Error())
	}

	return string(out)
}

//Handlers
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Service Running"))
}

func RunCommand(w http.ResponseWriter, r *http.Request) {
	var commbuffer bytes.Buffer

	commbuffer.WriteString(mux.Vars(r)["name"])
	commbuffer.WriteString(ParseArgs(r))
	commbuffer.WriteString(" | ConvertTo-Json")

	w.Write([]byte(fmt.Sprintf(exec_script(commbuffer.String()))))
}

func RunScript(w http.ResponseWriter, r *http.Request) {
	var commbuffer bytes.Buffer

	commbuffer.WriteString("&\"")
	commbuffer.WriteString(filepath.Join(config.GetString("ScriptFolder"), mux.Vars(r)["name"]))
	commbuffer.WriteString("\"")
	commbuffer.WriteString(ParseArgs(r))

	w.Write([]byte(fmt.Sprintf(exec_script(commbuffer.String()))))
}
