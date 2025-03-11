## Proje Adı: Karmaşık Sayılarla Aritmetik İşlemler (Complex Numbers Arithmetic)

Bu proje, Go (Golang) dilinde kullanıcıdan alınan iki **complex number** (karmaşık sayı) üzerinde temel aritmetik işlemleri gerçekleştiren bir uygulamadır. Projede, kullanıcı etkileşimi, girdi doğrulama (input validation), hata yönetimi (error handling) ve Go’nun yerleşik **fmt** (format) ile **log** (loglama) paketlerinin etkin kullanımı detaylı biçimde incelenmiştir. Bu proje, hem teorik kavramların pratiğe dökülmesi hem de programlama mantığının (programming logic) anlaşılması açısından kapsamlı bir örnek sunar.

### İçerik

- [Genel Bakış](#genel-bakış)
- [Özellikler](#özellikler)
- [Kullanım](#kullanım)
- [Kod Açıklaması](#kod-açıklaması)
  - [Paket İçe Aktarma (Import)](#paket-içe-aktarma-import)
  - [Ana Fonksiyon (Main Function)](#ana-fonksiyon-main-function)
  - [Karmaşık Sayı Girişi (Input)](#karmaşık-sayı-girişi-input)
  - [Yardımcı Fonksiyonlar](#yardımcı-fonksiyonlar)
  - [Aritmetik İşlemler (Arithmetic Operations)](#aritmetik-işlemler-arithmetic-operations)
- [Özet ve Teknik Notlar](#özet-ve-teknik-notlar)

### Genel Bakış

Bu projede, kullanıcı etkileşimi ile alınan iki karmaşık sayı üzerinde matematiksel aritmetik işlemler gerçekleştirilir. Uygulama sırasında aşağıdaki adımlar detaylandırılmıştır:

- **Girdi Alma:**

  - Kullanıcıdan alınan karmaşık sayı girişi, doğru formatta olup olmadığı **fmt.Scan** (girdi okuma) fonksiyonu kullanılarak kontrol edilir.
  - Yanlış formatta bir giriş yapıldığında, **log.Fatalf** (hata loglama) ile hata mesajı üretilerek uygulama sonlandırılır.

- **Çıktı Gösterimi:**
  - Kullanıcının girdiği karmaşık sayılar, ekrana formatlı biçimde yazdırılır.
  - Her karmaşık sayının **real** (reel) ve **imaginary** (sanal) bileşenleri ayrı ayrı ekrana basılır.
- **Aritmetik İşlemler:**
  - Toplama, çıkarma, çarpma ve bölme işlemleri gerçekleştirilir.
  - Bölme işleminde, bölenin sıfır olup olmadığı kontrol edilerek “sıfıra bölme” (division by zero) hatası önlenir.

Bu yapı, hem temel matematiksel işlem mantığını hem de Go dilindeki veri tiplerinin (data types) ve hata yönetiminin nasıl kullanılacağını örnekler.

### Özellikler

- **Kullanıcı Girişi (User Input):**
  - Kullanıcıdan **complex number** (karmaşık sayı) girişi alınır.
  - Giriş formatı, örneğin `3.4+4.5i` veya `(3.4+4.5i)` şeklinde olmalıdır.
- **Hata Yönetimi (Error Handling):**

  - Yanlış giriş durumlarında, **log.Fatalf** fonksiyonu kullanılarak hata mesajı yazdırılır ve program derhal sonlandırılır.
  - Bu, programın beklenmeyen durumlarda güvenli bir şekilde durmasını sağlar.

- **Aritmetik İşlemler (Arithmetic Operations):**

  - Toplama (addition), çıkarma (subtraction), çarpma (multiplication) ve bölme (division) işlemleri karmaşık sayılar üzerinde gerçekleştirilir.
  - Özellikle bölme işlemi öncesinde, sıfıra bölme kontrolü yapılarak potansiyel hatalı işlemler engellenir.

- **Çıktı Formatlama (Output Formatting):**
  - **fmt.Printf** kullanılarak hem karmaşık sayıların kendisi hem de bileşenleri (real, imaginary) düzenli biçimde ekrana basılır.
  - Çıktı mesajları, işlemlerin sonucunu ve işlemin hangi aşamada yapıldığını net olarak ifade eder.

### Kullanım

Projeyi yerel ortamınızda çalıştırmak için aşağıdaki adımları izleyebilirsiniz:

1. **Go dilinin (Golang) kurulu olduğundan emin olun.**
   - [Go Installation Guide (Kurulum Rehberi)](https://golang.org/doc/install) gibi kaynaklardan destek alabilirsiniz.
2. **Proje dosyasını klonlayın veya indirin.**
3. **Terminal veya komut satırında proje dizinine gidin.**
4. **Aşağıdaki komutu çalıştırın:**

   ```bash
   go run main.go
   ```

5. **Program, kullanıcıdan iki karmaşık sayı girişi bekleyecektir.**
   - Örnek giriş:
     - İlk karmaşık sayı: `3.4+4.5i`
     - İkinci karmaşık sayı: `2.1+3.2i`

### Kod Açıklaması

#### Paket İçe Aktarma (Import)

- **fmt:**

  - Formatlı giriş/çıkış işlemleri (input/output operations) için kullanılır.
  - Programın ekrana veri yazdırması ve kullanıcıdan veri alması bu paket sayesinde sağlanır.

- **log:**
  - Hata yönetimi (error handling) için kullanılır.
  - Giriş hatası gibi durumlarda, hata mesajının detaylı şekilde loglanmasını sağlar.

#### Ana Fonksiyon (Main Function)

- **Amaç:**
  - Programın giriş noktasıdır ve tüm temel işlemleri (user input, veri işleme, aritmetik hesaplamalar) koordine eder.
- **İşleyiş:**
  - **scanComplexNumber:** Fonksiyonu ile kullanıcıdan iki karmaşık sayı alınır.
  - **printComplexNumber** ve **printRealAndImagParts:** Fonksiyonları kullanılarak, alınan sayıların hem kendileri hem de ayrı ayrı bileşenleri (real ve imaginary) ekrana basılır.
  - **performArithmeticOperations:** Fonksiyonu, karmaşık sayılar üzerinde aritmetik işlemleri gerçekleştirir ve sonuçları ekrana yazdırır.

Bu yapı, programın ana işlevlerinin (core functionalities) modüler bir biçimde düzenlenmesine olanak tanır; böylece her fonksiyon belirli bir görevi üstlenir.

#### Karmaşık Sayı Girişi (Input)

- **scanComplexNumber Fonksiyonu:**
  - **İşlev:** Kullanıcıdan girdi alarak bu girdiyi `complex128` veri tipinde saklar.
  - **Detaylar:**
    - Kullanıcıya bir prompt mesajı gösterilir.
    - Kullanıcının girdiği değer, `fmt.Scan` kullanılarak okunur.
    - Giriş formatı uygun değilse, `log.Fatalf` fonksiyonu ile hata mesajı üretilir ve program sonlandırılır.
  - **Önemli Terimler:**
    - **fmt.Scan (girdi okuma):** Girdi alınması işlemini gerçekleştirir.
    - **log.Fatalf (hata fırlatma):** Hata durumunda programı durdurarak log kaydı oluşturur.

#### Yardımcı Fonksiyonlar

- **printComplexNumber Fonksiyonu:**

  - **İşlev:** Belirtilen etikete (label) uygun olarak karmaşık sayıyı ekrana yazdırır.
  - **Detaylar:**
    - `fmt.Printf` kullanılarak, karmaşık sayı formatlı biçimde yazdırılır.
    - Bu, sayının tüm bileşenlerini (hem reel hem sanal) tek satırda gösterir.

- **printRealAndImagParts Fonksiyonu:**
  - **İşlev:** Karmaşık sayının bileşenlerini ayırarak (real ve imaginary) ekrana yazdırır.
  - **Detaylar:**
    - `real()` ve `imag()` fonksiyonları kullanılarak, sayının sırasıyla reel ve sanal kısmı elde edilir.
    - Bu sayede kullanıcı, sayının hangi parçalardan oluştuğunu net bir şekilde görebilir.

#### Aritmetik İşlemler (Arithmetic Operations)

- **performArithmeticOperations Fonksiyonu:**
  - **İşlev:** İki karmaşık sayı üzerinde toplama, çıkarma, çarpma ve bölme işlemlerini gerçekleştirir.
  - **Detaylar:**
    - Her bir işlem sonucunu `fmt.Printf` kullanarak formatlı biçimde ekrana basar.
    - **Toplama (Addition):** `c1 + c2`
    - **Çıkarma (Subtraction):** `c1 - c2`
    - **Çarpma (Multiplication):** `c1 * c2`
    - **Bölme (Division):**
      - İşlem öncesinde `c2` değerinin sıfır olup olmadığı kontrol edilir.
      - Eğer `c2` sıfır ise, "Cannot divide by zero" (sıfıra bölme yapılamaz) mesajı verilir ve bölme işlemi gerçekleştirilmez.
  - **Önemli Kontrol:**
    - **Division by Zero (Sıfıra Bölme):** Programın hata almadan çalışması için kritik bir kontrol mekanizmasıdır.

### Özet ve Teknik Notlar

- **Kodun Amacı:**

  - Kullanıcıdan alınan iki karmaşık sayı ile temel aritmetik işlemlerin (addition, subtraction, multiplication, division) nasıl gerçekleştirileceğini göstermek.
  - Girdi doğrulama (input validation), hata yönetimi (error handling) ve Go dilinde modüler fonksiyon kullanımı gibi kavramları örneklemek.

- **Kullanılan Temel Terimler:**
  - **complex number (karmaşık sayı)**
  - **fmt.Scan (girdi okuma)**
  - **log.Fatalf (hata loglama)**
  - **Arithmetic Operations (aritmetik işlemler)**
  - **Division by Zero (sıfıra bölme)**
- **Örnek Kullanım Senaryosu:**
  - Kullanıcı, iki karmaşık sayı girer.
  - Program, girilen sayıların tamamını ve her bir sayının bileşenlerini (real ve imaginary) ekrana basar.
  - Toplama, çıkarma, çarpma ve bölme işlemleri yapılırken, özellikle bölme işlemi için sıfıra bölme kontrolü sağlanır.
  - Bu sayede program, hatalı giriş veya matematiksel hata durumlarını güvenli bir şekilde yönetir.
