package main

import (
	"log"

	tdlib "github.com/c0re100/gotdlib/client"
)

func (helper *Client) login() {
	authorizer := tdlib.ClientAuthorizer()
	go tdlib.CliInteractor(authorizer)

	authorizer.TdlibParameters <- GetTdParameters()

	var err error
	client, err := tdlib.NewClient(authorizer)
	if err != nil {
		log.Fatalln(err.Error())
	}

	me, err := client.GetMe()
	if err != nil {
		log.Fatalln(err.Error())
	}

	log.Println(CheckUsernameEmpty(me.Usernames) + " connected")

	helper.Client = client
	helper.ClientId = me.Id
}
