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
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"
)

//Runtime
func init() {
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
	config.SetDefault("ExitEnabled", true)
}

func main() {
	log.Info("Listening at http://" + config.GetString("Binding"))

	mx := mux.NewRouter()
	mx.HandleFunc("/", IndexHandler)

	if config.GetBool("ExitEnabled") {
		mx.HandleFunc("/exit", ExitHandler)
	}

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

func ExitHandler(w http.ResponseWriter, r *http.Request) {
	defer os.Exit(0)

	log.Info("Shutting Down...")
	w.Write([]byte(fmt.Sprintf("Shutting Down...")))
	time.Sleep(3000 * time.Millisecond)
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
