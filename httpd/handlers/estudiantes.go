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
var bdDaniel = model.Notas{
	Asignatura: "Base de datos",
	Notas:      9,
}
var bdVictor = model.Notas{
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
var lmDaniel = model.Notas{
	Asignatura: "Lenguaje de marcas",
	Notas:      6,
}
var lmVictor = model.Notas{
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
var folDaniel = model.Notas{
	Asignatura: "Fol",
	Notas:      7,
}
var folVictor = model.Notas{
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
var proDaniel = model.Notas{
	Asignatura: "Programacion",
	Notas:      5,
}
var proVictor = model.Notas{
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
var siDaniel = model.Notas{
	Asignatura: "Sistemas Informaticos",
	Notas:      7,
}
var siVictor = model.Notas{
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
var edDaniel = model.Notas{
	Asignatura: "Entornos de desarrollo",
	Notas:      6,
}
var edVictor = model.Notas{
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
var victor = model.Alumno{
	Id:     3,
	Nombre: "Victor Sanchez",
	Curso:  "1 DAM",
	Notas:  []model.Notas{bdVictor, lmVictor, folVictor, proVictor, siVictor, edVictor},
}
var daniel = model.Alumno{
	Id:     4,
	Nombre: "Daniel",
	Curso:  "1 DAM",
	Notas:  []model.Notas{bdDaniel, lmDaniel, folDaniel, proDaniel, siDaniel, edDaniel},
}

var listaAlumnos = []model.Alumno{sebas, rafa, victor, daniel}

func RequestHandler(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Credentials", "true")
	c.JSON(http.StatusOK, listaAlumnos)
}
