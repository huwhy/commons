package basemodel

type Term struct {
	Page   int `json:"page"`
	Size   int `json:"size"`
	LastId int `json:"lastId"`
}

func (term *Term) GetOffset() int {
	if term.LastId > 0 {
		return 0
	} else {
		if term.Page == 0 {
			term.Page = 1
		}
		return (term.Page - 1) * term.Size
	}
}
