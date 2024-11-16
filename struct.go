package main

import (
	tdlib "github.com/c0re100/gotdlib/client"
)

type Client struct {
	Client   *tdlib.Client
	ClientId int64
}

func GetTdParameters() *tdlib.SetTdlibParametersRequest {
	return &tdlib.SetTdlibParametersRequest{
		UseTestDc:           false,
		DatabaseDirectory:   "./tdlib-db",
		FilesDirectory:      "./tdlib-files",
		UseFileDatabase:     true,
		UseChatInfoDatabase: true,
		UseMessageDatabase:  true,
		UseSecretChats:      false,
		ApiId:               132712,
		ApiHash:             "e82c07ad653399a37baca8d1e498e472",
		SystemLanguageCode:  "en",
		DeviceModel:         "ProfileVideoHelper",
		SystemVersion:       "1.1",
		ApplicationVersion:  "1.1",
	}
}
