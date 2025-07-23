package dispatches

import "time"

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

type Status struct {
	Id         int       `json:"id"`
	Action     string    `json:"action"`
	Status     string    `json:"status"`
	DispatchId int       `json:"dispatch_id"`
	Error      string    `json:"error"`
	CreatedAt  time.Time `json:"created_at"`
	ModifiedAt time.Time `json:"modified_at"`
}

func New(nation string, title string, text string, category int, subcategory int) NewDispatch {
	return NewDispatch{
		nation,
		title,
		text,
		category,
		subcategory,
	}
}

func Edit(id int, title string, text string, category int, subcategory int) EditDispatch {
	return EditDispatch{
		id,
		title,
		text,
		category,
		subcategory,
	}
}
