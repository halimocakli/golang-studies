
# README

## Programın Amaçları ve Amacı (Purpose and Objective of the Program)
Bu program, Golang dilinde çeşitli "literal" (İngilizce: literal, Türkçe: sabit değer) türlerini göstermek ve bunları terminal üzerinden dışari yazdırmak amacıyla oluşturulmuştur. Program, aşağıdaki veri türleriyle çalışır:

- **Tam sayı (Integer) literal**
- **Ondalıklı (Floating Point) literal**
- **Karmaşık (Complex) sayı literal**
- **Mantıksal (Boolean) literal**
- **Karakter (Rune) literal**
- **Dize (String) literal**

Bu veri türleri, kod içinde birçok durumda kullanılan temel veri türleridir. Bu program, farklı şekillerde tanımlanan bu değerlerin, formatlı bir şekilde terminal üzerinde gösterilmesine olanak tanır.

---

## Programın Çalışma Prensibi (How the Program Works)
1. Her bir "literal" türü (integer, float, complex, boolean, rune, string) tanımlanır.
2. `fmt.Printf()` fonksiyonu kullanılarak terminal üzerine çıktı verilir.
3. "Escape sequence" (Kaçış dizisi) karakterleri kullanılarak farklı metin formatlamaları ve dizin yolları yazdırılır.

---

## Fonksiyonların Amaçları (Purpose of the Functions)
Bu programda yalnızca bir adet fonksiyon vardır:

### `func main()`
Bu fonksiyon, programın başlangıç noktısıdır. Şu işlevleri yerine getirir:

| **Kısım (Section)** | **Amaç (Purpose)** | **Açıklama (Explanation)** |
|---------------------|---------------------|-----------------------------------|
| **Integer Literals** | Tam sayı değerlerini göstermek | `int` türünde bir değer tanımlar ve yazdırır. |
| **Floating Point Literals** | Ondalıklı sayıları göstermek | `float64` türü bir değer tanımlar ve yazdırır. |
| **Complex Literals** | Karmaşık sayıları göstermek | `complex128` türünde iki karmaşık sayıyı tanımlar ve yazdırır. |
| **Boolean Literals** | Mantıksal değerleri göstermek | `bool` türünde iki mantıksal değer tanımlar ve yazdırır. |
| **Rune Literals** | Karakterleri göstermek | Unicode ve ASCII karakterleri tanımlar ve yazdırır. |
| **String Literals** | Dizeleri (string) göstermek | Kaçış dizileri ile ham ve normal dize yazımını tanımlar ve yazdırır. |
| **Escape Sequences** | Kaçış dizilerini göstermek | Yeni satır, tab ve geri alma gibi kaçış dizilerini gösterir. |
| **Binary, Hex, Octal Literals** | Sayı sistemlerini göstermek | İkili (binary), sekizli (octal) ve onaltılık (hex) sayı sistemlerini gösterir. |
| **Scientific Notation** | Bilimsel gösterimi göstermek | Üsülü gösterimle ondalık sayıları gösterir. |
| **File Path Printing** | Dosya yolunu yazdırmak | Dosya yollarını yazdırır ve kaçış dizisi hatalarını gösterir. |

---

## Koddaki Tüm Satırların Açıklamaları (Explanation of All Code Lines)
Her bir kod satırının açıklaması ve terminalde oluşan çıktılar aşağıda verilmiştir.

```go
// Integer Literal
var intLiteral int = 42
fmt.Printf("Integer Literal: %d\n", intLiteral) // Çıktı: Integer Literal: 42
```
Açıklama: Bu satır bir tamsayı değerini tanımlar ve formatlı çıktı ile yazdırır.

```go
// Floating Point Literal
var floatLiteral float64 = 3.14159
fmt.Printf("Floating Point Literal: %f\n", floatLiteral) // Çıktı: Floating Point Literal: 3.141590
```
Açıklama: Bu satır bir ondalıklı sayı değerini tanımlar ve yazdırır.

```go
// Boolean Literals
var isTrue bool = true
var isFalse bool = false
fmt.Printf("Boolean True: %t\n", isTrue) // Çıktı: Boolean True: true
fmt.Printf("Boolean False: %t\n", isFalse) // Çıktı: Boolean False: false
```
Açıklama: Mantıksal (boolean) değerleri tanımlar ve yazdırır.

```go
// Rune Literals
var runeLiteral rune = 'A'
fmt.Printf("Rune Literal: %c\n", runeLiteral) // Çıktı: Rune Literal: A
```
Açıklama: Bir karakteri (rune) tanımlar ve ASCII karşılığını yazdırır.

```go
// String Literals
var normalString string = "Hello, Go!"
var rawString string = `This is a raw string literal.`
fmt.Printf("Normal String: %s\n", normalString) // Çıktı: Normal String: Hello, Go!
fmt.Printf("Raw String: %s\n", rawString) // Çıktı: Raw String: This is a raw string literal.
```
Açıklama: Ham ve normal dize türlerini tanımlar ve yazdırır.

---
