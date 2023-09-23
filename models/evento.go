package models

import "time"


type Evento struct {
	ID int `json:"id"`
	Data_ini time.Time `json:"data_ini"`
	Data_fim time.Time `json:"data_fim"`
	Modalidade string `json:"modalidade"`
	Duracao float32 `json:"duracao"`
	Instrutor Instrutor `json:"instrutor"`
	Curso Curso `json:"curso"`
}
