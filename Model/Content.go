package model

type content struct {
	Data string
}

func newContent(data string) *content {
	return &content{
		Data: data,
	}
}
