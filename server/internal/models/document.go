package models

type Document struct {
	Id             int    `json:"id"`
	Name           string `json:"name"`
	Annotation     string `json:"annotation"`
	Link           string `json:"link"`
	PublishingDate string `json:"publishingDate"`
	Author         string `json:"author"`
}
