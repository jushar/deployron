package main

import (
	"bytes"
	"fmt"
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

	// Remove old sockets (in case the service crashed)
	os.Remove(config.Service.Unixsocket)

	// Start unix socket
	l, err := net.ListenUnix("unix", &net.UnixAddr{config.Service.Unixsocket, "unix"})
	if err != nil {
		panic(err)
	}

	// Update permissions
	// TODO: Make this more fine-grained
	os.Chmod(config.Service.Unixsocket, 0777)

	fmt.Println("Waiting for commands")

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
		var commandBuffer bytes.Buffer

		for _, line := range config.Service.Script {
			commandBuffer.WriteString(line)
			commandBuffer.WriteString("; ")
		}

		// Execute deploy script
		cmd := exec.Command("/bin/sh", "-c", commandBuffer.String())
		err := cmd.Run()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Executing the script failed: %s\n", err.Error())
		}
	}
}
