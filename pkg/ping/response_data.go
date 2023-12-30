package ping

import (
	"encoding/json"
)

type StatusResponse struct {
	Version struct {
		Name     string `json:"name"`
		Protocol int    `json:"protocol"`
	} `json:"version"`
	Players struct {
		Max    int `json:"max,omitempty"`
		Online int `json:"online"`
		Sample []struct {
			Name string `json:"name"`
			ID   string `json:"id"`
		} `json:"sample,omitempty"`
	} `json:"players,omitempty"`
	Description struct {
		Text string `json:"text"`
	} `json:"description"`
	EnforcesSecureChat bool   `json:"enforcesSecureChat,omitempty"`
	FavIcon            string `json:"favicon,omitempty"`
	PreviewsChat       bool   `json:"previewsChat,omitempty"`
}

type ResponseData struct {
	Version       string
	Protocol      int
	HidePlayers   bool
	MaxPlayers    int
	OnlinePlayers int
	Description   string
	Favicon       string
}

func (s *ResponseData) GetPingResponse() string {
	var response StatusResponse
	response.Version.Name = s.Version
	response.Version.Protocol = s.Protocol

	if !s.HidePlayers {
		response.Players.Max = s.MaxPlayers
		response.Players.Online = s.OnlinePlayers
	}

	response.Description.Text = s.Description
	response.FavIcon = s.Favicon

	responseJson, err := json.Marshal(response)
	if err != nil {
		return ""
	}

	return string(responseJson)
}
