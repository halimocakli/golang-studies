
# 📘 Go Format Specifiers Guide

Bu proje, Go dilinde kullanılan **tüm formatlama karakterlerini ve yer tutucularını** içeren kapsamlı bir örnek sunmaktadır. Formatlama karakterleri, özellikle `fmt.Printf`, `fmt.Sprintf` ve `fmt.Fprintf` gibi fonksiyonlarla kullanılır. Bu rehber, tamsayılar, gerçek sayılar, dizgeler, boolean değerler, pointerlar, mapler ve dilimler gibi birçok veri türü için nasıl biçimlendirme yapacağınızı açıklar.

---

## 📂 **Dosya Yapısı**
```
.
├── main.go        # Format karakterlerini açıklayan ve çalıştırılabilir bir Go programı
└── README.md      # Bu belgenin kendisi
```

---

## 🚀 **Nasıl Çalıştırılır?**

1. **Go kurulu olduğundan emin olun.** [Go Kurulum Kılavuzu](https://go.dev/doc/install)

2. **Bu projeyi klonlayın:**
   ```bash
   git clone <REPO_URL>
   cd <REPO_DIRECTORY>
   ```

3. **Programı çalıştırın:**
   ```bash
   go run main.go
   ```

4. **Çıktıyı terminalde görün:**
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
%p: 0xc000012080
%v: [1 2 3]
%+v: {Name:Go}
%T: int
```

---

## 📝 **Lisans**
Bu proje **MIT Lisansı** ile lisanslanmıştır.
