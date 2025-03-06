## 1. Stdin’den Değer Okuma (Input Reading)

Go dilinde kullanıcıdan veri almak için **fmt** paketindeki üç temel fonksiyon kullanılır:  
**Scan (tarama)**, **Scanf (biçimli tarama)** ve **Scanln (satır sonu tarama)**.  
Bu fonksiyonların kullanım özellikleri ve farkları aşağıdaki tabloda özetlenmiştir:

| **Fonksiyon** | **Açıklama**                                                                                                                         | **Kullanım Durumu**                                                                            | **Özellikler**                                                                                                                                                                         |
| ------------- | ------------------------------------------------------------------------------------------------------------------------------------ | ---------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `fmt.Scan`    | Standart giriş akışından (stdin – standart giriş) verileri, boşluk, tab veya newline gibi whitespace karakterlerine göre ayrıştırır. | Basit veri girişlerinde, veriler arasında boşluk varsa.                                        | Giriş sırasında girilen veriler arasında otomatik ayrıştırma yapılır; her bir veri ilgili değişkene atanır. Hata kontrolü yapılabilir, fakat burada örnekte hata kontrolü atlanmıştır. |
| `fmt.Scanf`   | Belirtilen format (format – biçim) doğrultusunda verileri okur.                                                                      | Giriş verilerinin belirli bir formatta (örn. "name: %s, age: %d") alınması gereken durumlarda. | Kullanıcıdan beklenen veri formatı önceden belirlenir; böylece hatalı girişler için format uyumsuzluğu yakalanabilir.                                                                  |
| `fmt.Scanln`  | Satır sonuna kadar verileri okur; boşluk karakterlerine göre ayrıştırma yapar.                                                       | Genellikle bir tam satırın okunması gereken durumlarda kullanılır.                             | Satır sonuna ulaşıldığında okuma işlemi tamamlanır; sonundaki boşluk karakterleri atlanır.                                                                                             |

**Ek Notlar:**

- Bu fonksiyonlar, kullanıcıdan alınan verileri ilgili değişkenlere aktarmak için **address of** operatörü (`&` – adres operatörü) kullanır.  
  Örneğin: `fmt.Scan(&a, &b)` ifadesinde, `a` ve `b` değişkenlerinin adresleri gönderilir.
- Gerçek uygulamalarda, kullanıcı hatalı veri girişi yapabilir. Bu durumlarda dönen hata (error) değeri kontrol edilerek uygun hata yönetimi (error handling – hata yönetimi) yapılmalıdır. Şimdilik, örneklerde hata kontrolü yapılmamaktadır.

### Örnek Kod

```go
package main

import "fmt"

func main() {
	var a int32
	var b int
	var c float64
	var d bool

	fmt.Print("Input values: ")
	// Kullanıcıdan alınan veriler, boşluk karakterlerine göre ayrıştırılarak ilgili değişkenlere atanır.
	_, err := fmt.Scan(&a, &b, &c, &d)
	if err != nil {
		fmt.Println("Veri okuma hatası:", err)
		return
	}
	fmt.Println("Okunan Değerler:", a, b, c, d)
}
```

**Açıklamalar:**

- `fmt.Scan` fonksiyonu, girilen verileri sırasıyla `a`, `b`, `c` ve `d` değişkenlerine atar.
- Dönüş değeri olarak, okunan öğe sayısı ve hata değeri döndürülür. Örnekte hata kontrolü eklenmiştir.

---

## 2. Fonksiyon Bildirimi ve Tanımlaması (Function Declaration & Definition)

Go’da bir **fonksiyonun bildirimi (declaration – bildirim)**, fonksiyon ismi, parametreleri ve geri dönüş türünü belirtir.  
**Fonksiyon tanımlaması (definition – tanımlama)** ise, bildirimin yanı sıra fonksiyonun gövdesinin (body – gövde) yazılmasıdır.  
Genellikle Go dilinde her iki kavram tek seferde yazılır.

### Detaylı Açıklamalar:

- **Parametreler:** Fonksiyon ismi parantez içerisinde, her bir parametre adı ve türü ile yazılır. Eğer ardışık parametrelerin türü aynı ise, tür sadece bir kez yazılır.
- **Geri Dönüş Türü:** Fonksiyon, tek veya çoklu değer döndürebilir. Çoklu dönüş değerleri parantez içinde, virgülle ayrılarak yazılır.
- **Return Deyimi:** Fonksiyon sonlandığında, `return` deyimi ile geri dönüş değeri belirtilir. Geri dönüş değeri olan fonksiyonlarda, tüm akış yollarında geçerli bir `return` deyimi bulunmalıdır; aksi halde derleme hatası meydana gelir.

### Örnek Kod: Basit Fonksiyon Tanımlaması

```go
package main

import "fmt"

func main() {
	result := sum() * 2
	fmt.Println("result =", result)
}

// sum fonksiyonu, kullanıcıdan iki tam sayı alır ve bu sayıların toplamını geri döndürür.
func sum() int {
	var a, b int
	fmt.Print("İki sayı giriniz: ")
	_, err := fmt.Scan(&a, &b)
	if err != nil {
		fmt.Println("Veri okuma hatası:", err)
		return 0 // Hata durumunda 0 değeri döndürülür.
	}
	// Hesaplama yapıldıktan sonra toplam değeri geri döndürülür.
	return a + b
}
```

**Ek Açıklamalar:**

- Fonksiyon bildirimi ve tanımlaması aynı yerde yapılmaktadır.
- Hata yönetimi eklenerek, veri okuma sırasında hata oluşması durumunda fonksiyonun 0 döndürmesi sağlanmıştır.

---

## 3. Geri Dönüş Değerleri ve Return Deyimi

Fonksiyonların çalışma sonlanırken geri dönüş değeri ile çağrıldığı noktaya bilgi göndermesi, fonksiyonların modüler ve yeniden kullanılabilir olmasını sağlar.  
Go’da her geri dönüş değeri `return` deyimiyle belirtilir.

### İki Kullanım Örneği:

1. **Direkt Geri Dönüş:**

   ```go
   func sum() int {
       var a, b int
       fmt.Print("İki sayı giriniz: ")
       fmt.Scan(&a, &b)
       return a + b
   }
   ```

2. **Değişkene Atama Yapılarak Geri Dönüş:**

   ```go
   func sum() int {
       var a, b int
       fmt.Print("İki sayı giriniz: ")
       fmt.Scan(&a, &b)
       total := a + b
       return total
   }
   ```

**Önemli Noktalar:**

- **Return Deyiminin Konumu:**  
  Fonksiyon içerisindeki tüm mantıksal akış yollarında bir `return` deyimi bulunmalıdır.  
  Örneğin, koşul ifadelerinde (if, switch) bazı durumlar için `return` atlanırsa derleyici hata verir.
- **Erken Çıkış (Early Return):**  
  Fonksiyon içerisinde belirli bir koşul gerçekleştiğinde, geri dönüş değeri olmaksızın fonksiyondan çıkmak için `return` deyimi kullanılabilir.

---

## 4. Çoklu Geri Dönüş Değerleri (Multiple Return Values)

Go dilinin güçlü özelliklerinden biri, fonksiyonların birden fazla değeri aynı anda geri döndürebilmesidir.  
Bu özellik, özellikle hata yönetimi, swapping (değer yer değiştirme) gibi senaryolarda büyük kolaylık sağlar.

### Örnek Kod: İki Değeri Yer Değiştirme (Swapping)

```go
package main

import "fmt"

func main() {
	var a, b int
	fmt.Print("İki sayı giriniz: ")
	fmt.Scan(&a, &b)

	x, y := swapped(a, b)
	fmt.Println("Yer değiştirilmiş değerler -> x =", x, ", y =", y)
}

func swapped(a, b int) (int, int) {
	// Gelen iki değerin yerlerini değiştirerek geri döndürür.
	return b, a
}
```

**Ek Açıklamalar:**

- Fonksiyon tanımında geri dönüş türleri parantez içinde belirtilmiştir.
- Ana fonksiyonda, dönen değerler sırasıyla `x` ve `y` değişkenlerine atanır.

---

## 5. İsimlendirilmiş Geri Dönüş Değerleri (Named Return Values)

Go’da fonksiyonların geri dönüş değerlerine isim verilebilir. Bu kullanım, özellikle fonksiyon içerisinde değerlerin hesaplanması ve sonunda bu isimlendirilmiş değerlerin döndürülmesi işlemini kolaylaştırır.

### Örnek Kod: İsimlendirilmiş Geri Dönüş Değerleri Kullanımı

```go
package main

import "fmt"

func main() {
	var a, b int
	fmt.Print("İki sayı giriniz: ")
	fmt.Scan(&a, &b)

	x, y := swapped(a, b)
	fmt.Println("Yer değiştirilmiş değerler -> x =", x, ", y =", y)
	fmt.Println("Toplam:", sum(a, b))
}

// sum fonksiyonunda geri dönüş değeri 'total' olarak isimlendirilmiştir.
func sum(a, b int) (total int) {
	total = a + b
	// İsimlendirilmiş geri dönüş değeri kullanıldığından, return ifadesinde değişken ismi belirtilmeden döndürme yapılır.
	return
}

func swapped(a, b int) (x int, y int) {
	x = b
	y = a
	return
}
```

**Detaylı Açıklamalar:**

- **İsimlendirilmiş Değişkenlerin Avantajları:**  
  Fonksiyon içinde yapılan hesaplamalarda bu değişkenlere doğrudan atama yapılabilir ve sonrasında `return` deyimi tek başına kullanılarak otomatik olarak geri dönüş yapılır.
- **Okunabilirlik:**  
  Bu yöntem, özellikle fonksiyonun neyi döndürdüğünü açıkça belirtmek istediğiniz durumlarda okunabilirliği artırır.

---

## 6. Fonksiyon Parametreleri ve Argümanlar

Fonksiyonların parantez içerisinde bildirilen değişkenlere **parametreler (parameters – parametreler)** denir;  
Fonksiyon çağrılırken bu parametrelere verilen değerlere ise **argümanlar (arguments – argümanlar)** denir.

### Detaylı Açıklamalar:

- **Aynı Türden Parametreler:**  
  Eğer birden fazla parametre aynı türden ise, tür sadece son parametreden sonra yazılır.

  Örnek:

  ```go
  func sum(a, b int) int {
      return a + b
  }
  ```

- **Fonksiyonun Kapsamı (Scope):**  
  Parametre değişkenleri, fonksiyonun başında tanımlanan yerel değişkenlere benzer şekilde, fonksiyon gövdesi boyunca geçerlidir.

- **Argüman İfadeleri:**  
  Fonksiyon çağrılarında verilen argümanlar önce hesaplanır, ardından fonksiyona iletilir.  
  Örneğin:
  ```go
  fmt.Print(sum(a+1, b*2))
  ```

### Örnek Kod: Parametre Kullanımı

```go
package main

import "fmt"

func main() {
	var a, b int
	fmt.Print("İki sayı giriniz: ")
	fmt.Scan(&a, &b)

	display(a, b)
	fmt.Println(a, "+", b, "=", sum(a, b))
}

func sum(a, b int) int {
	return a + b
}

func display(a int, b int) {
	fmt.Println("a =", a, ", b =", b)
}
```

---

## 7. Parametre Geçiş Yöntemleri: Call by Value ve Call by Reference

Go’da bir fonksiyona parametre geçişi iki farklı yöntemle yapılır:

### Call by Value (Değer Gönderimi)

- **Tanım:**  
  Fonksiyona, değişkenin **kopyası** gönderilir.
- **Özellik:**  
  Fonksiyon içerisinde yapılan değişiklikler orijinal değişkeni etkilemez.

#### Örnek Kod: Call by Value

```go
package main

import "fmt"

func main() {
	a := 10
	foo(a) // a'nın kopyası gönderilir.
	fmt.Println("main: a =", a) // main fonksiyonundaki a değeri değişmez.
}

func foo(a int) {
	a = a + 1
	fmt.Println("foo: a =", a)
}
```

### Call by Reference (Referans Gönderimi)

- **Tanım:**  
  Fonksiyona, değişkenin **adresi** gönderilir; yani, orijinal değerin bellekteki konumu iletildiği için fonksiyon içerisindeki değişiklikler doğrudan orijinal değişken üzerinde etkilidir.
- **Özellik:**  
  `&` operatörü kullanılarak değişkenin adresi gönderilir ve fonksiyon içerisinde pointer (`*`) kullanılarak değere erişilir.

#### Örnek Kod: Call by Reference

```go
package main

import "fmt"

func main() {
	a := 10
	foo(&a) // a'nın adresi gönderilir.
	fmt.Println("main: a =", a) // a, foo fonksiyonundaki değişiklikten etkilenir.
}

func foo(p *int) {
	*p = *p + 1
	fmt.Println("foo: a =", *p)
}
```

**Tablo: Call by Value vs. Call by Reference**

| **Özellik**                           | **Call by Value (Değer Gönderimi)**                                  | **Call by Reference (Referans Gönderimi)**                    |
| ------------------------------------- | -------------------------------------------------------------------- | ------------------------------------------------------------- |
| **Gönderilen Değer**                  | Değişkenin kopyası                                                   | Değişkenin adresi (`&` operatörü kullanılarak)                |
| **Orijinal Değişken Üzerindeki Etki** | Fonksiyon içerisinde yapılan değişiklik orijinal değişkene yansımaz  | Fonksiyon içerisindeki değişiklik, orijinal değişkeni etkiler |
| **Bellek Kullanımı**                  | Kopyalama işlemi yapılır (büyük veri yapılarında maliyetli olabilir) | Bellek adresi gönderildiği için kopyalama maliyeti düşüktür   |

---

## 8. Math Paketi ve Matematiksel İşlemler

Go dilinde matematiksel işlemler için **math** paketi geniş bir fonksiyon seti sunar.  
Sık kullanılan fonksiyonlardan bazıları şunlardır:

- **Sqrt (karekök – square root):** Bir sayının karekökünü alır.
- **Pow (üst alma – power):** Bir sayının üstel değerini hesaplar.
- **Log (logaritma – logarithm):** Doğal logaritmayı hesaplar.
- **Ceil (yukarı yuvarlama – ceiling):** Bir sayıyı, kendisinden büyük veya eşit en yakın tam sayıya yuvarlar.
- **Floor (aşağı yuvarlama – floor):** Bir sayıyı, kendisinden küçük veya eşit en yakın tam sayıya yuvarlar.
- **Round (yuvarlama – round):** Standart yuvarlama işlemini gerçekleştirir.
- **Max, Min (maksimum, minimum):** İki sayı arasından en büyük veya en küçüğü seçer.

### Sınıf Çalışması: Euclidean (Öklid) Uzaklığı Hesaplama

**Görev:**  
Parametreleri `float64` türünde olan iki noktanın koordinatlarını (x1, y1 ve x2, y2) alarak aralarındaki Euclidean (Öklid) uzaklığını hesaplayan bir fonksiyon yazınız.

#### Örnek Kod: Euclidean Distance

```go
package main

import (
	"fmt"
	"math"
)

func main() {
	distanceTest()
}

func distanceTest() {
	fmt.Print("Input coordinates (x1 y1 x2 y2): ")
	var x1, y1, x2, y2 float64
	_, err := fmt.Scan(&x1, &y1, &x2, &y2)
	if err != nil {
		fmt.Println("Veri okuma hatası:", err)
		return
	}
	fmt.Println("Euclidean Distance:", distance(x1, y1, x2, y2))
}

func distance(x1, y1, x2, y2 float64) float64 {
	// Euclidean distance = sqrt((x1-x2)^2 + (y1-y2)^2)
	return math.Sqrt(math.Pow(x1-x2, 2) + math.Pow(y1-y2, 2))
}
```

**Detaylı Açıklamalar:**

- `math.Pow` fonksiyonu ile (x1 - x2) ve (y1 - y2) farklarının karesi hesaplanır.
- `math.Sqrt` fonksiyonu, karesini alınan toplamın karekökünü bularak iki nokta arasındaki Öklid uzaklığını verir.
- Hata kontrolü ile, kullanıcı hatalı değer girişi yaptığında uygun mesaj gösterilebilir.

---

## 9. Pointer (Gösterici) Kullanımının Detayları

Go’da pointer’lar, bir değişkenin bellek adresini tutar.  
**`&` Operatörü:**

- Bir değişkenin adresini almak için kullanılır.
- Örneğin, `&a` ifadesi, `a` değişkeninin bellek adresini döndürür.

**`*` Operatörü (Dereference – gösterilen değeri alma):**

- Pointer’ın gösterdiği adresteki değeri almak için kullanılır.
- Örneğin, `*p` ifadesi, `p` pointer’ının işaret ettiği değeri döndürür.

**Örnek Kod: Pointer Kullanımı**

```go
package main

import "fmt"

func main() {
	a := 100
	p := &a // p, a'nın adresini tutar.
	fmt.Println("a'nın adresi:", p)
	fmt.Println("p'nin işaret ettiği değer:", *p)
}
```

**Ek Açıklamalar:**

- Pointer kullanımı, özellikle call by reference (referans ile geçiş) durumlarında, fonksiyonların orijinal veriler üzerinde değişiklik yapabilmesini sağlar.
- Pointer’lar, büyük veri yapılarının kopyalanmasının önüne geçerek bellek verimliliğini artırır.

---

## Sonuç ve Özet

Yukarıda detaylı olarak Go dilinde kullanıcıdan veri alma, fonksiyon bildirimleri ve tanımlamaları, geri dönüş değerleri, çoklu dönüş değerleri, isimlendirilmiş dönüş değerleri, parametre geçişi (call by value ve call by reference) ile pointer kullanımının yanı sıra matematiksel işlemler için math paketinin kullanımı ele alınmıştır.

**Önemli Noktalar:**

- **Stdin’den Veri Okuma:**

  - `fmt.Scan`, `fmt.Scanf` ve `fmt.Scanln` fonksiyonları, farklı giriş senaryoları için kullanılabilir.
  - Hata yönetimi eklenmesi, güvenli ve sağlam bir uygulama geliştirme süreci sağlar.

- **Fonksiyon Bildirimi ve Tanımlaması:**

  - Parametreler ve geri dönüş türleri doğru biçimde bildirilmelidir.
  - Her akış yolunda geçerli bir `return` deyiminin bulunması gerekmektedir.

- **Çoklu Geri Dönüş Değerleri:**

  - Fonksiyonlar, birden fazla değeri aynı anda döndürebilir.
  - İsimlendirilmiş dönüş değerleri okunabilirliği ve yönetimi kolaylaştırır.

- **Parametre Geçişi:**

  - Call by value, değerin kopyasını gönderirken; call by reference, adres gönderimi yaparak orijinal veriyi etkiler.
  - Pointer’lar, bu geçiş türlerinde kritik rol oynar.

- **Math Paketi ve Uygulamaları:**
  - Matematiksel hesaplamalar, `math.Sqrt`, `math.Pow` gibi fonksiyonlarla gerçekleştirilir.
  - Euclidean distance örneği, iki nokta arasındaki mesafenin hesaplanması açısından somut bir uygulamadır.
