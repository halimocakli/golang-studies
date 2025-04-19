package main

import (
	"fmt"
)

func main() {
	performIntegerOperations() // Tam sayı işlemlerini gerçekleştirir (Integer operations [Tam sayı işlemleri]).
	performComplexOperations() // Karmaşık sayı işlemlerini gerçekleştirir (Complex operations [Karmaşık işlemler]).
	showInvalidOperations()    // Geçersiz işlem uyarısı gösterir (Invalid operation reminder [Geçersiz işlem hatırlatması]).
}

func performIntegerOperations() {
	// İki tam sayı (integer numbers [tam sayılar]) için değişkenlerin tanımlanması.
	var num1, num2 int

	// Kullanıcıdan iki tam sayı girişi istenir (Input reading [Girdi alma]).
	fmt.Print("Input two integers (e.g., 3 5): ")
	// fmt.Scan ile kullanıcı girdisi alınır; hata kontrolü yapılarak doğru formatta veri girilmesi sağlanır.
	if _, err := fmt.Scan(&num1, &num2); err != nil {
		// Hatalı giriş durumunda kullanıcıya bilgilendirme yapılır ve fonksiyon sonlandırılır.
		fmt.Println("Invalid input. Please enter valid integers.")
		return
	}

	fmt.Printf("\nInteger Operations:\n")
	// Toplama işlemi (Addition [Toplama]):
	// İki tam sayının toplamı hesaplanarak ekrana yazdırılır.
	fmt.Printf("%d + %d = %d\n", num1, num2, num1+num2)
	// Çıkarma işlemi (Subtraction [Çıkarma]):
	// İki tam sayı arasındaki fark hesaplanır.
	fmt.Printf("%d - %d = %d\n", num1, num2, num1-num2)
	// Çarpma işlemi (Multiplication [Çarpma]):
	// İki tam sayının çarpımı hesaplanır.
	fmt.Printf("%d * %d = %d\n", num1, num2, num1*num2)

	// Bölme işlemi (Division [Bölme]) yapılmadan önce sıfır kontrolü gerçekleştirilir.
	if num2 != 0 {
		// Bölme işlemi doğru şartlar altında gerçekleştirilir ve sonuç ekrana yazdırılır.
		fmt.Printf("%d / %d = %d\n", num1, num2, num1/num2)
	} else {
		// Sıfıra bölme hatası durumunda kullanıcıya uyarı mesajı verilir.
		fmt.Printf("Division by zero is not allowed.\n")
	}

	// Modulo işlemi (Remainder [Kalan]) gerçekleştirilir:
	// % operatörü kullanılarak, iki tam sayı arasındaki kalan hesaplanır.
	fmt.Printf("%d %% %d = %d\n", num1, num2, num1%num2)
}

func performComplexOperations() {
	// İki karmaşık sayı (Complex numbers [Karmaşık sayılar]) için değişkenlerin tanımlanması.
	var complex1, complex2 complex128
	// Kullanıcıdan iki karmaşık sayı girişi istenir (Input reading [Girdi alma]).
	fmt.Print("\nInput two complex numbers (e.g., 3+4i 5-2i): ")
	// fmt.Scan kullanılarak kullanıcı girdisi alınır; doğru formatta veri girilmediğinde hata mesajı ekrana yazdırılır.
	if _, err := fmt.Scan(&complex1, &complex2); err != nil {
		// Hatalı giriş durumunda kullanıcı bilgilendirilir ve fonksiyon sonlandırılır.
		fmt.Println("Invalid input. Please enter valid complex numbers.")
		return
	}
	// Karmaşık sayı işlemlerinin sonuçlarının ekrana yazdırılması (Output display [Çıktı gösterimi]).
	fmt.Printf("\nComplex Operations:\n")
	// Karmaşık sayıların toplama işlemi (Addition of complex numbers [Toplama]):
	// Her iki karmaşık sayının reel ve imajiner kısımları ayrı ayrı toplanır.
	fmt.Printf("%v + %v = %v\n", complex1, complex2, complex1+complex2)
	// Karmaşık sayıların çıkarma işlemi (Subtraction of complex numbers [Çıkarma]):
	// İki karmaşık sayı arasındaki fark hesaplanır.
	fmt.Printf("%v - %v = %v\n", complex1, complex2, complex1-complex2)
}

func showInvalidOperations() {
	// Kullanıcıya, mod operatörünün (Modulo Operator [Mod Operatörü]) float veya karmaşık sayılar üzerinde kullanılamayacağını hatırlatan mesaj ekrana yazdırılır.
	// Bu bilgi, operatörün yalnızca integer tipinde geçerli olduğunu vurgular.
	fmt.Println("\nReminder: %% [Modulo Operator] cannot be used with Floats or Complex numbers.")
}
