
# Literals

## 📘 Programın Amacı (Purpose of the Program)
Bu program, Golang'de çeşitli türlerdeki (data types) literalleri (literals) ve kaçış dizilerini (escape sequences) kullanmayı ve nasıl çalıştıklarını göstermek için tasarlanmıştır. Program, her bir türün örneklerini konsola (console) yazdırarak, kullanıcıya bu türlerin nasıl tanımlandığını ve kullanıldığını açıkça gösterir.

---

## 🚀 Program Nasıl Çalışır? (How Does the Program Work?)
1. **Değişken Tanımlamaları (Variable Declarations):** Program, `int`, `float64`, `complex128`, `bool`, `rune`, `string` gibi veri türlerini (data types) kullanarak değişkenler oluşturur.
2. **Değişken Değerlerinin Gösterilmesi (Printing Variable Values):** `fmt.Printf()` ve `fmt.Println()` fonksiyonları, bu değişkenlerin değerlerini konsola yazdırmak için kullanılır.
3. **Kaçış Dizileri (Escape Sequences):** Program, özel karakterleri (special characters) göstermek için kaçış dizilerini kullanır. Örneğin, yeni satır (`\n`), sekme (`\t`), geri alma (`\b`) gibi diziler.
4. **Sayısal Temsil (Numerical Representations):** Program, onaltılık (hexadecimal), ikilik (binary) ve sekizlik (octal) sayı sistemlerini kullanarak değişkenleri tanımlar ve bu sayı sistemleri arasındaki farkları gösterir.
5. **Veri Türlerini Gösterme (Data Type Display):** `fmt.Printf("%T", variable)` kullanılarak değişkenlerin veri türleri ekrana yazdırılır.

---

## 🎯 Kodun Amacı (Purpose of the Code)
Bu kodun temel amacı, Golang'deki veri türlerini (data types), literalleri (literals), kaçış dizilerini (escape sequences) ve sayısal sistemleri (number systems) öğretmektir. Yeni başlayanlar için eğitici bir örnek olup, farklı veri türlerinin nasıl tanımlandığını ve bu türlerin konsola nasıl yazdırıldığını açıkça gösterir.

---

## 🛠️ Fonksiyonların Amacı (Purpose of the Functions)

### **`main()` Fonksiyonu (Main Function)**
Bu programda sadece `main()` fonksiyonu bulunmaktadır. Tüm işlemler ve veri türü örnekleri bu fonksiyonun içinde yer alır.

#### **Fonksiyon İçindeki Kod Blokları (Code Blocks in the Function)**

**1. Tamsayı (Integer) Literali**
```go
var intLiteral int = 42
fmt.Printf("Integer Literal: %d\n", intLiteral)
```
- **Amaç (Purpose):** Bir `int` türünde değişkeni tanımlamak ve değeri konsola yazdırmak.

**2. Ondalıklı Sayı (Floating Point) Literali**
```go
var floatLiteral float64 = 3.14159
fmt.Printf("Floating Point Literal: %f\n", floatLiteral)
```
- **Amaç (Purpose):** Bir `float64` türünde değişkeni tanımlamak ve değeri konsola yazdırmak.

**3. Karmaşık Sayılar (Complex Numbers) Literalleri**
```go
var complexLiteral1 complex128 = 3 + 4i
var complexLiteral2 complex128 = (2.2 + 1.1i)
fmt.Printf("Complex Literal 1: %f\n", complexLiteral1)
fmt.Printf("Complex Literal 2: %f\n", complexLiteral2)
```
- **Amaç (Purpose):** `complex128` veri türünde değişkenler tanımlamak ve değerlerini konsola yazdırmak.

**4. Boolean Literalleri (Boolean Literals)**
```go
var boolTrue bool = true
var boolFalse bool = false
fmt.Printf("Boolean Literal (true): %t\n", boolTrue)
fmt.Printf("Boolean Literal (false): %t\n", boolFalse)
```
- **Amaç (Purpose):** `bool` veri türünde `true` ve `false` değerlerini göstermek.

**5. Rune Literalleri (Rune Literals)**
```go
var runeLiteral1 rune = 'A'
var runeLiteral2 rune = '\n'
var runeLiteral3 rune = '\u4F60'
var runeLiteral4 rune = '\x41'
fmt.Printf("Rune Literal 1: %c\n", runeLiteral1)
fmt.Printf("Rune Literal 2: %q\n", runeLiteral2)
fmt.Printf("Rune Literal 3 (Unicode 你): %c\n", runeLiteral3)
fmt.Printf("Rune Literal 4 (Hexadecimal 'A'): %c\n", runeLiteral4)
```
- **Amaç (Purpose):** Unicode ve ASCII karakterlerinin nasıl gösterildiğini göstermek.

**6. Kaçış Dizileri (Escape Sequences)**
```go
fmt.Println("Escape Sequences:")
fmt.Println("Alert (Bell): \a")
fmt.Println("Backspace: ABC\bD")
fmt.Println("New Line: First Line\nSecond Line")
```
- **Amaç (Purpose):** Kaçış dizilerinin nasıl çalıştığını göstermek.
