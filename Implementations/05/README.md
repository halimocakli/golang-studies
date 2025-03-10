
# Golang Uygulaması

Bu proje, **NaN (Not-a-Number)**, **Infinity (sonsuzluk)** ve **log(0)** kavramlarını açıklayan bir Golang uygulamasıdır. Uygulama, bu durumların nasıl oluştuğunu ve nasıl kontrol edileceğini öğretir.

## 📘 **Proje Genel Bakış**
Bu uygulama, aşağıdaki işlemleri ele alır ve açıklar:

- **NaN (Not-a-Number) Değerleri**: Negatif bir sayının karekökü, (∞ - ∞) ve negatif bir sayının logaritması tanımsız olduğu için `NaN` (Not-a-Number) değerini üretir.
- **Infinity Değerleri**: `math.Log(0)` işlemi negatif sonsuzluk (-∞) döndürür.
- **Kontroller**: Bu NaN ve Infinity değerlerinin nasıl kontrol edileceği gösterilir.

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

3. **Sonuçları inceleyin**. NaN, Infinity ve log(0) durumlarının nasıl oluştuğunu ve nasıl kontrol edildiğini göreceksiniz.

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
   - **Negatif Sayı**: `math.Log(-10)` ifadesi, -10'un doğal logaritmasını hesaplamaya çalışır.
     - Negatif bir sayının logaritması tanımsız olduğu için `NaN` döner.
     - Bu değerin NaN olup olmadığı `math.IsNaN(x)` ile kontrol edilir.
   - **Log(0)**: `math.Log(0)` ifadesi, 0'ın doğal logaritmasını hesaplamaya çalışır.
     - 0'ın logaritması, negatif sonsuzluk (-∞) olarak tanımlanır.
     - `math.IsInf(x, 0)` fonksiyonu ile bu değerin sonsuz olup olmadığı kontrol edilir.

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

math.Log(0) = -Inf
logZero is negative Infinity!
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
| Logaritma (0)           | `math.Log(0)`             | 0'ın logaritması negatif sonsuzluk (-∞) döner.   |
| NaN Kontrolü            | `math.IsNaN(x)`           | Değerin NaN (Not-a-Number) olup olmadığını kontrol eder |
| Infinity Kontrolü       | `math.IsInf(x, 0)`        | Değerin sonsuz olup olmadığını kontrol eder (pozitif ya da negatif) |

---

## ❗ **Hata Yönetimi**

- NaN (Not-a-Number) değerlerini tespit etmek için `math.IsNaN(x)` fonksiyonu kullanılır.
- Infinity değerlerini tespit etmek için `math.IsInf(x, 0)` fonksiyonu kullanılır.
- Kareköklü, logaritmalı veya infinity işlemleri ile oluşan NaN ve sonsuzluk değerleri kullanıcıya bildirir.

---

## 📚 **Kullanılan Golang Kavramları**

- **`fmt` Paketi**: Kullanıcıdan giriş almak ve ekrana çıktı vermek için kullanılır.
- **`math` Paketi**: Logaritma, Infinity, NaN işlemleri için kullanılır.

---

## 📜 **Lisans**

Bu proje açık kaynaklıdır ve eğitim amaçlı serbestçe kullanılabilir.
