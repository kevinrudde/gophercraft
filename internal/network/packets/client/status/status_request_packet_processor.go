package status

import (
	"encoding/json"
	"errors"
	"github.com/kevinrudde/gophercraft/internal/event"
	"github.com/kevinrudde/gophercraft/internal/event/server"
	"github.com/kevinrudde/gophercraft/internal/network/packets/client/common"
	"github.com/kevinrudde/gophercraft/internal/network/packets/server/handshake"
	networkplayer "github.com/kevinrudde/gophercraft/internal/network/player"
	"reflect"
)

func ProcessStatusRequestPacket(connection *networkplayer.PlayerConnection, p common.ClientPacket) error {
	_, ok := p.(*StatusRequestPacket)
	if !ok {
		return errors.New("expected StatusRequestPacket, but got " + reflect.TypeOf(p).String())
	}

	serverListEvent := &server.ServerListPingEvent{}

	var err error

	err = event.Call(serverListEvent, func() {
		response := &handshake.StatusResponsePacket{
			Payload: serverListEvent.ResponseData,
		}

		err = connection.SendPacket(response)
	})

	return err
}

func GetListResponse() string {
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
	list.Description.Text = "Gophercraft"
	list.PreviewsChat = true

	json, err := json.Marshal(list)
	if err != nil {
		return ""
	}

	return string(json)
}
