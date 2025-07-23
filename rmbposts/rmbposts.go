package rmbposts

type RmbPost struct {
	Nation string `json:"nation"`
	Region string `json:"region"`
	Text   string `json:"text"`
}

func New(nation string, region string, text string) RmbPost {
	return RmbPost{
		nation,
		region,
		text,
	}
}
