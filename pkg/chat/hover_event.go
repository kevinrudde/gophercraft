package chat

import "encoding/json"

type HoverEvent struct {
	Action   string          `json:"action"`
	Contents json.RawMessage `json:"contents"`
	Value    Message         `json:"value"`
}
