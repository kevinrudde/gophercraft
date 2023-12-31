package main

import (
	"github.com/kevinrudde/gophercraft/internal/event"
	server2 "github.com/kevinrudde/gophercraft/internal/event/server"
	"github.com/kevinrudde/gophercraft/pkg/chat"
	"github.com/kevinrudde/gophercraft/pkg/minecraft"
	"github.com/kevinrudde/gophercraft/pkg/ping"
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
	event.AddListener[*server2.ServerListPingEvent](event.Highest, func(event *server2.ServerListPingEvent) {
		log.Println("ServerListPingEvent 1")
		event.ResponseData = GetListResponse("Gophercraft")
	})

	event.AddListener[*server2.ServerListPingEvent](event.Lowest, func(event *server2.ServerListPingEvent) {
		log.Println("ServerListPingEvent 2")
		event.ResponseData = GetListResponse("Gophercraft 2")
	})

	event.AddListener[*server2.ServerListPingEvent](event.Monitor, func(event *server2.ServerListPingEvent) {
		log.Println("ServerListPingEvent 3")
		event.ResponseData = GetListResponse("Gophercraft 3")
	})
}

func GetListResponse(description string) *ping.ResponseData {
	chatComponent := chat.Text("Gophercraft\n").WithColor("#fca903")
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
