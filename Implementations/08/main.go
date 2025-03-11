package main

import (
	"fmt" // Formatlı giriş/çıkış işlemleri için kullanılan paket (Input/Output işlemleri)
	"log" // Hata yönetimi ve loglama işlemleri için kullanılan paket (Error Logging)
)

func main() {
	// Kullanıcıdan ilk karmaşık sayıyı alma işlemi başlatılıyor.
	// prompt: "Enter the first complex number [like 3.4+4.5i or (3.4+4.5i)]: "
	// Burada, kullanıcıdan alınan değer complex128 tipinde saklanacaktır.
	complex1 := scanComplexNumber("Enter the first complex number [like 3.4+4.5i or (3.4+4.5i)]: ")

	// Kullanıcıdan ikinci karmaşık sayıyı alma işlemi başlatılıyor.
	// prompt: "Enter the second complex number [like 3.4+4.5i or (3.4+4.5i)]: "
	complex2 := scanComplexNumber("Enter the second complex number [like 3.4+4.5i or (3.4+4.5i)]: ")

	fmt.Print("\n-----------------------------------\n")

	// İlk karmaşık sayı, etiket ile birlikte ekrana yazdırılır.
	printComplexNumber("First complex number", complex1)
	// İkinci karmaşık sayı, etiket ile birlikte ekrana yazdırılır.
	printComplexNumber("Second complex number", complex2)

	fmt.Print("-----------------------------------\n")

	// İlk karmaşık sayının reel (real) ve sanal (imaginary) bileşenleri ayrı ayrı ekrana yazdırılır.
	printRealAndImagParts("Complex1", complex1)

	fmt.Print("-----------------------------------\n")

	printRealAndImagParts("Complex2", complex2)

	fmt.Print("-----------------------------------\n")

	// İki karmaşık sayı üzerinde aritmetik işlemler (toplama, çıkarma, çarpma, bölme) gerçekleştirilir.
	performArithmeticOperations("Complex1", "Complex2", complex1, complex2)
}

// scanComplexNumber fonksiyonu, kullanıcıdan bir karmaşık sayı girişi alır ve bunu complex128 tipinde döner.
// Fonksiyon, kullanıcıdan doğru formatta veri alınmasını garanti altına almak için hata kontrolü yapar.
func scanComplexNumber(prompt string) complex128 {
	var c complex128

	// Kullanıcıya, karmaşık sayı girişi için talimat içeren prompt mesajı gösterilir.
	fmt.Print(prompt)

	// Kullanıcıdan girdi alınır; alınan değer c değişkenine aktarılır.
	_, err := fmt.Scan(&c)
	if err != nil {
		// Eğer kullanıcı yanlış formatta bir giriş yaparsa, log.Fatalf kullanılarak hata mesajı verilir.
		// Bu, programın beklenmeyen girişlerden kaynaklanan hataları erken aşamada yakalamasını sağlar.
		log.Fatalf("Wrong input: %v", err)
	}
	return c
}

// printComplexNumber fonksiyonu, verilen karmaşık sayıyı formatlı bir şekilde ekrana yazdırır.
// Bu, sayının hem reel hem de sanal bileşenlerinin tek satırda gösterilmesini sağlar.
func printComplexNumber(label string, c complex128) {
	fmt.Printf("%s: %f\n", label, c)
}

// printRealAndImagParts fonksiyonu, verilen karmaşık sayının reel (real) ve sanal (imaginary) bileşenlerini ayrı ayrı ekrana yazdırır.
// Bu, kullanıcının sayının hangi parçalardan oluştuğunu daha iyi anlamasına yardımcı olur.
func printRealAndImagParts(label string, c complex128) {
	fmt.Printf("%s Real part: %f\n", label, real(c))
	fmt.Printf("%s Imaginary part: %f\n", label, imag(c))
}

// performArithmeticOperations fonksiyonu, iki karmaşık sayı üzerinde aritmetik işlemler gerçekleştirir.
// İşlemler: toplama (addition), çıkarma (subtraction), çarpma (multiplication) ve bölme (division).
func performArithmeticOperations(label1, label2 string, c1, c2 complex128) {
	// İşlem öncesinde, her iki karmaşık sayının ekrana yazdırılması.
	fmt.Printf("%s: %f, %s: %f\n", label1, c1, label2, c2)
	// Toplama işlemi gerçekleştirilir ve sonucu ekrana basılır.
	fmt.Printf("Sum: %f\n", c1+c2)
	// Çıkarma işlemi gerçekleştirilir ve sonucu ekrana basılır.
	fmt.Printf("Difference: %f\n", c1-c2)
	// Çarpma işlemi gerçekleştirilir ve sonucu ekrana basılır.
	fmt.Printf("Product: %f\n", c1*c2)

	// Bölme işlemi yapılmadan önce, bölenin (c2) sıfır olup olmadığı kontrol edilir.
	if c2 == 0 {
		// Eğer c2 sıfırsa, sıfıra bölme hatası önlenir ve kullanıcıya uyarı mesajı gösterilir.
		fmt.Println("Cannot divide by zero")
		return
	}

	// c2 sıfır değilse, bölme işlemi gerçekleştirilir ve sonucu ekrana yazdırılır.
	fmt.Printf("Quotient: %f\n", c1/c2)
}
