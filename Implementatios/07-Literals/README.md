
# Golang Literals and Escape Sequences

## Amaç (Purpose)
Bu kod, Go programlama dilinde (Golang) kullanılan farklı türdeki *literal* (sabit) değerleri ve *escape sequence* (kaçış dizileri) örnekleriyle birlikte tanıtmak için yazılmıştır. Bu kod, *integer literals* (tam sayı sabitleri), *floating-point literals* (ondalık sayı sabitleri), *complex literals* (karmaşık sayı sabitleri), *boolean literals* (mantıksal sabitler), *rune literals* (karakter sabitleri) ve *string literals* (metin sabitleri) gibi farklı veri türlerini kapsamaktadır. 

Ayrıca, *escape sequence* (kaçış dizileri) ve *numeric literals* (sayısal sabitler) için kullanılan heksadesimal (hexadecimal), binary (ikili) ve oktal (octal) formatları da ele alınmıştır.

---

## Kodun Tamamının Açıklaması (Code Explanation)

### Integer Literals (Tam Sayı Sabitleri)
```go
var intLiteral int = 42
fmt.Printf("Integer Literal: %d
", intLiteral)
```
Bu kısımda bir *integer literal* (tam sayı sabiti) tanımlanıp ekrana yazdırılmıştır. `intLiteral` değişkeni, değeri `42` olan bir tam sayıdır.

---

### Floating Point Literals (Ondalık Sabitler)
```go
var floatLiteral float64 = 3.14159
fmt.Printf("Floating Point Literal: %f
", floatLiteral)
```
Bu kısımda bir *floating point literal* (ondalık sabit) tanımlanıp ekrana yazdırılmıştır. Değer `3.14159` olarak atanır.

---

### Complex Literals (Karmaşık Sayı Sabitleri)
```go
var complexLiteral1 complex128 = 3 + 4i
var complexLiteral2 complex128 = (2.2 + 1.1i)
fmt.Printf("Complex Literal 1: %f
", complexLiteral1)
fmt.Printf("Complex Literal 2: %f
", complexLiteral2)
```
Bu kısımda *complex literals* (karmaşık sayı sabitleri) tanımlanır. Karmaşık sayı, bir reel (gerçek) ve bir de imajiner (hayali) bölüme sahiptir.

---

### Boolean Literals (Mantıksal Sabitler)
```go
var boolTrue bool = true
var boolFalse bool = false
fmt.Printf("Boolean Literal (true): %t
", boolTrue)
fmt.Printf("Boolean Literal (false): %t
", boolFalse)
```
Bu kısımda iki *boolean literals* (mantıksal sabit) kullanılır. `true` ve `false` olmak üzere iki değer bulunur.

---

### Rune Literals (Karakter Sabitleri)
```go
var runeLiteral1 rune = 'A'
var runeLiteral2 rune = '
'
var runeLiteral3 rune = '你'
var runeLiteral4 rune = 'A'
fmt.Printf("Rune Literal 1: %c
", runeLiteral1)
fmt.Printf("Rune Literal 2: %q
", runeLiteral2)
fmt.Printf("Rune Literal 3 (Unicode 你): %c
", runeLiteral3)
fmt.Printf("Rune Literal 4 (Hexadecimal 'A'): %c
", runeLiteral4)
```
*Rune literals* (karakter sabitleri) Go dilinde Unicode karakterlerini ifade eder. Bu kısımda farklı yöntemlerle karakterler tanımlanmıştır.

---

### Escape Sequences (Kaçış Dizileri)
```go
fmt.Println("Escape Sequences:")
fmt.Println("Alert (Bell): ")
fmt.Println("Backspace: ABCD")
fmt.Println("Form Feed: AB")
fmt.Println("New Line: First Line
Second Line")
fmt.Println("Carriage Return: ABC
123")
fmt.Println("Horizontal Tab: A	B")
fmt.Println("Vertical Tab: AB")
```
Bu kısımda farklı *escape sequences* (kaçış dizileri) gösterilmektedir. Bunlar metin çıktısını çeşitlendirmek için kullanılır.

---

## Escape Sequence Tablosu (Escape Sequence Table)

| Escape Sequence | Açıklama (Explanation)                        |
|-----------------|----------------------------------------------|
| `\a`           | Alarm (Bell) - Sesli uyarı verir             |
| `\b`           | Backspace - Bir karakter geri gider          |
| `\f`           | Form feed - Yeni bir sayfa başlatır          |
| `\n`           | New line - Yeni satıra geçer                 |
| `\r`           | Carriage return - Satır başına döner         |
| `\t`           | Horizontal tab - Yatay sekme yapar           |
| `\v`           | Vertical tab - Dikey sekme yapar             |
| `\\`           | Backslash - Ters bölüm işareti ekler         |
| `'`            | Single quote - Tek tırnak işareti ekler      |
| `\"`           | Double quote - Çift tırnak işareti ekler     |

---

## Programdan Alınan Çıktı (Program Output)
```
Integer Literal: 42
Floating Point Literal: 3.141590
Complex Literal 1: 3.000000+4.000000i
Complex Literal 2: 2.200000+1.100000i
Boolean Literal (true): true
Boolean Literal (false): false
Rune Literal 1: A
Rune Literal 2: '
'
Rune Literal 3 (Unicode 你): 你
Rune Literal 4 (Hexadecimal 'A'): A
Escape Sequences:
Alert (Bell): 
Backspace: ABCD
Form Feed: AB
New Line: First Line
Second Line
Carriage Return: 123
Horizontal Tab: A	B
Vertical Tab: AB
Backslash: \
Single Quote can be directly used in string literal: '\'
Double Quote: "
```
---
