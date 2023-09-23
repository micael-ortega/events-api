package models

import "time"

type Event struct {
	ID         int       `json:"id"`
	Begin_date time.Time `json:"begin_date"`
	End_date   time.Time `json:"end_date"`
	Modality   string    `json:"modality"`
	Duration   float32   `json:"duration"`
	Instructor Instructor       `json:"instructor"`
	Course     Course       `json:"course"`
}
