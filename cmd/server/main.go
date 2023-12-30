package main

import (
	"encoding/json"
	"fmt"
	"github.com/kevinrudde/gophercraft/internal/crypto"
	"github.com/kevinrudde/gophercraft/internal/event"
	server2 "github.com/kevinrudde/gophercraft/internal/event/server"
	"github.com/kevinrudde/gophercraft/internal/network/packets/client"
	"github.com/kevinrudde/gophercraft/internal/network/server"
	"log"
)

func main() {
	go Init()
	networkServer := server.NewServer(":25565")
	fmt.Println("Listening on :25565")

	client.InitializeClientPacketProcessors()

	log.Fatal(networkServer.Start())
}

func Init() {
	crypto.Init()
	Listen()
}

func Listen() {
	event.Listen(&server2.ServerListPingEvent{}, event.Lowest, func(message event.Message) {
		log.Println("ServerListPingEvent")
		serverListPingEvent := message.(*server2.ServerListPingEvent)
		serverListPingEvent.ResponseData = GetListResponse("Gophercraft")
	})

	event.Listen(&server2.ServerListPingEvent{}, event.Highest, func(message event.Message) {
		log.Println("ServerListPingEvent2")
		serverListPingEvent := message.(*server2.ServerListPingEvent)
		serverListPingEvent.ResponseData = GetListResponse("Gophercraft2")
	})

	event.Listen(&server2.ServerListPingEvent{}, event.Monitor, func(message event.Message) {
		log.Println("ServerListPingEvent3")
		serverListPingEvent := message.(*server2.ServerListPingEvent)
		serverListPingEvent.ResponseData = GetListResponse("Gophercraft3")
	})
}

func GetListResponse(description string) string {
	var list struct {
		Version struct {
			Name     string `json:"name"`
			Protocol int    `json:"protocol"`
		} `json:"version"`
		Players struct {
			Max    int `json:"max"`
			Online int `json:"online"`
			Sample []struct {
				Name string `json:"name"`
				ID   string `json:"id"`
			} `json:"sample,omitempty"`
		} `json:"players"`
		Description struct {
			Text string `json:"text"`
		} `json:"description"`
		EnforcesSecureChat bool   `json:"enforcesSecureChat"`
		FavIcon            string `json:"favicon,omitempty"`
		PreviewsChat       bool   `json:"previewsChat"`
	}

	list.Version.Protocol = 765
	list.Version.Name = "1.20.4"
	list.Players.Max = 100
	list.Players.Online = 0
	list.Description.Text = description
	list.PreviewsChat = true

	json, err := json.Marshal(list)
	if err != nil {
		return ""
	}

	return string(json)
}
