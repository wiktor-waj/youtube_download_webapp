package page

import (
	"html/template"
	"os"
)

type Page struct {
	Title       string
	Body        []byte
	DisplayBody template.HTML
}

// func (p *Page) save() error {
// 	filename := "data/" + p.Title + ".txt"
// 	return os.WriteFile(filename, p.Body, 0600)
// }

func LoadPage(title string) (*Page, error) {
	filename := "data/" + title
	body, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}
