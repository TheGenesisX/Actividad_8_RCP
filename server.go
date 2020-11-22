package main

import (
	"errors"
	"fmt"
	"net"
	"net/rpc"
)

var materias map[string]map[string]float64
var alumnos map[string]map[string]float64

// Informacion ...
type Informacion struct {
	NombreAlumno string
	Materia      string
	Calificacion float64
}

// Server ...
type Server struct{}

// AgregarCalificacion ...
func (server *Server) AgregarCalificacion(alumno Informacion, reply *string) error {
	if _, alumnoExists := alumnos[alumno.NombreAlumno]; alumnoExists {
		if _, materiaExists := alumnos[alumno.NombreAlumno][alumno.Materia]; materiaExists {

			calificacionActual := alumnos[alumno.NombreAlumno][alumno.Materia]

			if calificacionActual != 0 {
				return errors.New("Ya existe una calificacion registrada. No es posible modificarla")
			}
			alumnos[alumno.NombreAlumno][alumno.Materia] = alumno.Calificacion
			materias[alumno.Materia][alumno.NombreAlumno] = alumno.Calificacion
			*reply = "Calificacion registrada con exito."
			return nil
		}
		return errors.New("Materia no encontrada")
	}
	return errors.New("Alumno no encontrado")
}

// ObtenerPromedioIndividual ...
func (server *Server) ObtenerPromedioIndividual(NombreAlumno string, reply *float64) error {
	promedio, contadorMaterias := 0.0, 0.0

	if _, alumnoExists := alumnos[NombreAlumno]; alumnoExists {
		for _, value := range alumnos[NombreAlumno] {
			promedio += value
			contadorMaterias++
		}
		promedio /= contadorMaterias
		*reply = promedio
		return nil
	}
	return errors.New("Alumno no encontrado")
}

// ObtenerPromedioGrupal ... skip int no se usa. Explicacion en client.go
func (server *Server) ObtenerPromedioGrupal(skip int, reply *float64) error {
	promedios := make([]float64, len(alumnos))
	promedioGeneral := 0.0
	pos, contadorMaterias := 0, 0.0

	for i := 0; i < len(alumnos); i++ {
		promedios[i] = 0
	}

	for nombre := range alumnos {
		for _, calificacion := range alumnos[nombre] {
			promedios[pos] += calificacion
			contadorMaterias++
			// Obtenemos la suma de las calificaciones de cada alumno por individual.
		}
		promedios[pos] /= contadorMaterias
		pos++
		contadorMaterias = 0.0
	}

	for x := range promedios {
		promedioGeneral += promedios[x]
	}
	promedioGeneral /= float64(len(alumnos))
	*reply = promedioGeneral
	return nil
}

// ObtenerPromedioMateria ...
func (server *Server) ObtenerPromedioMateria(materia string, reply *float64) error {
	promedio, contadorAlumnos := 0.0, 0.0

	if _, materiaExists := materias[materia]; materiaExists {
		for _, value := range materias[materia] {
			promedio += value
			contadorAlumnos++
		}
		promedio /= contadorAlumnos
		*reply = promedio
		return nil
	}
	return errors.New("Materia no encontrada")
}

func server() {
	rpc.Register(new(Server))
	ln, err := net.Listen("tcp", ":9999")
	if err != nil {
		fmt.Println(err)
	}

	for {
		newClient, err := ln.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}
		go rpc.ServeConn(newClient)
	}
}

func main() {
	materias = make(map[string]map[string]float64)
	alumnos = make(map[string]map[string]float64)

	// Inicializacion de materias.
	materias["Programacion"] = map[string]float64{
		"Alejandro": 100,
		"Rafael":    100,
		"Efren":     100,
	}
	materias["Algoritmia"] = map[string]float64{
		"Alejandro": 100,
		"Rafael":    100,
		"Efren":     100,
	}
	materias["Concurrentes"] = map[string]float64{
		"Alejandro": 100,
		"Rafael":    100,
		"Efren":     100,
	}

	// Inicializacion de alumnos.
	alumnos["Alejandro"] = map[string]float64{
		"Programacion": 100,
		"Algoritmia":   100,
		"Concurrentes": 100,
	}
	alumnos["Rafael"] = map[string]float64{
		"Programacion": 100,
		"Algoritmia":   100,
		"Concurrentes": 100,
	}
	alumnos["Efren"] = map[string]float64{
		"Programacion": 100,
		"Algoritmia":   100,
		"Concurrentes": 100,
	}

	go server()

	var input string
	fmt.Scanln(&input)
}
