package models

import "time"

type Workplace struct {
	Company     Company   `json:"company_name"`
	Description string    `json:"description"`
	StartDate   time.Time `json:"start_date"`
	EndDate     time.Time `json:"end_date"`
}

type Workplaces []Workplace
