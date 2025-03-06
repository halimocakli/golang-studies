## 1. Tür (Type – tür) Kavramı ve Kategorileri

**Tanım:**  
Bir **type (tür – tür)**, bir değişkenin içerisindeki değerin hangi formatta saklanacağını ve bellekte ne kadar yer kaplayacağını belirler. Go dilinde türler, bellekte verimli yer kullanımını sağlamak ve kodun güvenliğini artırmak amacıyla farklı kategorilere ayrılır.

**Alt Kategoriler:**

- **Boolean types (boolean türleri – boolean türleri):**  
  Mantıksal değerler (true/false) için kullanılır.
- **Numeric types (sayısal türler – sayısal türler):**  
  Tamsayılar (integer/integral types – tamsayı türleri) ve gerçek sayılar (floating point types – kayan nokta türleri) ile karmaşık sayıları (complex types – karmaşık türler) kapsar.
- **String types (string türleri – string türleri):**  
  Metin verilerini saklamak için kullanılır.
- **Derived types (türetilmiş türler – türetilmiş türler):**  
  Array, slice, map, struct gibi veri yapılarını içerir.

**Ek Bilgiler:**

- **Pointer types (pointer türleri – işaretçi türleri):**  
  Pointer’lar, başka bir değişkenin bellekteki adresini saklar. Go’da pointer kullanımı, bellek yönetimini kolaylaştırır; ancak pointer aritmetiği desteklenmez.
- **Struct (struct – yapı) ve Interface (interface – arayüz) Türleri:**  
  Yapılar, verileri gruplamak için kullanılırken; arayüzler, davranışların (method – method) soyutlanmasını sağlar.
- **Fonksiyon Türleri (function types – fonksiyon türleri):**  
  Fonksiyonlar da birer değerdir ve değişkenlere atanabilir, parametre veya dönüş değeri olarak kullanılabilir.

---

## 2. Sayısal (Numeric – sayısal) Türler ve Özellikleri

Sayısal türler iki ana gruba ayrılır:

- **Tamsayı Türleri (Integer / Integral Types – tamsayı türleri):**
  - İşaretli (signed) ve işaretsiz (unsigned) türler bulunur.
  - İkiye tümleme (two's complement – ikiye tümleme) yöntemi kullanılarak saklanırlar.
  - Taşma (overflow – taşma) durumlarına karşı derleyici kontrolü bulunmaz; bu yüzden hesaplamalarda dikkat gerektirir.
- **Gerçek Sayı Türleri (Floating Point Types – gerçek sayı türleri):**
  - IEEE754 standardını temel alır.
  - Hesaplamalarda yuvarlama hataları ve kesinlik sorunları yaşanabilir; özellikle bilimsel hesaplamalarda dikkatli olunmalıdır.
  - **complex64 (complex64 – complex64):** Gerçek (real – gerçek) ve sanal (imaginary – sanal) kısımları `float32` türünden olan karmaşık sayıları temsil eder.
  - **complex128 (complex128 – complex128):** Gerçek (real – gerçek) ve sanal (imaginary – sanal) kısımları `float64` türünden olan karmaşık sayıları temsil eder.

Aşağıdaki tablo, Go’daki temel sayısal türlerin byte (bayt) uzunluklarını özetlemektedir:

| **Tür İsmi** | **Byte (Bayt)** | **Açıklama**                                                            |
| ------------ | --------------- | ----------------------------------------------------------------------- |
| int8         | 1               | 8-bit işaretli tamsayı                                                  |
| uint8        | 1               | 8-bit işaretsiz tamsayı                                                 |
| int16        | 2               | 16-bit işaretli tamsayı                                                 |
| uint16       | 2               | 16-bit işaretsiz tamsayı                                                |
| int32        | 4               | 32-bit işaretli tamsayı                                                 |
| uint32       | 4               | 32-bit işaretsiz tamsayı                                                |
| int64        | 8               | 64-bit işaretli tamsayı                                                 |
| uint64       | 8               | 64-bit işaretsiz tamsayı                                                |
| byte         | 1               | `uint8` için alias (takma ad – alias)                                   |
| rune         | 4               | `int32` için alias (takma ad – alias); Unicode karakterleri temsil eder |
| int          | 4/8             | İşletim sistemine bağlı olarak 4 veya 8 byte                            |
| uint         | 4/8             | İşletim sistemine bağlı olarak 4 veya 8 byte                            |
| float32      | 4               | 32-bit kayan noktalı sayı                                               |
| float64      | 8               | 64-bit kayan noktalı sayı                                               |
| complex64    | 8               | Gerçek ve sanal kısımları `float32` olan karmaşık sayı                  |
| complex128   | 16              | Gerçek ve sanal kısımları `float64` olan karmaşık sayı                  |
| bool         | 1               | Boolean (mantıksal – boolean) tür, true veya false değer alır           |

**Ek Örnek:**  
Aşağıdaki örnekte, tamsayı ve kayan nokta hesaplamaları ile birlikte taşma durumuna dikkat çekilmiştir:

```go
package main

import "fmt"

func main() {
	var a uint8 = 250
	var b uint8 = 10
	// a + b işlemi, uint8 sınırını aşarsa taşma meydana gelebilir.
	sum := a + b
	fmt.Println("uint8 toplamı:", sum) // Taşma örneği: beklenenden düşük değer verebilir

	var x float64 = 0.1
	var y float64 = 0.2
	result := x + y
	fmt.Println("float64 toplamı:", result) // Kayan nokta hesaplamalarında kesinlik farkları gözlemlenebilir
}
```

---

## 3. İfade (Expression – ifade) Kavramı

**Tanım:**  
İfade, sabitler, değişkenler ve operatörlerin (operators – operatörler) oluşturduğu dizilimlerdir. İfadeler, hesaplamaların sonucunu belirler ve programın akışında önemli rol oynar.

**Örnek İfade Türleri:**

- **Aritmetik İfadeler:**  
  Örneğin: `3 + 4 * 2`
- **Mantıksal (Boolean) İfadeler:**  
  Örneğin: `a > b && b != 0`
- **İşlev Çağrıları:**  
  Örneğin: `fmt.Println("Merhaba, Go!")`
- **Karmaşık İfadeler:**  
  Parantez kullanarak öncelik belirleme, fonksiyon çağrıları ve değişkenlerin birleşimi:  
  `result := (a + b) * (c - d)`

**Ek Örnek:**  
Aşağıdaki kod parçası, farklı ifade türlerini kullanarak koşullu bir işlem yapmaktadır:

```go
package main

import "fmt"

func main() {
	a, b := 10, 20
	// Karmaşık ifade: koşul ifadesi ve aritmetik işlem
	max := a
	if b > a {
		max = b
	}
	fmt.Println("Büyük olan:", max)
}
```

---

## 4. Değişken (Variable – değişken) Kavramları

**Temel Kavramlar:**

- **İsim (Name – isim):**  
  Değişkenin adını belirler. Kurallara uygun (rakamla başlamama, case-sensitive vb.) yazılmalıdır.
- **Tür (Type – tür):**  
  Değişkenin bellekte hangi formatta saklanacağını belirler.
- **Faaliyet Alanı (Scope – faaliyet alanı):**  
  Değişkenin kod içinde erişilebilir olduğu bölgedir.
- **Ömür (Storage Duration – ömür):**  
  Değişkenin bellekte yaratılmasından yok edilmesine kadar geçen süreyi ifade eder.

**Sıfır Değerler Tablosu:**  
Go, bildirilen değişkenlere otomatik olarak sıfır (zero) değeri atar. Aşağıdaki tablo, yaygın türlerin varsayılan sıfır değerlerini göstermektedir:

| **Tür**                                 | **Sıfır Değeri** |
| --------------------------------------- | ---------------- |
| bool                                    | false            |
| int, int8, int16, int32, int64          | 0                |
| uint, uint8, uint16, uint32, uint64     | 0                |
| float32, float64                        | 0                |
| complex64, complex128                   | 0+0i             |
| string                                  | "" (boş string)  |
| pointer, slice, map, channel, interface | nil              |

**Ek Örnek:**  
Type inference (tür çıkarımı) mekanizmasını vurgulamak amacıyla ek bir örnek:

```go
package main

import "fmt"

func main() {
	// Tür çıkarımı sayesinde, değişkenin türü otomatik belirlenir.
	num := 100          // num int tipindedir.
	str := "Go Öğren"   // str string tipindedir.
	flag := true        // flag bool tipindedir.
	fmt.Println(num, str, flag)
}
```

---

## 5. Nesne (Object – nesne), lvalue ve rvalue

**Nesne (Object – nesne):**  
Bellekte yer ayrılmış olan veri birimidir. Bir değişken tanımlandığında bellekte ona karşılık gelen bir nesne oluşturulur. Nesnelerin ömrü, garbage collection (çöp toplama – garbage collection) mekanizması ile yönetilir.

**lvalue (Sol Taraf Değeri – lvalue):**  
Bellekte belirli bir adresi olan, değişken gibi atanabilir ifadelerdir. Atama operatörünün sol tarafında yer alır.  
**rvalue (Sağ Taraf Değeri – rvalue):**  
Sadece değeri temsil eden, geçici ifadelerdir. rvalue ifadeler, doğrudan atama operatörünün sol tarafında kullanılamaz.

**Ek Örnek:**  
Aşağıdaki kod, lvalue ve rvalue arasındaki farkı göstermektedir:

```go
package main

import "fmt"

func main() {
	var x int = 5   // 'x' bir lvalue’dur.
	y := 10         // y de bir lvalue’dur.
	x = y           // Burada y’nin rvalue’si x'e atanır.
	// 5 = x       // Bu kullanım geçersizdir çünkü 5 bir rvalue'dir.
	fmt.Println("x:", x)
}
```

---

## 6. Değişken Initialization (İlk Değer Atama) ve Bildirim Yöntemleri

Go dilinde değişken bildirimi iki temel yöntemle yapılır:

1. **var Anahtar Sözcüğü ile Bildirim:**

   - Tür belirtilirse, ilk değer verilmediğinde otomatik olarak sıfır değeri atanır.
   - İlk değer verildiğinde, tür otomatik tespit edilir (type inference/deduction – tür çıkarımı).

   **Örnek:**

   ```go
   package main

   import "fmt"

   func main() {
       var a int       // 'a' için sıfır değeri (0) atanır
       var b = 15      // Tür otomatik olarak int olarak belirlenir
       fmt.Println("a:", a, "b:", b)
   }
   ```

2. **:= Operatörü ile Kısa Bildirim (Immediate Declaration – anında bildirim):**

   - Hem bildirim hem de ilk değer atama aynı anda yapılır.
   - Tür, sağlanan değere göre otomatik tespit edilir.

   **Örnek:**

   ```go
   package main

   import "fmt"

   func main() {
       a := 25        // 'a' bildirimi ve 25 değeri atanır
       b := a         // b, a'nın değerini alır
       fmt.Println("a:", a, "b:", b)
   }
   ```

3. **Çoklu Değişken Bildirimi:**

   - Aynı satırda, virgül kullanılarak aynı türde ya da karışık türlerde değişkenler bildirilebilir.

   **Örnek (Aynı Tür):**

   ```go
   package main

   import "fmt"

   func main() {
       var x, y, z int = 5, 10, 15
       fmt.Println(x, y, z)
   }
   ```

   **Örnek (Mixed Declaration – karışık türler):**

   ```go
   package main

   import "fmt"

   func main() {
       var a, b, c = 42, 3.14, "Go"
       fmt.Println(a, b, c)
   }
   ```

**Ek Bilgiler:**

- Değişken bildirimi sırasında, derleyici türü otomatik çıkarır ve bu sayede kod okunabilirliği artar.
- Bildirilen fakat kullanılmayan değişkenler hata (error) oluşturur.

---

## 7. Değişken İsimlendirme Kuralları

Go dilinde değişken isimleri için uyulması gereken temel kurallar şunlardır:

- **Başlangıç Karakteri:**
  - Değişken isimleri rakam ile başlayamaz; yalnızca alfabetik karakter veya alttire (underscore – alt çizgi) karakteri ile başlamalıdır.
- **Devamı:**
  - İsim, istenildiği kadar rakam ve harf içerebilir.
- **Case-Sensitive (Büyük/Küçük Harf Duyarlılığı – case sensitive):**
  - Değişken isimleri büyük/küçük harf duyarlıdır; örneğin, `deger` ile `Deger` farklı değişkenlerdir.
- **Unicode Desteği:**
  - Unicode karakterler kullanılabilir; ancak ASCII dışı karakterlerin kullanımı önerilmez.
- **İsimlendirme Konvansiyonları:**
  - Yerel değişkenler için genellikle camelCase (örneğin, `kullaniciAdi`) kullanılırken; paket dışı (exported) değişkenler için PascalCase (örneğin, `KullaniciAdi`) tercih edilir.
  - Gereksiz kısaltmalardan kaçınılmalı, anlamlı ve açıklayıcı isimler kullanılmalıdır.

---

## 8. Bloklar (Block – blok) ve Yerel Değişkenler (Local Variables – yerel değişkenler)

**Blok (Block – blok):**

- Bir fonksiyonun `{` ve `}` arasındaki kısmı bir bloktur.
- İç içe bloklar oluşturulabilir ve her blok kendi yerel değişkenlerini barındırır.

**Yerel Değişkenler (Local Variables – yerel değişkenler):**

- Bir blok içerisinde bildirilen değişkenler yerel değişkenlerdir.
- Bu değişkenlerin ömrü, bulundukları blokla sınırlıdır ve blok dışına çıkıldığında erişilemez hale gelir.

**Ek Örnek:**  
İç içe bloklarda yerel değişkenlerin nasıl çalıştığını gösteren ek bir örnek:

```go
package main

import "fmt"

func main() {
	{
		// İlk blok: 'a' yerel değişkeni oluşturulur.
		var a int = 100
		fmt.Println("İlk blok, a:", a)
		{
			// İç blok: 'b' adında yeni yerel değişken oluşturulur.
			var b int = 200
			fmt.Println("İç blok, b:", b)
		}
		// 'b' burada erişilemez.
	}
	// 'a' burada da erişilemez.
}
```

---

## 9. Shadowing (Masking – gölgeleme) Durumu

**Tanım:**  
İçiçe bloklarda aynı isimde yerel değişken bildirildiğinde, iç bloktaki değişken dış bloktaki değişkenin üzerine yazar (masking/shadowing). Bu durum, dış bloktaki değere erişimin kaybolmasına sebep olabilir.

**Ek Bilgiler:**

- Shadowing, özellikle büyük kod tabanlarında yanlış anlaşılmalara yol açabileceğinden dikkatli kullanılmalıdır.
- İsimlerin tekrarlanması durumunda, kod okunabilirliğini artırmak için farklı isimlendirme stratejileri (örneğin, daha açıklayıcı isimler) tercih edilebilir.

**Örnek:**

```go
package main

import "fmt"

func main() {
	var value int = 50
	fmt.Println("Dış blok value:", value)

	{
		// Shadowing: 'value' isminde yeni bir değişken oluşturuluyor.
		var value int = 100
		value += 20
		fmt.Println("İç blok value:", value)
	}

	// Dış bloktaki 'value' etkilenmemiştir.
	fmt.Println("Dış blok, tekrar value:", value)
}
```

---

## 10. Global Değişkenler (Global Variables – global değişkenler)

**Tanım:**  
Fonksiyonların dışında, paket içerisinde bildirilen değişkenler global değişkenlerdir. Aynı dosya içindeki tüm fonksiyonlar bu değişkenlere erişebilir. Global değişkenlere de sıfır değeri otomatik olarak atanır.

**Ek Bilgiler:**

- Global değişkenler, programın genel durumunu tutmak için kullanılır ancak aşırı kullanımı programın yönetilebilirliğini zorlaştırabilir.
- İyi tasarlanmış bir mimaride, global değişkenler minimum düzeyde tutulmalı ve mümkünse fonksiyonlara parametre olarak aktarılmalıdır.

**Örnek:**

```go
package main

import "fmt"

// Global değişken bildirimi
var globalCount int = 10

func main() {
	fmt.Println("Başlangıç globalCount:", globalCount)
	increase()
	fmt.Println("increase() sonrası globalCount:", globalCount)
	reset()
	fmt.Println("reset() sonrası globalCount:", globalCount)
}

func increase() {
	globalCount += 5
}

func reset() {
	globalCount = 0
}
```

---

## 11. Özet ve Sonuç

**Go (Golang) programlama dilinde temel kavramlar şu şekilde özetlenebilir:**

- **Type (tür – tür):**  
  Değişkenlerin bellekte hangi formatta ve ne kadar yer kaplayacağını belirler. Sayısal türler, boolean, string ve türetilmiş türler gibi kategorilere ayrılır. Ek olarak, pointer, struct, interface ve fonksiyon türleri de dilin esnekliğini artırır.

- **Expression (ifade – ifade):**  
  Sabitler, operatörler ve değişkenlerin birleşiminden oluşur. İfadeler, koşul ve döngü yapıları gibi kontrol akışlarında kritik rol oynar.

- **Variable (değişken – değişken):**  
  İsim, tür, scope (faaliyet alanı) ve storage duration (ömür) gibi kavramları içerir.

  - **Bildirim Yöntemleri:** `var` bildirimi, kısa bildirim (`:=`) ve çoklu değişken bildirimi.
  - **Sıfır Değerler:** Her değişken bildirildiğinde, türüne uygun sıfır değeri otomatik atanır.
  - **Type Inference:** Tür çıkarımı sayesinde, kod daha okunabilir ve yazımı kolaylaşır.

- **Blok (block – blok) ve Scope (faaliyet alanı):**  
  Bloklar, fonksiyon gövdeleri ve iç içe yapılar şeklinde tanımlanır. Yerel değişkenlerin ömrü, bulundukları blokla sınırlıdır.

  - **Shadowing (gölgeleme):** Aynı isimde değişken bildirimi durumunda iç bloktaki değişken, dış bloktaki değişkeni maskeleyebilir.

- **Global Variables (global değişkenler – global değişkenler):**  
  Paket genelinde erişilebilen değişkenlerdir. Global değişkenler, programın durumunu merkezi bir şekilde yönetmede kullanılır; ancak dikkatli kullanılmalıdır.

**Genel Değerlendirme:**  
Go, güçlü tip güvenliği, otomatik bellek yönetimi ve etkili eşzamanlılık (concurrency – eşzamanlılık) desteği sayesinde sistem programlama, ağ uygulamaları, CLI araçları ve mikroservis mimarileri gibi pek çok alanda verimli ve güvenilir çözümler sunar. Eklenen detaylar; türler, ifadeler, değişken bildirimi, isimlendirme kuralları, blok yapıları ve global değişkenlerin doğru kullanımı konularında derinlemesine bilgi sağlamaktadır. Bu bilgiler, Go dilinin temel özelliklerini anlamak ve pratik uygulamalarda karşılaşılabilecek durumlara hazırlıklı olmak açısından büyük önem taşır.
