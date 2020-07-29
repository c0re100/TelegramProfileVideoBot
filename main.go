package main

import (
	"github.com/c0re100/go-tdlib"
)

func main() {
	tdlib.SetLogVerbosityLevel(0)
	tdlib.SetFilePath("./errors.txt")

	tg := &Client{
		client: tdConfig(),
	}
	tg.login()
	go tg.checkNewNessage()

	// Prevent process exit
	theBestPreventExitWay := make(chan struct{})
	<-theBestPreventExitWay
}
