
# Complex Number Operations in Go

## 📘 Program Explanation

Bu program, kullanıcıdan iki karmaşık sayı (complex number) girmesini isteyen ve bu sayılarla temel aritmetik işlemleri gerçekleştiren bir Go (Golang) uygulamasıdır. Kullanıcı, "3.4+4.5i" formatında karmaşık sayılar girebilir. Bu program, kullanıcı girdilerini alır, karmaşık sayıların gerçek (real) ve sanal (imaginary) bileşenlerini ayırır ve bu bileşenlerle aritmetik işlemleri gerçekleştirir.

---

## 🎯 Programın Amacı

Bu program, Go programlama dilinde `complex128` türünün nasıl kullanılacağını öğretmek için oluşturulmuştur. Kullanıcı girdilerini okuma (`fmt.Scan`), hataları yönetme (`log.Fatalf`), ve temel matematiksel işlemleri (toplama, çıkarma, çarpma, bölme) gerçekleştirme konularında bir eğitim aracı olarak kullanılabilir.

---

## 📋 Program Akışı

1. Kullanıcıdan iki karmaşık sayı (complex number) girmesi istenir.
2. Girilen karmaşık sayılar ekranda gösterilir.
3. Her karmaşık sayının gerçek (real) ve sanal (imaginary) bileşenleri ayrı ayrı yazdırılır.
4. İki karmaşık sayı arasında toplama, çıkarma, çarpma ve bölme işlemleri gerçekleştirilir.

---

## 📂 Fonksiyonların Amacı

### 1️⃣ `main()`
Bu fonksiyon programın başlangıç noktasıdır. Kullanıcıdan karmaşık sayı girişini alır, girdileri işler ve aritmetik işlemleri gerçekleştirir.

### 2️⃣ `scanComplexNumber(prompt string) complex128`
**Amaç:** Kullanıcıdan karmaşık bir sayı (complex number) alır.  
**Parametre:** Kullanıcıya gösterilecek metin (prompt).  
**Dönüş Tipi:** `complex128`  
**Detay:** `fmt.Scan` kullanılarak kullanıcıdan giriş alınır. Hatalı giriş yapıldığında `log.Fatalf` ile program sonlandırılır.  

**Kod Parçası:**  
```go
func scanComplexNumber(prompt string) complex128 {
    var c complex128
    fmt.Print(prompt)
    _, err := fmt.Scan(&c)
    if err != nil {
        log.Fatalf("Wrong input: %v", err)
    }
    return c
}
```

### 3️⃣ `printComplexNumber(label string, c complex128)`
**Amaç:** Belirtilen etiketi (label) ve karmaşık sayıyı (complex number) ekrana yazdırır.  
**Parametreler:**  
- `label`: Yazdırılacak metin etiketi.  
- `c`: Yazdırılacak karmaşık sayı.  

**Kod Parçası:**  
```go
func printComplexNumber(label string, c complex128) {
    fmt.Printf("%s: %f\n", label, c)
}
```

### 4️⃣ `printRealAndImagParts(label string, c complex128)`
**Amaç:** Belirtilen karmaşık sayının (complex number) gerçek (real) ve sanal (imaginary) bileşenlerini ekrana yazdırır.  
**Parametreler:**  
- `label`: Yazdırılacak metin etiketi.  
- `c`: İşlenecek karmaşık sayı.  

**Kod Parçası:**  
```go
func printRealAndImagParts(label string, c complex128) {
    fmt.Printf("%s Real part: %f\n", label, real(c))
    fmt.Printf("%s Imaginary part: %f\n", label, imag(c))
}
```

### 5️⃣ `performArithmeticOperations(label1, label2 string, c1, c2 complex128)`
**Amaç:** İki karmaşık sayı (complex number) arasında toplama, çıkarma, çarpma ve bölme işlemlerini gerçekleştirir.  
**Parametreler:**  
- `label1`: Birinci karmaşık sayı etiketi.  
- `label2`: İkinci karmaşık sayı etiketi.  
- `c1`: Birinci karmaşık sayı.  
- `c2`: İkinci karmaşık sayı.  

**Kod Parçası:**  
```go
func performArithmeticOperations(label1, label2 string, c1, c2 complex128) {
    fmt.Printf("%s: %f, %s: %f\n", label1, c1, label2, c2)
    fmt.Printf("Sum: %f\n", c1+c2)
    fmt.Printf("Difference: %f\n", c1-c2)
    fmt.Printf("Product: %f\n", c1*c2)
    if c2 == 0 {
        fmt.Println("Cannot divide by zero")
        return
    }
    fmt.Printf("Quotient: %f\n", c1/c2)
}
```

---

## 🔧 Kullanılan Yazılım Terimleri

- **Complex Number (Karmaşık Sayı)**: Matematikte, bir gerçek (real) ve bir sanal (imaginary) bileşenden oluşan bir sayıdır.  
- **Real Part (Gerçek Bileşen)**: Karmaşık bir sayının reel kısmıdır.  
- **Imaginary Part (Sanal Bileşen)**: Karmaşık bir sayının sanal kısmıdır.  
- **Arithmetic Operations (Aritmetik İşlemler)**: Toplama, çıkarma, çarpma ve bölme işlemlerini ifade eder.  
- **Standard Input (Standart Giriş)**: Kullanıcıdan veri girişi alma işlemidir. Go'da `fmt.Scan` ile yapılır.  
- **Error Handling (Hata Yönetimi)**: Programda oluşabilecek hataları kontrol etme işlemidir. Bu programda `log.Fatalf` ile hatalı giriş yönetilir.  

---

## 🔍 Örnek Kullanım

```
Enter the first complex number [like 3.4+4.5i or (3.4+4.5i)]: 3+4i
Enter the second complex number [like 3.4+4.5i or (3.4+4.5i)]: 1+2i

-----------------------------------
First complex number: 3.000000+4.000000i
Second complex number: 1.000000+2.000000i
-----------------------------------
Complex1 Real part: 3.000000
Complex1 Imaginary part: 4.000000
-----------------------------------
Complex2 Real part: 1.000000
Complex2 Imaginary part: 2.000000
-----------------------------------
Complex1: 3.000000+4.000000i, Complex2: 1.000000+2.000000i
Sum: 4.000000+6.000000i
Difference: 2.000000+2.000000i
Product: -5.000000+10.000000i
Quotient: 2.200000-0.400000i
```

---

## 🧩 Sonuç

Bu program, Go dilinde karmaşık sayılarla aritmetik işlemleri öğretmek için bir araçtır. Kullanıcı girdisi alma, hata yönetimi ve aritmetik işlemler konularında pratik yapma fırsatı sunar. Go dilindeki `complex128` türü ile çalışmayı öğrenmek isteyen geliştiriciler için faydalı bir kaynaktır.
