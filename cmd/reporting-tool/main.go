package main

import (
	"bytes"
	"encoding/json"
	"log"

	"github.com/ReeceRose/home-network-proxy/internal/client"
	"github.com/ReeceRose/home-network-proxy/internal/consts"
	"github.com/ReeceRose/home-network-proxy/internal/utils"
)

type apiRequest struct {
	ExternalIP string
}

func main() {
	log.Default().Println("Reporting Tool v1 Started")

	client, err := client.NewClient()
	if err != nil {
		log.Fatalln("Failed to initiate a HTTP client")
	}
	data, _, err := client.Get("https://ifconfig.me")
	if err != nil {
		log.Fatalln("Failed to external IP from ifconfig.me. Error: " + err.Error())
	}

	payload := new(bytes.Buffer)
	json.NewEncoder(payload).Encode(apiRequest{ExternalIP: string(data)})

	data, statusCode, err := client.Post(utils.GetVariable(consts.API_URL)+"ip/", payload)
	if err != nil {
		log.Fatalln("Failed to send IP to API. Error: " + err.Error())
	}
	log.Default().Printf("Sent %s to API and got status code of: %d \n", string(data), statusCode)
}
