package configuration

import (
	"github.com/kevinrudde/gophercraft/internal/network/packets/server/configuration"
	networkplayer "github.com/kevinrudde/gophercraft/internal/network/player"
	"log"
)

func ProcessClientInformationPacket(connection *networkplayer.PlayerConnection, packet *ClientInformationPacket) error {
	featureFlags := make([]string, 1)
	featureFlags[0] = "minecraft:vanilla"

	featureFlagPacket := &configuration.FeatureFlagsPacket{FeatureFlags: featureFlags}

	log.Printf("Sending client information packet to server: %v", featureFlagPacket)

	err := connection.SendPacket(featureFlagPacket)
	if err != nil {
		return err
	}

	clientboundKnownPacksPacket := &configuration.ClientboundKnownPacksPacket{
		KnownPacks: []configuration.KnownPacks{
			{
				Namespace: "minecraft",
				ID:        "core",
				Version:   "1.21",
			},
		},
	}

	return connection.SendPacket(clientboundKnownPacksPacket)
}
