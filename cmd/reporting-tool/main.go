package main

import (
	"bytes"
	"encoding/json"
	"log"

	"github.com/ReeceRose/home-network-proxy/internal/client"
	"github.com/ReeceRose/home-network-proxy/internal/consts"
	"github.com/ReeceRose/home-network-proxy/internal/store"
	"github.com/ReeceRose/home-network-proxy/internal/types"
	"github.com/ReeceRose/home-network-proxy/internal/utils"
)

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

	ip := string(data)

	info := store.Instance().GetReportingToolAgentInformation()

	payload := new(bytes.Buffer)
	json.NewEncoder(payload).Encode(types.IP{
		ID:         info.ID.String(),
		ExternalIP: ip,
		UserId:     info.UserID,
	})
	headers := make(map[string]string)
	headers["Authorization"] = info.APIKey
	headers["UserID"] = info.UserID
	_, statusCode, err := client.Post(
		utils.GetVariable(consts.API_URL)+"/ip",
		payload,
		headers,
	)
	if err != nil {
		log.Fatalln("Failed to send IP to API. Error: " + err.Error())
	}
	log.Default().Printf("Sent IP of %s to API and got status code of: %d", ip, statusCode)
}
