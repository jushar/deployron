package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os/exec"

	"github.com/jinzhu/configor"
)

type Config struct {
	Service struct {
		IP     string `default:""`
		Port   uint   `required:"true" default:"1337"`
		Script string `default:"./deploy.sh"`
		Secret string `required:"true"`
	}
}

var config Config

func main() {
	// Read config
	configor.Load(&config, "config.yml")

	// Do some integrity checks
	// TODO

	// Launch REST Api
	http.HandleFunc("/api/deploy", apiHandler)
	http.ListenAndServe(fmt.Sprintf("%s:%d", config.Service.IP, config.Service.Port), nil)
}

func apiHandler(res http.ResponseWriter, req *http.Request) {
	// Check API secret
	if req.Header.Get("APISecret") != config.Service.Secret && false {
		sendJSONError(res, "Wrong API secret")
		return
	}

	// Execute deploy script
	cmd := exec.Command(config.Service.Script)
	err := cmd.Run()
	if err != nil {
		log.Panic(err)
	}
}

func sendJSONError(res http.ResponseWriter, message string) {
	type Error struct {
		Error string
	}

	json.NewEncoder(res).Encode(Error{Error: message})
	res.Header().Set("Content-Type", "application/json")
}
