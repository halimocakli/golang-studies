
# Golang Uygulaması

Bu proje, kullanıcının girdiği bir ondalık sayının yuvarlama, aşağı yuvarlama (floor) ve yukarı yuvarlama (ceil) işlemlerini gerçekleştiren temel matematiksel işlemleri Golang ile göstermektedir.

## 📘 **Proje Genel Bakış**
Bu basit komut satırı uygulaması, kullanıcının girdiği bir ondalık sayıyı alır ve Golang'in standart `math` paketini kullanarak bu sayının yuvarlanmış, aşağı yuvarlanmış ve yukarı yuvarlanmış değerlerini ekrana yazdırır.

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

3. **Ondalık bir sayı girin** (örneğin, `3.14`) ve çıktı değerlerini inceleyin.

---

## 🔧 **Gereksinimler**

- Golang (1.18 veya daha yüksek sürüm önerilir)

---

## 🛠️ **Kod Açıklaması**

Uygulama şu adımları izler:

1. **Girdi İşleme**: 
   - Kullanıcıdan bir ondalık sayı girmesini ister.
   - Girdiyi okur ve herhangi bir hata olup olmadığını kontrol eder.
   
2. **Matematiksel İşlemler**:
   - **Original**: Girdi olarak alınan sayıyı olduğu gibi gösterir.
   - **Rounded**: Girdiyi en yakın tam sayıya yuvarlar (`math.Round()`).
   - **Floored**: Girdiden küçük veya ona eşit olan en büyük tam sayıyı döndürür (`math.Floor()`).
   - **Ceiled**: Girdiden büyük veya ona eşit olan en küçük tam sayıyı döndürür (`math.Ceil()`).

---

## 📂 **Proje Yapısı**

```
├── main.go       # Ana Golang uygulama dosyası
```

---

## 📘 **Örnek Kullanım**

```bash
$ go run main.go
Input a float value (e.g., 3.14): 3.14

Original: 3.14
Rounded: 3.00
Floored: 3.00
Ceiled: 4.00
```

---

## ❗ **Hata Yönetimi**

- Kullanıcı geçersiz bir giriş (örneğin bir metin) girerse, uygulama bir hata mesajıyla çıkış yapar.

---

## 📚 **Kullanılan Golang Kavramları**

- **`fmt` Paketi**: Girdi ve çıktı işlemleri için kullanılır.
- **`math` Paketi**: Yuvarlama, aşağı yuvarlama ve yukarı yuvarlama işlemleri için kullanılır.
- **`log` Paketi**: Hataların yakalanması ve hatalı girişlerde uygulamanın sonlandırılması için kullanılır.

---

## 📜 **Lisans**

Bu proje açık kaynaklıdır ve eğitim amaçlı serbestçe kullanılabilir.
