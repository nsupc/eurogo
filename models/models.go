package models

import "time"

type TelegramList struct {
	Recruitment []Telegram `json:"recruitment"`
	Standard    []Telegram `json:"standard"`
}

type Telegram struct {
	Recipient string `json:"recipient"`
	Id        string `json:"id"`
}

type NewTelegram struct {
	Sender    string `json:"sender"`
	Recipient string `json:"recipient"`
	Id        string `json:"id"`
	Secret    string `json:"secret_key"`
	Type      string `json:"tg_type"`
}

type DeleteTelegram struct {
	Recipient string `json:"recipient"`
	Id        string `json:"id"`
}

type Dispatch struct {
	Id          int       `json:"id"`
	Nation      string    `json:"nation"`
	Category    int       `json:"category"`
	SubCategory int       `json:"subcategory"`
	Title       string    `json:"title"`
	Text        string    `json:"text"`
	CreatedBy   string    `json:"created_by"`
	ModifiedAt  time.Time `json:"modified_at"`
}

type NewDispatch struct {
	Nation      string `json:"nation"`
	Title       string `json:"title"`
	Text        string `json:"text"`
	Category    int    `json:"category"`
	SubCategory int    `json:"subcategory"`
}

type EditDispatch struct {
	Id          int    `json:"-"`
	Title       string `json:"title"`
	Text        string `json:"text"`
	Category    int    `json:"category"`
	SubCategory int    `json:"subcategory"`
}

type DispatchStatus struct {
	Id         int       `json:"id"`
	Action     string    `json:"action"`
	Status     string    `json:"status"`
	DispatchId int       `json:"dispatch_id"`
	Error      string    `json:"error"`
	CreatedAt  time.Time `json:"created_at"`
	ModifiedAt time.Time `json:"modified_at"`
}

type RmbPost struct {
	Nation string `json:"nation"`
	Region string `json:"region"`
	Text   string `json:"text"`
}

type Template struct {
	Id          string    `json:"id"`
	Nation      string    `json:"nation"`
	Tgid        int       `json:"tgid"`
	Key         string    `json:"key"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	ModifiedAt  time.Time `json:"modified_at"`
}

type NewTemplate struct {
	Nation      string `json:"nation"`
	Tgid        int    `json:"tgid"`
	Key         string `json:"key"`
	Description string `json:"description"`
}

type EditTemplate struct {
	Id          string `json:"-"`
	Nation      string `json:"nation"`
	Tgid        int    `json:"tgid"`
	Key         string `json:"key"`
	Description string `json:"description"`
}
