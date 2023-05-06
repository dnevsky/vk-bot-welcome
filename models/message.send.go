package models

type MessageSend struct {
	UserId   int    `json:"user_id"`
	RandomId int    `json:"random_id"`
	PeerId   int    `json:"peer_id"`
	Message  string `json:"message"`

	Keyboard Keyboard `json:"keyboard,omitempty"`

	Version     string `json:"v"`
	AccessToken string `json:"access_token"`
}
