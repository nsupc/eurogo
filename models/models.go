package models

type Telegram struct {
	Sender    string `json:"sender"`
	Recipient string `json:"recipient"`
	Id        string `json:"id"`
	Secret    string `json:"secret_key"`
	Type      string `json:"tg_type"`
}
