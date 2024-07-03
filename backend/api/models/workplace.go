package models

type Workplace struct {
	Company     Company `json:"company_name"`
	Description string  `json:"description"`
	StartDate   string  `json:"start_date"`
	EndDate     string  `json:"end_date"`
}

type Workplaces []Workplace
