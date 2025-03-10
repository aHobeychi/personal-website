package models

type Blog struct {
	Id            string   `json:"id"`
	Title         string   `json:"title"`
	Description   string   `json:"description"`
	Tags          []string `json:"tags"`
	PublishedDate string   `json:"publishedDate"`
	ExternalLink  string   `json:"externalLink"`
}
