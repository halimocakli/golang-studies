# Go Dilinde Literaller (Literals) Örneği

Bu proje, Go dilinde farklı literal türlerinin nasıl tanımlanıp kullanılacağını ve kod içerisinde bu sabit ifadelerin (literal expressions [sabit ifadeler]) nasıl yorumlandığını kapsamlı bir şekilde göstermektedir. Projede; integer literal (tamsayı sabiti), floating point literal (ondalık sabiti), complex literal (karmaşık sayı sabiti), boolean literal (mantıksal sabit), rune literal (karakter sabiti) ve string literal (metin sabiti) gibi temel veri tipleri ile birlikte, kaçış dizilerinin (escape sequences [kaçış dizileri]) doğru kullanımına da yer verilmiştir.

## İçerik

- [Genel Bakış](#genel-bakış)
- [Özellikler](#özellikler)
- [Kullanım](#kullanım)
- [Kod Açıklaması](#kod-açıklaması)
  - [Paket Tanımlaması (Package Declaration)](#paket-tanımlaması-package-declaration)
  - [Paket Dahil Etme (Import)](#paket-dahil-etme-import)
  - [Ana Fonksiyon (Main Function)](#ana-fonksiyon-main-function)
  - [Literal Türleri ve Kaçış Dizileri (Literals and Escape Sequences)](#literal-türleri-ve-kaçış-dizileri-literals-and-escape-sequences)
- [Özet ve Teknik Notlar](#özet-ve-teknik-notlar)

## Genel Bakış

Bu proje, Go dilinde literal ifadelerin (sabit değerlerin) doğrudan kod içerisinde tanımlanmasının ve bu değerlerin nasıl formatlanarak çıktıya dönüştürüldüğünün anlaşılmasını sağlamaktadır. Literaller, programlama dillerinde sabit verilerin kod içine gömülmesiyle ifade edilir; bu proje, özellikle sayı sistemlerinin (binary, octal, hexadecimal) ve bilimsel gösterimin (scientific notation [bilimsel gösterim]) yanı sıra, karakter ve metin ifadelerinde kullanılan kaçış dizilerinin (escape sequences [kaçış dizileri]) örnek kullanımlarını detaylandırmaktadır.

## Özellikler

- **Integer Literal (Tamsayı Sabiti):**

  - _Açıklama:_ `42` değeri, doğrudan tamsayı olarak kod içerisine yazılmıştır.
  - _Teknik Detay:_ Go derleyicisi, bu sabiti `int` türünde değerlendirir ve bellekte uygun biçimde saklar.

- **Floating Point Literal (Ondalık Sabiti):**

  - _Açıklama:_ `3.14159` değeri, ondalık sayı olarak tanımlanmıştır.
  - _Teknik Detay:_ Bu literal, `float64` türünde tanımlanmış olup, yüksek hassasiyet gerektiren hesaplamalarda kullanılır.

- **Complex Literal (Karmaşık Sayı Sabiti):**

  - _Açıklama:_ `3 + 4i` ve `(2.2 + 1.1i)` ifadeleri ile iki farklı karmaşık sayı tanımlanmıştır.
  - _Teknik Detay:_ Karmaşık sayılar, gerçek (real [gerçek]) ve sanal (imaginary [imajiner]) kısımlardan oluşur. Go dilinde `complex128` türü, karmaşık sayıların yüksek hassasiyetle işlenmesini sağlar.

- **Boolean Literal (Mantıksal Sabit):**

  - _Açıklama:_ `true` ve `false` değerleri, mantıksal (boolean [mantıksal]) ifadelerin tanımlanması örneklenmiştir.
  - _Teknik Detay:_ Bu değerler, koşul ifadelerinde (conditional statements [koşul ifadeleri]) ve kontrol akışlarında yaygın olarak kullanılır.

- **Rune Literal (Karakter Sabiti):**

  - _Açıklama:_ Farklı yöntemlerle (`'A'`, `'\n'`, `'\u4F60'`, `'\x41'`) tanımlanan rune değerleri, tek bir karakteri temsil eder.
  - _Teknik Detay:_ Rune'lar, Unicode karakter setini destekler ve `int32` temelli bir veri tipidir. Bu sayede çok çeşitli karakterler doğru şekilde işlenir.

- **String Literal (Metin Sabiti):**

  - _Açıklama:_ Çift tırnaklı string literal ile normal metin ifadesi, backtick ile tanımlanan raw string literal örneği karşılaştırmalı olarak sunulmuştur.
  - _Teknik Detay:_ Raw string literal'lerde (`` `...` ``) kaçış dizileri (escape sequences [kaçış dizileri]) etkisiz hale gelir; bu özellik, özellikle dosya yolları veya düzenli ifadeler yazılırken faydalıdır.

- **Kaçış Dizileri (Escape Sequences):**

  - _Açıklama:_ Çeşitli kaçış dizileri kullanılarak (örneğin, `\a`, `\b`, `\n`, `\t` vs.) metin içerisindeki özel karakterlerin nasıl temsil edileceği gösterilmiştir.
  - _Teknik Detay:_ Bu diziler, çıktıda istenmeyen karakterlerin doğru biçimde gösterilmesini ve metin düzeninin korunmasını sağlar.

- **Farklı Sayı Sistemleri:**

  - _Açıklama:_ Hexadecimal (ondalık sayı sisteminde 16 tabanlı), binary (ikili sayı sistemi) ve octal (sekizlik sayı sistemi) literal ifadeler, alt çizgi kullanımıyla okunabilirlik artırılarak örneklendirilmiştir.
  - _Teknik Detay:_ Bu literal türleri, farklı sayı sistemlerinde yazım kolaylığı sağlar ve özellikle donanım seviyesinde veya düşük seviye programlama gerektiren durumlarda önem kazanır.

- **Bilimsel Gösterim (Scientific Notation) ve Hexadecimal Floating Point Literal:**
  - _Açıklama:_ Avogadro sayısı gibi çok büyük değerlerin üstel biçimde yazılması ve hexadecimal floating point literal kullanımı örneklenmiştir.
  - _Teknik Detay:_ Bilimsel gösterim, çok büyük veya çok küçük sayıları daha okunabilir hale getirirken, hexadecimal floating point literal ise farklı tabanlarda yazılan ondalık sayıların hesaplamada kullanılmasını sağlar.

## Kullanım

Projeyi yerel makinenizde çalıştırmak için şu adımları takip edebilirsiniz:

1. **Go Kurulumu:**  
   Go dilinin sisteminizde yüklü olduğundan emin olun. Yükleme ve yapılandırma için [Go'nun resmi web sitesini](https://golang.org) ziyaret edebilirsiniz.

2. **Projeyi İndirin:**  
   Proje dosyasını klonlayın veya indirin.

3. **Terminal Açma:**  
   Terminal veya komut satırında proje dizinine gidin.

4. **Uygulamayı Çalıştırma:**  
   Aşağıdaki komutu kullanarak uygulamayı başlatın:

   ```bash
   go run main.go
   ```

5. **Çıktıyı İnceleme:**  
   Çıktıda, farklı literal türlerinin ve kaçış dizilerinin nasıl işlendiğine dair detaylı bilgiler ekrana basılacaktır.

## Kod Açıklaması

### Paket Tanımlaması (Package Declaration)

- **Açıklama:**  
  `package main` ifadesi, uygulamanın yürütülebilir (executable [çalıştırılabilir]) ana paketini tanımlar. Bu, Go programlarının temel yapı taşıdır.

### Paket Dahil Etme (Import)

- **Açıklama:**  
  `import "fmt"` ifadesi, formatlı giriş/çıkış işlemleri için Go'nun standart `fmt` paketini programa dahil eder. Bu paket, verilerin ekrana yazdırılması ve kullanıcıdan veri alınması gibi işlemlerde kullanılır.

### Ana Fonksiyon (Main Function)

- **Açıklama:**  
  `func main()` fonksiyonu, programın başlangıç noktasıdır. Tüm literal tanımlamaları, hesaplamalar ve çıktılar bu fonksiyon içerisinde gerçekleştirilir. Go dilinde `main` fonksiyonu, derleyici tarafından otomatik olarak çağrılır.

### Literal Türleri ve Kaçış Dizileri (Literals and Escape Sequences)

- **Sayı Literalleri:**

  - _Integer Literal (Tamsayı Sabiti):_  
    `var intLiteral int = 42` ifadesi ile 42 değeri, doğrudan integer türünde sabit olarak tanımlanmıştır. Bu tür sabit ifadeler, hesaplamalarda ve mantıksal kontrol yapılarında (conditional statements) sıklıkla kullanılır.

  - _Floating Point Literal (Ondalık Sabiti):_  
    `var floatLiteral float64 = 3.14159` ifadesi, ondalık sayılar için kullanılır. Yüksek hassasiyet gerektiren hesaplamalarda, bu tür sabit ifadeler büyük önem taşır.

  - _Complex Literal (Karmaşık Sayı Sabiti):_  
    `var complexLiteral1 complex128 = 3 + 4i` ve  
    `var complexLiteral2 complex128 = (2.2 + 1.1i)` ifadeleri ile karmaşık sayılar tanımlanır. Bu sayılar, hem gerçek hem de sanal bileşenler içerir. Karmaşık hesaplamalarda (complex arithmetic [karmaşık aritmetik]) ve mühendislik uygulamalarında kullanılır.

  - _Farklı Sayı Sistemleri:_  
    Farklı tabanlarda yazılmış sayı literal'leri, örneğin:

    - Hexadecimal: `0x1FC0`
    - Binary: `0b0001111111000000`
    - Octal: `0o17700`  
      Alt çizgi kullanımı (`0x1F_C0`) ile sayının okunabilirliği artırılır. Bu kullanım, özellikle büyük sayılarla çalışırken hata yapma riskini azaltır.

  - _Bilimsel Gösterim:_  
    `avogadro := 6.02e23` ifadesi, bilimsel gösterim kullanılarak çok büyük sayıların ifade edilmesini sağlar. Bu gösterim, üstel (exponential [üs]) biçimde yazım imkanı tanır.

  - _Hexadecimal Floating Point Literal:_  
    `hexFloat := 0xABp3` ifadesi, hexadecimal tabanlı ondalık sayıların tanımlanmasını sağlar. Bu tür literaller, belirli düşük seviyeli hesaplamalarda veya sabitlerin daha okunabilir biçimde ifade edilmesinde kullanılır.

- **Boolean Literals (Mantıksal Sabitler):**

  - `var boolTrue bool = true` ve  
    `var boolFalse bool = false` ifadeleri, mantıksal değerlerin (boolean values) nasıl tanımlanacağını göstermektedir. Mantıksal ifadeler, program akışını kontrol eden if-else yapıları ve döngülerde kritik rol oynar.

- **Rune Literals (Karakter Sabiti):**

  - Farklı yöntemlerle tanımlanan rune literal'lar, tek karakter değerlerini temsil eder.  
    Örnekler:

    - `'A'`: Doğrudan karakter ifadesi.
    - `'\n'`: Yeni satır karakteri.
    - `'\u4F60'`: Unicode standardına göre "你" karakteri.
    - `'\x41'`: Hexadecimal gösterimde "A" karakteri.

  - _Teknik Detay:_ Rune'lar, `int32` temelli veri tipleridir ve geniş karakter setlerinin (Unicode [Unicode]) desteklenmesi sayesinde uluslararası uygulamalarda güvenle kullanılabilir.

- **String Literals (Metin Sabiti):**
  - İki farklı string literal kullanımı örneklendirilmiştir:
    - Çift tırnak ("") içerisinde tanımlanan string literal: Kaçış dizileri aktif çalışır.
    - Raw string literal (`` ` ` ``): Kaçış dizileri etkisizdir, bu özellik özellikle dosya yolu ya da düzenli ifadeler yazarken hata yapma riskini azaltır.
- **Kaçış Dizileri (Escape Sequences):**
  - Çeşitli kaçış dizileri kullanılarak, özel karakterlerin nasıl temsil edileceği gösterilmiştir. Örneğin:
    - `\a`: Alarm sesi (alert).
    - `\b`: Bir karakter geri alma (backspace).
    - `\n`: Yeni satır (newline).
    - `\t`: Yatay sekme (horizontal tab).
    - Diğer kaçış dizileri, metin içindeki düzen ve formatın korunması için kullanılır.

## Özet ve Teknik Notlar

- **Kodun Amacı:**  
  Bu proje, Go dilinde literal ifadelerin (sabitlerin) ve kaçış dizilerinin (escape sequences) kullanımını detaylı olarak örnekleyerek, programlama dilinin veri tiplerini nasıl yorumladığını ve formatladığını öğretmeyi amaçlamaktadır.

- **Temel Terimler:**

  - **Literal (Sabit):** Kod içerisinde doğrudan tanımlanan, değişmeyen değerlerdir.
  - **Escape Sequence (Kaçış Dizisi):** Metin içerisinde özel karakterleri temsil etmek için kullanılan karakter dizileridir.
  - **Scientific Notation (Bilimsel Gösterim):** Çok büyük veya çok küçük sayıların üstel biçimde ifade edilmesidir.
  - **Hexadecimal, Binary, Octal:** Farklı sayı sistemlerinde değerleri ifade etmek için kullanılan notasyonlardır.

- **Önemli Notlar:**
  - Her literal, Go derleyicisinin (compiler [derleyici]) bu sabitleri nasıl işlediğini gösterir.
  - Literallerin doğru kullanımı, veri doğruluğu ve hata yönetimi açısından özellikle büyük projelerde kritik öneme sahiptir.
  - Bu proje, temel veri tiplerinin ve sayı sistemlerinin yanı sıra, uluslararası karakter desteği (Unicode) ve dosya yolu gibi pratik konularda da yol gösterici niteliktedir.
