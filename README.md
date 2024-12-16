
# Golang Operatörleri İçin Notlar

Bu not defteri, Golang operatörleri ile ilgili temel kavramları, operatörlerin sınıflandırılmasını, öncelik ve ilişki kurallarını ve yazılan kodla ilgili tüm önerileri içerir. Notlar, hem öğretici hem de pratik bilgileri kapsar.

## 1. **Operatörlerin Tanımı ve Sınıflandırılması**
- **Operatör (Operator)**: Bir işleme yol açan ve işlem sonucunda bir değer üreten sembollerdir.
- **Operand**: Bu operatörlerle işlem yapan ifadelerdir.
- **Sınıflandırma:**
  - **Aritmetik Operatörler (Arithmetic Operators)**: `+`, `-`, `*`, `/`, `%`
  - **Karşılaştırma Operatörleri (Comparison Operators)**: `==`, `!=`, `<`, `>`, `<=`, `>=`
  - **Mantıksal Operatörler (Logical Operators)**: `&&`, `||`, `!`
  - **Bitsel Operatörler (Bitwise Operators)**: `&`, `|`, `^`, `<<`, `>>`
  - **Özel Amaçlı Operatörler (Special Purpose Operators)**: `&`, `*`, `=`, `:=`

## 2. **Operatör Önceliği ve İşlem Sırası**
- Operatörler, öncelik düzeylerine göre işlenmektedir.
- **Aritmetik operatörlerin Öncelik Sıralaması**: `*`, `/`, `%` > `+`, `-`
- **Sağdan sola (right associative)** veya **soldan sağa (left associative)** olabilir. 
  - Aritmetik operatörler genellikle soldan sağa (left associative) çalışır.
  - Atama operatörleri (`=`) sağdan sola (right associative) çalışır.

## 3. **Golang Operatörleri Hakkında Detaylı Bilgiler**

- **Modulo Operatörü (%)**:
  - Yalnızca `int` türleri ile kullanılabilir.
  - `float64` veya `complex` türleriyle kullanılamaz, aksi halde derleme hatası alınır.
  - Matematiksel mod işleminden farklı olarak, Go'da negatif sayıların mod işlemi, operandın işaretiyle ilişkilidir.
  
- **Aritmetik Operatörler** (`+`, `-`, `*`, `/`):
  - Hem `int` hem de `float64` değerleriyle kullanılabilir.
  - `complex` değerleriyle de kullanılabilir. 
  - `/` operatörü bölme işleminde, eğer bölen 0 ise hata oluşturur.

- **Mantıksal Operatörler** (`&&`, `||`, `!`):
  - Mantıksal işlemler için kullanılır ve yalnızca `bool` türleriyle çalışır.
  - Kısa devre mantığı (short-circuiting) kullanılır. 

- **Karşılaştırma Operatörleri** (`==`, `!=`, `<`, `>`, `<=`, `>=`):
  - Bütün sayısal türlerle ve `string` ile kullanılabilir.

- **Bitsel Operatörler** (`&`, `|`, `^`, `<<`, `>>`):
  - Yalnızca `int` türleriyle kullanılabilir.

- **Atama Operatörleri** (`=`, `+=`, `-=`, `*=`, `/=`, `%=`):
  - Değişkenlere değer atamak veya mevcut değeri güncellemek için kullanılır.

## 4. **Final Versiyon Kod**

```go
package main

import (
	"fmt"
)

func main() {
	performIntegerOperations()
	performComplexOperations()
	showInvalidOperations()
}

func performIntegerOperations() {
	var num1, num2 int
	fmt.Print("Input two integers (e.g., 3 5): ")
	if _, err := fmt.Scan(&num1, &num2); err != nil {
		fmt.Println("Invalid input. Please enter valid integers.")
		return
	}

	fmt.Printf("
Integer Operations:
")
	fmt.Printf("%d + %d = %d
", num1, num2, num1+num2)
	fmt.Printf("%d - %d = %d
", num1, num2, num1-num2)
	fmt.Printf("%d * %d = %d
", num1, num2, num1*num2)

	if num2 != 0 {
		fmt.Printf("%d / %d = %d
", num1, num2, num1/num2)
	} else {
		fmt.Printf("Division by zero is not allowed.
")
	}

	fmt.Printf("%d %% %d = %d
", num1, num2, num1%num2)
}

func performComplexOperations() {
	var complex1, complex2 complex128
	fmt.Print("
Input two complex numbers (e.g., 3+4i 5-2i): ")
	if _, err := fmt.Scan(&complex1, &complex2); err != nil {
		fmt.Println("Invalid input. Please enter valid complex numbers.")
		return
	}
	fmt.Printf("
Complex Operations:
")
	fmt.Printf("%v + %v = %v
", complex1, complex2, complex1+complex2)
	fmt.Printf("%v - %v = %v
", complex1, complex2, complex1-complex2)
}

func showInvalidOperations() {
	fmt.Println("
Reminder: %% [Modulo Operator] cannot be used with Floats or Complex numbers.")
}
```

## 5. **Sonuç**
Bu not defteri, Golang operatörleri hakkında önemli bilgiler ve kod için iyileştirme önerileri içermektedir. Tüm bu öneriler, kodun okunabilirliğini, modülerleştirilmesini ve hata kontrolünü geliştirir.
