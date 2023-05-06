package models

type MessageNewObject struct {
	Message Message `json:"message"`
	// ClientInfo ClientInfo      `json:"client_info"` // упустим за ненадобностью
}

// большая часть данных упущена за ненадобностью
type Message struct {
	FromID int `json:"from_id"`

	Payload string `json:"payload"`
	PeerID  int    `json:"peer_id"`

	Text         string `json:"text"`
	UpdateTime   int    `json:"update_time"`
	MembersCount int    `json:"members_count"`
}
