package algoritmoCESAR

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"

	"github.com/fatih/color"
)

/*
	° Algoritmo Cesar

	° Autor: @Hackerman , @Valeria , @Gabriela

	° Fecha: 01/04/2023

	° Descripción: Este programa implementa el algoritmo Cesar para cifrar y descifrar mensajes.


	° Requerimientos:
		- Se le debe de pedir al usuario el mensaje que desaa cifrar
		- Se le debe de pedir el n de desplazamiento al usuario
		- Se debe de usar el mismo n para cifrar y descifrar
		- El algotirmo debe de cifrar
		- El algoritmo debe de descifrar
		- Mostrar el mensaje original
		- Mostrar el mensaje cifrado
		- Mostrar el mensaje descifrado
		- Total de costo operaciones elementales : 95
*/

// Funcion cifrar mensaje con el algoritmo Cesar #Operaciones elementales totales: 34
func encrypt(message string, n int) string { // Declare : 2
	// Convertimos el mensaje a minúsculas
	message = strings.ToLower(message) // Invocacion, asignacion, Declaracion : 3

	// Mapa de letras con acentos y diéresis y su correspondiente letra sin acentos ni diéresis
	var lettersMap = map[string]string{ // Asignacion, Declaracion: 2
		"á": "a",
		"é": "e",
		"í": "i",
		"ó": "o",
		"ú": "u",
		"ä": "a",
		"ë": "e",
		"ï": "i",
		"ö": "o",
		"ü": "u",
		"ñ": "n",
	}

	// Sustituimos las letras con acentos y diéresis por sus correspondientes letras sin acentos ni diéresis
	for k, v := range lettersMap { // Invoacion, Asignacion , Declaracion: 3
		message = strings.Replace(message, k, v, -1) // Invocacion, asignacion, Aritmetica: 3
	}

	// Inicializamos la variable para el mensaje cifrado
	var cryptedMessage string // Declaracion: 1

	// Proceso de cifrado del mensaje
	for _, r := range message { // Invoacion, Asignacion , Declaracion: 3
		// Verificamos si el carácter actual es una letra del alfabeto
		/*
			En esta parte del código se verifica si el carácter actual es una letra del alfabeto,
			utilizando la función unicode.IsLetter(r) de la librería unicode, que devuelve true si
			el carácter r es una letra y false en caso contrario.
		*/
		if unicode.IsLetter(r) { // Invocacion, asignacion: 2
			// Se busca la posicion de la letra en el alfabeto
			/*
				Luego, si el carácter es una letra del alfabeto, se busca su posición en el alfabeto restando
				 la posición en Unicode del carácter 'a' ('a' tiene el valor Unicode 97, 'b' el 98 y así sucesivamente).
				 El resultado de la resta es la posición de la letra en el alfabeto, donde 'a' es la posición 0, 'b' la 1, etc.
			*/
			var position int = int(r) - 'a' // a es la posición 0 en el alfabeto , invocacion, asignacion, declaracion, aritmetica: 4

			// Se cifra la letra
			/*
				El proceso de cifrado se puede expresar como una fórmula matemática:
				C(i) = (P(i) + n) mod 26
				donde:
				    C(i) es el carácter cifrado en la posición "i".
				    P(i) es el carácter original en la posición "i".
				    n es el número de posiciones que se desplaza cada letra en el alfabeto (la clave del cifrado).
				    mod 26 es una operación matemática que asegura que el resultado final esté dentro de los límites del alfabeto.
					Finalmente, se crea un nuevo rune con el carácter cifrado y se agrega al mensaje cifrado.
			*/
			var letterCrypted rune = rune((position+n)%26 + 'a') // invocacion, asignacion, declaracion, aritmetica: 6

			// Agregamos la letra cifrada al mensaje cifrado
			cryptedMessage += string(letterCrypted) // invocacion, asignacion: 2
		} else {
			// Si el carácter actual no es una letra del alfabeto, lo agregamos directamente al mensaje cifrado
			cryptedMessage += string(r) // invocacion, asignacion: 2
		}
	}

	// Se retorna el mensaje cifrado
	return cryptedMessage // asignacion: 1
}

// Funcion descifrar mensaje con el algoritmo Cesar #Operaciones elementales totales: 28
func decipher(message string, n int) string { // Declare : 2
	// Convertimos el mensaje a minúsculas
	message = strings.ToLower(message) // Invocacion, asignacion, Declaracion : 3

	// Inicializamos la variable para el mensaje descifrado
	var decryptedMessage string // Declaracion: 1

	// Proceso de descifrado del mensaje
	for _, r := range message { // Asignacion, Declaracion,Invoacion: 3
		// Verificamos si el carácter actual es una letra del alfabeto
		if unicode.IsLetter(r) { // Invocacion, asignacion: 2
			// Se busca la posicion de la letra en el alfabeto
			var position int = int(r) - 'a' //invocacion, asignacion, declaracion, aritmetica: 4

			// Se descifra la letra
			/*
				El proceso de descifrado se realiza de manera similar, pero en sentido contrario:
				P(i) = (C(i) - n) mod 26 + 26) mod 26 + 'a')
			*/
			var letterDecrypted rune = rune(((position-n)%26+26)%26 + 'a') // invocacion, asignacion, declaracion, aritmetica: 8
			// Agregamos la letra
			decryptedMessage += string(letterDecrypted) // invocacion, asignacion: 2
		} else {
			// Si el carácter actual no es una letra del alfabeto, lo agregamos directamente al mensaje descifrado
			decryptedMessage += string(r) // invocacion, asignacion: 2
		}
	}

	// Se retorna el mensaje descifrado
	return decryptedMessage // asignacion: 1
}

// Funcion principal interfaz #Operaciones elementales totales: 33
func InterfaceCESAR() {
	// Jump to line
	println() // Invocacion: 1

	// Title
	color.Magenta("*****Algoritmo Cesar*****") // Invocacion: 1
	println()                                  // Invocacion: 1

	// Se le pide al usuario el mensaje que desea cifrar
	color.Cyan("Ingrese el mensaje que desea cifrar:") // Invocacion: 1
	scanner := bufio.NewScanner(os.Stdin)              // Se crea un scanner para leer el mensaje, declaracion,invocacion,asignacion: 3
	scanner.Scan()                                     // Se lee el mensaje , invocacion, asignacion: 2
	message := scanner.Text()                          // Se guarda el mensaje en la variable word , invocacion, asignacion,delcaracion: 3
	println()                                          // Salto de linea , invocacion: 1
	// Se le pide al usuario el n de desplazamiento
	color.Cyan("Ingrese el n de desplazamiento:") // Invocacion: 1
	var n int                                     // Se crea la variable n , declaracion: 1
	fmt.Scanln(&n, "\n")                          // Se lee el n , invocacion, asignacion: 2
	println()                                     // Salto de linea , invocacion: 1

	// Se cifra el mensaje
	// Example "El veloz murciélago hindú comía feliz cardillo y kiwi. La cigüeña tocaba el saxofón detrás del palenque de paja"
	var messageEncrypt string = encrypt(message, n)          // Se cifra el mensaje , invocacion, asignacion, declaracion: 3
	var messageDecipher string = decipher(messageEncrypt, n) // Se descifra el mensaje , invocacion, asignacion, declaracion: 3

	// Resultados
	color.Magenta("\n ***RESULTADOS*** \n") // Invocacion: 1
	// Se muestra el mensaje
	fmt.Println("Mensaje Original:", message) // Invocacion,asignacion: 2
	println()                                 // Salto de linea , invocacion: 1
	// Se muestra el mensaje cifrado
	fmt.Println("Mensaje cifrado:", messageEncrypt) // Invocacion,asignacion: 2
	println()                                       // Salto de linea , invocacion: 1
	// Se muestra el mensaje descifrado
	fmt.Println("Mensaje descifrado:", messageDecipher) // Invocacion,asignacion: 2
}
