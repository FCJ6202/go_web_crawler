package model

// This structure of Content of page
type content struct {
	Data string
}

// this function will create a new Content
func newContent(data string) *content {
	return &content{
		Data: data,
	}
}
