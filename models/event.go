package models

import "time"

type Event struct {
	ID         int     `json:"id"`
	Begin_date string  `json:"begin_date"`
	End_date   string  `json:"end_date"`
	Modality   string  `json:"modality"`
	Duration   float32 `json:"duration"`
	Instructor int     `json:"instructor_id"`
	Course     int     `json:"course_id"`
}

type EventResponse struct{
	ID         int     `json:"id"`
	Begin_date time.Time  `json:"begin_date"`
	End_date   time.Time  `json:"end_date"`
	Modality   string  `json:"modality"`
	Duration   float32 `json:"duration"`
	Instructor Instructor     `json:"instructor"`
	Course     Course     `json:"course"`
}
