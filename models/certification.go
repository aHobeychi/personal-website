package models

type Certification struct {
	Name         string `json:"name"`
	Issuer       string `json:"issuer"`
	DateReceived string `json:"dateReceived"`
	Url          string `json:"url"`
}
