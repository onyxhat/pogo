package main

import (
	"bytes"
	"fmt"
	log "github.com/Sirupsen/logrus"
	"github.com/gorilla/mux"
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
	config.SetConfigName("config")
	config.AddConfigPath(".\\")
	config.ReadInConfig()

	config.SetDefault("Binding", ":8080")
	config.SetDefault("ScriptFolder", ".\\scripts\\")
}

func main() {
	mx := mux.NewRouter()

	mx.HandleFunc("/", IndexHandler)
	mx.HandleFunc("/exit", ExitHandler)
	mx.HandleFunc("/command/{name:\\S+}", RunShell)
	mx.HandleFunc("/script/{name:\\S+}", RunScript)

	log.Info("Listening at " + config.GetString("Binding"))
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

	log.Info("Shutting Down")
	w.Write([]byte(fmt.Sprintf("Shutting Down")))
	time.Sleep(3000 * time.Millisecond)
}

func RunShell(w http.ResponseWriter, r *http.Request) {
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
