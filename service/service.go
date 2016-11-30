package main

import (
	"bytes"
	"fmt"
	"net"
	"os"
	"os/exec"
	"syscall"

	"github.com/Jusonex/deployron/common"
)

var config *common.Config

func main() {
	// Read config
	config = common.MakeConfig("config.yml")

	// Do some integrity checks
	// Make sure config.yml is owned by root (as editing it would result in root privileges)
	statInfo, err := os.Stat("config.yml")
	if err != nil {
		panic(err)
	}
	if (statInfo.Mode()&0x1B) != 0 || statInfo.Sys().(*syscall.Stat_t).Uid != 0 { // 0x1B = ~0x1E4 = 774 base 8
		panic("Make sure config.yml is owned by root and group and other don't have write permissions")
	}

	// Remove old sockets (in case the service crashed)
	os.Remove(config.Service.Unixsocket)

	// Start unix socket
	l, err := net.ListenUnix("unix", &net.UnixAddr{config.Service.Unixsocket, "unix"})
	if err != nil {
		panic(err)
	}
	defer os.Remove(config.Service.Unixsocket)

	// Update permissions
	// TODO: Make this more fine-grained
	os.Chmod(config.Service.Unixsocket, 0777)

	fmt.Println("Waiting for commands")

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
		deployment := config.FindDeploymentByName(message.Parameter)

		if deployment == nil {
			fmt.Fprintf(os.Stderr, "Invalid deployment service name passed")
			return
		}

		var commandBuffer bytes.Buffer
		for _, line := range deployment.Script {
			commandBuffer.WriteString(line)
			commandBuffer.WriteString("; ")
		}

		// Prepare deploy script for execution
		cmd := exec.Command("su", "-s", "/bin/sh", "-c", commandBuffer.String(), deployment.User)

		// Redirect stdout, stderr
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr

		// Run deploy script
		err := cmd.Run()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Executing the script failed: %s\n", err.Error())
		}
	}
}
