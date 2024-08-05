package configuration

import "github.com/kevinrudde/gophercraft/internal/network"

type FeatureFlagsPacket struct {
	FeatureFlags []string
}

func (p *FeatureFlagsPacket) PacketId() int {
	return 0x0C
}

func (p *FeatureFlagsPacket) Write(buffer network.Buffer) error {
	return buffer.WriteStringSlice(p.FeatureFlags)
}
