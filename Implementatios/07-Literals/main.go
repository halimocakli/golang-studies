package main

import (
	"fmt"
)

func main() {
	var intLiteral int = 42
	fmt.Printf("Integer Literal: %d\n", intLiteral)

	var floatLiteral float64 = 3.14159
	fmt.Printf("Floating Point Literal: %f\n", floatLiteral)

	var complexLiteral1 complex128 = 3 + 4i
	var complexLiteral2 complex128 = (2.2 + 1.1i)
	fmt.Printf("Complex Literal 1: %f\n", complexLiteral1)
	fmt.Printf("Complex Literal 2: %f\n", complexLiteral2)

	var boolTrue bool = true
	var boolFalse bool = false
	fmt.Printf("Boolean Literal (true): %t\n", boolTrue)
	fmt.Printf("Boolean Literal (false): %t\n", boolFalse)

	var runeLiteral1 rune = 'A'
	var runeLiteral2 rune = '\n'
	var runeLiteral3 rune = '\u4F60'
	var runeLiteral4 rune = '\x41'
	fmt.Printf("Rune Literal 1: %c\n", runeLiteral1)
	fmt.Printf("Rune Literal 2: %q\n", runeLiteral2)
	fmt.Printf("Rune Literal 3 (Unicode 你): %c\n", runeLiteral3)
	fmt.Printf("Rune Literal 4 (Hexadecimal 'A'): %c\n", runeLiteral4)

	fmt.Println("Escape Sequences:")
	fmt.Println("Alert (Bell): \a")
	fmt.Println("Backspace: ABC\bD")
	fmt.Println("Form Feed: A\fB")
	fmt.Println("New Line: First Line\nSecond Line")
	fmt.Println("Carriage Return: ABC\r123")
	fmt.Println("Horizontal Tab: A\tB")
	fmt.Println("Vertical Tab: A\vB")
	fmt.Println("Backslash: \\")
	fmt.Println("Single Quote: '")
	fmt.Println("Double Quote: \"")

	var stringLiteral1 string = "Hello, Go!"
	var stringLiteral2 string = `This is a raw string literal. \n No escape sequences work here.`
	fmt.Printf("String Literal 1: %s\n", stringLiteral1)
	fmt.Printf("String Literal 2: %s\n", stringLiteral2)

	var hexRune rune = '\x41'
	var octalRune rune = '\101'
	var unicodeRune rune = '\u0041'
	fmt.Printf("Hexadecimal Rune: %c\n", hexRune)
	fmt.Printf("Octal Rune: %c\n", octalRune)
	fmt.Printf("Unicode Rune: %c\n", unicodeRune)
}
