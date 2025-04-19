package main

import (
	"fmt"
	"log"
)

// main fonksiyonu: Programın başlangıç noktasıdır (Entry point [giriş noktası])
// Bu fonksiyon, uygulamanın çalışmaya başladığı ilk yerdir ve runTest fonksiyonunu çağırarak test sürecini başlatır.
func main() {
	runTest() // Test sürecini başlatır
}

// runTest fonksiyonu: Test işlemlerinin organize edilmesi amacıyla tasarlanmıştır.
// Bu fonksiyon, uygulamanın temel test mantığını izole ederek, diğer test fonksiyonlarının çağrılmasını sağlar.
func runTest() {
	testSumThreeDigits() // Rakamların toplamı hesaplama testini çalıştırır
}

// testSumThreeDigits fonksiyonu: Kullanıcıdan sayı girişi alır ve rakamların toplamını hesaplar.
// Bu fonksiyon, kullanıcı etkileşimi, hata kontrolü ve negatif sayı işlemlerini içerir.
func testSumThreeDigits() {
	var num int // num değişkeni tanımlanır, tamsayı (integer [integer])

	// Kullanıcıdan sayı girişi istenir
	fmt.Print("Insert a number: ")

	// Kullanıcı girdisini okumaya çalışır; hata oluşursa, err değişkeni üzerinden hata yakalanır.
	// Hata durumunda, log.Fatalf fonksiyonu hata mesajı basar ve programı sonlandırır.
	_, err := fmt.Scan(&num)
	if err != nil {
		log.Fatalf("Error reading input: %v", err) // Hata yönetimi (error handling [hata yönetimi])
	}

	// Negatif sayı kontrolü: Eğer girilen sayı negatifse, abs fonksiyonu çağrılarak pozitif hale getirilir.
	// Bu sayede, rakamların toplamı hesaplarken negatif işaretin etkisi ortadan kaldırılır.
	if num < 0 {
		num = abs(num)
	}

	// sumThreeDigits fonksiyonu çağrılarak, sayının rakamlarının toplamı hesaplanır.
	// Hesaplanan toplam, fmt.Printf fonksiyonu ile ekrana yazdırılır.
	fmt.Printf("Sum of digits of the number is %d\n", sumThreeDigits(num))
}

// sumThreeDigits fonksiyonu: Bir sayının tüm rakamlarının toplamını hesaplar.
// Fonksiyon, while benzeri bir for döngüsü kullanarak sayıyı parçalara ayırır.
func sumThreeDigits(number int) int {
	var total int // Toplamın saklanacağı değişken (Accumulator [toplayıcı])

	// Sayı sıfır olana kadar döngü devam eder.
	// Her adımda, sayının son basamağı mod işlemi (number % 10) ile ayrılır.
	for number != 0 {
		// Her adımda, sayının son basamağı toplanarak total değişkenine eklenir.
		total = total + (number % 10)
		// Sayı, 10'a bölünerek en son basamak atılır ve döngü bir sonraki basamağa geçer.
		number = number / 10
	}

	// Hesaplanan rakamların toplamı döndürülür.
	return total
}

// abs fonksiyonu: Negatif sayıların pozitif hale çevrilmesini sağlar.
// Bu fonksiyon, yalnızca negatif sayı kontrolü yapıldığında çağrıldığından, basitçe değerin negatifini alır.
func abs(value int) int {
	return -value
}
