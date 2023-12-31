package chat

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
)

type generalMessage struct {
	Bold          bool `json:"bold,omitempty"`
	Italic        bool `json:"italic,omitempty"`
	Underlined    bool `json:"underlined,omitempty"`
	Strikethrough bool `json:"strikethrough,omitempty"`
	Obfuscated    bool `json:"obfuscated,omitempty"`

	Font  string `json:"font,omitempty"`
	Color string `json:"color,omitempty"`

	Insertion  string      `json:"insertion,omitempty"`
	ClickEvent *ClickEvent `json:"clickEvent,omitempty"`
	HoverEvent *HoverEvent `json:"hoverEvent,omitempty"`

	Translate string    `json:"translate"`
	With      []Message `json:"with,omitempty"`
	Extra     []Message `json:"extra,omitempty"`
}

type Message struct {
	generalMessage
	Text string `json:"text"`
}

type translatableMessage struct {
	generalMessage
	Text string `json:"text,omitempty"`
}

func (s Message) MarshalJson() ([]byte, error) {
	if s.Translate != "" {
		return json.Marshal(translatableMessage(s))
	}
	return json.Marshal(s)
}

func (s Message) UnmarshallJson(data []byte) error {
	data = bytes.TrimSpace(data)
	if len(data) == 0 {
		return io.EOF
	}

	switch data[0] {
	case '"':
		return json.Unmarshal(data, &s.Text)
	case '{':
		return json.Unmarshal(data, &s)
	case '[':
		return json.Unmarshal(data, &s.Extra)
	default:
		return fmt.Errorf("unknown message type %s", string(data[0]))
	}
}

func (s Message) Append(extraMessage ...Message) Message {
	s.Extra = append(s.Extra, extraMessage...)
	return s
}

func (s Message) WithColor(color string) Message {
	s.Color = color
	return s
}

func Text(text string) Message {
	return Message{Text: text}
}
