package main

import (
	"fmt"
)

func main() {
	// Tamsayılar
	fmt.Printf("%%d: %d\n", 42)
	fmt.Printf("%%b: %b\n", 42)
	fmt.Printf("%%b: %016b\n", 42)
	fmt.Printf("%%o: %o\n", 42)
	fmt.Printf("%%x: %x\n", 42)
	fmt.Printf("%%x: %08x\n", 42)
	fmt.Printf("%%X: %X\n", 42)
	fmt.Printf("%%c: %c\n", 65)
	fmt.Printf("%%q: %q\n", 65)

	// Gerçek Sayılar
	fmt.Printf("%%f: %f\n", 3.14159)
	fmt.Printf("%%e: %e\n", 1234.5678)
	fmt.Printf("%%g: %g\n", 1234.5678)

	// Dizgeler
	fmt.Printf("%%s: %s\n", "hello")
	fmt.Printf("%%q: %q\n", "hello")
	fmt.Printf("%%x: %x\n", "hello")

	// Pointer
	x := 42
	fmt.Printf("%%p: %p\n", &x)

	// Genel Yer Tutucular
	fmt.Printf("%%v: %v\n", []int{1, 2, 3})
	fmt.Printf("%%+v: %+v\n", struct{ Name string }{"Go"})
	fmt.Printf("%%T: %T\n", x)
}
