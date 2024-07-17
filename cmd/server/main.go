package main

import (
	"log"

	"github.com/kevinrudde/gophercraft/internal/event"
	"github.com/kevinrudde/gophercraft/internal/event/server"
	"github.com/kevinrudde/gophercraft/pkg/chat"
	"github.com/kevinrudde/gophercraft/pkg/minecraft"
	"github.com/kevinrudde/gophercraft/pkg/ping"
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
	event.AddListener(event.Highest, func(event *server.ServerListPingEvent) {
		log.Println("ServerListPingEvent 1")
		event.ResponseData = GetListResponse("Gophercraft")
	})

	event.AddListener(event.Lowest, func(event *server.ServerListPingEvent) {
		log.Println("ServerListPingEvent 2")
		event.ResponseData = GetListResponse("Gophercraft 2")
	})

	event.AddListener(event.Monitor, func(event *server.ServerListPingEvent) {
		log.Println("ServerListPingEvent 3")
		event.ResponseData = GetListResponse("Gophercraft 3")
	})
}

func GetListResponse(description string) *ping.ResponseData {
	chatComponent := chat.Text(description + "\n").WithColor("#fca903")
	chatComponent = chatComponent.Append(chat.Text("1.20.4").WithColor("#a503fc"))

	return &ping.ResponseData{
		Version:       "1.20.4",
		Protocol:      765,
		HidePlayers:   false,
		MaxPlayers:    100,
		OnlinePlayers: 0,
		Description:   chatComponent,
		Favicon:       "",
	}
}
