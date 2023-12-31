package server

import "github.com/kevinrudde/gophercraft/pkg/ping"

type ServerListPingEvent struct {
	ResponseData *ping.ResponseData
}
