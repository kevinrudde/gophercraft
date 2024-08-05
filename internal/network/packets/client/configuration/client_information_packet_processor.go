package configuration

import (
	"github.com/kevinrudde/gophercraft/internal/network/packets/server/configuration"
	networkplayer "github.com/kevinrudde/gophercraft/internal/network/player"
	"log"
)

func ProcessClientInformationPacket(connection *networkplayer.PlayerConnection, packet *ClientInformationPacket) error {
	featureFlags := make([]string, 1)
	featureFlags[0] = "minecraft:vanilla"

	response := &configuration.FeatureFlagsPacket{FeatureFlags: featureFlags}

	log.Printf("Sending client information packet to server: %v", response)

	return connection.SendPacket(response)
}
