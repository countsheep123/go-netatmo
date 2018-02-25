package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/countsheep123/go-netatmo"
)

func main() {
	username := os.Getenv("NETATMO_USERNAME")
	password := os.Getenv("NETATMO_PASSWORD")
	clientID := os.Getenv("NETATMO_CLIENT_ID")
	clientSecret := os.Getenv("NETATMO_CLIENT_SECRET")

	cli, err := netatmo.NewClient(&netatmo.Config{
		Username:     username,
		Password:     password,
		ClientID:     clientID,
		ClientSecret: clientSecret,
		Scopes:       []string{netatmo.ReadStation},
	})
	if err != nil {
		log.Fatal(err)
	}

	data, err := cli.Getstationsdata(&netatmo.StationDataRequest{
		DeviceID:     nil,
		GetFavorites: true,
	})
	if err != nil {
		log.Fatal(err)
	}

	b, err := json.Marshal(data)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(b))
}
