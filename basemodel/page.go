package basemodel

type Page struct {
	*Term
	Data interface{} `json:"data"`
}

func NewPage(term *Term, data interface{}) *Page {
	return &Page{term, data}
}
