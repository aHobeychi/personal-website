package models

type WorkExperience struct {
	JobTitle    string   `json:"jobTitle"`
	CompanyName string   `json:"companyName"`
	Description string   `json:"description"`
	StartDate   string   `json:"startDate"`
	EndDate     string   `json:"endDate"`
	Tags        []string `json:"tags"`
}
