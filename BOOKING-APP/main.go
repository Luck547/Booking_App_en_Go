package main

import (
	"fmt"
	"strings"
)

func main() {
	var appName = "Nombres de usuario para STTDR."
	const cantidaddeUsuariosDisponibles uint = 50
	var usuariosqueQuedan int = 50
	var bookings []string
	fmt.Printf("Los usuarios son %T, los que quedan son %T, el nombre de esta app es %T\n", cantidaddeUsuariosDisponibles, cantidaddeUsuariosDisponibles, appName)
	fmt.Printf("Bienvenido a la app de reserva de %v\n", appName)
	fmt.Printf("Tenemos un total de %v usuarios disponibles, y van quedando solo %v libres.\n", cantidaddeUsuariosDisponibles, usuariosqueQuedan)
	fmt.Println("Reserva tu nombre de usuario aqui.")
	for {
		var userName1 string
		var userName2 string
		var email string
		var usuarios int
		var nombresdeUsuarios string
		// Solicita nombre de usuario al usuario
		fmt.Println("tipea tu primer nombre")
		fmt.Scan(&userName1)
		fmt.Println("tipea tu apellido")
		fmt.Scan(&userName2)
		fmt.Println("tipea tu email")
		fmt.Scan(&email)
		fmt.Println("Cuantos usuarios reservas?")
		fmt.Scan(&usuarios)
		fmt.Println("Ingresa tu nombre de usuario a reservar")
		fmt.Scan(&nombresdeUsuarios)

		isValidName := len(userName1) >= 3 && len(userName2) >= 3
		isValidEmail := strings.Contains(email, "@")
		isValidUsuarios := usuarios > 0 && usuarios <= usuariosqueQuedan
		isValidNombresdeUsuarios := len(nombresdeUsuarios) >= 5
		//tipo := "Usuario publico" OR "Usuario privado"
		//isValidType := tipo == "Usuario publico" || tipo == "Usuario privado"
		if isValidName && isValidEmail && isValidUsuarios && isValidNombresdeUsuarios {
			usuariosqueQuedan = usuariosqueQuedan - usuarios
			bookings = append(bookings, userName1+" "+userName2)
			fmt.Printf("Gracias %v %v por reservar %v Nombres de usuario.\nEl primero es %v. Te llagara un correo de confirmacion a %v\n", userName1, userName2, usuarios, nombresdeUsuarios, email)
			fmt.Printf("Solo quedan %v nombres de usuario en %v,. Gracias!\n", usuariosqueQuedan, appName)
			var userName []string
			for _, booking := range bookings {
				var names = strings.Fields(booking)
				userName = append(userName, names[0])
			}
			fmt.Printf("These are all the bookings: %v\n", userName)
			if usuariosqueQuedan == 0 {
				// Finalizar programa
				fmt.Println("Se acabaron los nombres de usuario, reinicia el programa")
				break
			}
		} else {
			//fmt.Printf("Solo tenemos %v nombres de usuario, por eso no puedes reservar %v nombres de usuario\n", usuariosqueQuedan, usuarios)
			fmt.Println("Has ingresados datos invalidos, prueba de nuevo")
		}
	}
}
