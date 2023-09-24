package models

type AttendanceEvent struct {
	ID         int  `json:"id"`
	EventID    int  `json:"event_id"`
	AttendeeID int  `json:"attendee_id"`
	Status     bool `json:"status"`
}
