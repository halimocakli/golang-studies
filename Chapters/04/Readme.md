## 1. Yuvarlama Fonksiyonları: **Round**, **Floor** ve **Ceil**

_(Round – Yuvarlama, Floor – Aşağı Yuvarlama, Ceil – Yukarı Yuvarlama)_

**Açıklama:**

- **math.Round(a):**

  - Girilen ondalıklı (float64) sayıyı en yakın tam sayıya yuvarlar.
  - Yarı noktadaki değerler (örneğin, 2.5, 3.5) matematiksel olarak en yakın çift tam sayıya veya yukarı doğru yuvarlanarak sonuç elde edilebilir.
  - Finansal uygulamalarda, grafik hesaplamalarında veya istatistiksel işlemlerde kullanılabilir.

- **math.Floor(a):**

  - Verilen değerden küçük veya ona eşit olan en büyük tam sayıyı döndürür.
  - Bu fonksiyon, özellikle negatif sayılarla çalışırken dikkat gerektirir; çünkü örneğin, -3.2 değeri için sonuç -4 olacaktır.
  - Sayısal hesaplamalarda taban değeri belirlemek için tercih edilir.

- **math.Ceil(a):**
  - Verilen değerden büyük veya ona eşit olan en küçük tam sayıyı döndürür.
  - Özellikle negatif sayılar için, -3.2 değeri örneğinde sonuç -3 olacaktır.
  - Katman sayılarının belirlenmesi, bellek veya kaynak tahsisi gibi durumlarda kullanılabilir.

**Detaylı Örnek:**

```go
package main

import (
	"fmt"
	"math"
)

func main() {
	var a float64

	fmt.Print("Bir ondalıklı sayı giriniz: ")
	fmt.Scan(&a)

	fmt.Println("Round (Yuvarlama):", math.Round(a))
	fmt.Println("Floor (Aşağı Yuvarlama):", math.Floor(a))
	fmt.Println("Ceil (Yukarı Yuvarlama):", math.Ceil(a))
}
```

**Önemli Noktalar:**

- Kullanıcının girdiği değerin pozitif veya negatif olması, fonksiyonların çıktısını etkiler.
- Finansal hesaplamalarda doğru yuvarlama yönteminin seçimi önemlidir.
- Grafik hesaplamalarında pikselleştirme işlemlerinde doğru kesme/yuvarlama kritik rol oynar.

**Tablo: Yuvarlama Fonksiyonlarının Karşılaştırılması**

| Fonksiyon       | Açıklama                                       | Örnek (a = 3.7) | Örnek (a = -3.7) |
| --------------- | ---------------------------------------------- | --------------- | ---------------- |
| `math.Round(a)` | En yakın tam sayıya yuvarlar                   | 4               | -4               |
| `math.Floor(a)` | a değerinden küçük veya eşit en büyük tam sayı | 3               | -4               |
| `math.Ceil(a)`  | a değerinden büyük veya eşit en küçük tam sayı | 4               | -3               |

---

## 2. IEEE 754 Özel Değerler ve Gerçek Sayı Fonksiyonları

_(IEEE 754 – IEEE 754, NaN – Tanımsız, Infinity – Sonsuzluk)_

**Açıklama:**

- **IEEE 754 Standardı:**
  - Bilgisayarların gerçek sayıları nasıl temsil ettiğini belirleyen uluslararası bir standarttır.
  - Bu standart, sayıların yaklaşık değerlerini saklarken karşılaşılabilecek belirsizlikleri ve aşırı büyük/küçük değerleri yönetir.
- **Özel Değerler:**
  - **NaN (Not a Number – Tanımsız):**
    - Hesaplama sonucu matematiksel olarak tanımsız veya geçersiz olduğunda ortaya çıkar.
    - Örneğin, negatif bir sayının karekökünü almaya çalışmak `NaN` üretir.
  - **+Inf ve -Inf (Sonsuzluk Değerleri):**
    - Pay sıfır olmayan bir sayının sıfıra bölünmesi durumunda; pay pozitifse `+Inf`, negatifse `-Inf` döner.
  - Bu değerler, hata fırlatmak yerine işlem devam ettirilmesine olanak sağlar.

**Detaylı Örnek:**

```go
package main

import (
	"fmt"
	"math"
)

func main() {
	var a float64

	fmt.Print("Bir gerçek sayı giriniz: ")
	fmt.Scan(&a)

	// Negatif değerler için karekök hesaplaması tanımsızlık durumunu ortaya koyabilir.
	fmt.Println("Sqrt(", a, ") =", math.Sqrt(a))
	// Logaritma işlemi, sıfır veya negatif değerlerde farklı özel değerler üretebilir.
	fmt.Println("Log(", a, ") =", math.Log(a))
}
```

**Önemli Noktalar:**

- **Hata Yönetimi:**

  - Gerçek uygulamalarda, bu özel değerler (NaN, +Inf, -Inf) kontrol edilerek uygun hata yönetimi sağlanmalıdır.
  - Örneğin, kullanıcıdan alınan girdi negatifse, `math.Sqrt` çağrılmadan önce bir kontrol eklemek gerekebilir.

- **Uygulama Alanları:**
  - Bilimsel hesaplamalar, mühendislik uygulamaları ve finansal modellerde, bu özel durumlar sıklıkla karşılaşılır.

---

## 3. Bölme İşlemleri: Tamsayılar ve Gerçek Sayılar

_(Division – Bölme)_

### 3.1. Tamsayı Bölmesi

_(Integer Division – Tamsayı Bölmesi)_

**Açıklama:**

- Tamsayılar arasında yapılan bölme işlemlerinde, ondalık kısım atılır; sonuç tam sayı kısmı elde edilir.
- Bu işlem **truncation (kesme)** olarak adlandırılır.
- Örneğin, 7 / 3 işlemi 2 sonucunu verir, çünkü 7/3 = 2.333… kısmı atılır.

**Detaylı Örnek:**

```go
package main

import (
	"fmt"
)

func main() {
	var a, b int

	fmt.Print("İki tam sayı giriniz (örn: 7 3): ")
	fmt.Scan(&a, &b)

	c := a / b

	fmt.Println("Bölme sonucu c =", c)
}
```

**Önemli Noktalar:**

- Negatif sayılarda sonuç sıfır noktasına göre kesilir; örneğin, -7 / 3 sonucu -2 olacaktır.
- Tamsayı bölmesi, döngülerde indeks hesaplamalarında ve bellek yönetiminde sıkça kullanılır.

### 3.2. Sıfıra Bölme Durumları

#### a) Tamsayılarla Bölme

- **Sorun:**
  - Eğer bölen (b) sıfır ise, derleme zamanı değil çalışma zamanında (runtime) hata meydana gelir.
- **Çözüm:**
  - Bölme işlemi gerçekleştirilmeden önce bölenin sıfır olup olmadığı kontrol edilmelidir.

**Detaylı Örnek:**

```go
package main

import (
	"fmt"
	"os"
)

func main() {
	var a, b int

	fmt.Print("İki tam sayı giriniz: ")
	fmt.Scan(&a, &b)

	if b == 0 {
		fmt.Println("Hata: Sıfıra bölme yapılmaz!")
		os.Exit(1)
	}

	c := a / b
	fmt.Println("Bölme sonucu c =", c)
}
```

#### b) Gerçek Sayılarla Bölme

_(Floating-Point Division – Gerçek Sayı Bölmesi)_

- **Özellik:**
  - Gerçek sayıların bölünmesi sırasında, bölen sıfır olduğunda runtime hatası yerine IEEE 754 standardına göre özel değerler üretilir:
    - Eğer pay (a) sıfır ise sonuç `NaN` (tanımsız) döner.
    - Eğer pay pozitif ise `+Inf` (pozitif sonsuzluk) üretilir.
    - Eğer pay negatif ise `-Inf` (negatif sonsuzluk) üretilir.
- **Avantaj:**
  - Programın aniden durmasını engeller; hata kontrolü daha sonra yapılabilir.

**Detaylı Örnek:**

```go
package main

import (
	"fmt"
)

func main() {
	var a, b float64

	fmt.Print("İki gerçek sayı giriniz (örn: 5.0 0): ")
	fmt.Scan(&a, &b)

	c := a / b
	fmt.Println("Bölme sonucu c =", c)
}
```

**Önemli Noktalar:**

- Gerçek sayı bölmeleri, özellikle matematiksel hesaplamalarda ve simülasyonlarda hata yönetimi için dikkatle ele alınmalıdır.

---

## 4. **fmt** Paketi ve Formatlama (Formatting) İşlemleri

_(Printf – Biçimlendirme, Scanf – Biçimli Giriş)_

### 4.1. Printf Fonksiyonu ile Formatlama

**Açıklama:**

- **fmt.Printf** fonksiyonu, standart çıktıya (stdout) biçimlendirilmiş metin yazdırmak için kullanılır.
- İlk parametre, format string (biçim dizesi) olup, içerisine format belirteçleri (format specifiers) yerleştirilir.
- Örneğin:
  - `%T`: Değişkenin türünü gösterir.  
    _(Type – Tür)_
  - `%d`: Tamsayıları ondalık biçimde gösterir.  
    _(Decimal – Ondalık)_
  - `%f`: Gerçek sayıları sabit ondalık biçimde gösterir.  
    _(Floating – Gerçek)_
  - `%t`: Boolean (mantıksal) değerleri gösterir.  
    _(Boolean – Mantıksal)_
  - `%v`: Değişkenin varsayılan gösterimini yapar.

**Detaylı Örnek:**

```go
package main

import (
	"fmt"
)

func main() {
	var a float64 = 3.14159
	var b int = 42
	var c bool = true

	fmt.Printf("a'nın türü %T ve değeri %f\n", a, a)
	fmt.Printf("b'nin türü %T ve değeri %d\n", b, b)
	fmt.Printf("c'nin türü %T ve değeri %t\n", c, c)
}
```

**Ek Açıklamalar:**

- **Format Belirteçlerinin Esnekliği:**
  - `%v` belirteci, herhangi bir türdeki değeri en genel biçimde yazdırmak için oldukça kullanışlıdır.
  - Karmaşık veri yapılarında (örneğin, struct'lar) `%+v` kullanılarak alan isimleriyle birlikte çıktı alınabilir.

### 4.2. Hizalama ve Kesinlik (Precision) Ayarlamaları

**Açıklama:**

- Biçim dizesinde, sayılar için hizalama ve belirli basamak sayısını ayarlamak mümkündür:
  - `%02d`: Tamsayıyı en az iki basamak olacak şekilde, eksik basamakları sıfır ile doldurur.
  - `%.10f`: Gerçek sayının noktadan sonra 10 basamağını gösterir. Gerekli durumlarda bilimsel gösterime geçebilir.

**Detaylı Örnek:**

```go
package main

import "fmt"

func main() {
	var day, month, year int

	fmt.Print("Tarih değerlerini giriniz (gün ay yıl): ")
	fmt.Scan(&day, &month, &year)

	// Tarih, DD.MM.YYYY formatında biçimlendirilir.
	fmt.Printf("Biçimlendirilmiş Tarih: %02d.%02d.%04d\n", day, month, year)
}
```

**Ek Açıklamalar:**

- Bu tür biçimlendirme, raporlama, kullanıcı arayüzü ve dosya çıktılarında okunabilirliği artırır.
- Farklı format belirteçleri kombinasyonları ile karmaşık çıktılar üretilebilir.

### 4.3. Farklı Sayı Sistemlerinde Formatlama

**Açıklama:**

- Tamsayılar, farklı sayı sistemlerinde (base/radix) formatlanabilir:
  - `%d`: Ondalık (decimal – ondalık) sistem.
  - `%x` veya `%X`: Onaltılık (hexadecimal – onaltılık) sistem; `%X` büyük harflerle.
  - `%b`: İkili (binary – binary) sistem.
  - `%o` veya `%O`: Sekizlik (octal – sekizlik) sistem; `%O` kullanıldığında sayının önüne `0` eklenebilir.

**Detaylı Örnek:**

```go
package main

import "fmt"

func main() {
	var val int

	fmt.Print("Bir tam sayı giriniz: ")
	fmt.Scan(&val)

	fmt.Printf("Decimal (ondalık) format: %d\n", val)
	fmt.Printf("Hexadecimal (onaltılık) format: %x\n", val)
	fmt.Printf("Hexadecimal (büyük harf) format: %X\n", val)
	fmt.Printf("Binary (ikili) format: %b\n", val)
	fmt.Printf("Octal (sekizlik) format: %o\n", val)
	fmt.Printf("Octal (ön ekli) format: %O\n", val)
}
```

**Tablo: Printf Format Specifier'ları**

| Specifier | Açıklama                                     | Kullanım Alanı   |
| --------- | -------------------------------------------- | ---------------- |
| `%T`      | Değişkenin tür bilgisini gösterir            | Herhangi bir tür |
| `%d`      | Tamsayıyı ondalık biçimde gösterir           | int              |
| `%f`      | Gerçek sayıyı sabit ondalık biçimde gösterir | float32, float64 |
| `%t`      | Boolean (mantıksal) değerleri gösterir       | bool             |
| `%v`      | Genel (varsayılan) biçimde gösterim          | Herhangi bir tür |
| `%x`/`%X` | Onaltılık biçim                              | int              |
| `%b`      | İkili biçim                                  | int              |
| `%o`/`%O` | Sekizlik biçim                               | int              |

### 4.4. Kaçış (Escape) Karakterleri ve Yüzde (%) Karakteri

**Açıklama:**

- Yazı içinde özel karakterleri göstermek için kaçış dizileri kullanılır.
- Örneğin:
  - `\"` → String içerisinde çift tırnak eklemek için.
  - `\\` → Ters bölü (backslash) karakterini göstermek için.
- `%` karakteri de format belirteci olarak kullanıldığı için, gerçek yüzde işareti göstermek amacıyla `%%` şeklinde yazılır.

**Detaylı Örnek:**

```go
package main

import "fmt"

func main() {
	var midTermRatio, finalRatio int

	fmt.Print("Vize ve final oranlarını giriniz (örn: 40 60): ")
	fmt.Scan(&midTermRatio, &finalRatio)

	// Yüzde işareti göstermek için %% kullanılır.
	fmt.Printf("Vize oranı: %%%d, Final oranı: %%%d\n", midTermRatio, finalRatio)
}
```

**Ek Açıklamalar:**

- Kaçış dizileri, dosya yolları, metin şablonları ve kullanıcı arayüzü çıktılarında sıkça kullanılır.
- Yanlış kaçış dizisi kullanımı derleme hatalarına neden olabilir; bu nedenle dikkatli olunmalıdır.

---

## 5. **Scanf** ile Formatlı Giriş

_(Scanf – Biçimli Giriş)_

**Açıklama:**

- **fmt.Scanf** fonksiyonu, kullanıcıdan format belirteçleri kullanılarak giriş alır.
- Girişte, kullanıcıdan alınacak verilerin formatı, beklenen türler ve sırayla yazılmış olmalıdır.
- Bu fonksiyon, özellikle birden fazla değerin aynı anda okunması gereken durumlarda tercih edilir.

**Detaylı Örnek:**

```go
package main

import "fmt"

func main() {
	var midTermRatio, finalRatio int

	fmt.Print("Vize ve final oranlarını giriniz (örn: 40 60): ")
	// İki değeri boşluk ile ayrılmış olarak alır.
	fmt.Scanf("%d%d", &midTermRatio, &finalRatio)

	fmt.Printf("Vize oranı: %%%v, Final oranı: %%%v\n", midTermRatio, finalRatio)
}
```

**Ek Açıklamalar:**

- **Scanf** kullanırken, format dizesi ile girilen değerlerin sırası ve aralarındaki ayırıcılar çok önemlidir.
- Yanlış format kullanımı girişin hatalı okunmasına veya programın beklenmedik davranış sergilemesine neden olabilir.

---

## 6. Karmaşık Sayılar (Complex Numbers) ile İşlemler

_(Complex Numbers – Karmaşık Sayılar)_

**Açıklama:**

- Go dilinde, karmaşık sayılar **complex64** (daha düşük hassasiyet) ve **complex128** (daha yüksek hassasiyet) türleri ile tanımlanır.
- Karmaşık sayılar, matematiksel işlemler (toplama, çıkarma, çarpma, bölme) üzerinde doğrudan işlem yapılmasına imkan tanır.
- Girişte karmaşık sayıların yazım biçimi, parantezli veya parantezsiz olabilir; ancak girilen değerlerin formatı doğru yorumlanmalıdır.

**Detaylı Örnek:**

```go
package main

import "fmt"

func main() {
	runApp()
}

func runApp() {
	fmt.Print("İki karmaşık sayı giriniz (örn: (3+4i) (5+6i)): ")
	z1, z2 := readComplexNumbers()
	printResults(z1, z2)
}

func readComplexNumbers() (complex128, complex128) {
	var z1, z2 complex128
	// Kullanıcıdan iki karmaşık sayı okunur.
	fmt.Scan(&z1, &z2)
	return z1, z2
}

func printResults(z1, z2 complex128) {
	fmt.Printf("(%v) + (%v) = %v\n", z1, z2, z1+z2)
	fmt.Printf("(%v) - (%v) = %v\n", z1, z2, z1-z2)
	fmt.Printf("(%v) * (%v) = %v\n", z1, z2, z1*z2)
	fmt.Printf("(%v) / (%v) = %v\n", z1, z2, z1/z2)
}
```

**Ek Açıklamalar:**

- Karmaşık sayılarla yapılan işlemler, mühendislik, sinyal işleme ve fiziksel hesaplamalar gibi alanlarda sıklıkla kullanılır.
- İşlemlerin doğruluğu, girilen karmaşık sayıların formatına bağlıdır; bu nedenle kullanıcı girdileri doğrulanmalıdır.

---

## 7. Sabitler (Literals) ve Sabit İfadeler

_(Literals – Sabitler, Constant Expressions – Sabit İfadeler)_

**Açıklama:**

- Kod içerisinde doğrudan yazılan değerlere **literal** (sabit) denir.
- Go dilinde sabitler, değerin türüne göre otomatik olarak belirlenir; fakat bazı durumlarda explicit olarak tanımlanabilir.
- Sabitler aşağıdaki kategorilere ayrılır:
  - **Integer literals (Tamsayı sabitleri):**
    - Ondalık, ikili (0b), onaltılık (0x) ve sekizlik (0o) biçimlerde yazılabilir.
    - Okunabilirliği artırmak için sayılar arasında `_` (alt çizgi) kullanılabilir.
  - **Floating point literals (Gerçek sayı sabitleri):**
    - Ondalık biçimde (örneğin, 5.6) veya üstel gösterimle (örneğin, 6.02e23) yazılabilir.
  - **Complex literals (Karmaşık sabitler):**
    - Gerçek ve sanal kısımlar, `+` işareti ile birleşir; örneğin, `1+2i`.
  - **Boolean literals (Mantıksal sabitler):**
    - `true` veya `false` değerleri.
  - **Rune literals (Rune/karakter sabitleri):**
    - Tek tırnak içinde, örneğin, `'A'`, kaçış dizileri de kullanılabilir.
  - **String literals (Dize sabitleri):**
    - Çift tırnak içinde veya raw string (`` ` ``) olarak yazılabilir.

**Detaylı Örnek:**

```go
package main

import "fmt"

func main() {
	a := 10              // integer literal (örn: 10, 0b1010, 0x1F)
	b := 5.6             // floating point literal (örn: 5.6, 6.02e23)
	c := 1 + 2i          // complex literal (örn: 1+2i)
	d := (1 + 2i)        // parantez opsiyonel kullanım
	e := 3.4i            // sadece sanal kısmı belirten karmaşık literal
	f := true            // boolean literal

	fmt.Printf("a = %d\n", a)
	fmt.Printf("b = %f\n", b)
	fmt.Printf("c = %v\n", c)
	fmt.Printf("d = %v\n", d)
	fmt.Printf("e = %v\n", e)
	fmt.Printf("f = %t\n", f)
}
```

**Tablo: Sabit Türleri ve Örnekleri**

| Sabit Türü             | Örnek Kullanım                | Açıklama                                         |
| ---------------------- | ----------------------------- | ------------------------------------------------ |
| Integer Literal        | `10`, `0b1010`, `0x1F`        | Ondalık, binary, hexadecimal, octal formatlarda  |
| Floating Point Literal | `5.6`, `6.02e23`              | Standart ondalıklı ve üstel gösterimli biçimler  |
| Complex Literal        | `1+2i`, `(3+4i)`              | Gerçek ve sanal kısımların birleşimi             |
| Boolean Literal        | `true`, `false`               | Mantıksal değerler                               |
| Rune Literal           | `'a'`, `'\n'`                 | Tek karakter sabitleri, kaçış dizileri de mümkün |
| String Literal         | `"Hello"`, `` `Raw string` `` | Çift tırnak veya raw string biçiminde yazım      |

---

## 8. String ve Rune Literalleri ile Escape Karakterleri

_(Escape Sequences – Kaçış Dizileri)_

**Açıklama:**

- **Escape dizileri** sayesinde, string veya rune literal içinde özel karakterler (örn: newline, tab, tırnak işaretleri) doğru biçimde temsil edilir.
- Desteklenen kaçış dizileri:
  - `\a` → alert/bell (sinyal sesi)
  - `\b` → backspace (geri silme)
  - `\f` → form feed (sayfa sonu)
  - `\n` → newline (yeni satır)
  - `\r` → carriage return (satır başı)
  - `\t` → horizontal tab (yatay sekme)
  - `\v` → vertical tab (dikey sekme)
  - `\\` → backslash (ters bölü)
  - `\'` → single quote (tek tırnak; rune literalinde)
  - `\"` → double quote (çift tırnak; string literalinde)
- Dosya yolları veya düzenli ifadelerde (regex) raw string literal (`` ` ``) kullanımı hata riskini azaltır.

**Detaylı Örnek:**

```go
package main

import "fmt"

func main() {
	// Çift tırnak içinde escape dizileri kullanarak metin yazdırma.
	fmt.Print("\"Hello, World\"\n")
	// Tek tırnak işareti doğrudan kullanılabilir.
	fmt.Print("'Goodbye, World'\n")
}
```

**Dosya Yolu Örneği:**

- **Yanlış Kullanım:**  
  Ters bölü karakterlerinin kaçış dizileri olarak algılanması sonucu istenmeyen çıktılar oluşur.

  ```go
  package main

  import "fmt"

  func main() {
  	// Bu örnekte "\t" tab, "\n" ise yeni satır olarak yorumlanır.
  	fmt.Print("C:\test\numbers.txt")
  }
  ```

- **Doğru Kullanım Yöntemleri:**

  - **İki Ters Bölü Kullanımı:**

    ```go
    package main

    import "fmt"

    func main() {
    	// Her ters bölü için iki tane \\ yazılır.
    	fmt.Print("C:\\test\\numbers.txt")
    }
    ```

  - **Raw String Literal Kullanımı:**

    ```go
    package main

    import "fmt"

    func main() {
    	// Raw string literal sayesinde kaçış dizileri işlenmez.
    	fmt.Print(`C:\test\numbers.txt`)
    }
    ```

**Unicode Örneği:**

```go
package main

import "fmt"

func main() {
	// Unicode karakteri kullanılarak özel sembol gösterimi.
	fmt.Printf("Ankara.\u2606İstanbul\n")
}
```

---

## 9. Gerçek Sayı Sabitlerinde Üstel Gösterim

_(Exponential Notation – Üstel Gösterim)_

**Açıklama:**

- Gerçek sayı sabitleri, çok büyük veya çok küçük sayıları daha okunabilir hale getirmek için üstel (exponential) gösterimle yazılabilir.
- Örneğin, `6.02e23` ifadesi 6.02 × 10^23 anlamına gelir.
- Ayrıca, Go’da gerçek sayı sabitleri hexadecimal üstel biçimde de ifade edilebilir.
  - Örneğin, `0xABp3` ifadesinde `p` karakteri, üstel kısmı temsil eder.
  - Bu biçim, özellikle düşük seviyeli programlama ve bellek hesaplamalarında kullanılır.

**Detaylı Örnek:**

```go
package main

import "fmt"

func main() {
	a := 6.02e23

	fmt.Printf("a = %f\n", a)
	fmt.Printf("a'nın türü = %T\n", a)
}
```

**Ek Açıklamalar:**

- Üstel gösterim, bilimsel hesaplamalarda ve mühendislik uygulamalarında çok yaygın kullanılır.
- Bu gösterim, verinin okunabilirliğini ve anlaşılabilirliğini artırır.

---

## 10. Özet ve Sonuç

**Önemli Noktalar:**

- **Yuvarlama Fonksiyonları:**
  - `math.Round`, `math.Floor` ve `math.Ceil` fonksiyonları, farklı yuvarlama yöntemleri sunarak, kullanıcı girdilerine uygun matematiksel işlemler sağlar.
- **IEEE 754 Özel Değerler:**
  - Gerçek sayı işlemlerinde `NaN`, `+Inf`, `-Inf` gibi özel durumlar, hata yönetimi ve hesaplama doğruluğu açısından kritik öneme sahiptir.
- **Bölme İşlemleri:**
  - Tamsayı bölmesi, ondalık kısmı atarken; gerçek sayı bölmesi IEEE 754’e uygun özel değerler döndürür.
  - Sıfıra bölme durumlarında, tamsayı bölmesi runtime error üretirken, gerçek sayı bölmesi hata fırlatmadan özel değer döner.
- **Formatlama (fmt) İşlemleri:**
  - `Printf` ve `Scanf` fonksiyonları, format belirteçleri sayesinde verilerin doğru biçimde gösterilmesini ve okunmasını sağlar.
  - Hizalama, kesinlik ayarı, farklı sayı sistemleri ve kaçış dizileri gibi özellikler, çıktıların estetik ve işlevsel olarak düzenlenmesine olanak tanır.
- **Karmaşık Sayılar:**
  - Go, karmaşık sayılar üzerinde doğrudan aritmetik işlemler yapma imkanı sunar.
  - Bu özellik, mühendislik, sinyal işleme ve fiziksel hesaplamalar gibi alanlarda büyük avantaj sağlar.
- **Sabitler (Literals):**
  - Kod içerisinde doğrudan yazılan sabitler, farklı türlerde (integer, floating point, complex, boolean, rune, string) tanımlanabilir.
  - Okunabilirlik ve hata önleme amacıyla, sayılar arasında `_` kullanımı ve doğru biçimlendirme kurallarına dikkat edilmelidir.
- **Escape Karakterleri:**
  - String ve rune literal'larda özel karakterlerin doğru temsil edilmesi için kaçış dizileri kullanılır.
  - Dosya yolları gibi durumlarda raw string literal kullanımı, kaçış karakterlerinin neden olabileceği hataları önler.
- **Üstel Gösterim:**
  - Gerçek sayı sabitlerinin üstel ve hexadecimal üstel formatlarda yazılması, büyük verilerin okunabilirliğini artırır.

**Genel Değerlendirme:**

Bu detaylı açıklamalar, Golang’ın temel matematiksel işlemlerden formatlama tekniklerine, sabit tanımlamalardan kullanıcı girişine kadar geniş bir yelpazede nasıl kullanıldığını kapsamlı örnekler, tablolar ve açıklamalar eşliğinde sunmaktadır.

- **Kullanım Kolaylığı:** Basit sözdizimi sayesinde öğrenilmesi ve uygulanması kolaydır.
- **Performans:** Derleyici destekli yapısı ile yüksek performanslı uygulamalar geliştirmeye olanak tanır.
- **Esneklik:** Hem sistem seviyesinde hem de uygulama düzeyinde çözümler sunar; özellikle eşzamanlılık (concurrency) desteği ile modern uygulamalara uygun çözümler üretebilir.
- **Geniş Kapsam:** Yuvarlama, bölme, formatlama, karmaşık sayılar ve sabitler gibi konuların tümü, gerçek dünya uygulamalarında karşılaşılabilecek durumlar için detaylı olarak ele alınmıştır.
