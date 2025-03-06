## 1. Go Deyimleri (Statements – deyimler)

Go programlama dilinde, bir programın çalıştırılabilir parçası **statement (deyim – statement)** olarak adlandırılır. Go’da program akışı, çeşitli deyimlerin (if, for, switch, vs.) sıralı veya iç içe çalıştırılması ile gerçekleşir. Bu yapı, hem basit hem de karmaşık algoritmaların ve uygulama akışlarının oluşturulmasında temel rol oynar. Aşağıdaki başlıklar, Go’daki deyim türlerini ve her birinin işlevini detaylandırmaktadır:

- **Basit Deyimler (Simple Statements – basit deyimler):**

  - **Tanım:** Bir ifadenin sonuna noktalı virgül eklenip satır sonu ile ayrılmasıyla oluşur.
  - **Özellikler:**
    - Tek bir işlemi ifade eder; örneğin, bir değişkene değer atama veya bir fonksiyon çağrısı.
    - Go derleyicisi, satır sonlarını otomatik olarak noktalı virgül olarak yorumlar, bu nedenle kodda noktalı virgül yazmaya gerek kalmaz.
  - **Örnek:**
    ```go
    x = 5
    fmt.Println(x)
    ```
  - **Detaylı Açıklama:**
    - Basit deyimler, özellikle tek satırda gerçekleştirilebilen işlemler için kullanılır.
    - İfadelerin sade ve okunabilir olması, hata ayıklama sürecinde kolaylık sağlar.

- **Bileşik Deyimler (Compound Statements – bileşik deyimler):**

  - **Tanım:** Bir blok (block – blok) içerisine yazılmış, birden fazla deyimden oluşan yapılardır.
  - **Özellikler:**
    - Her zaman `{}` süslü parantezler içerisinde tanımlanır.
    - Fonksiyon gövdeleri, kontrol yapıların (if, for, switch) içerikleri bileşik deyim örnekleridir.
  - **Örnek:**
    ```go
    {
        fmt.Println("Merhaba")
        fmt.Println("Dünya")
    }
    ```
  - **Detaylı Açıklama:**
    - Bileşik deyimler, kodun mantıksal bölümlerini ayırarak okunabilirliği artırır.
    - İç içe bloklar kullanılarak daha karmaşık algoritmalar oluşturulabilir.

- **Bildirim Deyimleri (Declaration Statements – bildirim deyimleri):**

  - **Tanım:** Değişken, sabit veya fonksiyon gibi elemanların tanımlanmasını sağlayan deyimlerdir.
  - **Özellikler:**
    - Bellekte gerekli alanın ayrılmasını sağlar.
    - Derleyici, bildirimler sayesinde tip kontrolü yapar ve olası hataları erken aşamada yakalar.
  - **Örnek:**
    ```go
    var a int
    const pi = 3.14
    ```
  - **Detaylı Açıklama:**
    - Bildirim deyimleri, kodun başlangıcında ve fonksiyonların başında yer alır.
    - İyi tanımlanmış değişken bildirimleri, kodun anlaşılabilirliğini ve sürdürülebilirliğini artırır.

- **Kontrol Deyimleri (Control Statements – kontrol deyimleri):**

  - **Tanım:** Program akışını yönlendiren, koşul ve döngü yapıları gibi yapı taşlarını içerir.
  - **Özellikler:**
    - **if, for, switch, select** gibi yapılar bu kategoriye girer.
    - Her biri, belirli durumlara göre programın akışını değiştirme imkanı sunar.
  - **Örnek:**
    ```go
    if a > 0 {
        fmt.Println("Pozitif")
    }
    ```
  - **Detaylı Açıklama:**
    - Kontrol deyimleri, dinamik koşullara ve algoritmik yapılara göre esnek kod yazımına olanak tanır.
    - Uygulamanın hata yönetimi ve mantıksal akışında kritik rol oynar.

- **Boş Deyim (Empty/Null Statement – boş deyim):**
  - **Tanım:** Tek başına noktalı virgül (;) şeklinde yazılır ve çalıştırıldığında hiçbir işlem yapmaz.
  - **Özellikler:**
    - Genellikle otomatik eklenen boş deyimler, kodun çalışmasını etkilemez.
  - **Örnek:**
    ```go
    ;
    ```
  - **Detaylı Açıklama:**
    - Boş deyimler, özellikle döngü yapılarında, belirli bir durumda hiçbir işlem yapılmaması gerektiğinde kullanılabilir.

Aşağıdaki tablo, Go dilinde kullanılan deyim türlerini özetlemektedir:

| **Deyim Türü**                              | **Açıklama**                                                                             | **Örnek**                                          |
| ------------------------------------------- | ---------------------------------------------------------------------------------------- | -------------------------------------------------- |
| Basit Deyimler (Simple Statements)          | Tek bir ifadeyi içerir, satır sonu ile ayrılır.                                          | `x = 5`                                            |
| Bileşik Deyimler (Compound Statements)      | Birden fazla deyimi kapsayan blok yapılarıdır.                                           | `{ fmt.Println("Merhaba"); fmt.Println("Dünya") }` |
| Bildirim Deyimleri (Declaration Statements) | Değişken, sabit ya da fonksiyon tanımlamalarını içerir; bellekte yer ayrılmasını sağlar. | `var a int`                                        |
| Kontrol Deyimleri (Control Statements)      | Program akışını değiştiren yapılar (if, for, switch vb.) içerir.                         | `if a > 0 { ... }`                                 |
| Boş Deyim (Empty/Null Statement)            | Hiçbir işlem yapmayan, sadece noktalı virgül olarak yazılan deyimdir.                    | `;`                                                |

---

## 2. if Statement (if deyimi – if statement) Kullanımı

**if statement (if deyimi)**, Go dilinde en temel kontrol yapılarından biridir. Kod akışını belirli koşullara göre yönlendirmek için kullanılır. Aşağıda, if deyiminin yapısı, kullanım detayları ve ek notlar yer almaktadır:

- **Genel Biçim:**  
  Go’da if deyimi, koşul ifadesiyle başlar ve çalışacak kod bloğu mutlaka `{}` süslü parantezler içerisinde yazılır.

  ```go
  if <bool ifadesi> {
      // Koşul sağlandığında çalışacak deyimler
  } else {
      // Koşul sağlanmadığında çalışacak deyimler
  }
  ```

  _Not:_ Koşul ifadesi parantez içine alınmadan direkt yazılır. Bu durum, kodun sade ve okunabilir olmasını sağlar.

- **Detaylandırılmış Açıklama ve İpuçları:**

  - **Koşul İfadesi:**
    - İfade, **bool (boolean – bool)** türünde olmalıdır.
    - Eğer koşul ifadesinde bir değişken tanımlanacaksa, if ifadesinin başında tanımlama yapılabilir (örneğin, `if x := hesapla(); x > 0 { ... }`).
  - **Blok Yapısı:**
    - Her durumda, if bloğu bir bileşik deyim olarak yazılmalıdır.
    - Tek satırlık ifadelerde bile süslü parantezlerin kullanılması kod tutarlılığı açısından önemlidir.
  - **else-if Zinciri:**
    - Birden fazla koşulun kontrolü gerektiğinde, okunabilirliği artırmak için `else if` yapısı tercih edilmelidir.
  - **Nesne İçinde if Kullanımı:**
    - if deyimleri, fonksiyonlar veya diğer kontrol yapılarının içerisinde iç içe kullanılabilir.

- **Kod Örnekleri:**

  - **Temel if-else Kullanımı:**

    ```go
    package main

    import "fmt"

    func main() {
        var a int
        fmt.Print("Bir sayı giriniz: ")
        fmt.Scan(&a)

        if a%2 == 0 {
            fmt.Println("Çift sayı girdiniz!...") // even number
        } else {
            fmt.Println("Tek sayı girdiniz!...")   // odd number
        }

        fmt.Println("Tekrar yapıyor musunuz?")
    }
    ```

    _Açıklama:_

    - Kullanıcının girdiği sayının çift veya tek olması, if-else yapısıyla kontrol edilir.
    - `a%2 == 0` ifadesi, sayının 2 ile bölümünden kalan sıfır ise true döner.

  - **Basit if Kullanımı (else kısmı olmadan):**

    ```go
    package main

    import "fmt"

    func main() {
        var a int
        fmt.Print("Bir sayı giriniz: ")
        fmt.Scan(&a)

        if a%2 == 0 {
            a /= 2
        }

        fmt.Printf("a = %d\n", a)
        fmt.Println("Tekrar yapıyor musunuz?")
    }
    ```

    _Açıklama:_

    - Sadece çift sayı durumunda belirli bir işlem (bölme) yapılır.
    - else bloğunun olmaması, koşul sağlanmayan durumda farklı bir işlem yapılmayacağını gösterir.

  - **else if Kullanımı (daha okunabilir yapı):**

    ```go
    package main

    import "fmt"

    func main() {
        var a int
        fmt.Print("Bir sayı giriniz: ")
        fmt.Scan(&a)

        if a > 0 {
            fmt.Println("Pozitif sayı girdiniz")
        } else if a == 0 {
            fmt.Println("Sıfır girdiniz")
        } else {
            fmt.Println("Negatif girdiniz")
        }

        fmt.Println("Tekrar yapıyor musunuz?")
    }
    ```

    _Açıklama:_

    - Sayının pozitif, sıfır veya negatif olma durumları net olarak ayrılarak kontrol edilir.
    - `else if` yapısı, birbirine bağlı koşulların zincirlenmesiyle kodun okunabilirliğini artırır.

  - **Koşulların Ayrık Olması Durumunda (daha kötü bir teknik örneği):**

    ```go
    package main

    import "fmt"

    func main() {
        var a int
        fmt.Print("Bir sayı giriniz: ")
        fmt.Scan(&a)

        if a > 0 {
            fmt.Println("Pozitif sayı girdiniz")
        }
        if a == 0 {
            fmt.Println("Sıfır girdiniz")
        }
        if a < 0 {
            fmt.Println("Negatif girdiniz")
        }

        fmt.Println("Tekrar yapıyor musunuz?")
    }
    ```

    _Açıklama:_

    - Ayrık if deyimleri kullanılması, mantıksal hata riskini artırabilir çünkü aynı anda birden fazla if bloğunun çalışması mümkün olabilir.
    - Kodun bakım ve okunabilirliği açısından else-if zinciri tercih edilmelidir.

- **Ek Not:**
  - if deyiminde, koşul ifadesinin yanında değişken tanımlaması yapılabilir. Bu, özellikle kısa süreli kullanılacak değerlerde avantaj sağlar.
  - Örneğin:
    ```go
    if x := hesapla(); x > 10 {
        fmt.Println("x değeri 10'dan büyük")
    }
    ```

Aşağıdaki tablo, **if statement (if deyimi)** kullanımındaki farklı teknikleri özetlemektedir:

| **Yöntem**         | **Açıklama**                                               | **Avantajlar / Dezavantajlar**                                                |
| ------------------ | ---------------------------------------------------------- | ----------------------------------------------------------------------------- |
| if-else            | Koşul doğruysa bir blok, yanlışsa diğer blok çalıştırılır. | Okunabilirlik yüksek, mantıksal bütünlük sağlar.                              |
| else if zinciri    | Birden fazla koşulun kontrolü için zincirleme yapılır.     | Kodun okunabilirliği ve sürdürülebilirliği artar; hata yapma olasılığı düşer. |
| Ayrık if deyimleri | Her koşul için ayrı ayrı if kullanılması.                  | Koşulların çakışması ve istenmeyen çoklu çalıştırma riski; bakımı zor.        |

---

## 3. Boolean İfadeler (Boolean Expressions – boolean ifadeler) ve Kullanım İpuçları

Go dilinde koşul ifadeleri **bool (boolean – bool)** türünden olmalıdır. Bu durum, kontrol yapıların doğruluk değerine dayalı çalışmasını sağlar. Aşağıda boolean ifadelerle ilgili detaylı ipuçları yer almaktadır:

- **Doğrudan Kullanım:**

  - Boolean ifadeler if deyiminde doğrudan kullanılmalıdır.
  - Örneğin:
    ```go
    if even {
        fmt.Println("Even! (Çift!)")
    }
    ```
  - _Detay:_ `if even == true` yazımına gerek yoktur; bu durum kodu gereksiz yere uzatır.

- **Olumsuzlama Kullanımı:**

  - Koşulun tersine çevrilmesi için `!` operatörü kullanılır.
  - Örnek:
    ```go
    if !even {
        fmt.Println("Odd! (Tek!)")
    } else {
        fmt.Println("Even! (Çift!)")
    }
    ```
  - _Detay:_ `!` operatörü, mantıksal ifadenin tersini alır; bu sayede kodda fazladan `== false` kontrolü yapmaya gerek kalmaz.

- **Boolean İfadelerde Yazım Standartları:**

  - Kod okunabilirliğini artırmak adına, boolean ifadeleri sade tutmak önemlidir.
  - Karmaşık mantık ifadeleri, parantezler yardımıyla daha anlaşılır hale getirilebilir.

- **Örnek Kod:**

  ```go
  package main

  import "fmt"

  func main() {
      var a int
      fmt.Print("Bir sayı giriniz: ")
      fmt.Scan(&a)

      printEvenStatus(a%2 == 0)
  }

  func printEvenStatus(even bool) {
      if even { // if even ifadesi yeterlidir, if even == true yazmaya gerek yoktur.
          fmt.Println("Even! (Çift!)")
      } else {
          fmt.Println("Odd! (Tek!)")
      }
  }
  ```

- **Ek İpuçları:**
  - Boolean ifadelerde karmaşıklığı azaltmak için ara değişkenler kullanılabilir.
  - Okunabilirlik açısından, koşulların net ve anlaşılır olması kodun bakımını kolaylaştırır.

---

## 4. Predicate Fonksiyonlar: isEven ve isOdd Örnekleri

**Predicate fonksiyonlar (predicate functions – predicate fonksiyonlar)**, bir durumu kontrol edip sonuç olarak doğru (true) veya yanlış (false) değeri döndüren fonksiyonlardır. Genellikle isimlendirme olarak `is`, `has` gibi önekler kullanılır. Bu fonksiyonlar, kodda mantıksal kontrolü merkezileştirerek tekrar kullanım imkânı sunar.

- **Detaylı Açıklama:**

  - Predicate fonksiyonlar, özellikle bir koşulun defalarca kontrol edileceği durumlarda kod tekrarını azaltır.
  - Bu fonksiyonların okunabilir isimlerle tanımlanması, kodun ne yaptığı konusunda netlik sağlar.

- **İlk Aşama: isEven Fonksiyonu (Detaylı Yapı)**

  ```go
  package main

  import "fmt"

  func main() {
      var a int
      fmt.Print("Bir sayı giriniz: ")
      fmt.Scan(&a)

      if isEven(a) {
          fmt.Println("Çift sayı! (Even!)")
      } else {
          fmt.Println("Tek sayı! (Odd!)")
      }
  }

  func isEven(val int) bool {
      if val%2 == 0 {
          return true
      } else {
          return false
      }
  }
  ```

  _Açıklama:_

  - `isEven` fonksiyonu, girilen sayının çift olup olmadığını kontrol eder.
  - if-else bloğu kullanılarak durum net bir şekilde ayrılmıştır.

- **Geliştirilmiş Yazım:**

  ```go
  package main

  import "fmt"

  func main() {
      var a int
      fmt.Print("Bir sayı giriniz: ")
      fmt.Scan(&a)

      if isEven(a) {
          fmt.Println("Çift sayı! (Even!)")
      } else {
          fmt.Println("Tek sayı! (Odd!)")
      }
  }

  func isEven(val int) bool {
      if val%2 == 0 {
          return true
      }
      return false
  }
  ```

  _Açıklama:_

  - else bloğuna gerek kalmadan, koşul sağlanmazsa fonksiyon otomatik olarak false döndürür.
  - Bu yapı, gereksiz kod tekrarını önler.

- **En İyi Yazım:**

  ```go
  package main

  import "fmt"

  func main() {
      var a int
      fmt.Print("Bir sayı giriniz: ")
      fmt.Scan(&a)

      if isEven(a) {
          fmt.Println("Çift sayı! (Even!)")
      } else {
          fmt.Println("Tek sayı! (Odd!)")
      }
  }

  func isEven(val int) bool {
      return val%2 == 0
  }
  ```

  _Açıklama:_

  - Tek satırlık return ifadesiyle fonksiyonun özelliği en sade biçimde ifade edilmiş olur.
  - Kodun okunabilirliği artarken, gereksiz if-else yapılarından kaçınılır.

- **isOdd Fonksiyonu (isOdd – isOdd):**

  ```go
  package main

  import "fmt"

  func main() {
      var a int
      fmt.Print("Bir sayı giriniz: ")
      fmt.Scan(&a)

      if isOdd(a) {
          fmt.Println("Tek sayı! (Odd!)")
      } else {
          fmt.Println("Çift sayı! (Even!)")
      }
  }

  func isEven(val int) bool {
      return val%2 == 0
  }

  func isOdd(val int) bool {
      return !isEven(val)
  }
  ```

  _Açıklama:_

  - `isOdd` fonksiyonu, `isEven` fonksiyonunu tersine çevirerek çalışır; bu sayede kod tekrarından kaçınılır.
  - Fonksiyon isimleri, predicate fonksiyonlarda mantıksal kontrolün neye göre yapıldığını net bir şekilde belirtir.

Aşağıdaki tablo, predicate fonksiyonların geliştirilme sürecini özetlemektedir:

| **Fonksiyon Versiyonu**   | **Özellikler**                                                   | **Avantajlar**                                   |
| ------------------------- | ---------------------------------------------------------------- | ------------------------------------------------ |
| Detaylı if-else           | Koşul açıkça kontrol edilip, return ifadeleriyle sonuç döndürür. | Açık mantık, ancak kod tekrarı içerir.           |
| Geliştirilmiş if kontrolü | if bloğu sonrası doğrudan return kullanılır.                     | Kod uzunluğu azalır, okunabilirlik artar.        |
| Tek satırlık return       | Koşul ifadesi doğrudan return edilir.                            | En sade ve okunabilir yapı; kod tekrarını önler. |

---

## 5. Sınıf Çalışması: İkinci Dereceden Denklemin Köklerini Bulma

Bu örnekte, kullanıcıdan alınan katsayılarla (a, b, c) ikinci dereceden (quadratic equation – ikinci dereceden denklem) bir denklemin köklerini bulan bir program geliştirilmektedir. Bu örnekte, temel matematiksel işlemlerin yanı sıra kenar durum kontrolleri de göz önüne alınmıştır.

- **Algoritma Açıklaması:**

  - **Discriminant (Delta – delta):**

    - Delta, \( \Delta = b^2 - 4ac \) formülüyle hesaplanır.
    - Delta pozitif ise; iki farklı gerçek kök elde edilir.
    - Delta sıfır ise; çakışık (aynı) bir kök bulunur.
    - Delta negatif ise; gerçek sayı kümesi içinde kök bulunmaz (karmaşık kökler için farklı yöntemler gerekir).

  - **Edge Case (Kenar Durumları – edge case):**
    - Katsayı olarak girilen a değeri 0 ise; denklem ikinci dereceden olmaktan çıkar ve program uygun bir mesaj gösterir.
    - İsteğe bağlı olarak, karmaşık kök hesaplamaları için `math/cmplx` paketi kullanılabilir.

- **Kod Örneği:**

  ```go
  package main

  import (
      "fmt"
      "math"
  )

  func main() {
      runApp()
  }

  func runApp() {
      var a, b, c float64
      fmt.Print("Katsayıları giriniz (a b c): ")
      fmt.Scan(&a, &b, &c)

      // a'nın 0 olup olmadığı kontrol edilmeli; çünkü a = 0 ise denklem ikinci dereceden değildir.
      if a == 0 {
          fmt.Println("a değeri 0 olamaz, ikinci dereceden denklem değildir.")
          return
      }

      status, x1, x2 := findQuadraticEquationRoots(a, b, c)

      if status {
          fmt.Printf("x1 = %f, x2 = %f\n", x1, x2)
      } else {
          fmt.Println("Gerçek kök yok")
      }
  }

  func findQuadraticEquationRoots(a, b, c float64) (bool, float64, float64) {
      delta := b*b - 4*a*c

      if delta > 0 {
          sqrtDelta := math.Sqrt(delta)
          return true, (-b+sqrtDelta)/(2*a), (-b-sqrtDelta)/(2*a)
      }

      if delta == 0 {
          return true, -b/(2*a), -b/(2*a)
      }

      return false, 0, 0
  }
  ```

  _Açıklama:_

  - Kullanıcıdan alınan a, b ve c değerlerine göre delta hesaplanır.
  - Delta'nın değeri koşula göre değerlendirilir; gerçek kök varsa ekrana yazdırılır, yoksa hata mesajı gösterilir.
  - `if a == 0` kontrolü eklenerek, programın yanlış hesaplama yapması önlenir.

Aşağıdaki tablo, ikinci dereceden denklem çözüm sürecini özetlemektedir:

| **Durum**          | **Delta Değeri** | **Kökler**                                       | **Açıklama**                            |
| ------------------ | ---------------- | ------------------------------------------------ | --------------------------------------- |
| İki Gerçek Kök     | Δ > 0            | \( x\_{1,2} = \frac{-b \pm \sqrt{\Delta}}{2a} \) | Farklı iki gerçek kök bulunur.          |
| Çakışık Gerçek Kök | Δ = 0            | \( x*{1} = x*{2} = \frac{-b}{2a} \)              | Aynı değerde tek kök bulunur.           |
| Gerçek Kök Yok     | Δ < 0            | -                                                | Gerçek sayı kümesi içinde kök bulunmaz. |

- **Ek Detaylar:**
  - Denklem çözümü yapılırken, hesaplama sırasında oluşabilecek sayısal hataların önüne geçmek için kontroller yapılmalıdır.
  - Karmaşık kökler için ileride geliştirilebilecek alternatif yöntemler (örneğin, `cmplx.Sqrt` kullanımı) göz önüne alınabilir.

---

## 6. Sonuç ve Özet

Yukarıda Go dilindeki deyimlerin türleri, **if statement (if deyimi – if statement)** kullanım örnekleri, boolean ifadelerin en iyi kullanımı, predicate fonksiyon örnekleri ve ikinci dereceden denklem köklerini bulma gibi konular detaylı şekilde ele alınmıştır. Ek bilgilerle birlikte, her bölümde yer alan örnekler, geliştiricilerin gerçek dünyadaki uygulamalara nasıl uyarlayabileceğini gösteren pratik yaklaşımlar sunmaktadır.

**Anahtar Noktalar:**

- **Deyimler (Statements – deyimler):**

  - **Basit deyimler:** Tek satırlık ifadeler; hızlı ve net işlemler sağlar.
  - **Bileşik deyimler:** Birden fazla işlemi kapsayan bloklar; kodun mantıksal bütünlüğünü oluşturur.
  - **Bildirim deyimleri:** Değişken, sabit veya fonksiyon tanımlamalarıyla bellek yönetimi sağlar.
  - **Kontrol deyimleri:** if, for, switch gibi yapılarla program akışını yönlendirir.
  - **Boş deyim:** Hiçbir işlem yapmayan yapıdır; belirli durumlarda yer tutucu olarak kullanılabilir.

- **if Statement (if deyimi – if statement):**

  - Her zaman bileşik bloklar içinde kullanılmalı; koşul ifadesi sade ve doğrudan yazılmalıdır.
  - else-if zinciri, kodun okunabilirliğini artırır; koşulların net olarak ayrılması sağlanır.
  - Koşul ifadesinde kısa süreli değişken tanımlaması da yapılabilir.

- **Boolean İfadeler (Boolean Expressions – boolean ifadeler):**

  - Koşul ifadeleri doğrudan kullanılmalı, fazladan karşılaştırma operatörleri gereksiz yere eklenmemelidir.
  - `!` operatörü ile koşul tersine çevrilebilir; kodun sade kalması sağlanır.

- **Predicate Fonksiyonlar:**

  - `isEven` ve `isOdd` fonksiyonları, kod tekrarını önleyen ve mantıksal kontrolleri merkezileştiren örneklerdir.
  - Fonksiyonların sadeleştirilmesi, okunabilirliği ve sürdürülebilirliği artırır.

- **İkinci Dereceden Denklem Çözümü:**
  - Delta hesabı ve kenar durum kontrolleri (örneğin, a == 0) kritik öneme sahiptir.
  - Gerçek köklerin bulunamaması durumunda kullanıcıya uygun mesaj verilmelidir.
