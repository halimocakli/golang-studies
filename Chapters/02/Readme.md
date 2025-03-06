## 1. Boşluk Karakterleri (Whitespace Characters – Boşluk Karakterleri)

- **Tanım ve Örnekler:**

  - Boşluk karakterleri, klavyeden basıldığında görsel olarak boşluk hissi yaratan karakterlerdir.
  - Tipik örnekler: **SPACE** (SPACE – boşluk), **TAB** (TAB – sekme) ve **ENTER** (ENTER – satır sonu) karakterleridir.

- **Go’da Boşluk Kullanımının Özellikleri:**
  - Go, diğer bazı programlama dillerine göre sözdizimi (syntax – sözdizimi) açısından çok esnek değildir.
  - Yazım sırasında boşluk karakterlerinin yeri, kodun okunabilirliği ve yorumlanması açısından önem taşır.
  - Kodun otomatik formatlanmasını sağlayan **go fmt** (go fmt – kod formatlama aracı) gibi araçlar sayesinde, gereksiz veya tutarsız boşluk kullanımı otomatik olarak düzeltilebilir.

---

## 2. Yorum Satırları (Comment Lines – Yorum Satırları)

- **Yorum Satırlarının Tanımı ve Türleri:**

  - Yorum satırları, derleyicinin (compiler – derleyici) ve yorumlayıcının (interpreter – yorumlayıcı) göz ardı ettiği, kod hakkında açıklama yapan metinlerdir.
  - Go dilinde iki tür yorum satırı bulunur:
    - **Tek satırlık yorum:** `//` ile başlayan yorum satırları.
    - **Çok satırlı yorum:** `/* ... */` arasına yazılan yorumlar.

- **Önemli Notlar ve İpuçları:**

  - Yorum satırları, kodun anlaşılabilirliğini artırmak amacıyla kullanılmalıdır.
  - Gereksiz ve dağınık yorum satırları, kodun okunabilirliğini olumsuz etkileyebilir.
  - İç içe çok satırlı yorumlarda dikkatli olunmalıdır; Go’da iç içe yorum satırları desteklenmez.

- **Örnek:**

  ```go
  package main

  import "fmt"

  func main() {
      fmt.Println("Hello, World") // Bu satır yorum satırı olup derleyici tarafından göz ardı edilir.

      /*
         Bu da çok satırlı bir yorumdur.
         Burada yazılanlar derleme sürecinde kullanılmaz.
      */
  }
  ```

---

## 3. Bildirim (Declaration – Bildirim) ve Program Yapısı

- **Bildirim Nedir?**

  - Bildirim, bir ismin derleyiciye tanıtılmasıdır. Derleyici, bildirimler sayesinde kod içinde kullanılan isimleri (değişken, fonksiyon, paket vb.) tanır.

- **Go Programlarının Temel Yapısı:**

  - **Paket Bildirimi:** Her Go programı en az bir paket bildirimi ile başlar. Örneğin:

    ```go
    package main
    ```

    - Burada `main` paketi, programın çalıştırılabilir bir uygulama oluşturması için özel bir isimdir.

  - **Fonksiyon Bildirimi:**
    - Fonksiyon bildiriminin genel biçimi:
      ```go
      func <fonksiyon ismi>([parametreler]) [geri dönüş bilgisi] {
          // Fonksiyonun gövdesi
      }
      ```
    - **Örnek:**
      ```go
      func foo() {
          fmt.Println("foo")
      }
      ```
    - **Fonksiyon Çağrısı:**
      - Aynı paket içerisindeki fonksiyonlar, doğrudan çağrılabilir:
        ```go
        foo()
        ```
      - Farklı paketlerde yer alan fonksiyonlar, paket adı ile çağrılır:
        ```go
        test.Foo()
        ```

- **İsimlendirme Kuralları ve Önemli Notlar:**

  - **Export Edilebilirlik (Exported – Dışa Açık):**
    - Fonksiyon veya değişken ismi küçük harfle başlıyorsa, sadece aynı paket içerisinde kullanılabilir.
    - İlk harfi büyük olan isimler (örn. `Bar()`) dış paketlerden erişime açıktır.
  - **Kod Düzeni:**
    - Fonksiyonun gövdesi, fonksiyon parametre parantezinin hemen aynı satırında veya bir sonraki satırda yazılabilir; fakat okunabilirlik açısından genellikle aynı satırda açılış `{` karakteri tercih edilir.

- **Örnek Program:**

  ```go
  package main

  import "fmt"

  func foo() {
      fmt.Println("foo")
      tar()
  }

  func Bar() {
      fmt.Println("Bar")
  }

  func tar() {
      fmt.Println("tar")
  }

  func main() {
      fmt.Println("Hello, World")
      foo()
      Bar()
      fmt.Println("Goodbye, World")
  }
  ```

- **Ek Açıklamalar:**
  - **Entry Point (Giriş Noktası):** Go programlarında çalışma, `main()` fonksiyonu ile başlar. Bu fonksiyon programın "entry point" (giriş noktası) olarak kabul edilir.
  - **Fonksiyon Akışı:** Bir fonksiyon çağrıldığında, kod akışı çağrılan fonksiyonun içerisine geçer ve fonksiyon tamamlandığında çağrıldığı yere geri döner.

---

## 4. Standart Girdi/Çıktı ve Process Kavramı

- **Standart Dosya Akışları:**

  - **stdin (Standart Girdi – Klavye):** Genellikle kullanıcıdan alınan veriler için kullanılır.
  - **stdout (Standart Çıktı – Ekran):** Program çıktılarının görüntülendiği yerdir.
  - **stderr (Standart Hata – Hata Çıktısı):** Hata mesajlarının gönderildiği akıştır. Genellikle stdout ile aynı yere yönlendirilir.

- **Örnek Kullanım:**

  - `fmt.Println` ve `fmt.Print` fonksiyonları, ekrana yazı yazar.
    - `fmt.Println`: Yazının sonunda imleci yeni satıra geçirir.
    - `fmt.Print`: Yazının sonunda imleci aynı satırda bırakır.

- **Process (Process – İşlem) Kavramı:**
  - Bir program, işletim sistemi tarafından çalıştırıldığında "process" (işlem) olarak adlandırılır. Bu, programın hafızada çalışan halidir.

---

## 5. Taşınabilirlik (Portability – Taşınabilirlik) ve Derleme Süreci

- **Kod Taşınabilirliği (Code Portability):**

  - Farklı sistemlerde, kodda değişiklik yapmadan derlenebilmesi.
  - Go gibi dillerde, kodun belirli bir platforma özgü olmaması sağlanır.

- **Program Taşınabilirliği (Program Portability):**

  - Derlenmiş programın, farklı işletim sistemlerinde (Windows, Linux, macOS) çalışabilmesi.
  - Go, **cross-compilation** (çoklu platform desteği – cross-compilation) sayesinde tek bir derleme ile farklı platformlarda çalışabilecek binary dosyalar üretebilir.

- **Derleme İşlemleri:**
  - **go build:** Kaynak kodu derleyerek platforma özgü çalıştırılabilir (executable – çalıştırılabilir) dosya üretir.
  - **go run:** Derleme sürecini arka planda gerçekleştirir ve geçici binary oluşturup çalıştırır.

---

## 6. Sayı Sistemleri ve İkilik Temsili

### 6.1. Temel Sayı Sistemleri

- **10'luk Sistem (Decimal – Decimal):**

  - 0, 1, 2, 3, 4, 5, 6, 7, 8, 9 rakamları kullanılır.
  - Her basamak, 10’un kuvvetiyle çarpılarak toplam ifade edilir.
  - Örneğin:  
    123.25 = 1×10² + 2×10¹ + 3×10⁰ + 2×10⁻¹ + 5×10⁻²

- **2'lik Sistem (Binary – İkilik):**

  - 0 ve 1 rakamları kullanılır.
  - Her basamak, 2’nin kuvvetiyle çarpılarak sayı elde edilir.
  - Bir **bit (binary digit – bit)** en küçük bilgi birimidir; 8 bit ise 1 **byte (byte – bayt)** oluşturur.
  - Bitler, okunabilirlik açısından genellikle 4’erli gruplar halinde yazılır.

- **Tablo: 1 Byte İçin Örnek Değerler**

  | İkilik Gösterim | Decimal (10'luk) |
  | --------------- | ---------------- |
  | 0000 0000       | 0                |
  | 1111 1111       | 255              |

### 6.2. İşaretli ve İşaretsiz Tamsayılar

- **İşaretsiz (Unsigned – İşaretsiz):**

  - En sol bit işaret biti olarak kullanılmaz.
  - Değer aralığı: [0, 2ⁿ – 1] (n: bit sayısı).

- **İşaretli (Signed – İşaretli):**

  - En sol bit, sayının işaretini belirler (0: pozitif, 1: negatif).
  - Negatif sayıların temsili için **ikiye tümleyen (two's complement – ikiye tümleyen)** yöntemi kullanılır.
  - İşaretli aralık: [–2^(n–1), 2^(n–1) – 1].

- **Örnek (1 Byte – 8 bit):**

  - Maksimum pozitif: 0 111 1111 → +127
  - Maksimum negatif: 1 000 0000 → –128

- **İkiye Tümleyen Hesaplama Örneği:**
  - Pozitif +10 için:
    ```plaintext
    +10: 0000 1010
    ```
  - -10’u elde etmek için +10’un ikiye tümleyenini alın:
    1. Bire tümleyen (bit flip): 1111 0101
    2. 1 ekleyin: 1111 0110 → Bu, -10 değeridir.

### 6.3. Kayan Noktalı Sayıların Temsili

- **Sabit Noktalı Format (Fixed Point – Sabit Noktalı):**

  - Noktanın yeri sabittir. Örneğin 4 byte ayrılırsa, 2 byte tam sayı kısmı, 2 byte kesirli kısım olabilir.
  - Dezavantaj: Dinamik değildir; farklı büyüklükte sayılar için esneklik sağlamaz.

- **Kayan Noktalı Format (Floating Point – Kayan Noktalı):**
  - Sayının temsilinde üç bölüm vardır:
    - **İşaret biti (sign bit – işaret biti)**
    - **Mantis (mantis – mantis)**
    - **Üstel kısım (exponential part – üstel kısım)**
  - Günümüzde yaygın olarak **IEEE 754** standardı kullanılır:
    - **Short Real Format (4 byte – float32)**
    - **Long Real Format (8 byte – float64)**
    - **Extended Real Format (10 byte – bazı platformlarda)**
  - **Yuvarlama Hatası (Rounding Error – yuvarlama hatası):**
    - Kayan noktalı sayılar, bazı sayıları tam ifade edemediğinden yuvarlama hatası oluşabilir.
    - Bu nedenle finansal uygulamalar gibi hassas hesaplamalarda doğrudan float kullanmak yerine, örneğin **decimal** türü gibi alternatifler tercih edilir.

---

## 7. Karakter Kodlamaları ve Yazıların Temsili

- **Temel Kavramlar:**

  - Bilgisayarlar, yazıları ikilik sistemde sayılar olarak saklar. Her bir karakter, bir sayı (code point – kod noktası) ile temsil edilir.
  - Karakter tabloları, her karaktere karşılık gelen sayıları (glyph – glif) içerir.

- **ASCII ve Genişletilmiş ASCII:**

  - **ASCII (American Standard Code Information Interchange – ASCII):**
    - 7 bit’lik kodlama ile 128 karakter içerir.
  - **Code Page’ler:**
    - Farklı ülkelerin karakterlerini desteklemek için 8 bitlik genişletilmiş tablolar (örneğin, ISO 8859-9) kullanılmıştır.

- **UNICODE ve UTF-8:**
  - **UNICODE:**
    - Dünyanın tüm dillerini kapsayan standart bir karakter setidir.
    - ISO 10646 olarak da adlandırılır.
  - **UTF-8 Encoding (UTF-8 – UTF-8):**
    - En yaygın kullanılan Unicode kodlamasıdır.
    - ASCII karakterler 1 byte, diğer karakterler 2 ila 5 byte kullanır.
    - Türkçe karakterler genellikle 2 byte yer kaplar.
  - **Not:** Modern Go editörleri varsayılan olarak dosyaları UTF-8 formatında kaydeder.

---

## 8. Hexadecimal ve Octal (8'lik) Sayı Sistemleri

### 8.1. Hexadecimal Sistem (Hexadecimal – 16'lık Sistem)

- **Tanım:**

  - 16 sembol kullanılır: 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, A, B, C, D, E, F.
  - Her hex digit (hex digit – hex rakamı) 4 bit ile ifade edilir.

- **Hex to Binary Dönüşüm Tablosu:**

  | Hex Rakamı | İkilik (Binary – İkilik) |
  | ---------- | ------------------------ |
  | 0          | 0000                     |
  | 1          | 0001                     |
  | 2          | 0010                     |
  | 3          | 0011                     |
  | 4          | 0100                     |
  | 5          | 0101                     |
  | 6          | 0110                     |
  | 7          | 0111                     |
  | 8          | 1000                     |
  | 9          | 1001                     |
  | A          | 1010                     |
  | B          | 1011                     |
  | C          | 1100                     |
  | D          | 1101                     |
  | E          | 1110                     |
  | F          | 1111                     |

- **Örnek:**
  - Hex: `1FC9`  
    İkilik: `0001 1111 1100 1001`
  - Bilgisayar belleğinde 1 byte, 2 hex digit ile gösterilir. Örneğin:
    - En büyük pozitif 1 byte sayısı: `7F`
    - 4 byte içerisinde -1: `FFFFFFFF`

### 8.2. Octal Sistem (Octal – 8'lik Sistem)

- **Tanım:**

  - 8 sembol kullanılır: 0, 1, 2, 3, 4, 5, 6, 7.
  - Her octal digit (octal digit – sekizlik rakam) 3 bit ile ifade edilir.

- **Octal to Binary Dönüşüm Tablosu:**

  | Octal Rakamı | İkilik (Binary – İkilik) |
  | ------------ | ------------------------ |
  | 0            | 000                      |
  | 1            | 001                      |
  | 2            | 010                      |
  | 3            | 011                      |
  | 4            | 100                      |
  | 5            | 101                      |
  | 6            | 110                      |
  | 7            | 111                      |

- **Örnek:**
  - Octal: `476`  
    İkilik: `100 111 110`
  - İkilik sayı, sağdan üçer gruplandırılarak octal sisteme dönüştürülebilir.

---

## Ek Uygulama ve Kullanım Senaryoları

- **Go Programlarında Sayı Sistemlerinin Kullanımı:**

  - Go’da tamsayılar **int**, **int8**, **int16**, **int32** ve **int64** gibi türlerde tanımlanır. Her tür, yukarıda açıklanan sayı aralıklarını kullanır.
  - Sayı sabitleri, ondalık, hexadecimal (`0x` veya `0X` öneki) ve octal (`0` öneki) formatlarda yazılabilir.

- **Yorum Satırları ve Kod Belgelendirme:**

  - Kodun anlaşılabilirliğini artırmak için yorum satırları, fonksiyonların ve paketlerin ne amaçla kullanıldığını açıklar.
  - **godoc** aracı, yorum satırlarından otomatik olarak dokümantasyon oluşturur.

- **Derleme ve Çalıştırma Sürecinde Dikkat Edilmesi Gerekenler:**
  - **go fmt** kullanılarak kodun tutarlı boşluk ve biçimlendirme kurallarına uygun yazılması sağlanır.
  - Portability konusuna dikkat edilerek, cross-compilation ile farklı platformlarda çalışabilecek uygulamalar geliştirilebilir.

---

## Sonuç

Verilen metinde Go dilinin temel yapı taşları ve bilgisayar bilimlerinde kullanılan sayı sistemlerinin temelleri detaylandırılmıştır. Bu konuların her biri; kod yazım standartları, derleyici davranışları, veri temsili, hata ayıklama ve uygulama geliştirme süreçlerinde kritik rol oynamaktadır. Özellikle:

- **Boşluk karakterlerinin** ve **yorum satırlarının** doğru kullanımı, kodun okunabilirliğini ve bakımını kolaylaştırır.
- **Bildirimler, paket ve fonksiyon tanımları**, Go dilinin yapısal bütünlüğünü oluşturur.
- **Sayı sistemlerinin ve ikiye tümleyen yönteminin** iyi anlaşılması, özellikle düşük seviyeli programlama, bellek yönetimi ve hata kontrolü konularında önemlidir.
- **Kayan noktalı sayılar, yuvarlama hataları** ve **karakter kodlamaları**, modern uygulamalarda doğruluk ve güvenilirlik açısından kritik unsurlardır.
- **Hexadecimal ve octal sistemlerin** kullanımı, özellikle sistem programlama ve hata ayıklamada sıkça karşımıza çıkar.
