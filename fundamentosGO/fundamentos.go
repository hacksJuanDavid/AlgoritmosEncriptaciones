package fundamentosgo

import (
	"fmt"
	"strings"
)

// function content tipes variables in go
func contentVars() {
	//Title
	println("Variables")
	// string
	var name string = "John"
	// int
	var age int = 20
	// float
	var height float32 = 1.75
	// bool
	var isMarried bool = false
	// array
	var fruits = []string{"apple", "banana", "grape"}
	// slice
	var fruitsSlice = []string{"apple", "banana", "grape"}
	// map
	var fruitsMap = map[string]string{"apple": "apple", "banana": "banana", "grape": "grape"}

	// pointer
	var pointer *string = &name

	println(name)
	println(age)
	println(height)
	println(isMarried)
	println(strings.Join(fruits, ","))
	println(strings.Join(fruitsSlice, ","))

	for key, value := range fruitsMap {
		println(key, ":", value)
	}
	println(pointer, *pointer)

	// Jump to line
	println()
}

// Function contentOperatorsArithmetic
func contentOperatorsArithmetic() {
	//Title
	println("Arithmetic Operators")
	var a int = 10
	var b int = 3

	println("Suma:", a+b)
	println("Resta:", a-b)
	println("Multiplicacion:", a*b)
	println("Division:", a/b)
	println("Modulo:", a%b)
	// Jump to line
	println()
}

// Function contentOperatorsComparison
func contentOperatorsComparison() {
	//Title
	println("Comparison Operators")
	var a int = 10
	var b int = 3

	println("Iguales:", a == b)
	println("Diferentes:", a != b)
	println("Mayor que:", a > b)
	println("Menor que:", a < b)
	println("Mayor o igual que:", a >= b)
	println("Menor o igual que:", a <= b)
	// Jump to line
	println()
}

// Function contentOperatorsLogical
func contentOperatorsLogical() {
	//Title
	println("Logical Operators")
	var a bool = true
	var b bool = false

	println("AND:", a && b)
	println("OR:", a || b)
	println("NOT:", !a)
	// Jump to line
	println()

}

// Function contentOperatorsBitwise
func contentOperatorsBitwise() {
	//Title
	println("Bitwise Operators")
	var a int = 10
	var b int = 3

	println("AND:", a&b)
	println("OR:", a|b)
	println("XOR:", a^b)
	println("NOT:", ^a)
	println("Left Shift:", a<<b)
	println("Right Shift:", a>>b)
	// Jump to line
	println()

}

// Function contentOperatorsAssignment
func contentOperatorsAssignment() {
	//Title
	println("Assignment Operators")
	var a int = 10
	var b int = 3

	a += b
	println("Suma:", a)

	a -= b
	println("Resta:", a)

	a *= b
	println("Multiplicacion:", a)

	a /= b
	println("Division:", a)

	a %= b
	println("Modulo:", a)

	a &= b
	println("AND:", a)

	a |= b
	println("OR:", a)

	a ^= b
	println("XOR:", a)

	a &^= b
	println("AND NOT:", a)

	a <<= b
	println("Left Shift:", a)

	a >>= b
	println("Right Shift:", a)
	// Jump to line
	println()

}

// Function numbers integers
func numbersIntegers() {
	//Title
	println("Numbers Integers")

	// int8 -128 to 127
	var int8 int8 = 127

	// int16 -32768 to 32767
	var int16 int16 = 32767

	// int32 -2147483648 to 2147483647
	var int32 int32 = 2147483647

	// int64 -9223372036854775808 to 9223372036854775807
	var int64 int64 = 9223372036854775807

	// int -9223372036854775808 to 9223372036854775807
	var int int = 9223372036854775807

	// uint8 0 to 255
	var uint8 uint8 = 255

	// uint16 0 to 65535
	var uint16 uint16 = 65535

	// uint32 0 to 4294967295
	var uint32 uint32 = 4294967295

	// uint64 0 to 18446744073709551615
	var uint64 uint64 = 18446744073709551615

	// uint 0 to 18446744073709551615
	var uint uint = 18446744073709551615

	// byte 0 to 255
	var byte byte = 255

	// rune -2147483648 to 2147483647
	var rune rune = 2147483647

	// uintptr 0 to 18446744073709551615
	var uintptr uintptr = 18446744073709551615

	println("int8:", int8)
	println("int16:", int16)
	println("int32:", int32)
	println("int64:", int64)
	println("int:", int)
	println("uint8:", uint8)
	println("uint16:", uint16)
	println("uint32:", uint32)
	println("uint64:", uint64)
	println("uint:", uint)
	println("byte:", byte)
	println("rune:", rune)
	println("uintptr:", uintptr)
	// Jump to line
	println()
}

// Function numbers floats
func numbersFloats() {
	//Title
	println("Numbers Floats")

	// float32 1.18E-38 to 3.40E+38
	var float32 float32 = 3.40e+38

	// float64 2.23E-308 to 1.80E+308
	var float64 float64 = 0.0000000

	println("float32:", float32)
	println("float64:", float64)
	// Jump to line
	println()

}

// Function strings
func varsStrings() {
	//Title
	println("Strings")

	// string
	var string string = "Hello World"

	// string
	var string2 = ""

	println("string:", string)
	println("string2:", string2)
	// Jump to line
	println()

}

// Function bools
func varsBools() {
	//Title
	println("Bools")

	// bool
	var bool bool = true

	// bool
	var bool2 = false

	println("bool:", bool)
	println("bool2:", bool2)
	// Jump to line
	println()
}

// Functions complex numbers
func numbersComplex() {
	//Title
	println("Numbers Complex")

	// complex64
	var complex64 complex64 = 1 + 2i

	// complex128
	var complex128 complex128 = 1 + 2i

	println("complex64:", complex64)
	println("complex128:", complex128)
	// Jump to line
	println()
}

// Functions return 2 values
func return2values() (int, int) {
	//Title
	println("Return 2 values")
	return 1, 2
}

// Function types ciclics for
func typesCiclicsFor() {
	//Title
	println("Types ciclics for")
	// for
	for i := 0; i < 10; i++ {
		println("For 1:", i)
	}

	// for 2
	var i int = 0
	for i < 10 {
		println("for 2:", i)
		i++
	}

	// for 3
	var j int = 0
	for {
		println("for 3:", j)
		j++
		if j == 10 {
			break
		}
	}

	// Jump to line
	println()
}

// Functions using if else and switch
func contentIfElseSwitch() {
	//Title
	println("If else and switch")
	// if else
	var a int = 10
	var b int = 3
	if a > b {
		println("a > b")
	} else {
		println("a < b")
	}

	// switch
	switch a {
	case 1:
		println("a = 1")
	case 2:
		println("a = 2")
	case 3:
		println("a = 3")
	default:
		println("a != 1, 2, 3")
	}

	// switch 2 par and impar
	switch a % 2 {
	case 0:
		println("a is even")
	case 1:
		println("a is odd")
	}

	// switch 3
	switch {
	case a > 0:
		println("a is positive")
	case a < 0:
		println("a is negative")
	default:
		println("a is zero")
	}

	// Jump to line
	println()
}

// Function use keywords
func contentKeywords() {
	//Title
	println("Keywords")

	// Defer is used to ensure that a function call is performed later in a program's execution, usually for purposes of cleanup. defer is often used where e.g. ensure and finally would be used in other languages.
	defer println("defer")

	// Break is used to immediately exit a for, switch, or select statement.
	for i := 0; i < 10; i++ {
		if i == 5 {
			break
		}
		println("break:", i)
	}

	// Continue is used to skip the current iteration of the for loop.
	for i := 0; i < 10; i++ {
		if i == 5 {
			continue
		}
		println("continue:", i)
	}

	// Fallthrough is used to execute the next case in a switch statement, even if the current case does not match.
	switch i := 0; i {
	case 0:
		println("fallthrough 0")
		fallthrough
	case 1:
		println("fallthrough 1")
		fallthrough // fallthrough is not needed here
	default:
		println("fallthrough default")
	}

	// Jump to line
	println()

}

// Functions arrays and slices
func contentArraysAndSlices() {
	//Title
	println("Arrays and slices")

	// Array
	var array [5]int // declare array of 5 integers
	array[0] = 1     // assign value 1 to first element
	array[1] = 2     // assign value 2 to second element
	array[2] = 3
	array[3] = 4
	array[4] = 5

	// print array
	println("array:", array[0], array[1], array[2], array[3], array[4])

	// lentgh array
	println("lentgh array:", len(array))

	// capacity array
	println("capacity array:", cap(array))

	// Slice
	var slice []int          // declare slice of integers
	slice = append(slice, 1) // append value 1 to slice

	// Slice 2
	var slice2 []int = []int{1, 2, 3, 4, 5} // declare and initialize slice of integers
	fmt.Println("slice2:", slice2, "capacity:", cap(slice2), "lentgh:", len(slice2))
	// Methods of slices
	// append
	slice2 = append(slice2, 6)
	fmt.Println("slice2:", slice2, "capacity:", cap(slice2), "lentgh:", len(slice2))

	// append 2
	slice = append(slice2, 7, 8, 9, 10)
	fmt.Println("slice2:", slice2, "capacity:", cap(slice2), "lentgh:", len(slice2))

	// append 3
	slice2 = append(slice2, slice2...)
	fmt.Println("slice2:", slice2, "capacity:", cap(slice2), "lentgh:", len(slice2))

	// copy
	var slice3 []int = make([]int, 5)
	copy(slice3, slice2)

	// Beetwen slices
	var slice4 []int = slice2[1:3]
	fmt.Println("slice4:", slice4, "capacity:", cap(slice4), "lentgh:", len(slice4))

	// Next slice
	var slice5 []int = slice2[1:]
	fmt.Println("slice5:", slice5, "capacity:", cap(slice5), "lentgh:", len(slice5))

	// Previous slice
	var slice6 []int = slice2[:3]
	fmt.Println("slice6:", slice6, "capacity:", cap(slice6), "lentgh:", len(slice6))

	// Jump to line
	println()
}

// Is word palindrom lowercase and uppercase
func isPalindrom(word string) {
	//Title
	println("Is word palindrom lowercase and uppercase")

	word = strings.ToLower(word)

	var wordReverse string
	for i := len(word) - 1; i >= 0; i-- {
		wordReverse += string(word[i])
	}

	if word == wordReverse {
		println("Palindrom")
	} else {
		println("Not palindrom")
	}

	// Jump to line
	println()
}

// Function interface fundamentos
func InterfaceFundamentos() {
	contentVars()
	contentOperatorsArithmetic()
	contentOperatorsComparison()
	contentOperatorsLogical()
	contentOperatorsBitwise()
	contentOperatorsAssignment()
	numbersIntegers()
	numbersFloats()
	varsStrings()
	varsBools()
	numbersComplex()
	println(return2values())
	typesCiclicsFor()
	contentIfElseSwitch()
	contentKeywords()
	contentArraysAndSlices()
	isPalindrom("Ana")
}
