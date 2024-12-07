
# Golang Uygulaması

Bu proje, aşağıdaki işlemleri gerçekleştiren bir Golang uygulamasıdır:

## 📘 **Proje Genel Bakış**
Bu uygulama, Golang'de bölme işlemleriyle ilgili çeşitli özel durumları ele alır. Kod, bölme işlemlerinin sonucunu analiz eder ve NaN, +Infinity ve -Infinity durumlarını kontrol eder.

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

3. **Sonuçları inceleyin**. Bölme işlemlerinin nasıl çalıştığını ve özel durumların nasıl tespit edildiğini gözlemleyin.

---

## 🔧 **Gereksinimler**

- Golang (1.18 veya daha yüksek sürüm önerilir)

---

## 🛠️ **Kod Açıklaması**

Bu kod, bölme işlemlerinin farklı durumlarını kontrol etmek için aşağıdaki işlemleri içerir:

1. **Bölme İşlemleri**:
   - Numerator (pay) ve denominator (payda) listelerindeki her bir eleman için **num / denom** işlemi yapılır.
   - Bölme işleminin sonucunda aşağıdaki özel durumlar kontrol edilir:
     - **NaN**: Eğer pay ve payda sıfırsa `(0.0 / 0.0)`, sonuç NaN olur.
     - **Pozitif Sonsuzluk**: Eğer pozitif bir sayı 0.0'a bölünürse `(10.0 / 0.0)`, sonuç +Infinity olur.
     - **Negatif Sonsuzluk**: Eğer negatif bir sayı 0.0'a bölünürse `(-5.0 / 0.0)`, sonuç -Infinity olur.
     - **Normal Sonuç**: Diğer durumlarda normal bir sonuç döner.
   - Bu özel durumlar, `math.IsNaN()`, `math.IsInf(result, 1)` ve `math.IsInf(result, -1)` kullanılarak kontrol edilir.

2. **Kodun Gövdesi**:
   - Aşağıda, projenin `main.go` dosyasındaki tam kod bulunmaktadır:
   
```go
package main

import (
	"fmt"
)

func main() {
	// Tamsayılar
	fmt.Printf("%%d: %d\n", 42)
	fmt.Printf("%%b: %b\n", 42)
	fmt.Printf("%%o: %o\n", 42)
	fmt.Printf("%%x: %x\n", 42)
	fmt.Printf("%%X: %X\n", 42)
	fmt.Printf("%%c: %c\n", 65)
	fmt.Printf("%%q: %q\n", 65)

	// Gerçek Sayılar
	fmt.Printf("%%f: %f\n", 3.14159)
	fmt.Printf("%%e: %e\n", 1234.5678)
	fmt.Printf("%%g: %g\n", 1234.5678)

	// Dizgeler
	fmt.Printf("%%s: %s\n", "hello")
	fmt.Printf("%%q: %q\n", "hello")
	fmt.Printf("%%x: %x\n", "hello")

	// Pointer
	x := 42
	fmt.Printf("%%p: %p\n", &x)

	// Genel Yer Tutucular
	fmt.Printf("%%v: %v\n", []int{1, 2, 3})
	fmt.Printf("%%+v: %+v\n", struct{ Name string }{"Go"})
	fmt.Printf("%%T: %T\n", x)
}

```

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
Komutun çıktısı, bölme işlemlerinin sonuçlarını gösterecektir. Özel durumlar (NaN, +Infinity, -Infinity) hakkında bilgi sağlayacaktır.

---

## 📜 **Lisans**

Bu proje açık kaynaklıdır ve eğitim amaçlı serbestçe kullanılabilir.
