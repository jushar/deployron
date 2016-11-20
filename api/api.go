package main

import (
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"os"

	"github.com/Jusonex/docker-autodeploy/common"
	"github.com/jinzhu/configor"
)

var config common.Config
var serviceConn *net.UnixConn

func main() {
	// Read config
	configor.Load(&config, "config.yml")

	// Launch REST Api
	http.HandleFunc("/api/deploy", apiHandler)

	fmt.Printf("Listening on %s:%d\n", config.API.IP, config.API.Port)
	http.ListenAndServe(fmt.Sprintf("%s:%d", config.API.IP, config.API.Port), nil)
}

func apiHandler(res http.ResponseWriter, req *http.Request) {
	// Check API secret
	if req.URL.Query().Get("APISecret") != config.API.Secret {
		sendJSONError(res, "Wrong API secret")
		return
	}

	sendMessageToService("EXC_DEPLOY", "")
}

func sendJSONError(res http.ResponseWriter, message string) {
	type Error struct {
		Error string
	}

	json.NewEncoder(res).Encode(Error{Error: message})
	res.Header().Set("Content-Type", "application/json")
}

func sendMessageToService(identifier string, parameter string) {
	// Connect to unix socket
	serviceConn, err := net.DialUnix("unix", &net.UnixAddr{config.API.Unixsocket, "unix"}, &net.UnixAddr{config.Service.Unixsocket, "unix"})
	if err != nil {
		panic(err)
	}
	defer os.Remove(config.API.Unixsocket)

	serviceConn.Write(common.WriteMessage(&common.Message{Identifier: identifier, Parameter: parameter}))
}
