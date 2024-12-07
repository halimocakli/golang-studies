
# Golang Uygulaması

Bu proje, Golang'de yerleşik biçimlendirme karakterlerini ve bölme işlemlerini öğreten bir uygulamadır. Proje, **NaN (Not-a-Number)**, **Infinity (sonsuzluk)**, **bölme işlemleri** ve **özel durumların kontrolü** gibi kavramları açıklar.

## 📘 **Proje Genel Bakış**
Bu uygulama, aşağıdaki konuları ele alır:

- **Floating-point Bölme**: Ondalıklı sayılarla yapılan bölme işlemlerini açıklar.
- **Integer Bölme**: Tam sayılarla yapılan bölme işlemlerini açıklar.
- **NaN, Infinity ve Runtime Error**: Bu özel durumların nasıl oluştuğunu ve kontrol edileceğini gösterir.
- **Biçimlendirme Karakterleri**: Golang'de printf ve benzeri fonksiyonlarda kullanılan biçimlendirme karakterlerini açıklar.

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

3. **Sonuçları inceleyin**. Biçimlendirme karakterlerinin nasıl çalıştığını ve bölme işlemlerinin nasıl özel durumlar oluşturduğunu gözlemleyin.

---

## 🔧 **Gereksinimler**

- Golang (1.18 veya daha yüksek sürüm önerilir)

---

## 🛠️ **Kod Açıklaması**

Uygulama şu adımları izler:

1. **Bölme İşlemleri**:
   - Numerator (pay) ve denominator (payda) listelerindeki her bir eleman için **num / denom** işlemi yapılır.
   - Bölme işleminin sonucunda aşağıdaki özel durumlar kontrol edilir:
     - **NaN**: Eğer pay ve payda sıfırsa `(0.0 / 0.0)`, sonuç NaN olur.
     - **Pozitif Sonsuzluk**: Eğer pozitif bir sayı 0.0'a bölünürse `(10.0 / 0.0)`, sonuç +Infinity olur.
     - **Negatif Sonsuzluk**: Eğer negatif bir sayı 0.0'a bölünürse `(-5.0 / 0.0)`, sonuç -Infinity olur.
     - **Normal Sonuç**: Diğer durumlarda normal bir sonuç döner.
   - Bu özel durumlar, `math.IsNaN()`, `math.IsInf(result, 1)` ve `math.IsInf(result, -1)` kullanılarak kontrol edilir.

2. **Biçimlendirme Karakterleri**:
   - `fmt.Printf()` ile biçimlendirme karakterlerinin kullanımı gösterilir.
   - Bu karakterlerin nasıl çalıştığı açıklanır.

---

## 📂 **Proje Yapısı**

```
├── main.go       # Ana Golang uygulama dosyası
```

---

## 📘 **Örnek Kullanım**

```bash
$ go run main.go
```
Komutun çıktısı, bölme işlemlerinin sonuçlarını ve biçimlendirme karakterlerinin kullanımını gösterecektir.

---

## 🔍 **Kullanılan Golang Matematik Fonksiyonları**

| **İşlem**                | **Kullanılan Fonksiyon**  | **Açıklama**                                      |
|------------------------|--------------------------|-------------------------------------------------|
| Bölme (Ondalıklı)       | `num / denom`             | Ondalıklı sayılarla bölme işlemi.                 |
| NaN Kontrolü            | `math.IsNaN(x)`           | Değerin NaN (Not-a-Number) olup olmadığını kontrol eder |
| Pozitif Sonsuzluk Kontrolü | `math.IsInf(x, 1)`      | Değerin +Infinity (sonsuz) olup olmadığını kontrol eder |
| Negatif Sonsuzluk Kontrolü| `math.IsInf(x, -1)`     | Değerin -Infinity (sonsuz) olup olmadığını kontrol eder |

---

## 📚 **Golang Biçimlendirme Karakterleri**

### 1. Genel Yer Tutucular

| **Yer Tutucu** | **Açıklama**                   | **Örnek**                |
|----------------|---------------------------------|--------------------------|
| `%v`           | Varsayılan biçimlendirme        | `[1 2 3]`, `42`           |
| `%+v`          | Struct alanlarını yazdırır     | `{Name: Go}`              |
| `%#v`          | Go sözdiziminde biçimlendirir  | `[]int{1, 2, 3}`          |
| `%T`           | Değişkenin türünü yazdırır     | `int`, `string`           |
| `%%`           | `%` sembolünü yazdırır         | `%`                      |

### 2. Tamsayılar (Integers)

| **Yer Tutucu** | **Açıklama**                   | **Örnek**                |
|----------------|---------------------------------|--------------------------|
| `%d`           | Ondalık sayı (decimal)         | `42`                      |
| `%b`           | İkilik sayı (binary)           | `101010`                  |
| `%o`           | Sekizlik sayı (octal)          | `52`                      |
| `%x`           | Onaltılık sayı (hex) (küçük)   | `2a`                      |
| `%X`           | Onaltılık sayı (hex) (büyük)   | `2A`                      |
| `%c`           | ASCII karakteri                | `A`                       |
| `%q`           | ASCII karakterini tırnak içinde| `'A'`                     |

### 3. Gerçek Sayılar (Floats)

| **Yer Tutucu** | **Açıklama**                   | **Örnek**                |
|----------------|---------------------------------|--------------------------|
| `%f`           | Ondalık gösterim (float)       | `3.141590`                |
| `%e`           | Bilimsel gösterim (e)          | `1.234568e+03`            |
| `%E`           | Bilimsel gösterim (E)          | `1.234568E+03`            |
| `%g`           | `%f` veya `%e`'den kısa olanı  | `1234.57`                 |
| `%G`           | `%f` veya `%E`'den kısa olanı  | `1234.57`                 |

---

## 📜 **Lisans**

Bu proje açık kaynaklıdır ve eğitim amaçlı serbestçe kullanılabilir.
