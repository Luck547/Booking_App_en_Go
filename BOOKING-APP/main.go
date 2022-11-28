package main

import "fmt"

const cantidaddeUsuariosDisponibles uint = 50

var appName = "Nombres de usuario para STTDR."
var UsuariosqueQuedan uint = 50
var bookings = make([]UserData, 1)
var nombresdeUsuarios []string

type UserData struct {
	FirstName        string
	LastName         string
	Email            string
	NumbersOfTickets uint
}

func main() {

	greetUsers()

	for {
		var FirstName string
		var LastName string
		var Email string
		var NumbersOfTickets uint
		var nombresdeUsuarios string
		// Solicita nombre de usuario al usuario
		// printUserName1(bookings)
		fmt.Println("tipea tu primer nombre")
		fmt.Scan(&FirstName)
		fmt.Println("tipea tu apellido")
		fmt.Scan(&LastName)
		fmt.Println("tipea tu email")
		fmt.Scan(&Email)
		fmt.Println("Cuantos usuarios reservas?")
		fmt.Scan(&NumbersOfTickets)
		fmt.Println("Ingresa tu nombre de usuario a reservar")
		fmt.Scan(&nombresdeUsuarios)

		var isValidName, isValidEmail, isValidUsuarios, isValidNombresdeUsuarios = ValidateUserInput(FirstName, LastName, Email, NumbersOfTickets, nombresdeUsuarios, UsuariosqueQuedan)

		if isValidName && isValidEmail && isValidUsuarios && isValidNombresdeUsuarios {
			//bookUserName(UsuariosqueQuedan, NumbersOfTickets, bookings, FirstName, Lastname, Email, appName)
			// Call a function print userName1
			//printUserName1(bookings)

			fmt.Printf("These are all the bookings: %v\n", FirstName)
			if UsuariosqueQuedan == 0 {
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

func greetUsers() {
	fmt.Printf("Los usuarios son %T, los que quedan son %T, el nombre de esta app es %T\n", cantidaddeUsuariosDisponibles, cantidaddeUsuariosDisponibles, appName)
	fmt.Printf("Bienvenido a la app de reserva de %v\n", appName)
	fmt.Printf("Tenemos un total de %v usuarios disponibles, y van quedando solo %v libres.\n", cantidaddeUsuariosDisponibles, UsuariosqueQuedan)
	fmt.Printf("Bienvenido a la reserva de %v", appName)
	fmt.Println("Reserva tu nombre de usuario aqui.")
}

func printUserName1(bookings []string) []string {
	var FirstName []string
	for _, booking := range bookings {
		FirstName = append(FirstName, booking)
	}
	return FirstName
}

func bookUserName(bookings []string, FirstName string, Lastname string, Email string, appName string, UsuariosqueQuedan uint, NumbersOfTickets uint) {
	UsuariosqueQuedan = UsuariosqueQuedan - NumbersOfTickets

	// Create a map for a user
	var userData = UserData{
		FirstName:         FirstName,
		LastName:          Lastname,
		Email:             Email,
		UsuariosqueQuedan: UsuariosqueQuedan,
	}
	bookings = append(bookings, userData)
	fmt.Printf("Gracias %v %v por reservar %v Nombres de usuario.\nEl primero es %v. Te llagara un correo de confirmacion a %v\n", FirstName, Lastname, NumbersOfTickets, nombresdeUsuarios, Email)
	fmt.Printf("Solo quedan %v nombres de usuario en %v,. Gracias!\n", UsuariosqueQuedan, appName)
}

func sendTicket(NumbersOfTickets uint, FirstName string, Lastname string, Email string) {
	var ticket, _ = fmt.Printf(
		"Sending %v v% v% tickests\n",
		NumbersOfTickets,
		FirstName,
		Lastname,
	)
	fmt.Println()
	fmt.Printf("Sending tickets:\n %v to email address %v\n", ticket, Email)
	fmt.Println()
}
