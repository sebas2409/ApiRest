package handlers

import (
	"ApiRest/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

//base de datos
var bdRafa = model.Notas{
	Asignatura: "Base de datos",
	Notas:      8,
}
var bdMario = model.Notas{
	Asignatura: "Base de datos",
	Notas:      9,
}
var bdDiego = model.Notas{
	Asignatura: "Base de datos",
	Notas:      5,
}
var bdSebas = model.Notas{
	Asignatura: "Base de datos",
	Notas:      9,
}

//lenguaje de marcas
var lmRafa = model.Notas{
	Asignatura: "Lenguaje de marcas",
	Notas:      9,
}
var lmMario = model.Notas{
	Asignatura: "Lenguaje de marcas",
	Notas:      6,
}
var lmDiego = model.Notas{
	Asignatura: "Lenguaje de marcas",
	Notas:      10,
}
var lmSebas = model.Notas{
	Asignatura: "Lenguaje de marcas",
	Notas:      8,
}

//fol
var folRafa = model.Notas{
	Asignatura: "Fol",
	Notas:      7,
}
var folMario = model.Notas{
	Asignatura: "Fol",
	Notas:      7,
}
var folDiego = model.Notas{
	Asignatura: "Fol",
	Notas:      8,
}
var folSebas = model.Notas{
	Asignatura: "Fol",
	Notas:      8,
}

//programacion
var proRafa = model.Notas{
	Asignatura: "Programacion",
	Notas:      6,
}
var proMario = model.Notas{
	Asignatura: "Programacion",
	Notas:      5,
}
var proDiego = model.Notas{
	Asignatura: "Programacion",
	Notas:      8,
}
var proSebas = model.Notas{
	Asignatura: "Programacion",
	Notas:      10,
}

//sistemas informaticos
var siRafa = model.Notas{
	Asignatura: "Sistemas Informaticos",
	Notas:      7,
}
var siMario = model.Notas{
	Asignatura: "Sistemas Informaticos",
	Notas:      7,
}
var siDiego = model.Notas{
	Asignatura: "Sistemas Informaticos",
	Notas:      6,
}
var siSebas = model.Notas{
	Asignatura: "Sistemas Informaticos",
	Notas:      9,
}

//entornos de desarrollo
var edRafa = model.Notas{
	Asignatura: "Entornos de desarrollo",
	Notas:      7,
}
var edMario = model.Notas{
	Asignatura: "Entornos de desarrollo",
	Notas:      6,
}
var edDiego = model.Notas{
	Asignatura: "Entornos de desarrollo",
	Notas:      5,
}
var edSebas = model.Notas{
	Asignatura: "Entornos de desarrollo",
	Notas:      7,
}

//Alumnos
var sebas = model.Alumno{
	Id:     1,
	Nombre: "Juan Sebastián González",
	Curso:  "1 DAM",
	Notas:  []model.Notas{bdSebas, lmSebas, folSebas, proSebas, siSebas, edSebas},
}
var rafa = model.Alumno{
	Id:     2,
	Nombre: "Rafael Martinez",
	Curso:  "1 DAM",
	Notas:  []model.Notas{bdRafa, lmRafa, folRafa, proRafa, siRafa, edRafa},
}
var diego = model.Alumno{
	Id:     3,
	Nombre: "Diego Gomez",
	Curso:  "1 DAM",
	Notas:  []model.Notas{bdDiego, lmDiego, folDiego, proDiego, siDiego, edDiego},
}
var mario = model.Alumno{
	Id:     4,
	Nombre: "Mario Valverde",
	Curso:  "1 DAM",
	Notas:  []model.Notas{bdMario, lmMario, folMario, proMario, siMario, edMario},
}

var listaAlumnos = []model.Alumno{sebas, rafa, diego, mario}

func RequestHandler(c *gin.Context) {
	c.JSON(http.StatusOK, listaAlumnos)
}
