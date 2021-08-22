package main

import (
	"bytes"
	"encoding/json"
	"log"

	"github.com/ReeceRose/home-network-proxy/internal/client"
	"github.com/ReeceRose/home-network-proxy/internal/consts"
	"github.com/ReeceRose/home-network-proxy/internal/types"
	"github.com/ReeceRose/home-network-proxy/internal/utils"
)

func main() {
	log.Default().Println("Reporting Tool v1 Started")

	client, err := client.NewClient()
	if err != nil {
		log.Fatalln("Failed to initiate a HTTP client")
	}
	data, _, err := client.Get("https://ifconfig.me", false)
	if err != nil {
		log.Fatalln("Failed to external IP from ifconfig.me. Error: " + err.Error())
	}

	ip := string(data)

	payload := new(bytes.Buffer)
	json.NewEncoder(payload).Encode(types.IP{
		ExternalIP: ip,
	})

	_, statusCode, err := client.Post(
		utils.GetVariable(consts.API_URL)+"ip/",
		payload,
		true,
	)
	if err != nil {
		log.Fatalln("Failed to send IP to API. Error: " + err.Error())
	}
	log.Default().Printf("Sent IP of %s to API and got status code of: %d", ip, statusCode)
}
