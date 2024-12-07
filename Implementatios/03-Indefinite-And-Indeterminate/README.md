
# Golang Uygulaması

Bu proje, **NaN (Not-a-Number)** kavramını açıklayan ve çeşitli durumlarda NaN değerinin nasıl ortaya çıktığını gösteren bir Golang uygulamasıdır. Ayrıca, NaN değerlerinin nasıl kontrol edileceğini de öğretir.

## 📘 **Proje Genel Bakış**
Bu uygulama, aşağıdaki durumlarda NaN (Not-a-Number) değerlerini nasıl üretebileceğinizi ve bunları nasıl kontrol edebileceğinizi göstermektedir:

- **Karekök alma işlemi**: Negatif bir sayının karekökü alındığında (√-1), NaN oluşur.
- **Infinity işlemleri**: Sonsuzluk ile yapılan belirli işlemler NaN değerini üretir. Örnek: (+∞) - (+∞) = NaN.
- **Negatif sayıların logaritması**: Negatif bir sayının logaritması tanımlı değildir, bu nedenle NaN döner.

---

## 🚀 **Nasıl Çalıştırılır**

Bu uygulamayı çalıştırmak için şu adımları izleyin:

1. **Golang'in yüklü olduğundan emin olun**. Doğrulamak için şu komutu çalıştırın:
   ```bash
   go version
   ```

2. **Uygulamayı çalıştırın**:
   ```bash
   go run main.go
   ```

3. **Sonuçları inceleyin**. NaN değerlerinin nasıl ortaya çıktığını ve nasıl kontrol edildiğini göreceksiniz.

---

## 🔧 **Gereksinimler**

- Golang (1.18 veya daha yüksek sürüm önerilir)

---

## 🛠️ **Kod Açıklaması**

Uygulama şu adımları izler:

1. **Karekök İşlemi**: 
   - `math.Sqrt(-1)` ifadesi ile -1'in karekökü alınır.
   - Bu işlem, matematiksel olarak tanımsız olduğu için `NaN` döner.
   - `math.IsNaN(x)` fonksiyonu ile bu değerin NaN olup olmadığı kontrol edilir.

2. **Infinity İşlemi**:
   - `math.Inf(1)` ile pozitif sonsuzluk değeri (+∞) oluşturulur.
   - `(∞ - ∞)` işlemi, matematiksel olarak tanımsız olduğu için `NaN` döner.
   - Bu değerin NaN olup olmadığı `math.IsNaN(x)` ile kontrol edilir.

3. **Logaritma İşlemi**:
   - `math.Log(-10)` ifadesi, -10'un doğal logaritmasını hesaplamaya çalışır.
   - Negatif bir sayının logaritması tanımsız olduğu için `NaN` döner.
   - Bu değerin NaN olup olmadığı `math.IsNaN(x)` ile kontrol edilir.

---

## 📂 **Proje Yapısı**

```
├── main.go       # Ana Golang uygulama dosyası
```

---

## 📘 **Örnek Kullanım**

```bash
$ go run main.go

math.Sqrt(-1) = NaN
sqrtNeg is NaN (Not-a-Number)!

+Infinity - +Infinity = NaN
infMinusInf is NaN (Not-a-Number)!

math.Log(-10) = NaN
logNeg is NaN (Not-a-Number)!
```

---

## 🔍 **Kullanılan Golang Matematik Fonksiyonları**

| **İşlem**                | **Kullanılan Fonksiyon**  | **Açıklama**                                      |
|------------------------|--------------------------|-------------------------------------------------|
| Karekök (Negatif Sayı)  | `math.Sqrt(x)`            | Negatif bir sayının karekökü alınır, NaN döner.  |
| Sonsuzluk (+Infinity)   | `math.Inf(1)`             | Pozitif sonsuzluk oluşturur (+∞)                |
| Sonsuzluk (-Infinity)   | `math.Inf(-1)`            | Negatif sonsuzluk oluşturur (-∞)                |
| Sonsuzluk İşlemi        | `inf - inf`               | Sonsuzluk ile işlem yapar, NaN döner.           |
| Logaritma (Negatif Sayı)| `math.Log(x)`             | Negatif bir sayının logaritması NaN döner.      |
| NaN Kontrolü            | `math.IsNaN(x)`           | Değerin NaN (Not-a-Number) olup olmadığını kontrol eder |

---

## ❗ **Hata Yönetimi**

- NaN (Not-a-Number) değerlerini tespit etmek için `math.IsNaN(x)` fonksiyonu kullanılır.
- Kareköklü, logaritmalı veya infinity işlemleri ile oluşan NaN değerleri kullanıcıya bildirir.

---

## 📚 **Kullanılan Golang Kavramları**

- **`fmt` Paketi**: Kullanıcıdan giriş almak ve ekrana çıktı vermek için kullanılır.
- **`math` Paketi**: Logaritma, Infinity, NaN işlemleri için kullanılır.

---

## 📜 **Lisans**

Bu proje açık kaynaklıdır ve eğitim amaçlı serbestçe kullanılabilir.
