package configuration

import (
	"github.com/kevinrudde/gophercraft/internal/network"
)

type ClientInformationPacket struct {
	Locale              string
	ViewDistance        byte
	ChatMode            byte
	ChatColors          bool
	DisplayedSkinParts  byte
	MainHand            byte
	EnableTextFiltering bool
	AllowServerListings bool
}

func (p *ClientInformationPacket) From(buffer network.Buffer) error {
	var err error

	p.Locale, err = buffer.ReadString()
	if err != nil {
		return err
	}

	p.ViewDistance, err = buffer.ReadByte()
	if err != nil {
		return err
	}

	p.ChatMode, err = buffer.ReadByte()
	if err != nil {
		return err
	}

	p.ChatColors, err = buffer.ReadBool()
	if err != nil {
		return err
	}

	p.DisplayedSkinParts, err = buffer.ReadByte()
	if err != nil {
		return err
	}

	p.MainHand, err = buffer.ReadByte()
	if err != nil {
		return err
	}

	p.EnableTextFiltering, err = buffer.ReadBool()
	if err != nil {
		return err
	}

	p.AllowServerListings, err = buffer.ReadBool()
	if err != nil {
		return err
	}

	return err
}
