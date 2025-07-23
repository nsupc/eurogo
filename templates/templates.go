package templates

import "time"

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

func New(nation string, tgid int, key string, description string) NewTemplate {
	return NewTemplate{
		nation,
		tgid,
		key,
		description,
	}
}

func Edit(id string, nation string, tgid int, key string, description string) EditTemplate {
	return EditTemplate{
		id,
		nation,
		tgid,
		key,
		description,
	}
}
