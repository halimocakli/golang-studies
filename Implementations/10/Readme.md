# Go Dilinde Temel Aritmetik ve Karmaşık İşlemler Uygulaması

Bu proje, kullanıcı etkileşimi yoluyla alınan verilerle temel aritmetik (arithmetic [aritmetik]) ve karmaşık (complex [karmaşık]) sayı işlemlerini gerçekleştiren interaktif bir konsol uygulamasıdır. Uygulama; tam sayılar (integer [tam sayı]) üzerinde toplama, çıkarma, çarpma, bölme ve modulo işlemlerini; karmaşık sayılar (complex [karmaşık sayı]) üzerinde ise toplama ve çıkarma işlemlerini örneklemektedir. Ek olarak, kullanıcıya mod operatörünün (Modulo Operator [Mod Operatörü]) float veya karmaşık sayılar üzerinde kullanılamayacağını hatırlatmaktadır. Bu sayede hem temel matematiksel kavramlar hem de Go dilinde hata kontrolü (error handling [hata yönetimi]) ve kullanıcı girdisi alma yöntemleri kapsamlı olarak ele alınmaktadır.

## İçerik

- [Genel Bakış](#genel-bakış)
- [Özellikler](#özellikler)
- [Kullanım](#kullanım)
- [Kod Açıklaması](#kod-açıklaması)
  - [Paket Tanımlaması (Package Declaration)](#paket-tanımlaması-package-declaration)
  - [Paket Dahil Etme (Package Import)](#paket-dahil-etme-package-import)
  - [Ana Fonksiyon (Main Function)](#ana-fonksiyon-main-function)
  - [Tam Sayı İşlemleri (Integer Operations)](#tam-sayı-işlemleri-integer-operations)
  - [Karmaşık Sayı İşlemleri (Complex Operations)](#karmaşık-sayı-işlemleri-complex-operations)
  - [Geçersiz İşlem Hatırlatması (Invalid Operation Reminder)](#geçersiz-işlem-hatırlatması-invalid-operation-reminder)
- [Özet ve Teknik Notlar](#özet-ve-teknik-notlar)

## Genel Bakış

Bu uygulama, kullanıcı ile etkileşim kurarak sırasıyla iki tam sayı ve iki karmaşık sayı girdisi alır. Ardından bu girdiler üzerinde temel matematiksel işlemleri uygulayarak sonuçları ekrana yazdırır.  
Uygulamanın temel amacı şunlardır:

- **Kullanıcı Temelli İşlemler:**
  - Girdi doğrulama ve hata yönetimi mekanizmaları sayesinde, yanlış formatta veri girilmesi durumunda kullanıcıya bilgilendirici mesajlar sunulmaktadır.
- **Modüler Yapı:**
  - Uygulamadaki her bir işlem (tam sayılar, karmaşık sayılar ve hatırlatma mesajı) ayrı fonksiyonlar içerisinde ele alınmıştır. Bu yaklaşım, kodun okunabilirliğini ve bakımını kolaylaştırır.
- **Eğitimsel Amaç:**
  - Proje, Go dilinde temel veri tipleri, kullanıcı girdisi alma, kontrol yapıları ve fonksiyon kullanımı gibi konuları derinlemesine anlamak isteyen geliştiricilere örnek teşkil etmektedir.

## Özellikler

- **Kullanıcı Girdisi Alımı (Input Reading):**
  - İki tam sayı ve iki karmaşık sayı kullanıcıdan alınır. Bu işlem sırasında `fmt.Scan` fonksiyonu kullanılarak girdinin doğruluğu kontrol edilir.
- **Tam Sayı İşlemleri (Integer Operations):**
  - Toplama (addition [toplama]), çıkarma (subtraction [çıkarma]), çarpma (multiplication [çarpma]), bölme (division [bölme]) ve modulo (remainder [kalan]) işlemleri uygulanır.
  - **Bölme İşlemi:** Sıfıra bölme hatası önlemek için kontrol mekanizması yer almaktadır.
  - **Modulo İşlemi:** `%` operatörü kullanılarak kalan hesaplanır; bu operatör sadece tam sayılar için geçerlidir.
- **Karmaşık Sayı İşlemleri (Complex Operations):**
  - Kullanıcıdan alınan karmaşık sayılar üzerinde toplama (addition [toplama]) ve çıkarma (subtraction [çıkarma]) işlemleri gerçekleştirilir.
  - Karmaşık sayı işlemleri, Go dilinin yerleşik `complex128` veri tipini kullanır; bu, hem reel (real [gerçek]) hem de sanal (imaginary [imajiner]) kısımları içeren sayılar için tasarlanmıştır.
- **Hata Yönetimi (Error Handling):**
  - Girdi alma işlemlerinde hata kontrolü yapılır. Yanlış formatta veri girilmesi durumunda kullanıcı bilgilendirilir ve işlem kesintiye uğrar.
- **Ek Hatırlatma (Reminder):**
  - Kullanıcıya, mod operatörünün (Modulo Operator [Mod Operatörü]) float veya karmaşık sayılar üzerinde kullanılamayacağını hatırlatan ek bir mesaj sunulmaktadır.

## Kullanım

1. **Go Kurulumu (Installation):**

   - Sisteminizde [Go](https://golang.org/) dilinin yüklü olduğundan emin olun. Bu, uygulamanın derlenip çalıştırılması için gereklidir.

2. **Projeyi Çalıştırma (Running the Project):**

   - Proje dosyasını yerel ortamınıza klonlayın veya indirin.
   - Terminal veya komut satırında proje dizinine gidin.
   - Aşağıdaki komutu kullanarak uygulamayı çalıştırın:

     ```bash
     go run main.go
     ```

   - Program, sırasıyla kullanıcıdan iki tam sayı ve iki karmaşık sayı girişi isteyecektir.
   - Her adımda, işlemin doğru yapılması için gerekli kontroller (örneğin sıfıra bölme kontrolü) sağlanmaktadır.

3. **Kullanıcı Girdisi (User Input):**
   - Program ilk olarak iki tam sayı girişi bekler; örneğin: `3 5` şeklinde.
   - Ardından, iki karmaşık sayı girişi beklenir; örneğin: `3+4i 5-2i`.
   - Yanlış formatta giriş yapıldığında, hata mesajlarıyla kullanıcı bilgilendirilir ve ilgili fonksiyon çalışmayı sonlandırır.

## Kod Açıklaması

### Paket Tanımlaması (Package Declaration)

- **Açıklama:**
  - `package main` ifadesi, uygulamanın yürütülebilir bir program olduğunu belirtir.
  - Bu, programın giriş noktası (entry point [giriş noktası]) olarak tanımlanan `main()` fonksiyonunun bulunduğu pakettir.
  - **Detay:** Go dilinde her uygulama en az bir `main` paketine sahip olmalıdır; bu, derleyicinin hangi paketten başlayacağını belirtir.

### Paket Dahil Etme (Package Import)

- **Açıklama:**
  - `import "fmt"` ifadesi, formatlı giriş/çıkış işlemleri için gerekli olan **fmt** paketini dahil eder.
  - **Detay:** `fmt` paketi, hem kullanıcıdan veri almak (input [girdi]) hem de sonuçları ekrana yazdırmak (output [çıktı]) için vazgeçilmezdir.
  - Bu paket sayesinde, formatlı yazdırma işlemleri ve tarama (scanning [okuma]) fonksiyonları kullanılmaktadır.

### Ana Fonksiyon (Main Function)

- **Açıklama:**
  - `func main()` fonksiyonu, programın başlangıç noktasıdır.
  - Programın yürütülmesinde sırasıyla tam sayı işlemleri, karmaşık sayı işlemleri ve hatırlatma mesajının gösterilmesi bu fonksiyon içerisinde organize edilmiştir.
- **Detay:**
  - Her fonksiyon çağrısı, ilgili işlemleri modüler şekilde ayırarak kodun okunabilirliğini artırır ve hata ayıklama (debugging [hata ayıklama]) sürecini kolaylaştırır.

### Tam Sayı İşlemleri (Integer Operations)

- **Fonksiyon:** `performIntegerOperations()`
- **İşlev:**
  - Kullanıcıdan iki tam sayı girişi alır ve bu sayılar üzerinde toplama, çıkarma, çarpma, bölme ve modulo işlemlerini gerçekleştirir.
- **Detaylı Açıklama:**
  - **Girdi Alma (Input):**
    - `fmt.Scan` fonksiyonu ile kullanıcıdan alınan veriler `num1` ve `num2` değişkenlerine aktarılır.
    - Hata kontrolü yapılarak, eğer kullanıcı yanlış bir formatta veri girerse, hata mesajı ekrana yazdırılır.
  - **Aritmetik İşlemler (Arithmetic Operations):**
    - **Toplama (Addition [Toplama]):** İki sayının toplamı hesaplanır.
    - **Çıkarma (Subtraction [Çıkarma]):** İki sayı arasındaki fark bulunur.
    - **Çarpma (Multiplication [Çarpma]):** İki sayının çarpımı hesaplanır.
    - **Bölme (Division [Bölme]):**
      - Bölme işlemi yapılmadan önce `num2` değişkeninin sıfır olup olmadığı kontrol edilir.
      - Sıfıra bölme durumu önlenir; aksi halde program hata vermez ve kullanıcı bilgilendirilir.
    - **Modulo (Remainder [Kalan]):**
      - `%` operatörü kullanılarak, iki tam sayı arasındaki kalan (modulus) hesaplanır.
      - Not: Bu operatör sadece tam sayılar üzerinde geçerlidir, bu yüzden float veya karmaşık sayılar için kullanılamaz.

### Karmaşık Sayı İşlemleri (Complex Operations)

- **Fonksiyon:** `performComplexOperations()`
- **İşlev:**
  - Kullanıcıdan iki karmaşık sayı girişi alır ve bu sayılar üzerinde toplama ile çıkarma işlemleri gerçekleştirir.
- **Detaylı Açıklama:**
  - **Girdi Alma (Input):**
    - Kullanıcıdan alınan karmaşık sayılar `complex128` veri tipi ile tanımlanır.
    - Bu tip, hem reel hem de imajiner kısımları barındırarak, karmaşık sayı işlemlerinde yüksek hassasiyet sağlar.
  - **Aritmetik İşlemler (Arithmetic Operations):**
    - **Toplama (Addition [Toplama]):**
      - İki karmaşık sayının her bir bileşeni (real ve imaginary) ayrı ayrı toplanır.
    - **Çıkarma (Subtraction [Çıkarma]):**
      - Benzer şekilde, iki karmaşık sayının birbirinden çıkarılması işlemi gerçekleştirilir.
  - **Önemli Not:**
    - Karmaşık sayılar üzerinde çarpma veya bölme gibi daha ileri matematiksel işlemler uygulanmamıştır; bu durum, uygulamanın odaklandığı temel işlemleri vurgulamak içindir.

### Geçersiz İşlem Hatırlatması (Invalid Operation Reminder)

- **Fonksiyon:** `showInvalidOperations()`
- **İşlev:**
  - Kullanıcıya, mod operatörünün (Modulo Operator [Mod Operatörü]) float veya karmaşık sayılar üzerinde kullanılamayacağını hatırlatan bir mesaj gösterir.
- **Detaylı Açıklama:**
  - Bu bölüm, geliştiricilerin ve kullanıcıların karşılaşabileceği potansiyel hataları önceden bildirmekte ve kodun güvenli çalışmasını sağlamada ek bir bilgi sunmaktadır.
  - Bu tür uyarılar, özellikle yeni başlayanlar için hangi işlemlerin dil tarafından desteklenmediğini anlamada önemli rol oynar.

## Özet ve Teknik Notlar

- **Kodun Amacı:**
  - Kullanıcıdan alınan girdiler üzerinden temel tam sayı ve karmaşık sayı işlemlerinin nasıl gerçekleştirileceğini ve bu işlemlerin hangi durumlarda hata yönetimi gerektirdiğini örneklemek.
  - Uygulama, Go dilinin temel özelliklerini, veri tiplerini ve fonksiyonel programlama yaklaşımlarını göstermektedir.
- **Kullanılan Temel Terimler:**
  - **Integer (Tam Sayı):**
    - `int` veri tipi, tam sayı aritmetiği için kullanılır.
    - Temel matematiksel işlemler, hata kontrol mekanizmaları ve operatör kullanımları bu veri tipi üzerinde gerçekleştirilir.
  - **Complex (Karmaşık):**
    - `complex128` veri tipi, karmaşık sayı işlemlerinde hem reel hem de imajiner kısımlarıyla yüksek doğruluk sağlar.
    - Karmaşık sayılar üzerinde yapılan işlemler, matematiksel teorinin pratik uygulamalarını temsil eder.
  - **Input/Output (Girdi/Çıktı):**
    - `fmt.Scan` ile kullanıcı girdisi alınır.
    - `fmt.Printf` ile formatlı çıktı oluşturularak sonuçlar ekrana yazdırılır.
  - **Error Handling (Hata Yönetimi):**
    - Kullanıcıdan alınan verinin doğruluğu kontrol edilerek, yanlış formatta veri girildiğinde uygun hata mesajları gösterilir.
- **Teknik Notlar:**
  - Kod, Go dilinin kullanıcı etkileşimli uygulamalarda nasıl yapılandırılacağını göstermektedir.
  - Her fonksiyon, belirli bir işlemi modüler olarak yerine getirir; bu, kodun genişletilebilirliğini ve bakımını kolaylaştırır.
  - Uygulama, eğitimsel amaçlarla temel kavramların yanında hata yönetimi ve operatör sınırlamaları gibi önemli detayları da içerir.
