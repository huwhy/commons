package basemodel

import "fmt"

type Page struct {
	*Term
	Data interface{} `json:"data"`
}

func NewPage(term *Term, data interface{}) *Page {
	s := fmt.Sprint(data)
	if s == "nil" || s == "&[]" || s == "[]" {
		return &Page{term, []interface{}{}}
	}
	return &Page{term, data}
}
