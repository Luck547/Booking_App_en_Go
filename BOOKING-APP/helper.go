package main

import "strings"

func ValidateUserInput(FirstName string, Lastname string, Email string, NumbersOfTickets uint, nombresdeUsuarios string, usuariosqueQuedan uint) (bool, bool, bool, bool) {
	isValidName := len(FirstName) >= 3 && len(Lastname) >= 3
	isValidEmail := strings.Contains(Email, "@")
	isValidUsuarios := NumbersOfTickets > 0 && NumbersOfTickets <= usuariosqueQuedan
	isValidNombresdeUsuarios := len(nombresdeUsuarios) >= 5
	return isValidName, isValidEmail, isValidUsuarios, isValidNombresdeUsuarios
}
