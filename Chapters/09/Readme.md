## 1. Kontrol Deyimleri (Control Statements)

Go dilinde, akış kontrol deyimleri programın mantıksal işleyişinde önemli rol oynar. Bu deyimlerin doğru kullanımı, kodun okunabilirliğini, verimliliğini ve hata ayıklamasını kolaylaştırır. Aşağıdaki bölümlerde her bir kontrol deyimi detaylandırılmıştır.

### 1.1. break (break – break) Deyimi

- **Tanım ve Kullanım:**

  - **break** deyimi, **for** döngüsü, **switch** veya **select** bloğu içerisinde kullanıldığında o bloğun yürütülmesini anında sonlandırır.
  - Döngü içerisinde tek başına kullanıldığında, o döngünün sonraki iterasyonlarına geçmeden tamamen sonlandırılmasını sağlar.
  - **Labeled break (etiketli break):**
    - İç içe geçmiş döngülerde, yalnızca en içteki döngüyü değil, belirli bir dış döngüyü sonlandırmak için kullanılır.
    - Kullanım öncesinde, çıkılması istenen döngüye anlamlı bir etiket (label) tanımlanmalıdır. Etiket isimleri Go’nun değişken isimlendirme kurallarına uygun, açıklayıcı olmalıdır.

- **Detaylı Açıklamalar:**

  - **Kodun Okunabilirliği:**
    - Etiket kullanımı, kodun hangi noktada döngüden çıkılacağını netleştirir. Ancak, gereksiz etiket kullanımı kodu karmaşıklaştırabilir.
  - **Hata Yönetimi:**
    - Döngü içindeki hatalı durumların tespiti sırasında, hatalı durumda doğrudan döngüden çıkmak için break kullanılması, hata ayıklama sürecini basitleştirir.
  - **Performans:**
    - Döngü içerisindeki gereksiz işlemleri atlayarak zaman kazandırır.

- **Örnek Kod (Detaylandırılmış):**

  ```go
  package main

  import (
  	"fmt"
  	"time"
  )

  func main() {
  exitLoop: // Dış döngüden çıkmak için etiket tanımlandı
  	for i := 0; i < 20; i++ {
  		// İç döngü: k değeri 40'tan 0'a kadar azalır
  		for k := 40; k >= 0; k-- {
  			fmt.Printf("(%d, %d)\n", i, k)
  			// Koşul sağlanırsa, break ile exitLoop etiketli döngüden çıkılır
  			if (i+k)%6 == 0 {
  				break exitLoop
  			}
  			// İşlemler arasında küçük gecikme ekleyerek çıktı akışını yavaşlatıyoruz
  			time.Sleep(50 * time.Millisecond)
  		}
  	}
  	fmt.Println("Döngüden break ile çıkıldı.")
  }
  ```

- **Ek Notlar:**
  - Etiketli break kullanımı, iç içe geçmiş döngülerde istenilen dış döngüden çıkmak için idealdir.
  - Eğer sadece en içteki döngü sonlandırılacaksa, etiket kullanılmasına gerek yoktur.
  - **break** deyiminin, switch ve select bloklarında da kullanıldığı unutulmamalıdır.

---

### 1.2. continue (continue – continue) Deyimi

- **Tanım ve Kullanım:**

  - **continue** deyimi, döngünün o anki iterasyonunun kalan kısmını atlayıp, bir sonraki döngü adımına geçiş yapar.
  - Yalnızca **for** döngülerinde kullanılabilir.
  - **Labeled continue:**
    - Etiketli continue kullanılarak, iç içe döngülerde belirli bir dış döngünün sonraki iterasyonuna geçmek mümkündür; ancak, bu kullanım genellikle kodun okunabilirliğini zorlaştırabilir.

- **Detaylı Açıklamalar:**

  - **Kullanım Amacı:**
    - Döngü içerisinde belirli bir koşul sağlandığında, o iterasyonda yapılması gerekmeyen işlemleri atlayarak performansı artırır.
  - **Kod Yapısına Etkisi:**
    - Gereksiz işlemlerden kaçınarak, kodun mantıksal akışını daha net hale getirir.
  - **Etiketli Kullanım:**
    - Karmaşık iç içe döngülerde, istenmeyen durumlarda daha dış bir döngü iterasyonuna geçiş yapma imkanı sunar.

- **Örnek Kod (Basit continue Kullanımı):**

  ```go
  package main

  import "fmt"

  func main() {
  	for i := 0; i < 10; i++ {
  		// Eğer i değeri çiftse, kalan işlemler atlanır
  		if i%2 == 0 {
  			continue
  		}
  		// Sadece tek sayı değerleri yazdırılır
  		fmt.Printf("Tek sayı: %d\n", i)
  	}
  }
  ```

- **Ek Örnek (İç içe döngüde etiketli continue):**

  ```go
  package main

  import "fmt"

  func main() {
  outerLoop: // Dış döngüye ait etiket
  	for i := 0; i < 3; i++ {
  		for j := 0; j < 3; j++ {
  			// Eğer i değeri j değerine eşitse, dış döngünün sonraki iterasyonuna geç
  			if i == j {
  				continue outerLoop
  			}
  			fmt.Printf("i = %d, j = %d\n", i, j)
  		}
  	}
  }
  ```

- **Ek Notlar:**
  - **continue** deyimi, özellikle büyük veri setleriyle çalışan döngülerde gereksiz hesaplamaları önler.
  - Etiketli continue kullanımı dikkatli uygulanmalıdır; aksi halde kod akışı kafa karıştırıcı hale gelebilir.

---

### 1.3. goto (goto – goto) Deyimi

- **Tanım ve Kullanım:**

  - **goto** deyimi, programın akışını belirlenen bir etiket noktasına atlayarak değiştirir.
  - Genellikle hata durumlarında, kaynak temizliği (resource release) veya iç içe döngülerde çıkış sağlamak için kullanılır.
  - Ancak, yapılandırılmış kontrol akışı sağlanabildiğinde goto kullanımı önerilmez; kodun okunabilirliğini düşürebilir.

- **Detaylı Açıklamalar:**

  - **Kullanım Senaryoları:**
    - Çok nadir durumlarda, özellikle karmaşık hata yönetimi senaryolarında kullanılır.
    - Kernel veya sistem programlama gibi çok özel durumlarda tercih edilebilir.
  - **Alternatif Yapılar:**
    - Go dilinde defer, break ve continue gibi yapıların mevcut olması, goto kullanımını çoğu durumda gereksiz kılar.
  - **Kod Düzenine Etkisi:**
    - Yanlış kullanıldığında, kodun akışını takip etmek zorlaşır; bu yüzden sadece gerekli durumlarda ve iyi belgelenmiş etiketlerle kullanılmalıdır.

- **Örnek Kod:**

  ```go
  package main

  import "fmt"

  func main() {
  	// İç içe döngülerde belirli koşul sağlandığında goto kullanılarak döngüden çıkılır.
  	for i := 0; i < 20; i++ {
  		for k := 40; k >= 0; k-- {
  			fmt.Printf("(%d, %d)\n", i, k)
  			if (i+k)%6 == 0 {
  				goto exitLoop
  			}
  		}
  	}
  exitLoop:
  	fmt.Println("goto ile belirtilen etiket noktasına atlanıldı, döngü sonlandı.")
  }
  ```

- **Ek Notlar:**
  - **goto** kullanılırken, etiketin tanımlandığı yerin kodun mantıksal akışına uygun olması önemlidir.
  - Yapısal kontrol akışlarını sağlamak için öncelikle diğer deyimler (break, continue, defer) düşünülmelidir.

---

## 2. switch (switch – switch) Deyimi ve İlgili Kavramlar

Switch deyimi, çeşitli durumlar arasında koşul kontrolü yapmayı kolaylaştırır. Go dilinde switch deyimi, hem sabit ifadeler hem de boolean veya tip bazlı kontrol yapılabilmesini sağlar.

### 2.1. Expression Switch

- **Tanım:**
  - Sabit bir ifade üzerinden durum kontrolü yapılır.
  - İfadeye göre case’ler arasında seçim yapar.
- **Detaylı Açıklamalar:**
  - **Çoklu Karşılaştırma:**
    - Birden fazla değerin aynı case bloğunda değerlendirilmesi mümkündür (örneğin, `case 212, 216:`).
  - **Okunabilirlik:**
    - if-else zincirlerine göre daha okunabilir ve net bir yapı sunar.
- **Örnek Kod:**

  ```go
  package main

  import "fmt"

  func main() {
  	var code int
  	fmt.Print("Telefon kodunu giriniz: ")
  	fmt.Scan(&code)

  	switch code {
  	case 212, 216:
  		fmt.Println("İstanbul")
  	case 312:
  		fmt.Println("Ankara")
  	case 372:
  		fmt.Println("Zonguldak")
  	default:
  		fmt.Println("Geçersiz kod numarası!")
  	}
  	fmt.Println("Switch kontrolü tamamlandı.")
  }
  ```

- **Ek Açıklama:**
  - Expression switch, sabit değerlerle çalıştığından performans açısından optimize edilmiş bir yapıdır.

---

### 2.2. Fallthrough (fallthrough – fallthrough) Kullanımı

- **Tanım:**
  - Switch bloğunda, bir case çalıştıktan sonra, sonraki case’in de çalışmasını zorlar.
- **Detaylı Açıklamalar:**
  - **Varsayılan Davranış:**
    - Go’da varsayılan olarak bir case bloğu çalıştıktan sonra diğer case’lere geçiş yapılmaz.
    - Fallthrough ifadesi eklenirse, kontrol sonraki case bloğuna aktarılır, ancak koşul kontrolü yeniden yapılmaz.
  - **Kullanım Amacı:**
    - Birden fazla durumun aynı çıktıyı üretmesi isteniyorsa veya ara bir açıklama yapmak gerekiyorsa kullanılabilir.
- **Örnek Kod:**

  ```go
  package main

  import "fmt"

  func main() {
  	var code int
  	fmt.Print("Telefon kodunu giriniz: ")
  	fmt.Scan(&code)

  	switch code {
  	case 212:
  		fmt.Print("Avrupa ")
  		fallthrough // Koşul kontrolü yapılmadan bir sonraki case çalıştırılır
  	case 216:
  		fmt.Println("İstanbul")
  	case 312:
  		fmt.Println("Ankara")
  	case 372:
  		fmt.Println("Zonguldak")
  	default:
  		fmt.Println("Geçersiz kod numarası!")
  	}
  	fmt.Println("Fallthrough örneği tamamlandı.")
  }
  ```

- **Ek Notlar:**
  - Fallthrough kullanımı, özellikle birbirini takip eden durumlarda, ara açıklama eklenmek istenen senaryolarda işe yarar.
  - Dikkat edilmezse, beklenmeyen çıktılara yol açabileceğinden, kodun mantığına uygun olup olmadığı iyi değerlendirilmelidir.

---

### 2.3. Boolean (Koşul) İfadelerle Switch

- **Tanım:**
  - Switch deyiminde sabit bir ifade yazılmadan, her bir case ifadesi boolean (true/false) değeri döndüren ifadelerle değerlendirilir.
- **Detaylı Açıklamalar:**
  - **if-else Alternatifi:**
    - Bu kullanım, klasik if-else zincirlerinin yerine daha okunabilir bir yapı sunar.
  - **Dinamik Koşullar:**
    - İfadeler fonksiyon çağrıları, matematiksel hesaplamalar veya değişkenlerin durumuna göre değerlendirilebilir.
- **Örnek Kod:**

  ```go
  package main

  import "fmt"

  func main() {
  	var n int
  	fmt.Print("Bir sayı giriniz: ")
  	fmt.Scan(&n)

  	switch result := factorial(n); {
  	case result > 10:
  		fmt.Printf("%d! > 10\n", result)
  	case result < 5:
  		fmt.Printf("%d! <= 5\n", result)
  	default:
  		fmt.Printf("%d! is not in range\n", result)
  	}
  	fmt.Println("Boolean switch örneği tamamlandı.")
  }

  func factorial(n int) int {
  	result := 1
  	for i := 2; i <= n; i++ {
  		result *= i
  	}
  	return result
  }
  ```

- **Ek Notlar:**
  - Bu yapı, karmaşık koşulların tek bir switch bloğunda toplanması için idealdir.
  - Her case ifadesi, koşulun sağlanıp sağlanmadığını net olarak belirtir.

---

### 2.4. Switch Deyimi Kullanımının Karşılaştırması

Aşağıdaki tablo, switch deyiminin farklı kullanım türlerini özetler:

| **Switch Türü**       | **Açıklama**                                                  | **Kullanım Alanları**                              |
| --------------------- | ------------------------------------------------------------- | -------------------------------------------------- |
| **Expression Switch** | Sabit ifadeler üzerinden durum kontrolü yapılır.              | Menü seçenekleri, sabit değer karşılaştırmaları    |
| **Boolean Switch**    | Her case ifadesi boolean değer döndürür, if-else alternatifi. | Karmaşık koşullar, dinamik ifade değerlendirmeleri |
| **Type Switch**       | Değişkenin tipine göre durum kontrolü yapılır.                | Arayüzlerden gelen verilerin tip kontrolü          |

- **Ek Açıklama:**
  - Type switch, Go’da daha ileri düzeyde, interface tipler ile çalışırken kullanılır. Bu durumda bir değişkenin hangi somut tipe sahip olduğunu kontrol etmek mümkündür.

---

## 3. Menü Uygulamaları ve Kullanıcı Etkileşimi

Konsol tabanlı uygulamalarda menü uygulamaları, kullanıcı etkileşimini sağlamanın temel yollarından biridir. Aşağıdaki örneklerde if ve switch deyimleri kullanılarak farklı menü uygulaması senaryoları ele alınmıştır.

### 3.1. if ile Menü Uygulaması

- **Özellikler ve Avantajları:**
  - Kullanıcının girdiği değerin kontrolü if-else yapıları ile yapılır.
  - Değişkenlerin scope’u daraltılarak, her bir if bloğunda yeni değişkenler tanımlanabilir.
  - Küçük ve basit menü seçeneklerinde tercih edilir.
- **Detaylı Açıklama:**
  - if yapısı, her koşul için açıkça belirlenmiş adımları içerir.
  - Eğer seçenek dışındaki değerler girilirse, kullanıcıya hata mesajı verilerek tekrar giriş yapılması sağlanır.
- **Örnek Kod:**

  ```go
  package main

  import (
  	"fmt"
  	"os"
  )

  func main() {
  	runDemoApp()
  }

  func runDemoApp() {
  	for {
  		printMenu()
  		option := readOption()
  		// Seçenek kontrolü: Geçersiz seçeneklerde döngü devam eder.
  		if option < 1 || option > 5 {
  			fmt.Println("Geçersiz seçenek, lütfen tekrar deneyiniz.")
  			continue
  		}
  		// Çıkış seçeneğinde döngüden çıkılır.
  		if option == 5 {
  			fmt.Println("Teşekkürler! Program sonlanıyor.")
  			break
  		}
  		doForOption(option)
  	}
  }

  func doForOption(option int) {
  	if option == 1 {
  		insertProc()
  	} else if option == 2 {
  		searchProc()
  	} else if option == 3 {
  		deleteProc()
  	} else {
  		updateProc()
  	}
  }

  func insertProc() {
  	fmt.Println("----- Ekleme İşlemi Başladı -----")
  	// Burada veri ekleme işlemleri gerçekleştirilebilir.
  	fmt.Println("Veri eklendi.")
  	fmt.Println("---------------------------------")
  }

  func searchProc() {
  	fmt.Println("----- Arama İşlemi Başladı -----")
  	// Arama algoritması veya veri sorgulaması gerçekleştirilebilir.
  	fmt.Println("Arama tamamlandı.")
  	fmt.Println("---------------------------------")
  }

  func deleteProc() {
  	fmt.Println("----- Silme İşlemi Başladı -----")
  	// Silme işlemleri burada yapılır.
  	fmt.Println("Silme işlemi başarılı.")
  	fmt.Println("---------------------------------")
  }

  func updateProc() {
  	fmt.Println("----- Güncelleme İşlemi Başladı -----")
  	// Güncelleme işlemleri burada uygulanır.
  	fmt.Println("Güncelleme tamamlandı.")
  	fmt.Println("-------------------------------------")
  }

  func printMenu() {
  	fmt.Println("1. Ekle")
  	fmt.Println("2. Ara")
  	fmt.Println("3. Sil")
  	fmt.Println("4. Güncelle")
  	fmt.Println("5. Çıkış")
  	fmt.Print("Lütfen bir seçenek giriniz: ")
  }

  func readOption() int {
  	var option int
  	fmt.Scan(&option)
  	return option
  }
  ```

---

### 3.2. switch ile Menü Uygulaması

- **Özellikler ve Avantajları:**
  - Switch deyimi kullanılarak durum kontrolü daha net ve okunabilir hale getirilir.
  - Menü seçenekleri arasında ayrım yapmak için daha derli toplu bir yapı sunar.
  - Hata durumunda default bloğu ile kullanıcıya uyarı mesajı verilmesi sağlanır.
- **Detaylı Açıklama:**
  - Switch kullanımı, çok sayıda if-else durumuna göre daha kısa ve anlaşılır kod yazılmasını sağlar.
- **Örnek Kod:**

  ```go
  package main

  import (
  	"fmt"
  	"os"
  )

  func main() {
  	runDemoApp()
  }

  func runDemoApp() {
  	for {
  		printMenu()
  		option := readOption()
  		switch option {
  		case 1:
  			insertProc()
  		case 2:
  			searchProc()
  		case 3:
  			deleteProc()
  		case 4:
  			updateProc()
  		case 5:
  			exitProc()
  		default:
  			fmt.Println("Geçersiz seçenek, lütfen tekrar deneyiniz.")
  		}
  	}
  }

  func insertProc() {
  	fmt.Println("----- Ekleme İşlemi Başladı -----")
  	// Veri ekleme işlemleri burada yapılabilir.
  	fmt.Println("Veri eklendi.")
  	fmt.Println("---------------------------------")
  }

  func searchProc() {
  	fmt.Println("----- Arama İşlemi Başladı -----")
  	// Veri arama işlemleri burada yapılır.
  	fmt.Println("Arama tamamlandı.")
  	fmt.Println("---------------------------------")
  }

  func deleteProc() {
  	fmt.Println("----- Silme İşlemi Başladı -----")
  	// Silme işlemleri burada uygulanır.
  	fmt.Println("Silme işlemi başarılı.")
  	fmt.Println("---------------------------------")
  }

  func updateProc() {
  	fmt.Println("----- Güncelleme İşlemi Başladı -----")
  	// Güncelleme işlemleri burada gerçekleştirilir.
  	fmt.Println("Güncelleme tamamlandı.")
  	fmt.Println("-------------------------------------")
  }

  func exitProc() {
  	fmt.Println("Teşekkürler! Program sonlanıyor.")
  	os.Exit(0)
  }

  func printMenu() {
  	fmt.Println("1. Ekle")
  	fmt.Println("2. Ara")
  	fmt.Println("3. Sil")
  	fmt.Println("4. Güncelle")
  	fmt.Println("5. Çıkış")
  	fmt.Print("Lütfen bir seçenek giriniz: ")
  }

  func readOption() int {
  	var option int
  	fmt.Scan(&option)
  	return option
  }
  ```

- **Karşılaştırma Tablosu:**

  | **Yöntem** | **Avantajları**                                       | **Dezavantajları**                           |
  | ---------- | ----------------------------------------------------- | -------------------------------------------- |
  | if         | Basit yapılar, dar kapsamlı değişken scope’u sağlar   | Çok fazla koşul varsa okunabilirlik azalır   |
  | switch     | Durumlar net şekilde ayrılır, kod daha derli topludur | Sadece sabit veya boolean ifadelerle çalışır |

---

## 4. Tarih İşlemleri: isValidDate, dayOfYear ve dayOfWeek Fonksiyonları

Standart **time** paketini kullanmadan tarih doğrulama ve hesaplamaları yapmak, algoritmanın temel kavramlarını anlamada önemlidir. Bu fonksiyonlar tarihlerin geçerliliğini kontrol etme, yılın kaçıncı günü olduğunu ve haftanın hangi gününe denk geldiğini hesaplamaya yöneliktir.

### 4.1. isValidDate (isValidDate – geçerli tarih kontrolü)

- **Tanım:**

  - Girilen gün, ay ve yıl değerlerinin geçerli bir tarih oluşturup oluşturmadığını kontrol eder.
  - Yıl değeri 1900’den küçük olmamalı, ay değeri 1 ile 12 arasında ve gün değeri ilgili ayın maksimum gün sayısına uygun olmalıdır.

- **Detaylı Açıklamalar:**

  - **Ay Bazlı Gün Kontrolü:**
    - Her ayın gün sayısı farklıdır; örneğin Şubat 28 veya artık yıllarda 29 gün çeker.
  - **Artık Yıl Kontrolü:**
    - Artık yıllarda Şubat ayı 29 gün olduğu için, bu kontrol fonksiyon içerisinde netleştirilmelidir.
  - **Giriş Doğrulama:**
    - Kullanıcıdan alınan verinin sayısal olması, negatif sayıların kontrolü gibi ek doğrulamalar yapılabilir.

- **Örnek Kod:**

  ```go
  package main

  import "fmt"

  func main() {
  	runIsValidDateTest()
  }

  func runIsValidDateTest() {
  	for {
  		var d, m, y int
  		fmt.Print("Tarih değerlerini giriniz (gün ay yıl): ")
  		fmt.Scan(&d, &m, &y)

  		if isValidDate(d, m, y) {
  			fmt.Printf("%02d/%02d/%04d geçerli bir tarihtir.\n", d, m, y)
  		} else {
  			fmt.Println("Geçersiz tarih değeri!")
  		}
  		// Çıkmak için 0 0 0 giriniz
  		if d == 0 && m == 0 && y == 0 {
  			break
  		}
  	}
  }

  func isValidDate(day, month, year int) bool {
  	return (1 <= day && day <= 31) &&
  		(1 <= month && month <= 12) &&
  		(year >= 1900) &&
  		(day <= getDays(month, year))
  }

  func getDays(month, year int) int {
  	days := 31
  	switch month {
  	case 4, 6, 9, 11:
  		days = 30
  	case 2:
  		days = 28
  		if isLeapYear(year) {
  			days++
  		}
  	}
  	return days
  }

  func isLeapYear(year int) bool {
  	return (year%4 == 0 && year%100 != 0) || (year%400 == 0)
  }
  ```

---

### 4.2. dayOfYear (dayOfYear – yılın kaçıncı günü)

- **Tanım:**

  - Girilen tarihin, yılın kaçıncı günü olduğunu hesaplar.
  - Geçersiz tarih için -1 döner.

- **Detaylı Açıklamalar:**

  - **Birikimli Gün Hesabı:**
    - Her ayın gün sayısı toplanarak girilen gün değeri eklenir.
  - **Döngüsel Hesaplama:**
    - Aylar için bir döngü kullanılarak, girilen ay öncesindeki tüm günlerin toplamı hesaplanır.
  - **Ek Doğrulama:**
    - Tarih geçerliliği kontrol edilmeden hesaplama yapılırsa yanlış sonuçlar elde edilebilir; bu yüzden ilk kontrol önemlidir.

- **Örnek Kod:**

  ```go
  package main

  import "fmt"

  func main() {
  	runDayOfYearTest()
  }

  func runDayOfYearTest() {
  	for {
  		var d, m, y int
  		fmt.Print("Tarih değerlerini giriniz (gün ay yıl): ")
  		fmt.Scan(&d, &m, &y)

  		dof := dayOfYear(d, m, y)
  		if dof != -1 {
  			fmt.Printf("%02d/%02d/%04d, yılın %d. günü\n", d, m, y, dof)
  		} else {
  			fmt.Println("Geçersiz tarih değeri!")
  		}
  		if d == 0 && m == 0 && y == 0 {
  			break
  		}
  	}
  }

  func dayOfYear(day, month, year int) int {
  	if isValidDate(day, month, year) {
  		return getAccumulatedDays(month, year) + day
  	}
  	return -1
  }

  func getAccumulatedDays(month, year int) int {
  	total := 0
  	for m := 1; m < month; m++ {
  		total += getDays(m, year)
  	}
  	return total
  }
  ```

---

### 4.3. dayOfWeek (dayOfWeek – haftanın günü)

- **Tanım:**
  - 01.01.1900 tarihinden itibaren girilen tarihe kadar geçen gün sayısının 7’ye bölümünden kalana göre haftanın gününü belirler.
  - Sonuç; 0: Pazar, 1: Pazartesi, 2: Salı, …, 6: Cumartesi şeklinde yorumlanır.
- **Detaylı Açıklamalar:**
  - **Toplam Gün Hesabı:**
    - İlk olarak, girilen tarihin yıl içindeki sırası (dayOfYear) hesaplanır.
    - Daha sonra, 1900 yılından girilen yıla kadar olan yılların toplam gün sayısı eklenir.
  - **Mod İşlemi:**
    - Toplam gün sayısının 7’ye bölümünden kalan, haftanın hangi gününe denk geldiğini gösterir.
- **Örnek Kod:**

  ```go
  package main

  import "fmt"

  func main() {
  	runDayOfWeekTest()
  }

  func runDayOfWeekTest() {
  	for {
  		var d, m, y int
  		fmt.Print("Tarih değerlerini giriniz (gün ay yıl): ")
  		fmt.Scan(&d, &m, &y)

  		dow := dayOfWeek(d, m, y)
  		switch dow {
  		case 0:
  			fmt.Printf("%02d/%02d/%04d Pazar\n", d, m, y)
  		case 1:
  			fmt.Printf("%02d/%02d/%04d Pazartesi\n", d, m, y)
  		case 2:
  			fmt.Printf("%02d/%02d/%04d Salı\n", d, m, y)
  		case 3:
  			fmt.Printf("%02d/%02d/%04d Çarşamba\n", d, m, y)
  		case 4:
  			fmt.Printf("%02d/%02d/%04d Perşembe\n", d, m, y)
  		case 5:
  			fmt.Printf("%02d/%02d/%04d Cuma\n", d, m, y)
  		case 6:
  			fmt.Printf("%02d/%02d/%04d Cumartesi\n", d, m, y)
  		default:
  			fmt.Println("Geçersiz tarih!")
  		}
  		if d == 0 && m == 0 && y == 0 {
  			break
  		}
  	}
  }

  func dayOfWeek(day, month, year int) int {
  	if !isValidDate(day, month, year) {
  		return -1
  	}
  	dof := dayOfYear(day, month, year)
  	totalDays := dof
  	// 1900 yılından girilen yıla kadar olan günleri ekle
  	for y := 1900; y < year; y++ {
  		totalDays += 365
  		if isLeapYear(y) {
  			totalDays++
  		}
  	}
  	return totalDays % 7
  }
  ```

- **Ek Notlar:**
  - Tarih hesaplamalarında, özellikle artık yıl kontrolü çok önemlidir.
  - Bu yöntem, standart kütüphane kullanılmadan manuel hesaplama yapmanın temel algoritmasını öğretir.

---

## 5. Fonksiyonlarda Named Return Values (İsimlendirilmiş Geri Dönüş Değişkenleri)

- **Tanım ve Avantajlar:**
  - Fonksiyon tanımlanırken, geri dönüş değerlerine isim verilebilir.
  - Bu, fonksiyon içinde bu değişkenlerin doğrudan kullanılmasını, dokümantasyonun ve kod okunabilirliğinin artmasını sağlar.
- **Detaylı Açıklamalar:**
  - **Kullanım Kolaylığı:**
    - Fonksiyon içerisinde hesaplanan değerler, isimlendirilmiş geri dönüş değişkenleri ile saklanır; bu sayede return ifadesinde değerleri tek tek belirtmeye gerek kalmaz.
  - **Hata Yönetimi:**
    - Fonksiyonun herhangi bir yerinde hata durumuna göre geri dönüş değeri belirlenebilir; bu durum kodun akışını daha anlaşılır kılar.
  - **Örnek Senaryolar:**
    - Matematiksel işlemler, veri dönüşümleri veya API çağrılarında, fonksiyonun döndürdüğü değerlerin ne anlama geldiğini isimlendirmek yararlı olur.
- **Örnek Kod (Quadratic Equation Roots – İkinci Dereceden Denklemin Kökleri):**

  ```go
  package main

  import (
  	"fmt"
  	"math"
  )

  func main() {
  	runQuadraticDemo()
  }

  func runQuadraticDemo() {
  	var a, b, c float64
  	fmt.Print("Katsayıları giriniz (a, b, c): ")
  	fmt.Scan(&a, &b, &c)

  	// Fonksiyon, isimlendirilmiş dönüş değerleri (exists, x1, x2) ile kökleri hesaplar.
  	exists, x1, x2 := findQuadraticEquationRoots(a, b, c)
  	if exists {
  		fmt.Printf("x1 = %f, x2 = %f\n", x1, x2)
  	} else {
  		fmt.Println("Gerçek kök yok!")
  	}
  }

  // Fonksiyon tanımında geri dönüş değerlerine isim verildi: exists, x1, x2
  func findQuadraticEquationRoots(a, b, c float64) (exists bool, x1 float64, x2 float64) {
  	delta := b*b - 4*a*c
  	if delta >= 0 {
  		sqrtDelta := math.Sqrt(delta)
  		exists = true
  		x1 = (-b + sqrtDelta) / (2 * a)
  		x2 = (-b - sqrtDelta) / (2 * a)
  	}
  	// İsimlendirilmiş dönüş değerleri otomatik olarak döner.
  	return
  }
  ```

- **Ek Notlar:**
  - Fonksiyon içerisindeki isimlendirilmiş dönüş değerleri, kodun daha anlaşılır olmasını sağlar.
  - Özellikle kompleks hesaplamalar yapılan durumlarda, hangi değişkenin neyi temsil ettiğinin belirtilmesi hata olasılığını azaltır.

---

## Özet ve Sonuç

Go dilinde kontrol deyimleri ve ilgili yapıların kullanımı, yazılımın akışını doğru şekilde yönetmek için kritiktir. Aşağıdaki tablo, temel deyimler ve kullanım alanlarını özetlemektedir:

| **Keyword**                   | **Açıklama**                                                                                          | **Kullanım Alanları / Avantajlar**                                                                                     |
| ----------------------------- | ----------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------- |
| **break (break)**             | Döngü, switch veya select içerisindeki akışı sonlandırır.                                             | İç içe döngülerde dış döngüden çıkma, hata kontrolü, gereksiz işlemden kaçınma.                                        |
| **continue (continue)**       | Döngünün mevcut iterasyonunun kalanını atlayarak sonraki iterasyona geçiş yapar.                      | Büyük veri setlerinde performans artışı, gereksiz işlemlerin atlanması.                                                |
| **goto (goto)**               | Belirtilen etikete atlayarak program akışını değiştirir.                                              | Nadir hata durumları, kaynak temizliği veya iç içe döngülerde özel durumlar; yapılandırılmış alternatifler yetersizse. |
| **switch (switch)**           | Durum kontrolü için çoklu koşulları net şekilde ayırır; hem sabit hem de boolean ifadelerle çalışır.  | Menü uygulamaları, sabit değer karşılaştırmaları, dinamik koşulların değerlendirilmesi.                                |
| **fallthrough (fallthrough)** | Switch bloğunda bir case sonrası sonraki case’in de çalışmasını sağlar.                               | İkili durumlarda ara çıktılar, birden fazla durumun aynı işlemi yapması gereken senaryolar.                            |
| **Named Returns**             | Fonksiyon geri dönüş değerlerine isim vererek, kod okunabilirliğini ve hata kontrolünü kolaylaştırır. | Matematiksel işlemler, veri dönüşümleri, API sonuçları gibi karmaşık hesaplamalarda açıklık sağlar.                    |

**Önemli Kullanım Senaryoları:**

- **Kontrol deyimleri:** Döngülerin içinde veya switch bloklarında akışı düzenleyerek, hata durumları ve kullanıcı etkileşimlerinde belirli adımların doğru sırayla uygulanmasını sağlar.
- **switch yapısı:** Farklı durumların net bir şekilde ayrıştırılması için idealdir; özellikle sabit değer veya koşul ifadeleriyle yapılan kontrollerde okunabilirlik sunar.
- **goto:** İstisnai durumlarda, yapılandırılmış kontrol akışı sağlanamadığında kullanılır; ancak modern Go kodlamasında tercih edilmez.
- **İsimlendirilmiş geri dönüşler:** Fonksiyonların sonucunun ne anlama geldiğinin net bir şekilde belirtilmesi, hata yönetimini ve kod bakımını kolaylaştırır.
