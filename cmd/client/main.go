package main

import (
	"net"

	"github.com/AidanThomas/iris/internal/core/logger"
)

func main() {
	logger := logger.New()
	c, err := net.Dial("tcp", "127.0.0.1:8080")
	if err != nil {
		logger.Error(err.Error())
	}

	_, err = c.Write([]byte("HERE IS A MESSAGE\n"))
	if err != nil {
		logger.Error(err.Error())
	}
}
