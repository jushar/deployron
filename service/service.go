package main

import (
	"net"
	"os"

	"github.com/jinzhu/configor"
)

var config Config

func main() {
	// Read config
	configor.Load(&config, "config.yml")

	// Start unix socket
	l, err := net.ListenUnix("unix", &net.UnixAddr{config.Service.Unixsocket, "unix"})
	if err != nil {
		panic(err)
	}
	defer os.Remove()

}
