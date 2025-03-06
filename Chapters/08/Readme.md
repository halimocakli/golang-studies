## 1. Genel Bakış ve Temel Bilgiler

- **Döngüler (Loops – döngüler):**

  - Bir işin yinelenerek yapılmasını sağlayan kontrol yapılarıdır.
  - Tekrarlanan işlemler, kod tekrarını azaltır ve programın akışını daha düzenli hale getirir.

- **Go'da For Döngüsü (for loop – for döngüsü):**

  - Go dilinde yalnızca **for döngüsü (for loop – for döngüsü)** bulunmaktadır.
  - Diğer popüler döngü yapıları; örneğin, _while_ ya da _do-while_ döngüleri Go’da ayrı bir sözdizimiyle bulunmaz.
  - For döngüsü, esnek yapısı sayesinde hem klasik for yapısı hem de while benzeri davranışları destekler.

- **Önemli Notlar:**
  - **Başlangıç Koşulu (Initialization – başlatma):** Döngüye ilk girildiğinde bir kez çalışır.
  - **Koşul İfadesi (Condition – koşul):** Her iterasyondan önce kontrol edilir; false olduğunda döngü sonlanır.
  - **Artış/Değişim (Post – artış/azalış):** Her döngü adımının sonunda koşul kontrolünden önce çalıştırılır.
  - **Sonsuz Döngüler:** Koşul ifadesi true veya boş bırakılırsa, döngü sonsuza kadar devam edebilir. Bu durumda, döngüden çıkmak için **break (break – break)** ifadesi kullanılmalıdır.

---

## 2. For Döngüsünün Farklı Kullanım Biçimleri

### 2.1. While Benzeri For Döngüsü

- **Genel Yapı:**
  ```go
  for [boolean expression] {
      // Döngü gövdesi
  }
  ```
- **Açıklama:**
  - Döngüye girmeden önce koşul ifadesi kontrol edilir.
  - İfade `true` ise döngü çalışmaya başlar, aksi halde hiç girilmez.
  - Bu yapı, diğer dillerdeki **while loop (while döngüsü – while döngüsü)** kullanımına denk gelir.
- **Detaylı Bilgiler:**

  - Eğer ilk kontrol esnasında koşul `false` ise döngü gövdesi hiç çalıştırılmaz.
  - Döngü içerisinde değişkenlerin değeri kontrolü etkileyebileceği için dikkatli güncellenmelidir.

- **Örnek Kod:**

  ```go
  package main

  import "fmt"

  func main() {
      var n int
      fmt.Print("Input a number: ")
      fmt.Scan(&n)
      i := 0

      // 'i < n' ifadesi true olduğu sürece döngü çalışır.
      for i < n {
          fmt.Printf("%d ", i)
          i++ // Değişkenin güncellenmesi koşulun sağlıklı çalışmasını sağlar.
      }
      fmt.Println()
      fmt.Println("Tekrar yapıyor musunuz?")
  }
  ```

---

### 2.2. Klasik For Döngüsü: Init, Condition, Post

- **Genel Yapı:**
  ```go
  for initialization; condition; post {
      // Döngü gövdesi
  }
  ```
- **Açıklama:**
  - **Initialization (Başlatma):** Döngüye ilk girildiğinde bir kez çalışır.
  - **Condition (Koşul):** Her iterasyondan önce kontrol edilir; koşul sağlanmazsa döngü sonlanır.
  - **Post (Artış/Azalış):** Döngü gövdesindeki işlemler tamamlandıktan sonra, bir sonraki adım öncesinde çalışır.
- **Detaylı Bilgiler:**

  - Bu yapı, kodun okunabilirliğini artırır ve her iterasyon sonunda değişkenin otomatik güncellenmesini sağlar.
  - Koşul kontrolü, döngü başlangıcında ve her adımda yapılır; bu sayede hatalı durumların erken yakalanması sağlanır.

- **Örnek Kod:**

  ```go
  package main

  import "fmt"

  func main() {
      var n int
      fmt.Print("Bir sayı giriniz: ")
      fmt.Scan(&n)

      // 'i' değişkeni 0'dan başlayarak n'e kadar artar.
      for i := 0; i < n; i++ {
          fmt.Printf("%d ", i)
      }
      fmt.Println()
  }
  ```

---

### 2.3. Sonsuz Döngüler (Infinite Loops)

- **Tanım:**
  - Koşul ifadesi `true` olarak yazıldığında veya koşul ifadesi hiç belirtilmediğinde döngü sonsuza kadar devam eder.
- **Kullanım Durumları:**
  - Belirli bir koşul gerçekleşene kadar sürekli çalışması gereken işlemler için idealdir.
  - Sonsuz döngü, kontrol mekanizmaları (örneğin, break veya return) ile sonlandırılmalıdır.
- **Detaylı Bilgiler:**

  - Sonsuz döngülerin yanlışlıkla unutulması, programın kilitlenmesine veya beklenmeyen sonuçlara yol açabilir.
  - İyi yapılandırılmış sonsuz döngüler, beklenen bir kullanıcı girişi veya sinyal alındığında döngüden çıkmalıdır.

- **Örnek Kod – Koşullu Sonsuz Döngü:**

  ```go
  package main

  func main() {
      // 'for true' ifadesi sonsuz döngü oluşturur.
      for true {
          // Sonsuz döngü gövdesi (örneğin, sürekli çalışan sunucu işlemleri)
      }
  }
  ```

- **Örnek Kod – Koşulsuz Sonsuz Döngü:**

  ```go
  package main

  func main() {
      // 'for {}' ifadesi de sonsuz döngü oluşturur.
      for {
          // Döngü gövdesi
      }
  }
  ```

---

### 2.4. Döngüden Çıkış: Break Kullanımı

- **Break İfadesi (break – break):**
  - Döngüyü anında sonlandırmak için kullanılır.
  - Döngü gövdesi içerisinde belirli bir koşul gerçekleştiğinde, **break** ifadesi döngüden çıkışı sağlar.
- **Detaylı Bilgiler:**
  - Break, sadece en içteki döngüyü sonlandırır; çok katmanlı döngülerde dikkatli kullanılmalıdır.
  - Döngüden çıkmak yerine belirli durumlarda **continue (devam et – continue)** kullanımı da düşünülebilir; fakat burada break’e odaklanıyoruz.
- **Örnek Kod:**

  ```go
  package main

  import "fmt"

  func main() {
      runFindTotalApp()
  }

  func runFindTotalApp() {
      fmt.Println("Begin to enter values:")
      fmt.Printf("Total = %d\n", findTotal())
  }

  func findTotal() int {
      var val int
      total := 0

      for {
          fmt.Scan(&val)
          if val == 0 {
              break // Döngüden çıkış
          }
          total += val
      }

      return total
  }
  ```

---

### 2.5. Döngü Değişkenlerinin Kapsamı (Scope) ve Maskelenme (Shadowing)

- **Açıklama:**
  - For döngüsünde tanımlanan döngü değişkenlerinin kapsamı yalnızca döngü içerisindedir.
  - Eğer döngü dışında da aynı isimde bir değişken tanımlandıysa, döngü içerisindeki değişken dıştaki tanımı **maskeler (shadowing – maskelenme)**.
- **Detaylı Bilgiler:**
  - Kapsam yönetimi, kodun okunabilirliğini ve hata ayıklamasını etkiler.
  - Değişken isimlerinin çakışmaması için dikkatli olmak gerekir.
- **Örnek Kod – Hata Oluşan Durum:**

  ```go
  package main

  import "fmt"

  func main() {
      var n int
      fmt.Print("Bir sayı giriniz: ")
      fmt.Scan(&n)

      for i := 0; i < n; i++ {
          fmt.Printf("%d ", i)
      }
      // Aşağıdaki satır hata verir çünkü 'i' değişkeni döngü dışındadır.
      // fmt.Printf("i = %d", i)
      fmt.Println()
  }
  ```

- **Örnek Kod – Maskelenmeye Örnek:**

  ```go
  package main

  import "fmt"

  func main() {
      var n int
      fmt.Print("Bir sayı giriniz: ")
      fmt.Scan(&n)

      i := 67  // Döngü dışındaki 'i'
      for i := 0; i < n; i++ {  // Bu 'i', sadece döngü kapsamında geçerli olup dıştaki 'i'yi maskeler.
          fmt.Printf("%d ", i)
      }
      fmt.Printf("\ni = %d\n", i)  // Burada dışarıdaki 'i' (67) kullanılır.
  }
  ```

---

### 2.6. For Döngüsünün Esnek Kullanım Biçimleri

Go dilinde for döngüsü, klasik yapıdaki üç bölümün (init, condition, post) dışında esnek biçimlerde de kullanılabilir:

- **1. Kısım Boş Bırakılması:**

  - Başlatma işlemi döngüden önce tanımlanabilir.

  ```go
  package main

  import "fmt"

  func main() {
      var n int
      fmt.Print("Bir sayı giriniz: ")
      fmt.Scan(&n)

      i := 0
      // Başlatma kısmı döngü dışındaki 'i' kullanılarak yapıldı.
      for ; i < n; i++ {
          fmt.Printf("%d ", i)
      }
      fmt.Printf("\nAfter loop: i = %d\n", i)
  }
  ```

- **3. Kısım Boş Bırakılması:**

  - Artış işlemi döngü gövdesinde manuel olarak yapılır.

  ```go
  package main

  import "fmt"

  func main() {
      var n int
      fmt.Print("Bir sayı giriniz: ")
      fmt.Scan(&n)

      for i := 0; i < n; { // Post kısmı boş bırakıldı.
          fmt.Printf("%d ", i)
          i++ // Artış işlemi döngü gövdesinde gerçekleştirildi.
      }
  }
  ```

- **Koşul İfadesinin Boş Bırakılması:**

  - Koşul boş bırakıldığında döngü sonsuza kadar çalışır; okunabilirlik açısından, sonsuz döngülerde **for { ... }** kullanımı tercih edilir.

  ```go
  package main

  import "fmt"

  func main() {
      var n int
      fmt.Print("Bir sayı giriniz: ")
      fmt.Scan(&n)

      // 2. kısım boş bırakıldığı için sonsuz döngü oluşur.
      for i := 0; ; i++ {
          fmt.Printf("%d\n", i)
          if i >= n-1 {
              break // Sonsuz döngüden çıkmak için kontrol eklenir.
          }
      }
  }
  ```

- **Tablo: For Döngüsü Kullanım Varyasyonları**

  | **Kullanım Biçimi**                 | **Açıklama**                                                              | **Örnek Kullanım**                |
  | ----------------------------------- | ------------------------------------------------------------------------- | --------------------------------- |
  | `for condition { ... }`             | While döngüsü benzeri kullanım.                                           | `for i < n { ... }`               |
  | `for init; condition; post { ... }` | Klasik for döngüsü; başlatma, koşul ve artış/azalış işlemleri belirtilir. | `for i := 0; i < n; i++ { ... }`  |
  | `for { ... }`                       | Tüm kısımlar boş bırakılarak sonsuz döngü oluşturulur.                    | `for { ... }`                     |
  | `for init; condition; { ... }`      | Post kısmı boş bırakılarak döngü gövdesinde artış işlemi yapılır.         | `for i := 0; i < n; { ...; i++ }` |

---

## 3. Sınıf Çalışmaları: Uygulamalı Fonksiyonlar

Aşağıdaki örnekler, **for loop (for döngüsü – for döngüsü)** kullanılarak yazılmış çeşitli fonksiyon uygulamalarını göstermektedir. Her bir örnekte, fonksiyonun amacı, kullanılan algoritma ve kritik noktalar açıklanmıştır.

### 3.1. Basamak Sayısı Hesaplama: `digitsCount`

- **Amaç:**
  - Bir int türden sayının kaç basamaklı olduğunu belirlemek.
  - **Özel Durum:** Sıfır (0) için basamak sayısı 1 olarak kabul edilir.
- **Detaylı Açıklamalar:**

  - Döngü, sayının 0 olmadığı sürece her adımda sayıyı 10’a bölerek basamak sayısını artırır.
  - İki farklı yöntemle, koşul kullanılarak veya sonsuz döngüyle break kullanılarak yazılabilir.

- **Örnek Kod (While Benzeri Biçim):**

  ```go
  package main

  import "fmt"

  func main() {
      runDigitsCountTest()
  }

  func runDigitsCountTest() {
      for {
          var val int
          fmt.Print("Bir sayı giriniz: ")
          fmt.Scan(&val)
          fmt.Printf("%d sayısının basamak sayısı: %d\n", val, digitsCount(val))
          if val == 0 {
              break
          }
      }
      fmt.Println("Tekrar yapıyor musunuz?")
  }

  func digitsCount(val int) int {
      if val == 0 {
          return 1
      }
      count := 0
      for val != 0 {
          count++
          val /= 10
      }
      return count
  }
  ```

- **Örnek Kod (Break Kullanılarak):**

  ```go
  func digitsCount(val int) int {
      count := 0
      for {
          count++
          val /= 10
          if val == 0 {
              break
          }
      }
      return count
  }
  ```

---

### 3.2. Sayının Tersini Hesaplama: `reverseNumber`

- **Amaç:**
  - Girilen sayının basamaklarını tersine çevirerek yeni bir sayı oluşturmak.
- **Detaylı Açıklamalar:**

  - Döngü, sayının son basamağını almak için mod (%) operatörü kullanır.
  - Her adımda elde edilen basamak, mevcut sonuç değeri ile çarpılarak ve eklenerek ters sayı oluşturulur.

- **Örnek Kod:**

  ```go
  package main

  import "fmt"

  func main() {
      runReverseNumberTest()
  }

  func runReverseNumberTest() {
      for {
          var val int
          fmt.Print("Bir sayı giriniz: ")
          fmt.Scan(&val)
          fmt.Printf("%d sayısının tersi: %d\n", val, reverseNumber(val))
          if val == 0 {
              break
          }
      }
      fmt.Println("Tekrar yapıyor musunuz?")
  }

  func reverseNumber(val int) int {
      result := 0
      for val != 0 {
          result = result*10 + val%10
          val /= 10
      }
      return result
  }
  ```

---

### 3.3. Basamaklar Toplamı Hesaplama: `digitsSum`

- **Amaç:**
  - Girilen bir sayının basamaklarının toplamını hesaplamak.
  - Negatif sayıların basamak toplamı pozitif olarak elde edilir.
- **Detaylı Açıklamalar:**

  - Döngü, her adımda sayının son basamağını toplam değişkenine ekler ve sayıyı 10’a böler.
  - Sonuç negatif ise, ayrı bir fonksiyon ile pozitif hale çevrilir.

- **Örnek Kod:**

  ```go
  package main

  import "fmt"

  func main() {
      runDigitsSumTest()
  }

  func runDigitsSumTest() {
      for {
          var val int
          fmt.Print("Bir sayı giriniz: ")
          fmt.Scan(&val)
          fmt.Printf("%d sayısının basamakları toplamı: %d\n", val, digitsSum(val))
          if val == 0 {
              break
          }
      }
      fmt.Println("Tekrar yapıyor musunuz?")
  }

  func digitsSum(val int) int {
      total := 0
      for val != 0 {
          total += val % 10
          val /= 10
      }
      return absInt(total)
  }

  func absInt(val int) int {
      if val < 0 {
          return -val
      }
      return val
  }
  ```

---

### 3.4. Faktoriyel Hesaplama: `factorial`

- **Amaç:**
  - Girilen bir sayının faktoriyel değerini hesaplamak.
  - **Özel Durum:** Negatif değerler için 1 değeri döndürülür.
- **Detaylı Açıklamalar:**

  - Faktoriyel, 0! = 1 ve 1! = 1 şeklinde tanımlanır; 2! = 1 _ 2, 3! = 1 _ 2 \* 3 şeklinde artarak devam eder.
  - İki farklı yöntemle, artan ya da azalan iterasyon kullanılarak hesaplanabilir.

- **Örnek Kod – Artan Döngü Kullanımı:**

  ```go
  package main

  import "fmt"

  func main() {
      runFactorialTest()
  }

  func runFactorialTest() {
      for n := -1; n <= 10; n++ {
          fmt.Printf("%d! = %d\n", n, factorial(n))
      }
  }

  func factorial(n int) int {
      result := 1
      for i := 2; i <= n; i++ {
          result *= i
      }
      return result
  }
  ```

- **Örnek Kod – Azalan Döngü Kullanımı:**

  ```go
  func factorial(n int) int {
      result := 1
      for ; n > 1; n-- {
          result *= n
      }
      return result
  }
  ```

---

### 3.5. Fibonacci Sayıları Hesaplama

#### 3.5.1. n'inci Fibonacci Sayısını Hesaplama: `fibonacciNumber`

- **Amaç:**
  - Fibonacci serisinde, pozitif bir n değeri için n'inci sayıyı döndürmek.
- **Fibonacci Serisi:**
  - 0, 1, 1, 2, 3, 5, 8, 13, 21, … şeklinde devam eder.
- **Detaylı Açıklamalar:**
  - İlk iki Fibonacci sayısı özel olarak ele alınır (n ≤ 2 durumunda).
  - Döngü, önceki iki değeri toplayarak sonraki Fibonacci sayısını hesaplar.
- **Örnek Kod:**

  ```go
  package main

  import "fmt"

  func main() {
      runFibonacciNumberTest()
  }

  func runFibonacciNumberTest() {
      for {
          var n int
          fmt.Print("n değerini giriniz: ")
          fmt.Scan(&n)
          if n <= 0 {
              break
          }
          fmt.Printf("%d. Fibonacci sayısı: %d\n", n, fibonacciNumber(n))
      }
      fmt.Println("Tekrar yapıyor musunuz?")
  }

  func fibonacciNumber(n int) int {
      if n <= 2 {
          return n - 1
      }
      prev1, prev2 := 1, 0
      value := prev1 + prev2
      for i := 3; i < n; i++ {
          prev2 = prev1
          prev1 = value
          value = prev1 + prev2
      }
      return value
  }
  ```

#### 3.5.2. Bir Sayıdan Büyük İlk Fibonacci Sayısını Bulma: `nextFibonacciNumber`

- **Amaç:**
  - Girilen bir değerden büyük ilk Fibonacci sayısını bulmak.
- **Detaylı Açıklamalar:**
  - Fibonacci serisi artarak devam ettiğinden, döngü içinde bir sonraki sayı bulunana kadar hesaplama devam eder.
- **Örnek Kod:**

  ```go
  package main

  import "fmt"

  func main() {
      runNextFibonacciNumberTest()
  }

  func runNextFibonacciNumberTest() {
      for {
          var val int
          fmt.Print("Bir sayı giriniz: ")
          fmt.Scan(&val)
          fmt.Printf("%d sayısından büyük ilk Fibonacci sayısı: %d\n", val, nextFibonacciNumber(val))
          if val == 0 {
              break
          }
      }
      fmt.Println("Tekrar yapıyor musunuz?")
  }

  func nextFibonacciNumber(val int) int {
      if val < 0 {
          return 0
      }
      prev1, prev2 := 1, 0
      var next int
      for {
          next = prev1 + prev2
          if next > val {
              return next
          }
          prev2 = prev1
          prev1 = next
      }
  }
  ```

---

### 3.6. Asal Sayı Kontrolü ve n'inci Asal Sayıyı Bulma

#### 3.6.1. Asal Olup Olmadığını Kontrol Etme: `isPrime`

- **Amaç:**
  - Girilen bir sayının asal olup olmadığını kontrol etmek.
- **Detaylı Açıklamalar:**
  - Temel yaklaşımda, 2'den başlayarak sayının yarısına kadar bölünebilirliği kontrol edilir.
  - İkinci yaklaşımda, Eratosten eleme prensibine dayalı olarak, sayının kareköküne kadar kontrol yapılır; bu yöntem daha verimli çalışır.
- **Örnek Kod – Temel Yaklaşım:**

  ```go
  package main

  import "fmt"

  func main() {
      runIsPrimeTest()
  }

  func runIsPrimeTest() {
      for val := 1; val <= 100; val++ {
          if isPrime(val) {
              fmt.Printf("%d ", val)
          }
      }
      fmt.Println()
      fmt.Println(isPrime(1_000_003))
      fmt.Println("Tekrar yapıyor musunuz?")
  }

  func isPrime(val int) bool {
      if val <= 1 {
          return false
      }
      for i := 2; i <= val/2; i++ {
          if val%i == 0 {
              return false
          }
      }
      return true
  }
  ```

- **Örnek Kod – Eratosten Eleme Yaklaşımı:**

  ```go
  func isPrime(val int) bool {
      if val <= 1 {
          return false
      }
      if val%2 == 0 {
          return val == 2
      }
      if val%3 == 0 {
          return val == 3
      }
      if val%5 == 0 {
          return val == 5
      }
      if val%7 == 0 {
          return val == 7
      }
      for i := 11; i*i <= val; i += 2 {
          if val%i == 0 {
              return false
          }
      }
      return true
  }
  ```

#### 3.6.2. n'inci Asal Sayıyı Bulma: `prime`

- **Amaç:**
  - Pozitif bir n değeri için, n'inci asal sayıyı bulmak.
- **Detaylı Açıklamalar:**
  - Döngü, 2’den başlayarak her sayıyı kontrol eder ve asal olanların sayısını sayar.
  - İstenen n değerine ulaşıldığında o asal sayı döndürülür.
- **Örnek Kod:**

  ```go
  package main

  import "fmt"

  func main() {
      runPrimeTest()
  }

  func runPrimeTest() {
      for {
          var n int
          fmt.Print("Bir sayı giriniz: ")
          fmt.Scan(&n)
          if n <= 0 {
              break
          }
          fmt.Printf("%d. asal sayı: %d\n", n, prime(n))
      }
      fmt.Println("Tekrar yapıyor musunuz?")
  }

  func prime(n int) int {
      count := 0
      val := 2
      for {
          if isPrime(val) {
              count++
          }
          if count == n {
              return val
          }
          val++
      }
  }
  ```

  _Not:_ Yukarıdaki örnekte kullanılan `isPrime` fonksiyonu, daha verimli olan Eratosten eleme yaklaşımını temel alır.

---

## 4. Özet ve Sonuç

- **For Döngüsü (for loop – for döngüsü) Esnekliği:**
  - Go dilinde tüm döngü ihtiyaçları için tek bir anahtar kelime (for) kullanılır.
  - Hem _while benzeri_ yapıda, hem klasik for yapısında hem de sonsuz döngü şeklinde kullanılabilir.
  - **Önemli Noktalar:**
    - **Başlatma (Initialization – başlatma):** Döngüye girişte bir kez çalışır.
    - **Koşul (Condition – koşul):** Her adımda kontrol edilir; sağlanmazsa döngü sonlanır.
    - **Artış/Azalış (Post – artış/azalış):** Her iterasyonda otomatik güncelleme sağlar.
- **Kritik Konular:**

  - **Döngüden Çıkış:** Sonsuz döngülerde **break (break – break)** ifadesinin doğru konumlandırılması önemlidir.
  - **Değişken Kapsamı (Scope):** Döngü içerisindeki değişkenlerin kapsamı, dışarıdaki değişkenlerle çakışmaları engellemek için iyi yönetilmelidir.
  - **Performans:** Özellikle asal sayı, Fibonacci ve faktoriyel gibi hesaplamalarda kullanılan algoritmalar, büyük sayılar için performans etkilerini göz önüne almalıdır.

- **Go’nun Kullanım Alanları:**

  - **Sistem Programlama ve Ağ Programlama:** Döngüler, düşük seviye işlemler, veri aktarımı ve eşzamanlılık (concurrency – eşzamanlılık) uygulamalarında kritik rol oynar.
  - **Web Servisleri ve Mikroservisler:** Sonsuz döngü yapıları, sunucu dinleme işlemlerinde ve arka plan işlemlerinde yaygın olarak kullanılır.
  - **CLI Araçları ve Otomasyon:** Kullanıcıdan sürekli veri alıp işleme döngüleri, etkileşimli uygulamalarda tercih edilir.

- **Sonuç:**  
  Go dilinde **for döngüsü (for loop – for döngüsü)**, basit sözdizimi, esnek kullanımı ve güçlü kontrol yapıları sayesinde modern yazılım geliştirme ihtiyaçlarına cevap verir. Yukarıda verilen detaylı açıklamalar, kod örnekleri ve ek açıklamalar; for döngüsünün farklı biçimlerini, kullanım varyasyonlarını ve pratik uygulamalarını net şekilde ortaya koymaktadır.
