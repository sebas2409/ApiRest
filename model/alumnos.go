package model

type Alumno struct {
	Id     int     `json:"id"`
	Nombre string  `json:"nombre"`
	Curso  string  `json:"curso"`
	Notas  []Notas `json:"notas"`
}
