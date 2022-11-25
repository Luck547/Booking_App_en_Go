package main

import (
	"fmt"
)

func main() {
	var appName = "Nombres de usuario para STTDR."
	const cantidaddeUsuariosDisponibles uint = 50
	var usuariosqueQuedan uint = 50

	fmt.Printf("Los usuarios son %T, los que quedan son %T, el nombre de esta app es %T\n", cantidaddeUsuariosDisponibles, cantidaddeUsuariosDisponibles, appName)

	fmt.Printf("Bienvenido a la app de reserva de %v \n", appName)
	fmt.Printf("Tenemos un total de %v usuarios disponibles, y van quedando solo %v libres.\n", cantidaddeUsuariosDisponibles, usuariosqueQuedan)
	fmt.Println("Reserva tu nombre de usuario aqui.")

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

	usuariosqueQuedan = usuariosqueQuedan - cantidaddeUsuariosDisponibles
	fmt.Printf("Gracias %v %v por reservar %v Nombres de usuario. El primero es %v. Te llagara un correo de confirmacion a %v\n", userName1, userName2, usuarios, nombresdeUsuarios, email)

	fmt.Printf("Solo quedan %v nombres de usuario en %v", usuariosqueQuedan, appName)

}
