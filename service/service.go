package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/exec"

	"github.com/Jusonex/docker-autodeploy/common"
	"github.com/jinzhu/configor"
)

var config common.Config

func main() {
	// Read config
	configor.Load(&config, "config.yml")

	// Start unix socket
	l, err := net.ListenUnix("unix", &net.UnixAddr{config.Service.Unixsocket, "unix"})
	if err != nil {
		panic(err)
	}

	// Update permissions
	// TODO: Make this more fine-grained
	os.Chmod(config.Service.Unixsocket, 0777)

	defer os.Remove(config.Service.Unixsocket)

	// Wait for commands
	for {
		// Accept incoming connection
		conn, err := l.AcceptUnix()
		if err != nil {
			panic(err)
		}

		// Read from stream
		var buf [256]byte
		_, err = conn.Read(buf[:])
		if err != nil {
			panic(err)
		}

		// Parse message
		message := common.ReadMessage(buf)
		processMessage(message)

		// Close connection
		fmt.Printf("Received command: %s\n", message.Identifier)
		conn.Close()
	}
}

func processMessage(message *common.Message) {
	switch message.Identifier {
	case "EXC_DEPLOY":
		// Execute deploy script
		cmd := exec.Command(config.Service.Script)
		err := cmd.Run()
		if err != nil {
			log.Panic(err)
		}
	}
}
