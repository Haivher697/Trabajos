package main

import (
	"fmt"
)

func main() {

	var nom string
	metros := 0.00
	coversion := 0.00

	fmt.Println("Ingrese su Nombre: ")
	fmt.Scanln(&nom)

	fmt.Printf("Hola " + nom + " Convierte tus Metros a Kilometros, ")
	fmt.Printf("Ingresa Metros:")
	fmt.Scanln(&metros)

	coversion = metros / 1000

	fmt.Println(nom, "Su Conversión es: ", coversion, "KM hasta luego :)")
}
