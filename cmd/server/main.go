package main

import (
	"encoding/json"
	"github.com/kevinrudde/gophercraft/internal/event"
	server2 "github.com/kevinrudde/gophercraft/internal/event/server"
	"github.com/kevinrudde/gophercraft/pkg/minecraft"
	"log"
)

func main() {
	minecraft.Instance.Init()
	Listen()

	log.Println("Listening on :25565")
	err := minecraft.Instance.Start(":25565")

	if err != nil {
		log.Fatal(err)
	}
}

func Listen() {
	event.Listen[*server2.ServerListPingEvent](event.Highest, func(event *server2.ServerListPingEvent) {
		log.Println("ServerListPingEvent 1")
		event.ResponseData = GetListResponse("Gophercraft")
	})

	event.Listen[*server2.ServerListPingEvent](event.Lowest, func(event *server2.ServerListPingEvent) {
		log.Println("ServerListPingEvent 2")
		event.ResponseData = GetListResponse("Gophercraft 2")
	})

	event.Listen[*server2.ServerListPingEvent](event.Monitor, func(event *server2.ServerListPingEvent) {
		log.Println("ServerListPingEvent 3")
		event.ResponseData = GetListResponse("Gophercraft 3")
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
