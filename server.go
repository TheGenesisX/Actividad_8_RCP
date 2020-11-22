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
	calificacionActual := alumnos[alumno.NombreAlumno][alumno.Materia]

	if calificacionActual != 0 {
		return errors.New("Ya existe una calificacion registrada. No es posible modificarla")
	}
	alumnos[alumno.NombreAlumno][alumno.Materia] = alumno.Calificacion
	materias[alumno.Materia][alumno.NombreAlumno] = alumno.Calificacion
	*reply = "Calificacion registrada con exito."
	return nil
}

// ObtenerPromedioIndividual ...
func (server *Server) ObtenerPromedioIndividual(NombreAlumno string, reply *float64) error {
	promedio := 0.0
	contador := 0.0

	for _, value := range alumnos[NombreAlumno] {
		promedio += value
		contador++
	}

	if contador > 0 {
		promedio /= contador
		*reply = promedio
		return nil
	}
	return errors.New("No es posible obtener el promedio. Division entre cero")
}

// ObtenerPromedioGrupal ...
func (server *Server) ObtenerPromedioGrupal(reply *float64) error {
	promedios := make([]float64, len(alumnos))
	promedioGeneral := 0.0
	contador := 0

	for i := 0; i < len(alumnos); i++ {
		promedios[i] = 0
	}

	for nombre := range alumnos {
		for _, calificacion := range alumnos[nombre] {
			promedios[contador] += calificacion
			// Obtenemos la suma de las calificaciones de cada alumno por individual.
		}
		if promedios[contador] == 0 {
			return errors.New("No es posible obtener el promedio. Division entre cero")
		} // Si no sumamos nada, no podemos calcular promedio por problema de division entre cero.
		promedios[contador] /= float64(len(materias))
		contador++
	}
	contador = 0

	for x := range promedios {
		promedioGeneral += promedios[x]
	}
	promedioGeneral /= float64(len(materias))
	*reply = promedioGeneral
	return nil
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
		"Alejandro": 0,
		"Rafael":    0,
		"Efren":     0,
	}
	materias["Algoritmia"] = map[string]float64{
		"Alejandro": 0,
		"Rafael":    0,
		"Efren":     0,
	}
	materias["Concurrentes"] = map[string]float64{
		"Alejandro": 0,
		"Rafael":    0,
		"Efren":     0,
	}

	// Inicializacion de alumnos.
	alumnos["Alejandro"] = map[string]float64{
		"Programacion": 100,
		"Algoritmia":   50,
		"Concurrentes": 90,
	}
	alumnos["Rafael"] = map[string]float64{
		"Programacion": 0,
		"Algoritmia":   0,
		"Concurrentes": 0,
	}
	alumnos["Efren"] = map[string]float64{
		"Programacion": 0,
		"Algoritmia":   0,
		"Concurrentes": 0,
	}

	go server()

	var input string
	fmt.Scanln(&input)
}
