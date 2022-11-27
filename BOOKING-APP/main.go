package main

import (
	fmt "fmt"
	"strings"
)

func main() {
	var appName = "Nombres de usuario para STTDR."
	const cantidaddeUsuariosDisponibles uint = 50
	var usuariosqueQuedan int = 50
	var bookings []string

	greetUsers(appName, cantidaddeUsuariosDisponibles, usuariosqueQuedan)

	for {
		var userName1 string
		var userName2 string
		var email string
		var usuarios int
		var nombresdeUsuarios string
		// Solicita nombre de usuario al usuario
		printUserName1(bookings)
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

		isValidName, isValidEmail, isValidUsuarios, isValidNombresdeUsuarios := validateUserInput(userName1, userName2, email, usuarios, nombresdeUsuarios, usuariosqueQuedan)

		if isValidName && isValidEmail && isValidUsuarios && isValidNombresdeUsuarios {
			usuariosqueQuedan = usuariosqueQuedan - usuarios
			bookings = append(bookings, userName1+" "+userName2)
			fmt.Printf("Gracias %v %v por reservar %v Nombres de usuario.\nEl primero es %v. Te llagara un correo de confirmacion a %v\n", userName1, userName2, usuarios, nombresdeUsuarios, email)
			fmt.Printf("Solo quedan %v nombres de usuario en %v,. Gracias!\n", usuariosqueQuedan, appName)
			// Call a function print userName1
			printUserName1(bookings)

			fmt.Printf("These are all the bookings: %v\n", userName1)
			if usuariosqueQuedan == 0 {
				// Finalizar programa
				fmt.Println("Se acabaron los nombres de usuario, reinicia el programa")
				break
			}
		} else {
			if !isValidName {
				fmt.Println("El nombre es invalido")
			}
			if !isValidEmail {
				fmt.Println("El EMAIL es invalido")
			}
			if !isValidUsuarios {
				fmt.Println("El numero de usuarios es invalido")
			}
			if !isValidNombresdeUsuarios {
				fmt.Println("El nombre de usuario es invalido")
			} else {
				fmt.Println("Reservado")
			}

		}
	}
}

func greetUsers(appName string, cantidaddeUsuariosDisponibles uint, usuariosqueQuedan int) {
	fmt.Printf("Los usuarios son %T, los que quedan son %T, el nombre de esta app es %T\n", cantidaddeUsuariosDisponibles, cantidaddeUsuariosDisponibles, appName)
	fmt.Printf("Bienvenido a la app de reserva de %v\n", appName)
	fmt.Printf("Tenemos un total de %v usuarios disponibles, y van quedando solo %v libres.\n", cantidaddeUsuariosDisponibles, usuariosqueQuedan)
	fmt.Printf("Bienvenido a la reserva de %v", appName)
	fmt.Println("Reserva tu nombre de usuario aqui.")
}

func printUserName1(bookings []string) []string {
	var userName1 []string
	for _, booking := range bookings {
		var names = strings.Fields(booking)
		userName1 = append(userName1, names[0])
	}
	return userName1
}

func validateUserInput(userName1 string, userName2 string, email string, usuarios int, nombresdeUsuarios string, usuariosqueQuedan int) (bool, bool, bool, bool) {
	isValidName := len(userName1) >= 3 && len(userName2) >= 3
	isValidEmail := strings.Contains(email, "@")
	isValidUsuarios := usuarios > 0 && usuarios <= usuariosqueQuedan
	isValidNombresdeUsuarios := len(nombresdeUsuarios) >= 5
	return isValidName, isValidEmail, isValidUsuarios, isValidNombresdeUsuarios
	//tipo := "Usuario publico" OR "Usuario privado"
	//isValidType := tipo == "Usuario publico" || tipo == "Usuario privado"
}
