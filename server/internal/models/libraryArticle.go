package models

type LibraryArticle struct {
	Id              int      `json:"id"`
	Name            string   `json:"name"`
	Annotation      string   `json:"annotation"`
	Link            string   `json:"link"`
	PublishingDate  string   `json:"publishingDate"`
	Lang            string   `json:"lang"`
	UDK             string   `json:"udk"`
	PublisherObject string   `json:"publisherObject"`
	Publisher       string   `json:"publisher"`
	Supervisor      string   `json:"supervisor"`
	Authors         []string `json:"authors"`
}
