package basemodel

type Page struct {
	*Term `json:"-"`
	Data  interface{} `json:"data"`
}

func NewPage(term *Term, data interface{}) *Page {
	return &Page{term, data}
}
