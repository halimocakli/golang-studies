package main

import (
	"fmt"
)

func main() {
	// 🟢 Integer Literals (tamsayı sabitleri)
	var intLiteral int = 42
	fmt.Printf("Integer Literal: %d\n", intLiteral)

	// 🟢 Floating Point Literals (gerçek sayı sabitleri)
	var floatLiteral float64 = 3.14159
	fmt.Printf("Floating Point Literal: %f\n", floatLiteral)

	// 🟢 Complex Literals (karmaşık sayılar)
	var complexLiteral1 complex128 = 3 + 4i
	var complexLiteral2 complex128 = (2.2 + 1.1i) // Parantezli biçim
	fmt.Printf("Complex Literal 1: %f\n", complexLiteral1)
	fmt.Printf("Complex Literal 2: %f\n", complexLiteral2)

	// 🟢 Boolean Literals (bool sabitleri)
	var boolTrue bool = true
	var boolFalse bool = false
	fmt.Printf("Boolean Literal (true): %t\n", boolTrue)
	fmt.Printf("Boolean Literal (false): %t\n", boolFalse)

	// 🟢 Rune Literals (rune sabitleri)
	var runeLiteral1 rune = 'A'
	var runeLiteral2 rune = '\n'     // Newline karakteri
	var runeLiteral3 rune = '\u4F60' // Unicode karakteri 你 (ni)
	var runeLiteral4 rune = '\x41'   // ASCII 'A' hexadecimal
	fmt.Printf("Rune Literal 1: %c\n", runeLiteral1)
	fmt.Printf("Rune Literal 2: %q\n", runeLiteral2)
	fmt.Printf("Rune Literal 3 (Unicode 你): %c\n", runeLiteral3)
	fmt.Printf("Rune Literal 4 (Hexadecimal 'A'): %c\n", runeLiteral4)

	// 🟢 Escape Sequences (kaçış dizileri)
	fmt.Println("Escape Sequences:")
	fmt.Println("Alert (Bell): \a")                  // Bell (çalışma zamanında duyulabilir)
	fmt.Println("Backspace: ABC\bD")                 // Backspace (C karakterini siler)
	fmt.Println("Form Feed: A\fB")                   // Form feed (bazı terminallerde etkisizdir)
	fmt.Println("New Line: First Line\nSecond Line") // Satır atlama (LF)
	fmt.Println("Carriage Return: ABC\r123")         // İmleci satır başına taşır
	fmt.Println("Horizontal Tab: A\tB")              // Yatay tab (TAB)
	fmt.Println("Vertical Tab: A\vB")                // Dikey tab (çok nadiren kullanılır)
	fmt.Println("Backslash: \\")                     // Backslash (\)
	fmt.Println("Single Quote: '")                   // Tek tırnak (rune için kullanılır)
	fmt.Println("Double Quote: \"")                  // Çift tırnak (string için)

	// 🟢 String Literals (dizgi sabitleri)
	var stringLiteral1 string = "Hello, Go!"
	var stringLiteral2 string = `This is a raw string literal. \n No escape sequences work here.`
	fmt.Printf("String Literal 1: %s\n", stringLiteral1)
	fmt.Printf("String Literal 2: %s\n", stringLiteral2)

	// 🟢 Hexadecimal, Octal ve Unicode Gösterimleri
	var hexRune rune = '\x41'       // Hexadecimal ('A')
	var octalRune rune = '\101'     // Octal ('A')
	var unicodeRune rune = '\u0041' // Unicode ('A')
	fmt.Printf("Hexadecimal Rune: %c\n", hexRune)
	fmt.Printf("Octal Rune: %c\n", octalRune)
	fmt.Printf("Unicode Rune: %c\n", unicodeRune)
}
