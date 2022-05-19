package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Notas struct {
	Asignatura string `json:"asignatura"`
	Notas      int    `json:"nota"`
}

type Alumno struct {
	Id     int     `json:"id"`
	Nombre string  `json:"nombre"`
	Curso  string  `json:"curso"`
	Notas  []Notas `json:"notas"`
}

//base de datos
var bdRafa = Notas{
	Asignatura: "Base de datos",
	Notas:      8,
}
var bdMario = Notas{
	Asignatura: "Base de datos",
	Notas:      9,
}
var bdDiego = Notas{
	Asignatura: "Base de datos",
	Notas:      5,
}
var bdSebas = Notas{
	Asignatura: "Base de datos",
	Notas:      9,
}

//lenguaje de marcas
var lmRafa = Notas{
	Asignatura: "Lenguaje de marcas",
	Notas:      9,
}
var lmMario = Notas{
	Asignatura: "Lenguaje de marcas",
	Notas:      6,
}
var lmDiego = Notas{
	Asignatura: "Lenguaje de marcas",
	Notas:      10,
}
var lmSebas = Notas{
	Asignatura: "Lenguaje de marcas",
	Notas:      8,
}

//fol
//programacion
//sistemas informaticos
//entornos de desarrollo

var sebas = Alumno{
	Id:     1,
	Nombre: "Juan Sebastián González",
	Curso:  "1 DAM",
	Notas:  []Notas{bdSebas, lmSebas},
}
var rafa = Alumno{
	Id:     2,
	Nombre: "Rafael Martinez",
	Curso:  "1 DAM",
	Notas:  []Notas{bdRafa, lmRafa},
}

var listaAlumnos = []Alumno{sebas, rafa}

func RequestHandler(c *gin.Context) {
	c.JSON(http.StatusOK, listaAlumnos)
}
