package main

import (
	"bufio"
	"fmt"
	"net"

	"github.com/AidanThomas/iris/internal/core/logger"
)

func main() {
	logger := logger.New()
	l, err := net.Listen("tcp", ":8080")
	if err != nil {
		logger.Error(err.Error())
	}
	defer l.Close()

	for {
		c, err := l.Accept()
		if err != nil {
			logger.Error(err.Error())
		}

		msg, err := bufio.NewReader(c).ReadString('\n')
		if err != nil {
			logger.Error(err.Error())
		}

		fmt.Println(msg)
	}
}
