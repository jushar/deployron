package main

import (
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"os"

	"github.com/Jusonex/deployron/common"
	"github.com/gorilla/mux"
)

var config *common.Config
var serviceConn *net.UnixConn

func main() {
	// Create config
	config = common.MakeConfig("config.yml")

	// Launch REST API
	router := mux.NewRouter()
	router.HandleFunc("/deploy/{name}", apiHandler)
	http.Handle("/", router)

	fmt.Printf("Listening on %s:%d\n", config.API.IP, config.API.Port)
	http.ListenAndServe(fmt.Sprintf("%s:%d", config.API.IP, config.API.Port), nil)
}

func apiHandler(res http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	deployName := params["name"]
	deployment := config.FindDeploymentByName(deployName)

	// Check if there is a deployment with this name
	if deployment == nil {
		sendJSONError(res, "Unknown deployment")
		return
	}

	// Check API secrets
	if req.URL.Query().Get("APISecret") != deployment.Secret {
		sendJSONError(res, "Wrong deploy secret")
		return
	}

	sendMessageToService("EXC_DEPLOY", deployName)
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
