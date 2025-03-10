# Go Dilinde Matematiksel Özel Durumlar Örneği

Bu proje, Go dilinde matematiksel hesaplamalarda karşılaşılan özel durumları (NaN, Sonsuzluk) nasıl yöneteceğimizi göstermektedir. Projede, IEEE 754 standardına uygun olarak tanımsız ve sonsuz değerlerin nasıl üretildiği ve kontrol edildiği detaylı teknik terimlerle açıklanmıştır.

## İçerik

- [Genel Bakış](#genel-bakış)
- [Özellikler](#özellikler)
- [Kullanım](#kullanım)
- [Kod Açıklaması](#kod-açıklaması)
  - [Paket Tanımlaması](#paket-tanımlaması)
  - [Paket Dahil Etme](#paket-dahil-etme)
  - [Ana Fonksiyon](#ana-fonksiyon)
  - [Matematiksel İşlemler ve Kontroller](#matematiksel-işlemler-ve-kontroller)
- [Özet ve Teknik Notlar](#özet-ve-teknik-notlar)

## Genel Bakış

Bu proje, Go dilinde matematiksel işlemler sırasında ortaya çıkabilecek geçersiz hesaplamalar ve belirsiz değerlerin nasıl yönetilebileceğini ortaya koymaktadır. Özellikle, negatif sayıların karekökü ve logaritması gibi işlemlerde elde edilen **NaN (Not-a-Number [Sayı Değil])** ve **Sonsuzluk (Infinity/Sonsuzluk)** değerleri üzerinde durulmaktadır.

## Özellikler

- Negatif sayıların karekökünün hesaplanması (NaN üretimi)
- Pozitif sonsuzluktan pozitif sonsuzluğun çıkarılması (NaN üretimi)
- Negatif sayıların logaritmasının hesaplanması (NaN üretimi)
- Sıfırın logaritmasının hesaplanması (Negatif sonsuzluk üretimi)
- Özel durumların tespiti için `math.IsNaN` ve `math.IsInf` fonksiyonlarının kullanımı

## Kullanım

Projeyi yerel ortamınızda çalıştırmak için aşağıdaki adımları izleyebilirsiniz:

1. Go dilinin sisteminizde yüklü olduğundan emin olun.
2. Proje dosyasını klonlayın veya indirin.
3. Terminal veya komut satırında proje dizinine gidin.
4. Aşağıdaki komutu çalıştırarak uygulamayı başlatın:

   ```bash
   go run main.go
   ```

Ekrana matematiksel işlemlerin sonuçları ve durum tespitlerine ilişkin mesajlar yazdırılacaktır.

## Kod Açıklaması

### Paket Tanımlaması

- **Açıklama:**  
  `package main` ifadesiyle, yürütülebilir bir uygulamanın ana paketi tanımlanır.

### Paket Dahil Etme

- **Açıklama:**  
  `fmt` paketi formatlı giriş/çıkış işlemleri için, `math` paketi ise matematiksel fonksiyonların kullanımı için dahil edilmiştir.

### Ana Fonksiyon

- **Açıklama:**  
  `func main()` fonksiyonu, programın başlangıç noktası olup, tüm matematiksel hesaplamaların ve kontrollerin gerçekleştirildiği yerdir.

### Matematiksel İşlemler ve Kontroller

- **Negatif Sayının Karekökü (`math.Sqrt(-1)`):**

  - **İşlem:** Negatif bir sayının karekökü hesaplanır.
  - **Teknik Detay:** Negatif sayıların karekökü tanımsızdır ve IEEE 754 standardına göre **NaN** değeri döner.
  - **Kontrol:** `math.IsNaN(sqrtNeg)` ile elde edilen değerin NaN olup olmadığı kontrol edilir.

- **Pozitif Sonsuzluk İşlemleri:**

  - **İşlem:** `math.Inf(1)` ile pozitif sonsuzluk (+Infinity [+Sonsuzluk]) oluşturulur.
  - **Teknik Detay:** Pozitif sonsuzluktan pozitif sonsuzluk çıkarılması belirsizlik içerdiğinden **NaN** değeri üretir.
  - **Kontrol:** `math.IsNaN(infMinusInf)` kullanılarak belirsizlik sonucu ortaya çıkan NaN değeri tespit edilir.

- **Negatif Sayının Logaritması (`math.Log(-10)`):**

  - **İşlem:** Negatif bir sayının logaritması hesaplanır.
  - **Teknik Detay:** Negatif sayıların logaritması tanımsızdır ve bu nedenle **NaN** değeri üretir.
  - **Kontrol:** `math.IsNaN(logNeg)` ile sonuç kontrol edilir.

- **Sıfırın Logaritması (`math.Log(0)`):**
  - **İşlem:** Sıfırın logaritması hesaplanır.
  - **Teknik Detay:** Matematiksel tanım gereği, sıfırın logaritması **Negatif Sonsuzluk (negative Infinity [negatif Sonsuzluk])** sonucunu üretir.
  - **Kontrol:** `math.IsInf(logZero, 0)` ifadesi kullanılarak, hem pozitif hem de negatif sonsuzluk kontrolü yapılır.

Aşağıda, kod üzerinde detaylı yorumlar içeren hali yer almaktadır:

```go
package main

import (
	"fmt"
	"math"
)

func main() {
	// Negatif bir sayının karekökünü hesaplama.
	// Gerçek sayı sisteminde tanımsız olduğundan IEEE 754'e uygun olarak NaN üretir.
	sqrtNeg := math.Sqrt(-1)

	fmt.Printf("math.Sqrt(-1) = %v\n", sqrtNeg) // Beklenen çıktı: NaN

	// Üretilen değerin NaN olup olmadığını kontrol etme (hata yönetimi açısından kritik)
	if math.IsNaN(sqrtNeg) {
		fmt.Println("sqrtNeg is NaN (Not-a-Number)!")
	}

	fmt.Print("\n")

	// Pozitif sonsuzluk (+Infinity) oluşturma;
	// IEEE 754 standardına göre sonsuzluk kavramı uygulanır.
	inf := math.Inf(1)
	// Pozitif sonsuzluktan pozitif sonsuzluk çıkarılması belirsizlik içerdiğinden NaN üretir.
	infMinusInf := inf - inf

	fmt.Printf("(+Infinity) - (+Infinity) = %v\n", infMinusInf) // Beklenen çıktı: NaN

	// İşlem sonucunun NaN olup olmadığını kontrol etme
	if math.IsNaN(infMinusInf) {
		fmt.Println("infMinusInf is NaN (Not-a-Number)!")
	}

	fmt.Print("\n")

	// Negatif bir sayının logaritmasını hesaplama.
	// Matematiksel olarak tanımsız olduğundan NaN sonucu verir.
	logNeg := math.Log(-10)

	fmt.Printf("math.Log(-10) = %v\n", logNeg) // Beklenen çıktı: NaN

	// Hesaplamadan elde edilen sonucun NaN olup olmadığını kontrol etme
	if math.IsNaN(logNeg) {
		fmt.Println("logNeg is NaN (Not-a-Number)!")
	}

	fmt.Print("\n")

	// Sıfırın logaritmasının hesaplanması.
	// Matematiksel tanım gereği negatif sonsuzluk (negative Infinity) sonucu üretir.
	logZero := math.Log(0)

	fmt.Printf("math.Log(0) = %v\n", logZero)

	// logZero değerinin sonsuzluk (Infinity) olup olmadığını kontrol etme.
	// 0 parametresi, her iki yönlü sonsuzluğu kapsar.
	if math.IsInf(logZero, 0) {
		fmt.Println("logZero is negative Infinity!")
	}
}
```

## Özet ve Teknik Notlar

- **Kodun Amacı:**

  - Matematiksel fonksiyonların geçersiz veya tanımsız durumlarda nasıl davranacağını göstermek.
  - IEEE 754 standardına uygun olarak NaN ve sonsuzluk değerlerinin nasıl üretildiğini, kontrol edildiğini ve hata yönetiminin nasıl uygulandığını örneklemek.

- **Kullanılan Temel Terimler:**
  - **Paket Tanımlaması (Package Declaration):** Uygulamanın ana paketinin belirlenmesi.
  - **Paket Dahil Etme (Package Import):** Gerekli kütüphanelerin programa eklenmesi.
  - **Ana Fonksiyon (Main Function):** Programın çalışmaya başladığı giriş noktası.
  - **IEEE 754 Standardı:** Floating point hesaplamalarda kullanılan standart; NaN ve Sonsuzluk gibi kavramları tanımlar.
  - **NaN (Not-a-Number):** Tanımsız matematiksel işlemler sonucunda elde edilen özel değer.
  - **Infinity (Sonsuzluk):** Belirli matematiksel işlemler sonucu ortaya çıkan özel durum.
