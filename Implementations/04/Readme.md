# Go Dilinde Sayı Yuvarlama Uygulaması

Bu repository, kullanıcıdan alınan ondalıklı (float) bir değerin, matematiksel yuvarlama işlemleri kullanılarak farklı sonuçlarının hesaplanması ve ekrana yazdırılmasını sağlayan basit bir Go uygulamasıdır. Uygulama, Go dilinin standart kütüphaneleri olan `fmt` (formatting - biçimlendirme), `log` (logging - loglama) ve `math` (mathematical functions - matematiksel fonksiyonlar) paketlerini kullanmaktadır. Bu proje, özellikle gerçek dünya senaryolarında kullanıcı girdisi işleme, hata yönetimi ve matematiksel hesaplamaların nasıl gerçekleştirildiğini detaylı bir şekilde anlamaya yönelik örnek bir uygulamadır.

## İçerik

- [Genel Bakış](#genel-bakış)
- [Özellikler](#özellikler)
- [Kullanım](#kullanım)
- [Kod Açıklaması](#kod-açıklaması)
  - [Paket Tanımlaması (Package Declaration)](#paket-tanımlaması-package-declaration)
  - [Paket Dahil Etme (Package Import)](#paket-dahil-etme-package-import)
  - [Ana Fonksiyon (Main Function)](#ana-fonksiyon-main-function)
  - [Kullanıcı Girdisi ve Hata Yönetimi (User Input and Error Handling)](#kullanıcı-girdisi-ve-hata-yönetimi-user-input-and-error-handling)
  - [Matematiksel İşlemler (Mathematical Operations)](#matematiksel-işlemler-mathematical-operations)
- [Özet ve Teknik Notlar](#özet-ve-teknik-notlar)

## Genel Bakış

Bu proje, kullanıcının terminalden girdiği ondalıklı bir sayıyı alarak; bu sayının orijinal değeri, en yakın tam sayıya yuvarlanmış hali, aşağıya yuvarlanmış hali ve yukarıya yuvarlanmış hali hesaplanıp, ekrana yazdırılmasını sağlamaktadır.

- **Kullanıcı Etkileşimi (User Interaction - Kullanıcı Etkileşimi):** Kullanıcının girdiği verinin doğru biçimde alınması ve işlenmesi vurgulanmıştır.
- **Matematiksel İşlemler (Mathematical Operations - Matematiksel İşlemler):** `math.Round`, `math.Floor` ve `math.Ceil` fonksiyonları kullanılarak, farklı yuvarlama yöntemleri detaylı şekilde ele alınmıştır.
- **Hata Yönetimi (Error Handling - Hata Yönetimi):** Kullanıcı girdisinde oluşabilecek hataların nasıl kontrol edildiği ve yönetildiği örneklenmiştir.

## Özellikler

- **Kullanıcı Girdisi (User Input):** Terminal üzerinden ondalıklı sayı girişi sağlanır.
- **Orijinal Değerin Yazdırılması:** Kullanıcının girdiği sayının orijinal hali, formatlı olarak gösterilir.
- **Yuvarlanmış Değer (Rounded):** `math.Round` fonksiyonu ile sayının en yakın tam sayıya yuvarlanmış hali hesaplanır.
- **Aşağı Yuvarlanmış Değer (Floored):** `math.Floor` fonksiyonu, sayıyı aşağıya doğru en yakın tam sayıya indirger.
- **Yukarı Yuvarlanmış Değer (Ceiled):** `math.Ceil` fonksiyonu, sayıyı yukarıya doğru en yakın tam sayıya çıkartır.
- **Hata Yönetimi (Error Handling):** Girdi alınırken oluşabilecek hatalar, `log.Fatalf` kullanılarak anında raporlanır ve uygulama güvenli bir şekilde sonlandırılır.
- **Formatlama ve Çıktı Yönetimi (Formatting and Output Management):** `fmt.Printf` fonksiyonu ile sonuçlar belirli formatta (örneğin, 2 ondalık basamak) sunulur.

## Kullanım

Projeyi yerel ortamınızda çalıştırmak için aşağıdaki adımları izleyin:

1. Sisteminizde [Go](https://golang.org/) dilinin yüklü olduğundan emin olun.
2. Repository’i klonlayın veya dosyaları indirin.
3. Terminal veya komut satırında proje dizinine gidin.
4. Aşağıdaki komutu çalıştırarak uygulamayı başlatın:

   ```bash
   go run main.go
   ```

5. Terminalde, "Input a float value (e.g., 3.14):" mesajı görünecektir. Belirtilen formatta bir sayı girin. Uygulama, girilen değerin orijinal hali ile birlikte yuvarlanmış, aşağıya yuvarlanmış ve yukarıya yuvarlanmış sonuçları hesaplayarak ekrana yazdıracaktır.

## Kod Açıklaması

### Paket Tanımlaması (Package Declaration)

- **Kod:** `package main`
- **Açıklama:**
  - Bu satır, dosyanın ana yürütülebilir (executable - çalıştırılabilir) program olduğunu belirtir.
  - Go'da her çalıştırılabilir program `main` paketinde yer almalıdır; bu yapı, derleyiciye uygulamanın başlangıç noktasını belirtir.

### Paket Dahil Etme (Package Import)

- **Kod:**
  ```go
  import (
      "fmt"
      "log"
      "math"
  )
  ```
- **Açıklama:**
  - **`fmt` (formatting - biçimlendirme):**
    - Terminalden giriş almak ve çıktıları formatlamak için kullanılır.
    - Örneğin, `fmt.Print` ve `fmt.Printf` fonksiyonları ile verilerin okunması ve yazdırılması sağlanır.
  - **`log` (logging - loglama):**
    - Uygulamada oluşabilecek hataların kaydedilmesi ve raporlanması için kullanılır.
    - `log.Fatalf` gibi fonksiyonlarla hata meydana geldiğinde, hata mesajı basılarak programın çalışması sonlandırılır.
  - **`math` (mathematical functions - matematiksel fonksiyonlar):**
    - Sayılar üzerinde matematiksel işlemler gerçekleştirmek için geniş bir fonksiyon yelpazesi sunar.
    - Bu örnekte, `math.Round`, `math.Floor` ve `math.Ceil` fonksiyonları kullanılarak farklı yuvarlama işlemleri yapılmaktadır.

### Ana Fonksiyon (Main Function)

- **Kod:** `func main() { ... }`
- **Açıklama:**
  - Go programında çalıştırılacak tüm kodların başladığı yerdir.
  - `main` fonksiyonu, programın giriş noktası olup, sırasıyla kullanıcı girdisini alma, hata kontrolü yapma ve matematiksel işlemleri uygulama görevlerini üstlenir.
  - Bu fonksiyon içerisinde, kodun akışı belirgin ve düzenlidir; her işlem adım adım açıklanmıştır.

### Kullanıcı Girdisi ve Hata Yönetimi (User Input and Error Handling)

- **Kullanıcı Girdisi Alma:**
  - `fmt.Print("Input a float value (e.g., 3.14): ")` ifadesi ile terminale kullanıcıdan ondalıklı bir sayı girmesi istenir.
  - `fmt.Scan(&number)` fonksiyonu, kullanıcı tarafından girilen değeri okur ve `number` değişkenine atar.
  - Bu işlem, veri tiplerinin uyumuna dikkat edilerek yapılır; burada `number` değişkeni `float64` olarak tanımlanmıştır.
- **Hata Yönetimi:**
  - Girdi alma işleminde hata meydana gelirse, bu hata `err` değişkenine atanır.
  - `if err != nil { ... }` bloğu içerisinde, hata kontrolü yapılır.
    - Eğer hata mevcutsa, `log.Fatalf` fonksiyonu ile hata mesajı ekrana yazdırılır ve program güvenli bir şekilde sonlandırılır.
  - Bu yöntem, özellikle gerçek dünya uygulamalarında kullanıcıdan gelen hatalı verinin yönetilmesi ve uygulamanın beklenmedik durumlarda çökmesinin önlenmesi açısından kritiktir.

### Matematiksel İşlemler (Mathematical Operations)

- **Orijinal Değerin Yazdırılması:**
  - `fmt.Printf("\nOriginal: %.2f\n", number)` ifadesi kullanılarak, kullanıcının girdiği orijinal sayı, 2 ondalık basamakla formatlanarak ekrana yazdırılır.
  - Bu, kullanıcıya girdisinin doğru alındığını ve işleme tabi tutulduğunu gösterir.
- **Yuvarlanmış Değer (Rounded):**
  - `math.Round(number)` fonksiyonu, girilen sayıyı en yakın tam sayıya yuvarlar.
    - Bu işlem, sayının kesirli kısmının 0.5 veya daha büyük olması durumunda yukarı yuvarlanmasını, aksi halde aşağı yuvarlanmasını sağlar.
  - Çıktı, yuvarlanmış değerin net bir biçimde gösterilmesi için `fmt.Printf` ile formatlanır.
- **Aşağı Yuvarlanmış Değer (Floored):**
  - `math.Floor(number)` fonksiyonu, sayıyı aşağıya doğru en yakın tam sayıya indirger.
  - Bu fonksiyon, özellikle negatif sayılarda veya kesin alt sınır gerektiren durumlarda kullanışlıdır.
- **Yukarı Yuvarlanmış Değer (Ceiled):**
  - `math.Ceil(number)` fonksiyonu, sayıyı yukarıya doğru en yakın tam sayıya çıkartır.
  - Bu işlem, sayının kesirli kısmını tamamen göz ardı ederek, daima bir üst tam sayıya ulaşmayı garanti eder.
- **Çıktıların Yazdırılması:**
  - Her bir matematiksel işlemin sonucu, `fmt.Printf` fonksiyonu kullanılarak belirli formatta ekrana yazdırılır.
  - Formatlama sırasında kullanılan `%.2f` ifadesi, sonucu 2 ondalık basamakla gösterir; bu, kullanıcıya daha okunabilir ve düzenli bir çıktı sunar.

## Özet ve Teknik Notlar

- **Kodun Amacı:**
  - Kullanıcının terminalden girdiği ondalıklı bir sayının, farklı matematiksel yuvarlama işlemleri ile nasıl işlenebileceğini örneklemek.
  - Bu örnek, gerçek dünya uygulamalarında kullanıcı girdisi işleme, hata yönetimi ve matematiksel hesaplamaların uygulanması konularında temel bir anlayış kazandırmayı amaçlamaktadır.
- **Temel Fonksiyonlar:**
  - **`fmt.Scan` (Input Reading - Girdi Okuma):** Kullanıcıdan veri alma işlemini gerçekleştirir.
  - **`log.Fatalf` (Error Logging - Hata Loglama):** Hata durumunda hata mesajı gösterir ve programı sonlandırır.
  - **`math.Round` (Rounding - Yuvarlama):** Sayıyı en yakın tam sayıya yuvarlar; kesirli kısmın büyüklüğüne göre yukarı veya aşağı yönlü yuvarlama yapar.
  - **`math.Floor` (Flooring - Aşağı Yuvarlama):** Sayıyı aşağıya doğru yuvarlayarak, sayının kendisinden küçük veya eşit en yakın tam sayıyı verir.
  - **`math.Ceil` (Ceiling - Yukarı Yuvarlama):** Sayıyı yukarıya doğru yuvarlayarak, sayının kendisinden büyük veya eşit en yakın tam sayıyı verir.
- **Teknik Terimler:**
  - **Executable (Çalıştırılabilir):** Derlendiğinde doğrudan çalıştırılabilen program.
  - **Standard Library (Standart Kütüphane):** Go dilinde, temel işlemleri gerçekleştirmek için sağlanan yerleşik kütüphaneler.
  - **Error Handling (Hata Yönetimi):** Programda meydana gelen hataların belirlenmesi, raporlanması ve güvenli şekilde sonlandırılmasını sağlayan yöntem.
  - **Floating Point Arithmetic (Ondalık Hesaplamalar):** Ondalık sayıların işlenmesinde kullanılan matematiksel yöntemler.
- **Ek Not:**
  - Bu örnek, özellikle kullanıcı girdisinin doğru alınması, matematiksel fonksiyonların kullanımı ve hata kontrol mekanizmalarının uygulanması açısından temel bir rehber niteliğindedir.
  - Her adımda, kullanılan fonksiyonların işlevleri detaylandırılarak, hem teorik hem de pratik kullanım örnekleri üzerinden açıklama getirilmiştir.

```

```
