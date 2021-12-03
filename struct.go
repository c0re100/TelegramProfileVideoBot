package main

import "github.com/c0re100/go-tdlib"

type Client struct {
	client   *tdlib.Client
	clientId int64
}

func tdConfig() *tdlib.Client {
	return tdlib.NewClient(tdlib.Config{
		APIID:                  "132712",
		APIHash:                "e82c07ad653399a37baca8d1e498e472",
		SystemLanguageCode:     "en",
		DeviceModel:            "ProfileVideoHelper",
		SystemVersion:          "1.0",
		ApplicationVersion:     "1.0",
		UseMessageDatabase:     true,
		UseFileDatabase:        true,
		UseChatInfoDatabase:    true,
		UseTestDataCenter:      false,
		DatabaseDirectory:      "./tdlib-db",
		FileDirectory:          "./tdlib-files",
		IgnoreFileNames:        false,
		EnableStorageOptimizer: true,
	})
}
