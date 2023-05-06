package models

type Keyboard struct {
	OneTime bool       `json:"one_time"`
	Buttons [][]Button `json:"buttons"`
}

type Button struct {
	Action ButtonText `json:"action"`
	Color  string     `json:"color,omitempty"`
}

type ButtonText struct {
	Type    string `json:"type"`
	Payload string `json:"payload"`
	Label   string `json:"label"`
}

type Action struct {
	Type    string `json:"type"`
	Payload string `json:"payload,omitempty"`
	AppID   int    `json:"app_id,omitempty"`
	OwnerID int    `json:"owner_id,omitempty"`
	Hash    string `json:"hash,omitempty"`
	Label   string `json:"label,omitempty"`
}

func (k *Keyboard) AddButton(row int, button Button) {
	if len(k.Buttons) <= row {
		newRow := make([]Button, 0)
		k.Buttons = append(k.Buttons, newRow)
	}
	k.Buttons[row] = append(k.Buttons[row], button)
}
