package telegrams

type Type int

const (
	Undefined Type = iota
	Standard
	Recruitment
)

func (t Type) String() string {
	switch t {
	case 0:
		return ""
	case 1:
		return "standard"
	case 2:
		return "recruitment"
	}
	return ""
}

func (t Type) MarshalText() ([]byte, error) {
	return []byte(t.String()), nil
}

type List struct {
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
	Type      Type   `json:"tg_type"`
}

type DeleteTelegram struct {
	Recipient string `json:"recipient"`
	Id        string `json:"id"`
}

func New(sender string, recipient string, id string, secret string, ttype Type) NewTelegram {
	return NewTelegram{
		sender,
		recipient,
		id,
		secret,
		ttype,
	}
}

func Delete(recipient string, id string) DeleteTelegram {
	return DeleteTelegram{
		recipient,
		id,
	}
}
