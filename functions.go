package main

import (
	"bytes"
	log "github.com/Sirupsen/logrus"
	"net/http"
	"net/url"
	"os/exec"
	"strings"
)

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
