
# Literals and Constants in Go

## 📘 Introduction

Bu dosya, Golang'deki **sabitler (literals)** ve **sabit ifadeler (constant expressions)** kavramlarına odaklanmaktadır. Sabitler, bir programda doğrudan yazılan değerlerdir ve birçok türde olabilirler. Bu dosya, Golang'deki tüm sabit türlerini, kurallarını ve ilgili örnekleri açıklar.

---

## 📚 Literals (Sabit Türleri)

Go'daki sabit türleri aşağıdaki gibidir:

- **Integer Literals (Tam sayı sabitleri)**: Tamsayı değerleri (`42`, `100`, `0b1010` vb.).
- **Floating Point Literals (Gerçek sayı sabitleri)**: Ondalık sayılar (`3.14`, `-0.5` vb.).
- **Complex Literals (Karmaşık sayı sabitleri)**: Karmaşık sayılar (`3+4i`, `(1.1 + 2.2i)` vb.).
- **Boolean Literals (Boolean sabitleri)**: `true` ve `false` sabitleridir.
- **Rune Literals (Rune sabitleri)**: Tek tırnak içinde tek bir karakter (`'a'`, `'中'`, `'
'`, `'你'` vb.).
- **String Literals (String sabitleri)**: Çift tırnakla belirtilen dizgeler (`"hello"`, `"world"` vb.).

---

## 📘 Full Example (Tam Kod Örneği)

Aşağıda Golang'de **sabitler (literals)**, **sabit ifadeler (constant expressions)**, **rune** ve **escape sequence** kullanımına dair eksiksiz bir örnek bulunmaktadır.

```go
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
	fmt.Println("Single Quote: \'")
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
```

---

## 📘 Çıktı (Output)

Programın çıktısı aşağıdaki gibidir:

```
Integer Literal: 42
Floating Point Literal: 3.141590
Complex Literal 1: 3.000000+4.000000i
Complex Literal 2: 2.200000+1.100000i
Boolean Literal (true): true
Boolean Literal (false): false
Rune Literal 1: A
Rune Literal 2: '\n'
Rune Literal 3 (Unicode 你): 你
Rune Literal 4 (Hexadecimal 'A'): A
Escape Sequences:
Alert (Bell): 
Backspace: ABD
Form Feed: A
B
New Line: First Line
Second Line
Carriage Return: 123
Horizontal Tab: A	B
Vertical Tab: A
B
Backslash: \
Single Quote: '
Double Quote: "
String Literal 1: Hello, Go!
String Literal 2: This is a raw string literal. \n No escape sequences work here.
Hexadecimal Rune: A
Octal Rune: A
Unicode Rune: A
```

---

## 🧩 Sonuç

Bu dosya, Go'daki **sabit türleri (literal types)** ve **sabit ifadeler (constant expressions)** konularını kapsayan kapsamlı bir kılavuzdur. Bu örnekler, Go'daki `complex`, `rune`, `string` ve diğer sabit türlerini nasıl kullanacağınızı gösterir. Ayrıca, kaçış dizileri (escape sequences) ile ilgili bilgiler sağlar. Bu konuları anlamak, Go'da daha etkili ve verimli kod yazmanıza olanak tanır.
