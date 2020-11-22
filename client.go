package main

import (
	"fmt"
	"net/rpc"
)

// Informacion ...
type Informacion struct {
	NombreAlumno string
	Materia      string
	Calificacion float64
}

func client() {
	newClient, err := rpc.Dial("tcp", "127.0.0.1:9999")
	if err != nil {
		fmt.Println(err)
		return
	}

	var op uint64

	for {
		fmt.Println("1) Agregar calificación de una materia.")
		fmt.Println("2) Mostrar el promedio de un alumno.")
		fmt.Println("3) Mostrar el promedio general.")
		fmt.Println("4) Mostrar el promedio de una materia.")
		fmt.Println("0) Salir.")
		fmt.Scanln(&op)

		switch op {
		case 1:
			var alumno Informacion
			var reply string

			fmt.Print("Nombre del alumno: ")
			fmt.Scanln(&alumno.NombreAlumno)

			fmt.Println("Materia: ")
			fmt.Scanln(&alumno.Materia)

			fmt.Println("Calificacion: ")
			fmt.Scanln(&alumno.Calificacion)

			err = newClient.Call("Server.AgregarCalificacion", alumno, &reply)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println(reply)
			}

		case 2:
			var nombreAlumno, reply float64

			fmt.Print("Nombre del alumno: ")
			fmt.Scanln(&nombreAlumno)

			err = newClient.Call("Server.ObtenerPromedioIndividual", nombreAlumno, &reply)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("Promedio:", reply)
			}

		case 3:
			var reply float64

			err = newClient.Call("Server.ObtenerPromedioGrupal", 0, &reply)
			// Como no necesito un primer parametro, lo "evito" mandando un cero...que no se usa.
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("Promedio general:", reply)
			}

		case 4:
			var materia string
			var reply float64

			fmt.Println("Materia: ")
			fmt.Scanln(&materia)

			err = newClient.Call("Server.ObtenerPromedioMateria", materia, &reply)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("Promedio:", reply)
			}

		case 0:
			fmt.Println("Saliendo del cliente.")
			return

		default:
			fmt.Println("Opcion no valida.")
		}
	}
}

func main() {
	client()
}
