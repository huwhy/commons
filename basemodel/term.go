package basemodel

type Term struct {
	Page      int  `json:"page"`
	Size      int  `json:"size"`
	LastId    int  `json:"lastId"`
	Total     int  `json:"total"`
	TotalPage int  `json:"totalPage"`
	QueryAll  bool `json:"queryAll"` //查询全部数据
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

func (term *Term) GetTotalPage() int {
	var totalPage = 0
	if term.Size <= 0 {
		return 0
	}
	if term.Total%term.Size > 0 {
		totalPage += 1
	}
	totalPage += term.Total / term.Size
	return totalPage
}
