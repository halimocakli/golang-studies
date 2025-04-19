# Go Dilinde Rakamlar Toplamı Hesaplama Uygulaması

Bu repository, kullanıcıdan alınan bir tamsayının rakamlarının toplamını hesaplayan bir Go (Go programming language [Go programlama dili]) uygulamasıdır. Uygulama, temel Go dil özelliklerini (main function [ana fonksiyon], for loop [döngü], if statement [if deyimi]) ve standart kütüphaneleri (fmt, log) kullanarak geliştirilmiştir. Aşağıda, her fonksiyonun ve algoritmanın detaylı açıklamaları ile projenin işleyişi kapsamlı bir şekilde anlatılmıştır.

---

## İçindekiler

- [Genel Bakış](#genel-bakış)
- [Özellikler](#özellikler)
- [Kullanım](#kullanım)
- [Kod Açıklaması](#kod-açıklaması)
  - [Ana Fonksiyon ve Test Süreci](#ana-fonksiyon-ve-test-süreci)
  - [Kullanıcı Girdisi ve Hata Yönetimi](#kullanıcı-girdisi-ve-hata-yönetimi)
  - [Rakamların Toplamının Hesaplanması](#rakamların-toplamının-hesaplanması)
  - [Negatif Sayıların İşlenmesi](#negatif-sayıların-işlenmesi)
- [Özet ve Teknik Notlar](#özet-ve-teknik-notlar)

---

## Genel Bakış

Bu proje, Go dilinde geliştirilen basit bir uygulamadır. Kullanıcıdan alınan tamsayının rakamlarının toplamını hesaplamak amacıyla oluşturulan bu uygulama, aşağıdaki önemli teknik konseptleri içerir:

- **Kullanıcı Girdisi Alma:**  
  Program, konsoldan kullanıcı girdisi alarak çalışır. Bu işlem, `fmt.Scan` fonksiyonu kullanılarak gerçekleştirilir. Bu yöntem, kullanıcıdan dinamik veri alma ve programın çalışma zamanında karar verme imkanı sağlar.

- **Hata Yönetimi:**  
  Girdi işlemi sırasında oluşabilecek hatalar, `log.Fatalf` kullanılarak yakalanır. Bu sayede, beklenmedik durumlarda programın güvenli bir şekilde sonlandırılması sağlanır.

- **Algoritma:**  
  Girilen sayının rakamlarının toplamı, klasik bir bölme (division) ve mod alma (modulus) algoritması kullanılarak hesaplanır. Bu yöntem, her seferinde sayının son basamağını ayrıştırıp toplam değere ekleyip, sayıyı güncelleyerek çalışır.

- **Negatif Sayı İşlemleri:**  
  Negatif sayı girilmesi durumunda, uygulama bu değeri pozitif hale getirerek işlemeye devam eder. Bu, `abs` fonksiyonuyla sağlanır. Böylece, rakamların toplamı her zaman doğru hesaplanır.

---

## Özellikler

- **Kullanıcı Girdisi (User Input [Kullanıcı Girdisi]):**

  - Konsoldan girilen tamsayı, `fmt.Scan` fonksiyonu ile okunur.
  - Girdi alma işlemi, kullanıcı ile etkileşim kurarak dinamik veri toplanmasını sağlar.

- **Hata Yönetimi (Error Handling [Hata Yönetimi]):**

  - Giriş sırasında oluşan hatalar `log.Fatalf` fonksiyonu ile yakalanır.
  - Hata anında, program çalışmasını durdurur ve hata mesajı basılarak sorunun detayları kullanıcıya iletilir.

- **Negatif Sayıların İşlenmesi (Handling Negative Numbers [Negatif Sayıların İşlenmesi]):**

  - Kullanıcı negatif bir sayı girdiğinde, `abs` fonksiyonu çağrılarak sayı pozitif hale getirilir.
  - Bu durum, rakamların toplamı hesaplamasında negatif işaretin hesaba katılmaması için gereklidir.

- **Rakamların Toplamının Hesaplanması (Digit Summation [Rakamların Toplamının Hesaplanması]):**
  - `sumThreeDigits` fonksiyonu, verilen tamsayının her basamağını mod alma ve bölme işlemleriyle ayırarak toplamını hesaplar.
  - Algoritma, sayının her basamağını tek tek işleyip toplam değeri güncelleyerek çalışır.

---

## Kullanım

Projeyi yerel ortamınızda çalıştırmak için aşağıdaki adımları takip ediniz:

1. **Go Kurulumu (Go Installation [Go Kurulumu]):**

   - Sisteminizde Go dilinin yüklü olduğundan emin olun.
   - Go'nun resmi web sitesinden (https://golang.org) indirme ve kurulum adımlarını uygulayabilirsiniz.

2. **Projeyi İndirme (Clone/Download [Projeyi İndirme]):**

   - Repository’i klonlayarak veya zip dosyası olarak indirerek yerel bilgisayarınıza alın.

3. **Projeyi Çalıştırma (Run the Project [Projeyi Çalıştırma]):**

   - Terminal veya komut satırında proje dizinine gidin.
   - Aşağıdaki komutu kullanarak uygulamayı çalıştırın:
     ```bash
     go run main.go
     ```

4. **Girdi Sağlama (Provide Input [Girdi Sağlama]):**
   - Konsol ekranında sayı girmeniz istenecek.
   - Girilen sayı, rakamlarına ayrılarak toplam değeri hesaplanacak ve ekrana yazdırılacaktır.

---

## Kod Açıklaması

Aşağıda, her bir fonksiyon ve algoritmanın detaylı açıklamaları yer almaktadır:

### Ana Fonksiyon ve Test Süreci

- **`main` Fonksiyonu (Main Function [Ana Fonksiyon]):**

  - Programın çalışmaya başladığı ilk yerdir.
  - `main` fonksiyonu, programın giriş noktası olarak `runTest` fonksiyonunu çağırır.
  - Bu yapı, modüler test sürecini başlatmak için kullanılır.

- **`runTest` Fonksiyonu:**
  - Test süreçlerini organize eder.
  - İçerisinde `testSumThreeDigits` fonksiyonu çağrılarak rakamların toplamı hesaplama süreci başlatılır.
  - Bu yöntem, fonksiyonların ayrıştırılmış şekilde test edilmesini ve olası hata durumlarının izole edilmesini sağlar.

---

### Kullanıcı Girdisi ve Hata Yönetimi

- **`testSumThreeDigits` Fonksiyonu:**

  - **Kullanıcı Girdisi Alma:**

    - Konsol üzerinden kullanıcıdan tamsayı girişi istenir.
    - `fmt.Scan` fonksiyonu ile girilen değer okunur ve `num` değişkenine atanır.

  - **Hata Yönetimi:**

    - Girdi alma işlemi sırasında hata oluşursa, `err` değişkeni ile hata bilgisi elde edilir.
    - `log.Fatalf` kullanılarak hata mesajı basılır ve program güvenli bir şekilde sonlandırılır.
    - Bu, programın beklenmedik durumlarda çökmesini önler.

  - **Negatif Sayı Kontrolü:**

    - Girilen sayı negatif ise, `if` kontrol bloğu devreye girer.
    - Negatif değer, `abs` fonksiyonu aracılığıyla pozitif hale getirilir.
    - Bu işlem, rakamların toplamının doğru hesaplanabilmesi için gereklidir.

  - **Hesaplama ve Çıktı:**
    - `sumThreeDigits` fonksiyonu çağrılarak, sayının rakamlarının toplamı hesaplanır.
    - Hesaplanan toplam, `fmt.Printf` fonksiyonu ile ekrana yazdırılır.

---

### Rakamların Toplamının Hesaplanması

- **`sumThreeDigits` Fonksiyonu:**

  - **Algoritma Açıklaması:**

    - Fonksiyon, parametre olarak aldığı `number` değişkeni üzerinde çalışır.
    - Bir **for loop (döngü [döngü])** kullanılarak, sayı sıfır olana kadar her adımda son basamak ayrıştırılır.

  - **Adım Adım İşleyiş:**

    - **Modulus İşlemi (number % 10):**
      - Sayının son basamağı, mod alma işlemi ile elde edilir.
      - Bu işlem, basamak değerlerinin ayrı ayrı toplanmasını sağlar.
    - **Toplamın Güncellenmesi:**
      - Elde edilen basamak değeri, `total` değişkenine eklenir.
      - Toplam değer her adımda güncellenir.
    - **Bölme İşlemi (number = number / 10):**
      - Sayı, 10’a bölünerek son basamağı kaldırılır.
      - Böylece, sıradaki basamak işleme alınır.

  - **Fonksiyon Sonu:**
    - Döngü tamamlandığında, tüm basamaklar toplanmış olur.
    - Hesaplanan toplam değer, fonksiyondan döndürülür.

---

### Negatif Sayıların İşlenmesi

- **`abs` Fonksiyonu:**
  - **Amaç:**
    - Negatif sayıların mutlak değerini (absolute value [mutlak değer]) elde etmek için kullanılır.
  - **İşleyiş:**
    - Girdi olarak alınan negatif sayı, `-value` ifadesiyle pozitif hale çevrilir.
    - Bu fonksiyon, yalnızca negatif değerler için çağrıldığından basitçe negatif işareti kaldırır.
  - **Not:**
    - Genel bir absolute value fonksiyonu, pozitif sayılar için de kontrol içerecek şekilde yazılabilir; ancak bu örnekte, sadece negatif sayı kontrolü yapıldığı için basitleştirilmiştir.

---

## Özet ve Teknik Notlar

- **Algoritma (Algorithm [Algoritma]):**

  - Sayının rakamlarının toplamını hesaplamak için, her adımda mod işlemi (number % 10) ile son basamak ayrılır, toplam değere eklenir ve sayı 10’a bölünerek güncellenir.
  - Bu süreç, sayının tamamı sıfır olana kadar tekrarlanır.

- **Error Handling (Hata Yönetimi):**

  - Giriş işlemi sırasında olası hatalar `fmt.Scan` kullanılarak kontrol edilir.
  - Hata durumunda, `log.Fatalf` ile detaylı hata mesajı basılarak program sonlandırılır.

- **Standart Kütüphaneler (Standard Library [Standart Kütüphaneler]):**

  - `fmt` paketi, formatlı giriş/çıkış işlemleri için kullanılır.
  - `log` paketi, hata durumlarını yönetmek ve programın güvenli bir şekilde sonlandırılmasını sağlamak için kullanılır.

- **Kullanılan Go Dil Yapıları:**
  - **Control Structures (Kontrol Yapıları [Control Structures]):**
    - `if` deyimleri, koşulların kontrolü için kullanılır.
    - `for` döngüsü, tekrar eden işlemler ve algoritmanın adım adım ilerlemesi için hayati öneme sahiptir.
  - **Functions (Fonksiyonlar [Functions]):**
    - Kodun modüler ve okunabilir olmasını sağlar.
    - Her fonksiyon, belirli bir görevi yerine getirmek üzere tasarlanmıştır.

Bu proje, Go dilinin temel yapı taşlarını kullanarak basit bir uygulama geliştirme sürecini detaylı olarak ortaya koymaktadır. Kullanıcı girdisinin alınması, hata yönetiminin uygulanması, negatif sayıların işlenmesi ve rakamların toplamının hesaplanması gibi konular, Go dilinin gerçek dünya uygulamaları için nasıl kullanılabileceğini göstermektedir.

## Özet

Bu Go uygulaması, kullanıcıdan alınan bir tamsayının rakamlarının toplamını hesaplamak amacıyla geliştirilmiştir.

- **Kullanıcı Girdisi:** Konsol üzerinden alınan veri, `fmt.Scan` fonksiyonu ile okunur.
- **Hata Yönetimi:** Giriş sırasında oluşabilecek hatalar, `log.Fatalf` ile yakalanarak güvenli bir şekilde program sonlandırılır.
- **Algoritma:** Her adımda mod işlemi (number % 10) ve bölme işlemi (number / 10) kullanılarak sayının basamakları ayrıştırılır ve toplanır.
- **Negatif Sayı İşlemleri:** Negatif girilen sayılar, `abs` fonksiyonu aracılığıyla pozitif hale getirilir, böylece toplam hesaplaması doğru gerçekleştirilir.
