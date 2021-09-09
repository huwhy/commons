package basemodel

type Page struct {
	*Term `json:"-"`
	Total int         `json:"total"`
	Data  interface{} `json:"data"`
}

func NewPage(term *Term, data interface{}) *Page {
	return &Page{term, 0, data}
}

func (p *Page) GetTotalPage() int {
	var totalPage int = 0
	if p.Total%p.Size > 0 {
		totalPage += 1
	}
	totalPage += p.Total / p.Size
	return totalPage
}
