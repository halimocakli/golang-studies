# Proje Adı: Floating-Point Bölme İşlemleri ve Özel Durum Kontrolleri

**(Floating-Point Division and Special Case Handling)**

Bu proje, Go (Golang) dilinde, float64 veri tipiyle yapılan bölme işlemleri sırasında ortaya çıkabilecek özel durumları (NaN, +Inf, -Inf) tespit etmek ve yönetmek amacıyla hazırlanmıştır. Projede, iki farklı float64 dizisi kullanılarak paylar (numerators) ve paydalar (denominators) üzerinde yapılan bölme işlemlerinin sonucu, matematiksel durum kontrolleri ile birlikte ekrana yazdırılmaktadır. Ayrıca, yorum satırlarıyla çalışma zamanı hatası (RUNTIME ERROR) üretebilecek örnek bir kod parçası da yer almaktadır.

### İçerik

- [Genel Bakış](#genel-bakış)
- [Özellikler](#özellikler)
- [Kullanım](#kullanım)
- [Kod Açıklaması](#kod-açıklaması)
  - [Paket İçe Aktarma (Import)](#paket-içe-aktarma-import)
  - [Ana Fonksiyon (Main Function)](#ana-fonksiyon-main-function)
  - [Veri Dizileri ve Döngü Kullanımı](#veri-dizileri-ve-döngü-kullanımı)
  - [Özel Durum Kontrolleri (Special Case Handling)](#özel-durum-kontrolleri-special-case-handling)
  - [Yorum Satırlarıyla Açıklama](#yorum-satırlarıyla-açıklama)
- [Özet ve Teknik Notlar](#özet-ve-teknik-notlar)

### Genel Bakış

Bu projede, Go dilinde yapılan floating-point (ondalık) bölme işlemleri sırasında aşağıdaki durumlar ele alınmaktadır:

- **NaN (Not a Number):**
  - Hesaplama sonucu tanımsız ise, örneğin 0/0 gibi durumlarda.
- **+Inf (Positive Infinity):**
  - Sonucun pozitif sonsuzluk olması durumu.
- **-Inf (Negative Infinity):**
  - Sonucun negatif sonsuzluk olması durumu.
- **Normal Bölme İşlemi:**
  - Geçerli bölme işlemlerinin sonucu, yukarıdaki özel durumlara uymadığı durumlarda.

Bu yapı, matematiksel hesaplamaların doğru kontrol edilmesi, hata yönetiminin sağlanması ve uygulama akışının güvenilir olması açısından örnek teşkil eder.

### Özellikler

- **Veri Dizileri (Data Arrays):**
  - `numerators` dizisi, bölme işlemi için pay (numerator) değerlerini içerir.
  - `denominators` dizisi, bölme işlemi için payda (denominator) değerlerini içerir.
- **Floating-Point Bölme (Floating-Point Division):**
  - `result := num / denom` ifadesiyle, her bir çift için bölme işlemi gerçekleştirilir.
- **Özel Durum Kontrolleri (Special Case Handling):**
  - `math.IsNaN(result)` ile NaN durumu kontrol edilir.
  - `math.IsInf(result, 1)` ile pozitif sonsuzluk,  
    `math.IsInf(result, -1)` ile negatif sonsuzluk tespit edilir.
- **Hata ve Uyarı Mesajları (Error and Warning Messages):**
  - Hesaplanan sonuçlara göre, özel durum mesajları ekrana yazdırılır.
- **Runtime Error Örneği:**
  - Yorum satırı içerisinde, sıfıra bölme sonucu çalışma zamanı hatası (RUNTIME ERROR) üretecek örnek kod parçası bulunmaktadır.

### Kullanım

Projeyi yerel ortamınızda çalıştırmak için şu adımları izleyebilirsiniz:

1. **Go dilinin (Golang) kurulu olduğundan emin olun.**
2. **Proje dosyasını klonlayın veya indirin.**
3. **Terminal veya komut satırında proje dizinine gidin.**
4. **Aşağıdaki komutu çalıştırın:**

   ```bash
   go run main.go
   ```

5. **Program, her bölme işleminin sonucunu, özel durum kontrolleriyle birlikte ekrana yazdıracaktır.**

### Kod Açıklaması

#### Paket İçe Aktarma (Import)

- **fmt:**
  - Formatlı giriş/çıkış işlemleri (Input/Output operations) için kullanılır.
- **math:**
  - Matematiksel işlemler ve özel durum kontrolleri (Mathematical operations and special case handling) için kullanılır.

#### Ana Fonksiyon (Main Function)

- **Amaç:**
  - Programın giriş noktasıdır. Ana fonksiyonda, tanımlı iki float64 dizisindeki pay ve payda değerleri üzerinden döngü oluşturulur ve her bir işlem sonucu özel durum kontrolleriyle birlikte ekrana yazdırılır.
- **İşleyiş:**
  - `for` döngüsü kullanılarak, her iki dizinin elemanları sırasıyla işlenir.
  - Her iterasyonda, `result := num / denom` işlemi yapılır ve elde edilen sonuç, matematiksel özel durum kontrolleriyle değerlendirilir.

#### Veri Dizileri ve Döngü Kullanımı

- **numerators (Paylar):**
  - Bölme işlemi için üst kısım değerlerini içerir.
- **denominators (Paydalar):**
  - Bölme işlemi için alt kısım değerlerini içerir.
- **Döngü (Loop):**
  - `for i := 0; i < len(numerators); i++ { ... }` yapısı ile her bir çift üzerinde işlem gerçekleştirilir.

#### Özel Durum Kontrolleri (Special Case Handling)

- **NaN Kontrolü:**
  - `math.IsNaN(result)` ifadesi, hesaplama sonucunun NaN (Not a Number) olup olmadığını kontrol eder.
- **Pozitif Sonsuzluk Kontrolü:**
  - `math.IsInf(result, 1)` ifadesi, sonucun pozitif sonsuzluk (+Inf) olup olmadığını belirler.
- **Negatif Sonsuzluk Kontrolü:**
  - `math.IsInf(result, -1)` ifadesi, sonucun negatif sonsuzluk (-Inf) olup olmadığını belirler.
- **Normal Sonuç Durumu:**
  - Yukarıdaki özel durumların dışında kalan sonuç, formatlı biçimde ekrana yazdırılır.

#### Yorum Satırlarıyla Açıklama

- **Runtime Error Örneği:**

  - Kod içerisinde yorum satırı olarak bulunan aşağıdaki parça, sıfıra bölme (division by zero) nedeniyle çalışma zamanı hatası (RUNTIME ERROR) üretecektir:

    ```go
    /*
    		c := 10 / 0
    		fmt.Println("c =", c)
    */
    ```

  - Bu örnek, hatalı matematiksel işlemlerin nasıl bir hata üreteceğini göstermektedir.

### Özet ve Teknik Notlar

- **Kodun Amacı:**

  - Float64 veri tipiyle yapılan bölme işlemlerinde ortaya çıkabilecek matematiksel özel durumları (NaN, +Inf, -Inf) tespit etmek ve yönetmek.
  - Floating-point işlemlerde, matematiksel hesaplamaların kontrolünün sağlanması, hata yönetiminin uygulanması ve doğru bilgilendirici çıktı verilmesi hedeflenmiştir.

- **Kullanılan Temel Terimler:**
  - **Floating-Point Division (ondalık bölme)**
  - **NaN (Not a Number)**
  - **Infinity (Sonsuzluk)**
  - **math.IsNaN (NaN kontrolü)**
  - **math.IsInf (sonsuzluk kontrolü)**
- **Örnek Kullanım Senaryosu:**
  - İki farklı float64 dizisi üzerinden gerçekleştirilen bölme işlemleri, çeşitli durumları (geçerli bölme, sıfıra bölme, negatif ve pozitif sonsuzluk) simüle eder.
  - Her durumda, ilgili özel durum kontrol fonksiyonları kullanılarak sonuç ekrana yazdırılır.
