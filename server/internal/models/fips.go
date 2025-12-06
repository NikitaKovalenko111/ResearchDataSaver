package models

type FipsContent struct {
	Id             int      `json:"id"`
	Name           string   `json:"name"`
	Link           string   `json:"link"`
	Type           string   `json:"type"`
	Annotation     string   `json:"annotation"`
	Registration   string   `json:"registration"`
	PublishingDate string   `json:"publishingDate"`
	Applicant      string   `json:"applicant"`
	Address        string   `json:"address"`
	Authors        []string `json:"authors"`
}
