package main

import (
	"errors"
	"fmt"
	"net"
	"net/rpc"
)

var materias map[string]map[string]float64
var alumnos map[string]map[string]float64

// Server ...
type Server struct{}

// AgregarCalificacion ...
func (server *Server) AgregarCalificacion(nombreAlumno string, materia string, calificacion float64, reply *string) error {
	calificacionActual := alumnos[nombreAlumno][materia]

	if calificacionActual != 0 {
		return errors.New("Ya existe una calificacion registrada. No es posible modificarla.")
	} else {
		alumnos[nombreAlumno][materia] = calificacion
		materias[materia][nombreAlumno] = calificacion
		fmt.Println(alumnos)
		fmt.Println(materias)
		*reply = "Calificacion registrada con exito."
		return nil
	}
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
		"Programacion": 0,
		"Algoritmia":   0,
		"Concurrentes": 0,
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
