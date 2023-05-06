package models

import "encoding/json"

type Response struct {
	Ts      string       `json:"ts"`
	Updates []GroupEvent `json:"updates"`
	Failed  int          `json:"failed"`
}

type GroupEvent struct {
	Type    string          `json:"type"`
	Object  json.RawMessage `json:"object"`
	GroupID int             `json:"group_id"`
	EventID string          `json:"event_id"`
	V       string          `json:"v"`
}
