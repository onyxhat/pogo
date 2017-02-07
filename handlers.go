package main

import (
	"bytes"
	"fmt"
	"github.com/gorilla/mux"
	config "github.com/spf13/viper"
	"net/http"
	"path/filepath"
)

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
