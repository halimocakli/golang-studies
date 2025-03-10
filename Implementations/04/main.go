package main

import (
	"fmt"  // Formatlı giriş/çıkış işlemleri için gerekli; kullanıcı arayüzü (UI - Kullanıcı Arayüzü) ile etkileşim sağlar.
	"log"  // Hata yönetimi ve loglama için gerekli; program çalışma sırasında oluşabilecek kritik hataların kaydedilmesi ve raporlanması amacıyla kullanılır.
	"math" // Matematiksel işlemler için gerekli fonksiyonları içerir; yuvarlama, trigonometrik işlemler gibi birçok hesaplama burada yer alır.
)

func main() {
	// Kullanıcıya ondalıklı bir sayı girmesi için mesaj gösterilir.
	fmt.Print("Input a float value (e.g., 3.14): ")

	// Kullanıcıdan alınacak ondalıklı sayı için değişken tanımlanır.
	// float64 veri tipi, yüksek doğruluk gerektiren hesaplamalar için tercih edilir.
	var number float64

	// Kullanıcı girdisini okuma; fmt.Scan fonksiyonu, terminalden veri okur.
	_, err := fmt.Scan(&number)
	if err != nil {
		// Hata oluştuğunda, log.Fatalf fonksiyonu hata mesajı ile birlikte programı sonlandırır.
		// Bu yöntem, beklenmeyen veri girişlerinin program akışını bozmasını önlemek için kullanılır.
		log.Fatalf("Error reading input: %v", err)
	}

	// Kullanıcının girdiği orijinal değeri, 2 ondalık basamak gösterilerek ekrana yazdırılır.
	fmt.Printf("\nOriginal: %.2f\n", number)

	// math.Round fonksiyonu, sayıyı en yakın tam sayıya yuvarlar (rounding - yuvarlama).
	// Bu fonksiyon, kesirli kısmın 0.5 ve üzeri olması durumunda bir üst tam sayıya, daha düşükse alt tam sayıya yuvarlar.
	fmt.Printf("Rounded: %.2f\n", math.Round(number))

	// math.Floor fonksiyonu, sayıyı aşağıya doğru en yakın tam sayıya yuvarlar (flooring - aşağı yuvarlama).
	// Bu işlem, sayının her durumda kesirli kısmını göz ardı ederek, sayının alt sınırını belirler.
	fmt.Printf("Floored: %.2f\n", math.Floor(number))

	// math.Ceil fonksiyonu, sayıyı yukarıya doğru en yakın tam sayıya yuvarlar (ceiling - yukarı yuvarlama).
	// Bu fonksiyon, kesirli kısım ne olursa olsun daima bir üst tam sayıya ulaşmayı garanti eder.
	fmt.Printf("Ceiled: %.2f\n", math.Ceil(number))
}
