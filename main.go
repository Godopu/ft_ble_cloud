package main

import (
	"errors"
	"fmt"
	"ftblecloud/api"
	"ftblecloud/config"
	"os"
)

func main() {
	if _, err := os.Stat("./config.properties"); errors.Is(err, os.ErrNotExist) {
		config.CreateInitFile()
	}

	config.LoadConfig()

	fmt.Println(config.Params["computing_delay"])

	srv := api.NewWebServer(":9910")
	if err := srv.ListenAndServe(); err != nil {
		panic(err)
	}
}
