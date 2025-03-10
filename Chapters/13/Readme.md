## 1. Go'da Yazılarla İşlemler

**Genel Bilgi:**

- **Veri Türü:**
  - Go dilinde yazılar, `string` (string – yazı) veri türüyle temsil edilir.
  - Yazılar UTF-8 encoding (UTF-8 kodlaması) biçiminde saklanır; bu nedenle her karakterin bellek üzerindeki byte (bayt) sayısı farklılık gösterebilir.
- **Immutable Yapı:**
  - `string` türü immutable (immutable – değiştirilemez) olduğu için, bir string üzerinde doğrudan değişiklik yapılamaz.
  - Herhangi bir düzenleme, mevcut string üzerinde işlem yapmadan yeni bir string nesnesi oluşturur.
- **Kullanım Alanları:**
  - Kullanıcı girdilerinin alınması, metin tabanlı dosya okuma/yazma, veri iletişimi ve web uygulamalarında verilerin saklanması gibi pek çok alanda kullanılır.

**Ek Açıklamalar:**

- **Bellek Yönetimi:**
  - Immutable yapı, aynı string literal (string literal – yazı sabiti) tekrar kullanıldığında bellek paylaşımını mümkün kılar, bu da performans açısından avantaj sağlar.
- **UTF-8 Esnekliği:**
  - UTF-8 kodlaması sayesinde, Latin alfabesi dışındaki karakterler (örneğin, Türkçe karakterler, Çince, Arapça) doğru şekilde saklanır ve işlenir.

**Örnek Kod:**

```go
package main

import "fmt"

func main() {
	s1 := "ankara"              // string literal kullanımı
	var s2 string = "istanbul"   // açık tip tanımlaması
	// Her iki yazı, standart çıktı aracılığıyla ekrana basılır.
	fmt.Println("s1:", s1)
	fmt.Println("s2:", s2)
	fmt.Printf("Formatted: City: %s\n", s1)
}
```

**Kullanım Senaryoları:**

- **Kullanıcı Arayüzü:** Girdi/çıktı işlemlerinde yazılar doğrudan kullanıcıya mesaj iletmek için kullanılır.
- **Veri İletişimi:** API’ler ve web servislerinde JSON, XML gibi formatlarda yazılar önemli rol oynar.

---

## 2. String Literal (SL – string literal) Kullanımı

**Açıklama:**

- **Tanım:**
  - İki çift tırnak (`"`) arasında yazılan yazılar string literal (string literal – yazı sabiti) olarak adlandırılır.
- **Özellikler:**
  - Bu yazılar derleme zamanında sabit (constant) olarak değerlendirilir ve çalışma zamanında değiştirilemezler.
  - Print, Println, Printf gibi fonksiyonlara argüman olarak geçildiğinde, ilgili string’in içeriği standart çıktıya (stdout – standart çıktı) basılır.

**Ek Açıklamalar:**

- **Kullanım Kolaylığı:**
  - String literal’ler, doğrudan sabit metin tanımlamak için ideal olup, hata yapma olasılığını azaltır.
- **Formatlama:**
  - `fmt.Printf` fonksiyonunda `%s` format karakteri ile yazı biçimlendirme işlemleri kolayca yapılabilir.

**Örnek Kod:**

```go
package main

import "fmt"

func main() {
	s := "Hello, Go!" // Basit string literal
	fmt.Println(s)
	// Formatlı çıktı örneği:
	fmt.Printf("City: %s\n", s)
}
```

---

## 3. Escape Sequence ve Raw String Literals

### 3.1. Escape Sequence

**Açıklama:**

- **Escape Sequence (Kaçış Dizisi):**
  - `\` karakteriyle başlayan özel karakter dizileridir.
  - Örneğin; `\n` yeni satır, `\t` tab, `\\` ters bölü (backslash) gibi.
- **Kullanım Durumu:**
  - Dosya yolları, metin formatlama ve özel karakterlerin gösterilmesi gereken durumlarda kullanılır.

**Örnek Kod:**

```go
package main

import "fmt"

func main() {
	path1 := "C:\\test\\names.txt"  // Her ters bölü için çift ters bölü kullanılır.
	fmt.Println("Escaped Path:", path1)
}
```

### 3.2. Raw String Literals

**Açıklama:**

- **Tanım:**
  - Raw string literals, `` ` `` (backtick – ters tırnak) karakterleri arasında tanımlanır.
  - Bu literal’ler içerisinde escape sequence’ler işlenmez; yazıldığı gibi saklanır.
- **Özellikleri:**
  - Çok satırlı yazılar oluşturmak için idealdir.
  - Kod içerisinde düzenli biçimlendirme yapmak, örneğin SQL sorguları veya HTML içeriği tanımlamak için kullanılabilir.

**Örnek Kod:**

```go
package main

import "fmt"

func main() {
	// Raw string literal örneği:
	path2 := `C:\test\names.txt`
	fmt.Println("Raw Path:", path2)

	// Çok satırlı raw string literal:
	multiline := `Bu bir
çok satırlı
yazıdır.`
	fmt.Println("Multiline Text:\n", multiline)
}
```

**Karşılaştırma Tablosu:**

| **Özellik**             | **Escape Sequence Kullanımı**                      | **Raw String Literal Kullanımı**                    |
| ----------------------- | -------------------------------------------------- | --------------------------------------------------- |
| **Yazım Şekli**         | `"C:\\test\\names.txt"`                            | `` `C:\test\names.txt` ``                           |
| **Escape İşlemi**       | Escape sequence’ler işlenir (örn: `\n` yeni satır) | Escape sequence’ler olduğu gibi, işlenmeden yazılır |
| **Çok Satırlı Yazılar** | Zor ve karmaşıktır                                 | Doğrudan çok satırlı yazı oluşturmak mümkündür      |

---

## 4. Yazının Uzunluğunu Ölçme (len Fonksiyonu)

**Açıklama:**

- `len()` fonksiyonu, verilen string’in byte cinsinden uzunluğunu döner.
- **Dikkat Edilmesi Gereken:**
  - UTF-8 kodlaması nedeniyle, özellikle Unicode karakterlerde gerçek karakter sayısı ile byte sayısı farklı olabilir.

**Örnek Kod:**

```go
package main

import "fmt"

func main() {
	var str string = "merhaba"
	// "merhaba" ifadesi 7 karakter olsa da, ASCII karakterler tek byte olduğu için byte uzunluğu 7’dir.
	fmt.Printf("Byte uzunluğu: %d\n", len(str))
}
```

**Örnek Durumlar:**

- **Kullanıcı Girdisi:** Kullanıcının girdiği metnin uzunluğunu ölçmek, örneğin şifre minimum uzunluk kontrolü gibi işlemlerde kullanılır.
- **Dosya İşlemleri:** Dosyadan okunan metinlerin uzunluğunu kontrol etmek için de yararlıdır.

| **Fonksiyon** | **Açıklama**                              | **Örnek Kullanım** |
| ------------- | ----------------------------------------- | ------------------ |
| `len()`       | String'in byte cinsinden uzunluğunu verir | `len("ankara")`    |

---

## 5. Yazı Üzerinde Döngü ve Karakterlere Erişim

### 5.1. Range Döngüsü

**Açıklama:**

- `for range` döngüsü, string üzerinde rune (karakter – karakter) bazında iterasyon yapar.
- Döngü sırasında her iterasyonda indeks (int) ve karakter (rune) elde edilir.
- **Avantajları:**
  - UTF-8 karakterlerin her biri doğru şekilde ele alınır; böylece çok byte’lık karakterler bozulmaz.

**Örnek Kod:**

```go
package main

import "fmt"

func main() {
	var str string = "merhaba"
	for i, c := range str {
		// c değeri rune olduğundan, %c formatı ile karakter olarak basılır.
		fmt.Printf("index: %d, karakter: %c\n", i, c)
	}
}
```

### 5.2. Byte Bazlı Erişim

**Açıklama:**

- İndeks operatörü (`[]`) kullanılarak string’in belirli bir byte’ına doğrudan erişilir.
- **Önemli Not:**
  - Bu yöntem, sadece standart ASCII karakterler için güvenlidir. UTF-8 karakterlerde birden fazla byte kullanıldığı için istenmeyen sonuçlar ortaya çıkabilir.

**Örnek Kod:**

```go
package main

import "fmt"

func main() {
	var str string = "merhaba"
	// Her byte'ı tek tek işleyerek basmak:
	for i := 0; i < len(str); i++ {
		fmt.Printf("%c ", str[i])
	}
	fmt.Println()
}
```

| **Yöntem**       | **Açıklama**                                                     | **Avantaj/Dezavantaj**                          |
| ---------------- | ---------------------------------------------------------------- | ----------------------------------------------- |
| Range Döngüsü    | UTF-8 uyumlu, her karakteri (rune) doğru şekilde iterasyon yapar | Unicode karakterlerde güvenilir, okunabilir kod |
| İndeks Operatörü | Belirli byte’lara erişim sağlar; ASCII için uygundur             | UTF-8 karakterlerde hatalı sonuç verebilir      |

---

## 6. String Birleştirme (Concatenation)

**Açıklama:**

- String’ler, `+` operatörü kullanılarak kolayca birleştirilebilir.
- Birleştirme işlemleri sırasında, string’lerin immutability (değiştirilemezlik) özelliği nedeniyle yeni string nesneleri oluşturulur.

**Örnek Kod:**

```go
package main

import "fmt"

func main() {
	s1 := "ankara"
	s2 := "istanbul"
	combined := s1 + " - " + s2  // İki yazının araya tire ve boşluk eklenerek birleştirilmesi
	fmt.Println("Birleştirilmiş:", combined)
}
```

**Kullanım Notları:**

- Büyük metinlerin birleştirilmesinde, sürekli `+` operatörü kullanımı performans sorunlarına neden olabilir. Bu durumlarda `strings.Builder` tercih edilebilir.

---

## 7. String Karşılaştırmaları

### 7.1. Eşitlik ve Eşitsizlik

**Açıklama:**

- `==` ve `!=` operatörleri kullanılarak iki string’in tamamen aynı olup olmadığı kontrol edilir.
- Bu karşılaştırmalar, karakter dizisinin byte bazında eşitliğine bakar.

**Örnek Kod:**

```go
package main

import "fmt"

func main() {
	s1 := "ankara"
	s2 := "ankara"
	if s1 == s2 {
		fmt.Println("Aynı yazı")
	} else {
		fmt.Println("Farklı yazılar")
	}
}
```

### 7.2. Lexicographical (Sözlük) Karşılaştırma

**Açıklama:**

- `<`, `>`, `<=`, `>=` operatörleri ile string’ler sözlük sırasına göre karşılaştırılır.
- Bu işlem, karakterlerin ASCII veya Unicode sıralamasına göre gerçekleştirilir.

**Örnek Kod:**

```go
package main

import "fmt"

func main() {
	s1 := "apple"
	s2 := "banana"
	if s1 < s2 {
		fmt.Printf("%s < %s\n", s1, s2)
	} else if s1 == s2 {
		fmt.Printf("%s == %s\n", s1, s2)
	} else {
		fmt.Printf("%s > %s\n", s1, s2)
	}
}
```

| **Operatör** | **Açıklama**                                               | **Örnek**            |
| ------------ | ---------------------------------------------------------- | -------------------- |
| `==`         | İki string’in tam olarak aynı olup olmadığını kontrol eder | `"a" == "a"`         |
| `<`          | Sözlük sırasına göre karşılaştırma yapar                   | `"apple" < "banana"` |

---

## 8. strings Paketindeki Fonksiyonlar

Go’nun standart kütüphanesinde bulunan `strings` (strings – yazılar) paketi, string işlemlerinde kullanılacak çok sayıda yardımcı fonksiyon içerir.

### 8.1. strings.Compare

**Açıklama:**

- İki string’i lexicographical (sözlük sırası) olarak karşılaştırır.
- Dönüş değeri:
  - `-1`: İlk string ikinci string’den küçüktür.
  - `0`: İki string eşittir.
  - `1`: İlk string ikinci string’den büyüktür.

**Örnek Kod:**

```go
package main

import (
	"fmt"
	"strings"
)

func main() {
	result := strings.Compare("apple", "banana")
	fmt.Println("Compare Sonucu:", result)
}
```

### 8.2. strings.Contains

**Açıklama:**

- Bir string’in içinde belirli bir alt string’in (substring – alt yazı) bulunup bulunmadığını kontrol eder.
- Dönüş değeri boolean (true/false) olur.

**Örnek Kod:**

```go
package main

import (
	"fmt"
	"strings"
)

func main() {
	if strings.Contains("merhaba dünya", "dünya") {
		fmt.Println("Yazı içerisinde 'dünya' var.")
	} else {
		fmt.Println("Yazı içerisinde 'dünya' yok.")
	}
}
```

### 8.3. strings.Index ve strings.LastIndex

**Açıklama:**

- `Index`: Bir alt string’in ilk geçtiği konumun indeksini döner; bulunamazsa `-1` döner.
- `LastIndex`: Bir alt string’in sondan itibaren ilk geçtiği konumun indeksini döner.

**Örnek Kod:**

```go
package main

import (
	"fmt"
	"strings"
)

func main() {
	text := "merhaba merhaba"
	fmt.Println("Index:", strings.Index(text, "haba"))
	fmt.Println("LastIndex:", strings.LastIndex(text, "haba"))
}
```

### 8.4. strings.Repeat

**Açıklama:**

- Belirtilen string’i, belirlenen sayı kadar tekrarlar ve yeni bir string oluşturur.

**Örnek Kod:**

```go
package main

import (
	"fmt"
	"strings"
)

func main() {
	repeated := strings.Repeat("go", 3) // "go" string'i 3 kez tekrarlar: "gogogo"
	fmt.Println("Repeat:", repeated)
}
```

| **Fonksiyon**       | **Açıklama**                                    | **Örnek Kullanım**                |
| ------------------- | ----------------------------------------------- | --------------------------------- |
| `strings.Compare`   | İki string’i sözlük sırası ile karşılaştırır    | `strings.Compare("a", "b")`       |
| `strings.Contains`  | Belirtilen alt string’in varlığını kontrol eder | `strings.Contains("hello", "ll")` |
| `strings.Index`     | İlk geçiş indeksini döner, bulunamazsa -1       | `strings.Index("hello", "l")`     |
| `strings.LastIndex` | Son geçiş indeksini döner, bulunamazsa -1       | `strings.LastIndex("hello", "l")` |
| `strings.Repeat`    | String’i belirtilen sayı kadar tekrarlar        | `strings.Repeat("a", 5)`          |

---

## 9. Sınıf Çalışması: Padding (PadLeading/PadTrailing) Fonksiyonları

**Açıklama:**

- **Amaç:**
  - Bir yazıyı (string – yazı) istenen uzunluğa getirmek için baştan veya sondan belirli karakter (rune – karakter) eklenmesi.
  - Örneğin, `PadLeading("ankara", 9, 'x')` ifadesi `"xxxankara"` sonucunu üretir.
- **Detaylar:**
  - Eğer `newLen` parametresi, verilen string’in uzunluğundan küçük veya eşitse, orijinal string döndürülür.
  - İki farklı versiyon:
    - **PadLeading:** Yazıyı baştan besler.
    - **PadTrailing:** Yazıyı sondan besler.
  - Ek olarak, sadece boşluk karakteri ile padlemek için özel fonksiyonlar (`PadLeadingSpace`, `PadTrailingSpace`) tanımlanabilir.

**Örnek Kod:**

```go
package main

import (
	"fmt"
	"strings"
)

func PadLeading(s string, newLen int, c rune) string {
	length := len(s)
	if newLen <= length {
		return s
	}
	// Belirlenen karakterin (c) tekrar sayısı: newLen - mevcut uzunluk
	return strings.Repeat(string(c), newLen-length) + s
}

func PadTrailing(s string, newLen int, c rune) string {
	length := len(s)
	if newLen <= length {
		return s
	}
	return s + strings.Repeat(string(c), newLen-length)
}

func main() {
	fmt.Println(PadLeading("ankara", 9, 'x'))  // Çıktı: xxxankara
	fmt.Println(PadTrailing("ankara", 9, 'x')) // Çıktı: ankaraxxx
}
```

| **Fonksiyon** | **Açıklama**                              | **Örnek Çıktı** |
| ------------- | ----------------------------------------- | --------------- |
| `PadLeading`  | Yazıyı baştan verilen karakterle doldurur | "xxxankara"     |
| `PadTrailing` | Yazıyı sondan verilen karakterle doldurur | "ankaraxxx"     |

---

## 10. Sınıf Çalışması: Reverse (Yazıyı Tersine Çevirme) Fonksiyonu

**Açıklama:**

- **Amaç:**
  - Bir yazının karakter sırasını tersine çevirerek yeni bir string elde etmek.
- **Yaklaşım Seçenekleri:**
  - **Naif Yaklaşım:** Her iterasyonda yeni bir string oluşturulur; küçük metinlerde kabul edilebilir, ancak büyük metinlerde performans problemi yaratır.
  - **Verimli Yaklaşım:** `strings.Builder` (strings.Builder – string derleyici) kullanılarak bellek tahsisi minimize edilir.

**Naif Yaklaşım Örneği:**

```go
func ReverseNaive(s string) string {
	var rev string = ""
	for i := len(s) - 1; i >= 0; i-- {
		rev += string(s[i])
	}
	return rev
}
```

**strings.Builder Kullanarak Verimli Yaklaşım:**

```go
package main

import (
	"fmt"
	"strings"
)

func Reverse(s string) string {
	var builder strings.Builder
	// Döngü, string’in son karakterinden başlayıp ilk karaktere kadar ilerler.
	for i := len(s) - 1; i >= 0; i-- {
		// WriteByte, her seferinde tek bayt ekler; ASCII karakterler için uygundur.
		builder.WriteByte(s[i])
	}
	return builder.String()
}

func main() {
	fmt.Println("Tersine:", Reverse("golang"))
}
```

| **Yöntem**      | **Açıklama**                                               | **Avantajları / Dezavantajları**                   |
| --------------- | ---------------------------------------------------------- | -------------------------------------------------- |
| Naif Yaklaşım   | Her adımda yeni string nesnesi oluşturur                   | Küçük metinlerde anlaşılır, büyük metinlerde yavaş |
| strings.Builder | Bellek tahsisini optimize eder, tek nesne üzerinden toplar | Büyük metinlerde performans artışı sağlar          |

---

## 11. strings.Builder Kullanımının Avantajları

**Açıklama:**

- **Neden Kullanılır?**
  - Go'da string'ler immutable (değiştirilemez) olduğu için, birleştirme veya düzenleme işlemleri sırasında sürekli yeni nesne üretimi yapılır.
  - `strings.Builder` bu sorunu aşmak için tasarlanmıştır; belleği önceden tahsis eder ve yazı ekleme işlemlerinde performansı artırır.

**Kullanım Örneği:**

```go
package main

import (
	"fmt"
	"strings"
)

func main() {
	var builder strings.Builder
	// WriteString metodu ile string parçaları eklenir.
	builder.WriteString("Go ")
	builder.WriteString("dilinde ")
	builder.WriteString("string işlemleri.")
	fmt.Println(builder.String())
}
```

**Özellikler:**

- **Bellek Verimliliği:** Tek seferlik bellek tahsisi ile tekrar eden eklemelerde kopyalama maliyetini azaltır.
- **Kolay Kullanım:** WriteString, WriteRune gibi metotlarla desteklenir.

---

## 12. Trim (Kırpma) Fonksiyonları

**Açıklama:**

- Yazının başındaki ve sonundaki gereksiz boşluklar veya belirli karakterler kaldırılabilir.
- **Temel Fonksiyonlar:**
  - **TrimSpace:** Yazının başındaki ve sonundaki tüm white-space (boşluk – boşluk karakterleri) karakterlerini temizler.
  - **TrimLeft / TrimRight:** Yazının solundan veya sağından, belirtilen karakter seti içinde bulunan karakterleri temizler.
  - **TrimPrefix / TrimSuffix:** Yazının başından veya sonundan belirtilen alt yazıyı kaldırır.
  - **Trim:** Yazının her iki tarafından da belirtilen karakterleri temizler.

**Örnek Kod:**

```go
package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "      \tankara istanbul\t\t  "
	fmt.Printf("Orijinal:(%s)\n", s)
	fmt.Printf("TrimSpace:(%s)\n", strings.TrimSpace(s))
	fmt.Printf("TrimLeft:(%s)\n", strings.TrimLeft(s, " \t"))
	fmt.Printf("TrimRight:(%s)\n", strings.TrimRight(s, " \t"))
}
```

**Kullanım Senaryoları:**

- Kullanıcı girdilerindeki gereksiz boşlukları temizleme, veri doğrulama işlemleri.

---

## 13. Case Dönüşümü Fonksiyonları

**Açıklama:**

- Yazı içerisindeki karakterlerin tümü büyük veya küçük harfe dönüştürülebilir.
- **Fonksiyonlar:**
  - **ToUpper / ToLower:** Tüm karakterleri büyük/küçük harfe çevirir.
  - **ToUpperSpecial / ToLowerSpecial:** Belirli dil kuralları (örneğin, Türkçe için unicode.TurkishCase) doğrultusunda dönüşüm yapar.

**Örnek Kod:**

```go
package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	text := "istanbul"
	fmt.Println("Upper:", strings.ToUpper(text))
	fmt.Println("Lower:", strings.ToLower(text))

	// Türkçe özel dönüşüm:
	fmt.Println("Turkish Upper:", strings.ToUpperSpecial(unicode.TurkishCase, "istanbul"))
	fmt.Println("Turkish Lower:", strings.ToLowerSpecial(unicode.TurkishCase, "İSTANBUL"))
}
```

**Kritik Noktalar:**

- Türkçe gibi dillerde, "i/I" ve "ı/İ" dönüşümleri farklılık gösterdiğinden, özel dönüşüm fonksiyonlarının kullanılması önemlidir.

---

## 14. strconv Paketi ile String Dönüşümleri

**Açıklama:**

- **Numeric to String:**
  - `Itoa`, `FormatInt`, `FormatFloat` gibi fonksiyonlar, sayısal veriyi string'e çevirir.
- **String to Numeric:**
  - `Atoi`, `ParseInt`, `ParseFloat` gibi fonksiyonlar, string içerisindeki sayısal değeri uygun sayısal türe çevirir.
- **Örnek Kullanım:**
  - Kullanıcı girdisinin sayıya dönüştürülmesi, dosya okuma işlemleri sırasında string'in sayısal forma aktarılması.

**Örnek Kod:**

```go
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Bir sayı giriniz: ")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	val, err := strconv.Atoi(input)
	if err != nil {
		fmt.Println("Geçersiz sayı!")
		return
	}
	fmt.Printf("%d * %d = %d\n", val, val, val*val)
}
```

**Notlar:**

- `bufio` paketi, büyük metinlerin ve kullanıcı girdilerinin verimli okunmasını sağlar.
- `Atoi` fonksiyonu, `ParseInt` fonksiyonunun özel bir versiyonudur.

---

## 15. unicode Paketi ile Rune İşlemleri

**Açıklama:**

- `unicode` (unicode – karakter seti) paketi, bir rune’un (karakter – karakter) özelliklerini test etmek ve dönüştürmek için yardımcı fonksiyonlar sunar.
- **Örnek Fonksiyonlar:**
  - `IsDigit`, `IsLetter`, `IsSpace`, `IsPunct` gibi predicate fonksiyonlar;
  - `ToUpper`, `ToLower` gibi dönüşüm fonksiyonları.

**Örnek Kod:**

```go
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Bir yazı giriniz: ")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	for _, c := range input {
		if unicode.IsDigit(c) {
			fmt.Printf("%c -> DIGIT\n", c)
		} else if unicode.IsSpace(c) {
			fmt.Printf("%c -> WS\n", c)
		} else if unicode.IsPunct(c) {
			fmt.Printf("%c -> PUNCT\n", c)
		} else {
			fmt.Printf("%c -> LETTER\n", c)
		}
	}
}
```

**Ek Örnek:**

- Belirli karakter türlerini filtreleyip, her karakterin türünü belirterek çıktı vermek, kullanıcı girdilerinin doğrulanması ve analizinde kullanışlıdır.

---

## 16. Bellek Optimizasyonu: String Literal Paylaşımı

**Açıklama:**

- Go’da aynı string literal (örneğin, `"ankara"`) birden fazla yerde kullanıldığında, runtime (çalışma zamanı) bu literal’i tek bir nesne olarak saklar.
- **Avantaj:**
  - Bellek kullanımını azaltır ve aynı veriye erişim durumunda gereksiz kopyalamaları önler.

**Örnek Kod:**

```go
package main

import "fmt"

func main() {
	s1 := "ankara"
	s2 := "ankara"
	// s1 ve s2 farklı değişkenler olsa da, aynı literal adresine referans verebilir.
	fmt.Printf("&s1 = %p, &s2 = %p\n", &s1, &s2)
}
```

**Not:**

- Bu örnek, değişkenlerin adreslerini gösterir. String literal’lerin gerçek verisi, çalışma zamanında tek bir bellek alanında saklanır.

---

## 17. Sınıf Çalışması: Pangram Fonksiyonları

**Tanım:**

- **Pangram:** Bir dilin alfabesindeki tüm harflerin en az bir kere kullanıldığı cümledir.
- **Örnekler:**
  - **İngilizce:** "The quick brown fox jumps over the lazy dog"
  - **Türkçe:** "Pijamalı hasta yağız şoföre çabucak güvendi"

**Detaylar:**

- Yazı küçük harfe çevrilip, alfabedeki her harf için kontrol yapılır.
- Türkçe için `ToLowerSpecial` kullanılarak, dilin özel dönüşüm kuralları uygulanır.

**Örnek Kod:**

```go
package main

import (
	"fmt"
	"strings"
	"unicode"
)

func IsPangramEN(s string) bool {
	return isPangram(strings.ToLower(s), "abcdefghijklmnopqrstuvwxyz")
}

func IsPangramTR(s string) bool {
	// Türkçe alfabesi: abcçdefgğhıijklmnoöprsştuüvyz
	return isPangram(strings.ToLowerSpecial(unicode.TurkishCase, s), "abcçdefgğhıijklmnoöprsştuüvyz")
}

func isPangram(s, alphabet string) bool {
	for _, c := range alphabet {
		if !strings.Contains(s, string(c)) {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println("EN Pangram Test 1:", IsPangramEN("The quick brown fox jumps over the lazy dog"))
	fmt.Println("EN Pangram Test 2:", IsPangramEN("The quick brown fo jumps over the lazy dog"))
	fmt.Println("TR Pangram Test 1:", IsPangramTR("Pijamalı hasta yağız şoföre çabucak güvendi"))
	fmt.Println("TR Pangram Test 2:", IsPangramTR("Pijamalı hasta yağız şföre çabucak güvendi"))
}
```

**Açıklama:**

- Her bir alfabe harfi, yazı içerisinde arama yapılarak eksik olup olmadığı kontrol edilir.
- Bu yöntem, hem İngilizce hem de Türkçe gibi farklı alfabeler için kullanılabilir.

---

## 18. Sınıf Çalışması: Palindrome (Palindrom) Fonksiyonu

**Tanım:**

- **Palindrom:** Yazının, alfabe dışı karakterler (boşluk, noktalama vb.) hariç ters çevrildiğinde orijinal haliyle aynı olması durumudur.
- **Örnekler:**
  - "Anastas mum satsana"
  - "Ey Edip Adana'da pide ye"
  - "Ali Papila"

**Detaylar:**

- Yazı içerisindeki sadece harfler ve rakamlar seçilerek, diğer karakterler elenir.
- Harfler küçük harfe çevrilir; bu sayede büyük/küçük harf farkı ortadan kalkar.
- Filtrelenmiş yazı, ters çevrilip orijinal haliyle karşılaştırılır.

**Örnek Kod:**

```go
package main

import (
	"fmt"
	"strings"
	"unicode"
)

func IsPalindrome(s string) bool {
	// Filtreleme: Sadece harf ve rakamları içerecek şekilde al.
	var filtered []rune
	for _, r := range s {
		if unicode.IsLetter(r) || unicode.IsDigit(r) {
			filtered = append(filtered, unicode.ToLower(r))
		}
	}

	// Ters çevrilen dizi ile orijinal dizinin karşılaştırılması:
	n := len(filtered)
	for i := 0; i < n/2; i++ {
		if filtered[i] != filtered[n-1-i] {
			return false
		}
	}
	return true
}

func main() {
	tests := []string{
		"Anastas mum satsana",
		"Ey Edip Adana'da pide ye",
		"Ali Papila",
		"Merhaba Dünya",
	}
	for _, t := range tests {
		fmt.Printf("'%s' palindrom mu? %v\n", t, IsPalindrome(t))
	}
}
```

**Ek Açıklamalar:**

- Filtreleme aşaması, sadece alfanümerik karakterlerin işleme dahil edilmesini sağlar.
- Bu yöntem, noktalama işaretleri, boşluklar veya diğer semboller yüzünden hatalı karşılaştırmayı engeller.

---

## Özet ve Sonuç

- **Go'da Yazı İşlemleri:**
  - Go dilinde yazılar (string – yazı) UTF-8 kodlamasıyla saklanır ve immutable (değiştirilemez) olduklarından, herhangi bir değişiklik yeni bir nesne oluşturur.
- **Temel Fonksiyonlar ve Paketler:**
  - `len()`, `for range`, `+` operatörü gibi temel araçlar, string işlemlerinde yaygın olarak kullanılır.
  - `strings` ve `unicode` paketleri, yazı üzerinde gelişmiş işlemler yapmayı sağlar.
- **Performans ve Bellek Yönetimi:**
  - Yoğun string işlemlerinde `strings.Builder` gibi yapıların kullanımı, performansı artırır.
  - Aynı string literal’in paylaşılması, bellek verimliliği sağlar.
- **Uygulama Örnekleri:**
  - Pangram ve palindrom kontrolü gibi fonksiyonlar, yazı işlemleri konusundaki yöntemlerin pratikte nasıl uygulanabileceğini göstermektedir.
- **Dil Desteği ve Özel Dönüşümler:**
  - `ToUpperSpecial` ve `ToLowerSpecial` gibi fonksiyonlarla, Türkçe gibi dillerin özel dönüşüm kuralları doğru şekilde uygulanır.
