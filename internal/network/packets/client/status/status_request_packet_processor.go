package status

import (
	"errors"
	"github.com/kevinrudde/gophercraft/internal/network/packets/client/common"
	"github.com/kevinrudde/gophercraft/internal/network/packets/server/handshake"
	networkplayer "github.com/kevinrudde/gophercraft/internal/network/player"
	"log"
	"reflect"
)

func ProcessStatusRequestPacket(connection *networkplayer.PlayerConnection, p common.ClientPacket) error {
	_, ok := p.(*StatusRequestPacket)
	if !ok {
		return errors.New("expected StatusRequestPacket, but got " + reflect.TypeOf(p).String())
	}

	log.Println("sending status response packet")

	response := handshake.StatusResponsePacket{
		Payload: "{\n    \"version\": {\n        \"name\": \"1.19.4\",\n        \"protocol\": 762\n    },\n    \"players\": {\n        \"max\": 100,\n        \"online\": 5,\n        \"sample\": [\n            {\n                \"name\": \"thinkofdeath\",\n                \"id\": \"4566e69f-c907-48ee-8d71-d7ba5aa00d20\"\n            }\n        ]\n    },\n    \"description\": {\n        \"text\": \"Hello world\"\n    },\n    \"favicon\": \"data:image/png;base64,<data>\",\n    \"enforcesSecureChat\": true,\n    \"previewsChat\": true\n}",
	}

	return connection.SendPacket(&response)
}
