package main

import (
	"bufio"
	"os"

	"examples.com/packages/algoritmoCESAR"
	"examples.com/packages/algoritmoRSA"
	"github.com/fatih/color"
)

// Function menu principal
// Función menu principal
func menu() {
	for {
		println()
		color.Magenta("********Menu********")
		color.Yellow("1. Algoritmo Cesar")
		color.Green("2. Algoritmo RSA")
		color.Red("3. Salir")
		println()

		// Pedir opción al usuario
		color.Cyan("Ingresa la opción que deseas: ")
		scanner := bufio.NewScanner(os.Stdin) // Se crea un scanner para leer la opción
		scanner.Scan()                        // Se lee la opción
		opcion := scanner.Text()              // Se guarda la opción en la variable opcion
		println()                             // Salto de línea

		// Opciones
		switch opcion {
		case "1":
			algoritmoCESAR.InterfaceCESAR()
		case "2":
			algoritmoRSA.InterfaceRSA()
		case "3":
			println("Saliendo...")
			return // Sale de la función menu() y termina el programa
		default:
			println("Opción no válida")
		}
	}
}

// function main
func main() {
	// Llamar a la función menu
	menu()
}
