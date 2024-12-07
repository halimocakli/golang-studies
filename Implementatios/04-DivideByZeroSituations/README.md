
# Golang Uygulaması

Bu proje, **Bölme İşlemleri** ve **Özel Durumlar (NaN, +Infinity, -Infinity, Runtime Error)** üzerine bir Golang uygulamasıdır. Uygulama, **floating-point** (ondalıklı) ve **integer** (tam sayı) bölme işlemlerinin farklarını açıklar ve özel durumların nasıl ortaya çıktığını gösterir.

## 📘 **Proje Genel Bakış**
Bu uygulama, aşağıdaki işlemleri ele alır ve açıklar:

- **Floating-point Bölme**: Ondalıklı sayılarla yapılan bölme işlemleri. `0.0` ile bölme özel durumlar oluşturabilir.
- **Tam Sayı Bölme**: Tam sayılarla yapılan bölme işlemleri. `0` ile bölme **Runtime Error** ile sonuçlanır.
- **Özel Durumlar**: NaN (Not-a-Number), +Infinity (pozitif sonsuzluk) ve -Infinity (negatif sonsuzluk) durumlarının nasıl oluştuğunu gösterir.
- **Kontroller**: NaN ve sonsuzluk değerlerinin nasıl kontrol edileceği gösterilir.

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

3. **Sonuçları inceleyin**. Bölme işlemleri, NaN, Infinity ve Runtime Error durumlarının nasıl oluştuğunu göreceksiniz.

---

## 🔧 **Gereksinimler**

- Golang (1.18 veya daha yüksek sürüm önerilir)

---

## 🛠️ **Kod Açıklaması**

Uygulama şu adımları izler:

1. **Floating-point Bölme İşlemi**:
   - Numerator (pay) ve denominator (payda) listelerindeki her bir eleman için **num / denom** işlemi yapılır.
   - Bölme işleminin sonucunda aşağıdaki özel durumlar kontrol edilir:
     - **NaN**: Eğer pay ve payda sıfırsa `(0.0 / 0.0)`, sonuç NaN olur.
     - **Pozitif Sonsuzluk**: Eğer pozitif bir sayı 0.0'a bölünürse `(10.0 / 0.0)`, sonuç +Infinity olur.
     - **Negatif Sonsuzluk**: Eğer negatif bir sayı 0.0'a bölünürse `(-5.0 / 0.0)`, sonuç -Infinity olur.
     - **Normal Sonuç**: Diğer durumlarda normal bir sonuç döner.
   - Bu özel durumlar, `math.IsNaN()`, `math.IsInf(result, 1)` ve `math.IsInf(result, -1)` kullanılarak kontrol edilir.

2. **Tam Sayı Bölme İşlemi**:
   - Eğer `10 / 0` gibi bir ifade çalıştırılırsa, **Runtime Error: integer divide by zero** hatası alınır.
   - Bu hatanın önüne geçmek için `0` ile bölme yapmaktan kaçınılmalıdır.

---

## 📂 **Proje Yapısı**

```
├── main.go       # Ana Golang uygulama dosyası
```

---

## 📘 **Örnek Kullanım**

```bash
$ go run main.go

Floating-point division: 10.00 / 2.00 = 5.00
Floating-point division: 0.00 / 0.00 = NaN (Not a Number)
Floating-point division: -5.00 / 0.00 = -Inf (Negative Infinity)
Floating-point division: 10.00 / 0.00 = +Inf (Positive Infinity)
Floating-point division: 0.00 / 10.00 = 0.00
```

---

## 🔍 **Kullanılan Golang Matematik Fonksiyonları**

| **İşlem**                | **Kullanılan Fonksiyon**  | **Açıklama**                                      |
|------------------------|--------------------------|-------------------------------------------------|
| Bölme (Ondalıklı)       | `num / denom`             | Ondalıklı sayılarla bölme işlemi.                 |
| NaN Kontrolü            | `math.IsNaN(x)`           | Değerin NaN (Not-a-Number) olup olmadığını kontrol eder |
| Pozitif Sonsuzluk Kontrolü | `math.IsInf(x, 1)`      | Değerin +Infinity (sonsuz) olup olmadığını kontrol eder |
| Negatif Sonsuzluk Kontrolü| `math.IsInf(x, -1)`     | Değerin -Infinity (sonsuz) olup olmadığını kontrol eder |
| Bölme (Tam Sayı)        | `10 / 0`                  | Tam sayılarla bölme, **Runtime Error** üretir.     |

---

## ❗ **Özel Durumlar ve Hatalar**

| **İşlem**                | **Sonuç**               | **Açıklama**                                 |
|------------------------|------------------------|---------------------------------------------|
| **10 / 2**              | 5.00                   | Normal bölme işlemi.                        |
| **0 / 0**               | NaN                    | 0'ın 0'a bölümü NaN olarak tanımlanır.      |
| **-5 / 0**              | -Infinity              | Negatif bir sayının 0'a bölümü -∞ döner.    |
| **10 / 0**              | +Infinity              | Pozitif bir sayının 0'a bölümü +∞ döner.    |
| **10 / 0** (tam sayı)   | Runtime Error          | Tam sayılarda 0'a bölme **Runtime Error** üretir.|

---

## 📚 **Kullanılan Golang Kavramları**

- **`fmt` Paketi**: Çıktı işlemleri için kullanılır.
- **`math` Paketi**: NaN, Infinity ve bölme işlemleri ile ilgili özel durumları kontrol etmek için kullanılır.

---

## 📜 **Lisans**

Bu proje açık kaynaklıdır ve eğitim amaçlı serbestçe kullanılabilir.
