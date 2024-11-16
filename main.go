package main

import (
	tdlib "github.com/c0re100/gotdlib/client"
)

func main() {
	tdlib.SetLogLevel(1)
	tdlib.SetFilePath("./errors.txt")

	helper := &Client{}
	helper.login()
	go helper.checkNewNessage()

	// Prevent process exit
	theBestPreventExitWay := make(chan struct{})
	<-theBestPreventExitWay
}
