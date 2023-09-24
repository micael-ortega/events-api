package models

type Attendee struct {
	ID      int `json:"id"`
	Name    int `json:"name"`
	CPF     int `json:"cpf"`
	Role    int `json:"role"`
	Company int `json:"Company"`
	Board   int `json:"board"`
	Branch  int `json:"branch"`
}
