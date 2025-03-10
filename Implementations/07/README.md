# Go Format Specifiers Example

Bu proje, Go programlama dilinde `fmt.Printf` fonksiyonunu kullanarak farklı biçimlendirme (format) karakterlerinin nasıl kullanılacağını gösterir.

## 📘 Amaç

Bu proje, tamsayılar, gerçek sayılar ve dizgeler (string) için kullanılabilecek format karakterlerinin örneklerini sunar. Kodun amacı, bu karakterlerin Go'da nasıl çalıştığını anlamaktır.

## 📂 Dosya Yapısı

```
📁 main.go - Format karakterlerini örneklerle açıklayan ana Go dosyası
```

## 🚀 Kullanım

Bu programı çalıştırmak için aşağıdaki adımları takip edebilirsiniz:

1. Go yüklü olduğundan emin olun. Eğer yüklü değilse, [Go'yu buradan indirin](https://go.dev/doc/install).
2. Terminal veya komut istemcisinde `main.go` dosyasının bulunduğu dizine gidin.
3. Aşağıdaki komutu çalıştırın:

```bash
go run main.go
```

Bu komut, programın çıktısını terminalde gösterecektir.

---

## 🧑‍💻 Örnek Çıktı

Aşağıdaki örnek, programın terminalde ürettiği çıktıya benzer olabilir:

```
%d: 42
%b: 101010
%o: 52
%x: 2a
%X: 2A
%c: A
%q: 'A'
%f: 3.141590
%e: 1.234568e+03
%g: 1234.57
%s: hello
%q: "hello"
%x: 68656c6c6f
```

---

## 📝 Format Karakterleri Tablosu

### 1️⃣ Genel Yer Tutucular (General Placeholders)

| Yer Tutucu | Açıklama                         | Örnek            |
|------------|---------------------------------|------------------|
| `%v`       | Varsayılan biçimlendirme         | `[1 2 3]`, `42`   |
| `%+v`      | Struct alanlarını yazdırır       | `{Name: Go}`     |
| `%#v`      | Go sözdiziminde biçimlendirir   | `[]int{1, 2, 3}`  |
| `%T`       | Değişkenin türünü yazdırır       | `int`, `string`  |
| `%%`       | `%` sembolünü yazdırır           | `%`              |

---

### 2️⃣ Tamsayılar (Integers)

| Yer Tutucu | Açıklama                         | Örnek             |
|------------|---------------------------------|-------------------|
| `%d`       | Ondalık sayı (decimal)           | `42`              |
| `%b`       | İkilik sayı (binary)             | `101010`          |
| `%o`       | Sekizlik sayı (octal)            | `52`              |
| `%x`       | Onaltılık sayı (hexadecimal) (küçük harf) | `2a`          |
| `%X`       | Onaltılık sayı (hexadecimal) (büyük harf) | `2A`          |
| `%c`       | ASCII karakteri olarak yazdırır  | `A`               |
| `%q`       | ASCII karakterini tırnak içinde yazdırır | `'A'`         |

---

### 3️⃣ Gerçek Sayılar (Floats)

| Yer Tutucu | Açıklama                         | Örnek             |
|------------|---------------------------------|-------------------|
| `%f`       | Ondalık gösterim (float)         | `3.141590`        |
| `%e`       | Bilimsel gösterim (e)            | `1.234568e+03`    |
| `%E`       | Bilimsel gösterim (E)            | `1.234568E+03`    |
| `%g`       | `%f` veya `%e`'den kısa olanı seçer | `1234.57`      |
| `%G`       | `%f` veya `%E`'den kısa olanı seçer | `1234.57`      |

---

Bu `README.md` dosyası, format karakterleri hakkında temel bilgileri ve `main.go` dosyasının nasıl çalıştırılacağını sunmaktadır.

## 📚 Daha Fazla Bilgi

Go'nun `fmt` paketine dair daha fazla bilgiye [Go'nun resmi belgesinden](https://pkg.go.dev/fmt) ulaşabilirsiniz.

