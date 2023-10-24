package model

import (
	"time"
)

// It is structure of Page in this site
type Page struct {
	URL         string
	LastCrawled time.Time
	Content     *content
}

// This function used for create a new Page
func NewPage(url string, data string) *Page {
	return &Page{
		URL:         url,
		LastCrawled: time.Now(),
		Content:     newContent(data),
	}
}

// This method is used for check that page is crossed oldPagetime or not.
// It means if till now page spent more than oldPagetime minute than this have to true that indicates that Page became older.
func (p *Page) Crosstime(oldPagetime float64) bool {
	return time.Since(p.LastCrawled).Minutes() > oldPagetime
}
