package model

import (
	"time"
)

type Page struct {
	URL         string
	LastCrawled time.Time
	Content     *content
}

func NewPage(url string, data string) *Page {
	return &Page{
		URL:         url,
		LastCrawled: time.Now(),
		Content:     newContent(data),
	}
}

func (p *Page) Crosstime(oldPagetime float64) bool {
	return time.Since(p.LastCrawled).Minutes() > oldPagetime
}
