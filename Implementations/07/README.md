# Go Dilinde Formatlı Çıktı Örneği

Bu proje, Go dilinde **fmt.Printf** fonksiyonunun (formatlı yazdırma - formatted output) farklı format belirteçlerini (format specifiers - format belirteçleri) kullanarak çeşitli veri tiplerinin (data types - veri tipleri) çıktısını nasıl elde edebileceğimizi göstermektedir. Kod, temel Go programlama yapısı ile hazırlanmış olup; **integer (tamsayılar - integers)**, **floating-point (gerçek sayılar - floating-point numbers)**, **string (dizgeler - strings)**, **pointer (işaretçiler - pointers)** ve **general placeholder (genel yer tutucular - general placeholders)** kullanım örneklerini detaylandırmaktadır. Bu örnek, hem çıktı biçimlendirme hem de hata ayıklama (debugging - hata ayıklama) süreçlerinde kullanılabilecek teknikleri içermektedir.

## İçerik

- [Genel Bakış](#genel-bakış)
- [Özellikler](#özellikler)
- [Kullanım](#kullanım)
- [Kod Açıklaması](#kod-açıklaması)
  - [Tamsayılar (Integers)](#tamsayılar-integers)
  - [Gerçek Sayılar (Floating-Point Numbers)](#gerçek-sayılar-floating-point-numbers)
  - [Dizgeler (Strings)](#dizgeler-strings)
  - [Pointer (İşaretçi)](#pointer-işaretçi)
  - [Genel Yer Tutucular (General Placeholders)](#genel-yer-tutucular-general-placeholders)
- [Özet ve Teknik Notlar](#özet-ve-teknik-notlar)

## Genel Bakış

Bu örnek, Go programlama dilinde **fmt.Printf** fonksiyonunun sunduğu farklı formatlama seçeneklerini kullanarak verilerin nasıl düzenli bir biçimde ekrana yazdırılacağını göstermektedir. Kodda yer alan format belirteçleri, özellikle büyük projelerde çıktının okunabilirliğini artırmak, veri doğrulaması (data validation - veri doğrulaması) yapmak ve hata ayıklama (debugging - hata ayıklama) süreçlerini kolaylaştırmak amacıyla kullanılır. Proje, Go'nun yerleşik özelliklerini öğrenmek isteyen geliştiriciler (developers - geliştiriciler) için kapsamlı bir kaynak sunmaktadır.

## Özellikler

- **Tamsayılar (Integers):**

  - **`%d`:** Tamsayıyı ondalık (decimal) sistemde gösterir. Bu format, günlük hesaplamalarda kullanılan temel gösterimdir.
  - **`%b`:** Tamsayıyı ikilik (binary) sistemde gösterir. Binary gösterim, düşük seviye (low-level - düşük seviye) programlama ve bit manipülasyonu işlemlerinde önemlidir.
  - **`%016b`:** 16 karakter genişliğinde, eksik yerleri sıfır (0) ile doldurulmuş ikilik gösterim sağlar. Bu, özellikle sabit genişlikte verilerin görüntülenmesi gereken durumlarda kullanılır.
  - **`%o`:** Tamsayıyı sekizlik (octal) sistemde gösterir. Bazı durumlarda, dosya izinlerinin (file permissions - dosya izinleri) gösterimi gibi alanlarda octal format tercih edilir.
  - **`%x` / `%X`:** Tamsayıyı onaltılık (hexadecimal) sistemde, sırasıyla küçük ve büyük harflerle gösterir. Hexadecimal gösterim, bellek adresleri (memory addresses - bellek adresleri) ve renk kodlamaları gibi durumlarda yaygın olarak kullanılır.
  - **`%c`:** Tamsayıyı karşılık geldiği Unicode karakter (Unicode character - Unicode karakter) olarak gösterir.
  - **`%q`:** Tamsayıyı tek tırnak içine alınmış karakter çıktısı şeklinde sunar; bu, çıktının daha okunabilir hale gelmesi için kullanışlıdır.

- **Gerçek Sayılar (Floating-Point Numbers):**

  - **`%f`:** Sabit nokta (fixed-point - sabit nokta) gösterimi sağlar. Özellikle para birimi gibi hassas hesaplamalarda tercih edilebilir.
  - **`%e`:** Bilimsel gösterim (scientific notation - bilimsel gösterim) sunar. Büyük veya çok küçük sayılar için okunabilirliği artırır.
  - **`%g`:** En uygun biçimde (compact form - kompakt biçim) gösterim yapar; değer büyüklüğüne göre %f veya %e formatını otomatik olarak seçer.

- **Dizgeler (Strings):**

  - **`%s`:** Ham (raw - ham) dizge çıktısı verir. Düz metin (plain text - düz metin) gösterimleri için kullanılır.
  - **`%q`:** Dizgeyi çift tırnak içine alarak çıktı verir. Bu, metin verilerinin sınırlarını belirgin hale getirir.
  - **`%x`:** Dizgedeki her bir karakterin hexadecimal (onaltılık - hexadecimal) çıktısını verir. Özellikle karakterlerin kod değerlerini (ASCII codes - ASCII kodları) incelemek için yararlıdır.

- **Pointer (İşaretçi):**

  - **`%p`:** Değişkenin bellek adresini hexadecimal biçimde gösterir. Bu, bellek yönetimi (memory management - bellek yönetimi) ve hata ayıklama (debugging - hata ayıklama) işlemlerinde kritik öneme sahiptir.

- **Genel Yer Tutucular (General Placeholders):**
  - **`%v`:** Değerin varsayılan gösterimini sağlar. Herhangi bir veri tipinde genel çıktı almak için kullanılır.
  - **`%+v`:** Struct (yapı - struct) verilerini alan adlarıyla birlikte gösterir; bu, veri yapılarının (data structures - veri yapıları) daha ayrıntılı incelenmesini sağlar.
  - **`%T`:** Değişkenin veri tipini (type - tip) belirtir. Programın çalışma zamanında (runtime - çalışma zamanı) veri tiplerinin kontrol edilmesi için yararlıdır.

## Kullanım

1. **Go'nun Yüklü Olduğundan Emin Olun:**  
   Sisteminizde Go dilinin kurulu olduğunu doğrulayın. (Go installation - Go kurulumu)

2. **Projeyi İndirin veya Klonlayın:**  
   Bu kodu içeren dosyayı yerel bilgisayarınıza indirin veya Git deposundan klonlayın.

3. **Terminal veya Komut Satırı:**  
   Proje dizinine geçiş yapın.

4. **Uygulamayı Çalıştırın:**

   ```bash
   go run main.go
   ```

5. **Çıktıyı İnceleyin:**  
   Terminalde, her veri tipi için formatlı çıktılar gözlemleyebilir, format belirteçlerinin nasıl çalıştığını detaylı olarak inceleyebilirsiniz.

## Kod Açıklaması

### Tamsayılar (Integers)

- **`%d`:** Tamsayıyı ondalık (decimal) sistemde gösterir.
  - **Ayrıntı:** Bu format, sayısal verinin standart gösterimidir ve günlük hesaplamalarda en yaygın kullanılan biçimdir.
- **`%b`:** Tamsayıyı ikilik (binary) sistemde gösterir.
  - **Ayrıntı:** Özellikle düşük seviye programlamada (low-level programming) bit düzeyinde işlemler için gereklidir.
- **`%016b`:** 16 karakter genişliğinde, eksik yerleri sıfır (0) ile doldurulmuş ikilik gösterim sağlar.
  - **Ayrıntı:** Sabit genişlikte veri çıktısı alınması gereken durumlarda (örneğin, sabit uzunluklu binary string'ler) kullanılır.
- **`%o`:** Tamsayıyı sekizlik (octal) sistemde gösterir.
  - **Ayrıntı:** Bazı sistem ve uygulamalarda (örneğin, dosya izinleri) octal format kullanımı yaygındır.
- **`%x` / `%X`:** Tamsayıyı onaltılık (hexadecimal) sistemde, sırasıyla küçük ve büyük harflerle gösterir.
  - **Ayrıntı:** Bellek adreslerinin veya renk kodlarının gösterimi gibi alanlarda hexadecimal format sıkça kullanılır.
- **`%c`:** Tamsayıyı karşılık geldiği Unicode karakter olarak gösterir.
  - **Ayrıntı:** Bu format, tamsayının temsil ettiği karakterin (örneğin, 65 -> 'A') ekranda gösterilmesini sağlar.
- **`%q`:** Tamsayıyı tırnak içine alınmış karakter çıktısı şeklinde sunar.
  - **Ayrıntı:** Çıktının sınırlarını belirgin hale getirmek ve karakterin doğru şekilde ayrıştırılmasını sağlamak amacıyla kullanılır.

### Gerçek Sayılar (Floating-Point Numbers)

- **`%f`:** Sabit nokta (fixed-point) gösterimi sağlar.
  - **Ayrıntı:** Bu format, ondalıklı sayıların insan tarafından okunabilir biçimde gösterilmesini sağlar.
- **`%e`:** Bilimsel (exponential) gösterim sunar.
  - **Ayrıntı:** Çok büyük veya çok küçük sayılar için kompakt bir gösterim sunar; örneğin, 1234.5678 sayısını `1.234568e+03` şeklinde ifade eder.
- **`%g`:** En uygun biçimde gösterim yapar.
  - **Ayrıntı:** Değerin büyüklüğüne göre otomatik olarak %f veya %e formatını seçerek, çıktıyı daha okunabilir kılar.

### Dizgeler (Strings)

- **`%s`:** Ham dizge çıktısı verir.
  - **Ayrıntı:** Dizgenin orijinal (raw - ham) halini olduğu gibi yazdırır.
- **`%q`:** Dizgeyi çift tırnak içine alarak çıktı verir.
  - **Ayrıntı:** Bu yöntem, dizgenin başlangıç ve bitiş sınırlarını belirginleştirir; özellikle JSON benzeri veri gösterimlerinde yararlıdır.
- **`%x`:** Dizgedeki her bir karakterin hexadecimal gösterimini verir.
  - **Ayrıntı:** Her karakterin ASCII veya Unicode kod değerini incelemek isteyen geliştiriciler için faydalıdır.

### Pointer (İşaretçi)

- **`%p`:** Değişkenin bellek adresini hexadecimal biçimde gösterir.
  - **Ayrıntı:** Bu format, bellek yönetimi (memory management) ve hata ayıklama (debugging) süreçlerinde, değişkenlerin fiziksel bellek konumlarını görmek için kullanılır.

### Genel Yer Tutucular (General Placeholders)

- **`%v`:** Değerin varsayılan gösterimini sağlar.
  - **Ayrıntı:** Herhangi bir veri tipinde genel çıktı almak için esnek bir formattır.
- **`%+v`:** Struct verilerini alan adlarıyla birlikte gösterir.
  - **Ayrıntı:** Veri yapılarının içeriğini ve ilgili alan adlarını daha açıklayıcı bir şekilde görüntülemek için kullanılır.
- **`%T`:** Değişkenin veri tipini belirtir.
  - **Ayrıntı:** Programın çalışma zamanında (runtime) veri tiplerinin kontrolü ve doğrulanması açısından kritik öneme sahiptir.

## Özet ve Teknik Notlar

- **Kodun Amacı:**  
  Bu örnek, Go dilinde formatlı çıktı elde etmek için kullanılan **fmt.Printf** fonksiyonunun farklı formatlama seçeneklerini (format specifiers) kapsamlı bir şekilde örneklemektedir. Çıktıların okunabilirliğini artırmak, veri doğrulaması (data validation) ve hata ayıklama (debugging) işlemlerinde kullanılan yöntemler detaylandırılmıştır.

- **Teknik Terimler ve Açıklamaları:**
  - **Integer (Tamsayı):** Tam sayı değerlerini ifade eder; genellikle sayısal hesaplamalarda kullanılır.
  - **Floating-Point (Gerçek Sayı):** Ondalıklı sayı değerlerini temsil eder; bilimsel ve mali hesaplamalarda önemlidir.
  - **String (Dizge):** Metin verilerini saklamak ve göstermek için kullanılır.
  - **Pointer (İşaretçi):** Bellek adresi bilgisini tutar; bellek yönetimi ve düşük seviye programlama için kritik bir kavramdır.
  - **Struct (Yapı):** Birden fazla veri tipini bir arada tutan, karmaşık veri yapıları oluşturmak için kullanılan yapı taşlarıdır.
  - **Format Specifiers (Format Belirteçleri):** Çıktı biçimlendirmesinde kullanılan, veri tiplerine uygun formatları belirleyen sembollerdir.

Bu örnek, Go dilinde verilerin nasıl biçimlendirileceğini ve ekrana nasıl yazdırılacağını öğrenmek isteyen geliştiriciler için kapsamlı bir rehberdir. Proje, hem basit hem de ileri seviye veri gösterim tekniklerini içermekte ve uygulamanın farklı senaryolarda nasıl davranacağını detaylandırmaktadır.
