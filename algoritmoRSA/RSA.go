/*
° Algoritmo RSA

° Autor: @Hackerman , @Valeria , @Gabriela

° Fecha: 31/03/2023

° Descripción: Este programa implementa el algoritmo RSA para cifrar y descifrar mensajes.

° Uso: go run algoritmoRSA.go

° Total de costo operaciones elementales : 57
*/
package algoritmoRSA

import (
	"bufio"
	"crypto/rand"
	"crypto/rsa"
	"encoding/base64"
	"os"

	"github.com/fatih/color"
)

// Funcion para cifrar un mensaje con RSA #Operaciones elementales totales: 14
func cifrarRSA(mensaje string, clavePublica rsa.PublicKey) string { // Declaracion, Invocacion: 3
	// Convertir mensaje a bytes de 0 a 255
	mensajeBytes := []byte(mensaje) // Asignacion, Declaracion, Invocacion, Acesso a estructura: 4
	// Cifrar mensaje
	cifradoBytes, _ := rsa.EncryptPKCS1v15(rand.Reader, &clavePublica, mensajeBytes) // Declaracion, Invocacion, Asignacion: 3
	// Convertir mensaje cifrado a base64
	cifradoBase64 := base64.StdEncoding.EncodeToString(cifradoBytes) // Declaracion, Invocacion, Asignacion: 3

	//fmt.Println("Mensaje: ", mensajeBytes)
	//fmt.Println("Cifrado: ", cifradoBytes)

	// Retornar mensaje cifrado
	return cifradoBase64 // Asignacion: 1
}

// Funcion para descifrar un mensaje con RSA #Operaciones elementales totales: 13
func descifrarRSA(mensaje string, clavePrivada rsa.PrivateKey) string { // Declaracion, Invocacion: 3
	// Convertir mensaje a bytes
	mensajeBytes, _ := base64.StdEncoding.DecodeString(mensaje) // Declaracion, Invocacion, Asignacion: 3
	// Descifrar mensaje
	descifradoBytes, _ := rsa.DecryptPKCS1v15(rand.Reader, &clavePrivada, mensajeBytes) // Declaracion, Invocacion, Asignacion: 3
	// Convertir mensaje descifrado a string
	descifradoString := string(descifradoBytes) // Declaracion, Invocacion, Asignacion: 3
	// Retornar mensaje descifrado
	return descifradoString // Asignacion: 1
}

// Funcion para mostrar la interfaz del algoritmo RSA #Operaciones elementales totales: 30
func InterfaceRSA() {
	// Salto de línea
	println() // invocacion: 1

	// Titulo
	color.Magenta("********Algoritmo RSA********") // invocacion: 1

	// Pedir mensaje al usuario
	color.Cyan("Ingresa el mensaje a cifrar: ") // invocacion: 1
	scanner := bufio.NewScanner(os.Stdin)       // Se crea un scanner para leer el mensaje  // Declaracion, Invocacion, Asignacion: 3
	scanner.Scan()                              // Se lee el mensaje // invocacion: 1
	mensaje := scanner.Text()                   // Se guarda el mensaje en la variable word // Declaracion, Invocacion, Asignacion: 3
	println()                                   // Salto de línea // invocacion: 1

	// Generar claves RSA
	clavePrivada, _ := rsa.GenerateKey(rand.Reader, 2048) // Declaracion, Invocacion, Asignacion: 3
	clavePublica := clavePrivada.PublicKey                // Declaracion,Asignacon,Acceso a estructura: 3

	// Cifrar mensaje
	mensajeCifrado := cifrarRSA(mensaje, clavePublica) // Declaracion, Invocacion, asignacion: 3
	// Descifrar mensaje
	mensajeDescifrado := descifrarRSA(mensajeCifrado, *clavePrivada) // Declaracion, Invocacion, asignacion: 3

	// Resultados
	color.Magenta("\n ***RESULTADOS*** \n") // invocacion: 1
	// Imprimir resultados
	println("Mensaje Original:", mensaje)             // invocacion: 1
	println()                                         // Salto de línea
	println("Mensaje Cifrado:", mensajeCifrado)       // invocacion: 1
	println()                                         // Salto de línea
	println("Mensaje Descifrado:", mensajeDescifrado) // invocacion: 1
} // <- Eliminar la coma y cerrar la declaración de función
