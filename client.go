package main

import (
	"fmt"
	"net/rpc"
)

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
			var nombreAlumno, materia, reply string
			var calificacion float64

			fmt.Print("Nombre del alumno: ")
			fmt.Scanln(&nombreAlumno)

			fmt.Println("Materia: ")
			fmt.Scanln(&materia)

			fmt.Println("Calificacion: ")
			fmt.Scanln(&calificacion)

			err = newClient.Call("Server.AgregarCalificacion", nombreAlumno, materia, calificacion, &reply)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println(reply)
			}

		// case 2:
		// 	var nombreAlumno string
		// 	var promedio float64

		// 	fmt.Print("Nombre del alumno: ")
		// 	fmt.Scanln(&nombreAlumno)

		// 	// err = newClient.Call("Server.ObtenerPromedioIndividual",
		// 	//						nombreAlumno, &promedio)
		// 	if err != nil {
		// 		fmt.Println(err)
		// 	} else {
		// 		// fmt.Println("Promedio:", promedio)
		// 	}

		// case 3:
		// 	var promedio float64

		// 	// err = newClient.Call("Server.ObtenerPromedioGrupal", &promedio)
		// 	if err != nil {
		// 		fmt.Println(err)
		// 	} else {
		// 		// fmt.Println("Promedio:", promedio)
		// 	}

		// case 4:
		// 	var materia string
		// 	var promedio float64

		// 	fmt.Println("Materia: ")
		// 	fmt.Scanln(&materia)

		// 	// err = newClient.Call("Server.ObtenerPromedioMateria", &promedio)
		// 	if err != nil {
		// 		fmt.Println(err)
		// 	} else {
		// 		// fmt.Println("Promedio:", promedio)
		// 	}

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
