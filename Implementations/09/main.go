package main

import (
	"fmt" // fmt paketi, formatlı giriş/çıkış işlemleri (I/O operations [Girdi/Çıktı İşlemleri]) için kullanılır.
)

func main() {
	// ------------------------------
	// Integer Literal (Tamsayı Sabiti)
	// ------------------------------
	// Aşağıdaki ifade, 42 sayısını int türünde sabit bir değer olarak tanımlar.
	var intLiteral int = 42
	// %d format belirteci, integer değerleri ekrana basmak için kullanılır.
	fmt.Printf("Integer Literal: %d\n", intLiteral)

	// ------------------------------
	// Floating Point Literal (Ondalık Sabiti)
	// ------------------------------
	// Bu ifade, 3.14159 değerini float64 türünde ondalık sayı olarak tanımlar.
	var floatLiteral float64 = 3.14159
	// %f format belirteci, ondalık sayıların çıktıda gösterilmesinde kullanılır.
	fmt.Printf("Floating Point Literal: %f\n", floatLiteral)

	// ------------------------------
	// Complex Literal (Karmaşık Sayı Sabiti)
	// ------------------------------
	// Karmaşık sayılar, gerçek (real) ve sanal (imaginary) kısımlar içerir.
	// Aşağıda iki farklı karmaşık sayı örneği tanımlanmıştır.
	var complexLiteral1 complex128 = 3 + 4i
	var complexLiteral2 complex128 = (2.2 + 1.1i)
	// %f format belirteci karmaşık sayılarda da gerçek kısmı göstermek için kullanılır.
	fmt.Printf("Complex Literal 1: %f\n", complexLiteral1)
	fmt.Printf("Complex Literal 2: %f\n", complexLiteral2)

	// ------------------------------
	// Boolean Literal (Mantıksal Sabit)
	// ------------------------------
	// Mantıksal değerler true veya false olarak tanımlanır.
	var boolTrue bool = true
	var boolFalse bool = false
	// %t format belirteci, boolean değerlerin çıktıda gösterilmesini sağlar.
	fmt.Printf("Boolean Literal (true): %t\n", boolTrue)
	fmt.Printf("Boolean Literal (false): %t\n", boolFalse)

	// ------------------------------
	// Rune Literal (Karakter Sabiti)
	// ------------------------------
	// Rune, Go dilinde tek karakteri temsil eden veri tipidir (int32 tabanlıdır).
	// Farklı yöntemlerle rune tanımlama örnekleri:
	var runeLiteral1 rune = 'A'      // Doğrudan karakter ifadesi
	var runeLiteral2 rune = '\n'     // Yeni satır karakteri (escape sequence: newline)
	var runeLiteral3 rune = '\u4F60' // Unicode gösterimiyle tanımlanmış karakter: 你
	var runeLiteral4 rune = '\x41'   // Hexadecimal gösterimde tanımlanmış karakter: A
	// %c format belirteci, rune değerini karakter olarak çıktı verir.
	fmt.Printf("Rune Literal 1: %c\n", runeLiteral1)
	// %q format belirteci, karakteri tek tırnak içinde gösterir.
	fmt.Printf("Rune Literal 2: %q\n", runeLiteral2)
	fmt.Printf("Rune Literal 3 (Unicode 你): %c\n", runeLiteral3)
	fmt.Printf("Rune Literal 4 (Hexadecimal 'A'): %c\n", runeLiteral4)

	// ------------------------------
	// Kaçış Dizileri (Escape Sequences) Örnekleri
	// ------------------------------
	// Aşağıdaki örneklerde, farklı kaçış dizilerinin nasıl kullanıldığı gösterilmiştir.
	fmt.Println("Escape Sequences:")
	fmt.Println("Alert (Bell): \a")                                           // \a: Alarm sesi (alert), terminalde sistem tarafından tanımlı bir uyarı sesi çalınabilir.
	fmt.Println("Backspace: ABC\bD")                                          // \b: Backspace (bir karakter geri alma), yazı düzenlemede kullanılır.
	fmt.Println("Form Feed: A\fB")                                            // \f: Form feed (sayfa başı), çıktıda yeni sayfa oluşturma amaçlı kullanılabilir.
	fmt.Println("New Line: First Line\nSecond Line")                          // \n: Newline (yeni satır), metin düzenini sağlamak için kullanılır.
	fmt.Println("Carriage Return: ABC\r123")                                  // \r: Carriage return (satır başına dönüş), mevcut satırı yeniden yazmak için kullanılabilir.
	fmt.Println("Horizontal Tab: A\tB")                                       // \t: Horizontal tab (yatay sekme), sütun hizalama için idealdir.
	fmt.Println("Vertical Tab: A\vB")                                         // \v: Vertical tab (dikey sekme), metin içindeki satırlar arasında dikey boşluk oluşturur.
	fmt.Println("Backslash: \\")                                              // \\: Backslash (ters eğik çizgi), metin içerisinde ters eğik çizgiyi göstermek için kullanılır.
	fmt.Println("Single Quote can be directly used in  string literal: '\\'") // Tek tırnak, string literal içerisinde doğrudan kullanılabilir.
	fmt.Println("Double Quote: \"")                                           // \" : Çift tırnak, string içerisinde çift tırnak karakterini göstermek için kaçış dizisi kullanılır.

	// ------------------------------
	// String Literal (Metin Sabiti)
	// ------------------------------
	// İki farklı string literal türü örneği:
	// 1. Çift tırnaklı string literal: Escape sequence'ler aktif çalışır.
	// 2. Raw string literal (backtick ile tanımlanır): Escape sequence'ler etkisizdir.
	var stringLiteral1 string = "Hello, Go!"
	var stringLiteral2 string = `This is a raw string literal. \n No escape sequences work here.`
	// %s format belirteci, string değerleri çıktıya basmak için kullanılır.
	fmt.Printf("String Literal 1: %s\n", stringLiteral1)
	fmt.Printf("String Literal 2: %s\n", stringLiteral2)

	// ------------------------------
	// Farklı Biçimlerde Rune Tanımlamaları
	// ------------------------------
	// Aşağıdaki örnekler, farklı numeral sistemlerinden (hexadecimal, octal, Unicode) rune tanımlamalarını göstermektedir.
	var hexRune rune = '\x41'       // Hexadecimal formatta tanımlanmış rune; 0x41, 'A' karakterine eşittir.
	var octalRune rune = '\101'     // Octal formatta tanımlanmış rune; 101 (octal) değeri yine 'A' karakterini temsil eder.
	var unicodeRune rune = '\u0041' // Unicode formatında tanımlanmış rune; \u0041, 'A' karakterini temsil eder.
	fmt.Printf("Hexadecimal Rune: %c\n", hexRune)
	fmt.Printf("Octal Rune: %c\n", octalRune)
	fmt.Printf("Unicode Rune: %c\n", unicodeRune)

	// ------------------------------
	// Farklı Sayı Sistemlerinde Literal İfadeler
	// ------------------------------
	// Aşağıdaki ifadeler, hexadecimal, binary ve octal sayı sistemlerinden literal ifadeleri örneklendirir.
	a := 0x1FC0             // Hexadecimal literal: 0x1FC0 ifadesi, ondalık olarak 8128 değerine karşılık gelir.
	b := 0b0001111111000000 // Binary literal: İkili sayı sistemi ile tanımlanmıştır.
	c := 0o17700            // Octal literal: Sekizlik sayı sistemi kullanılarak tanımlanmıştır. '0o' öneki ile belirtilir.
	d := 0o17700            // Octal literal: Büyük 'O' harfi ile de kullanılabilir, aynı değeri ifade eder.
	fmt.Printf("Hexadecimal a: %d\n", a)
	fmt.Printf("Binary b: %d\n", b)
	fmt.Printf("Octal c: %d\n", c)
	fmt.Printf("Octal d: %d\n", d)

	// ------------------------------
	// Alt Çizgi Kullanımıyla Okunabilirlik Artırılmış Literaller
	// ------------------------------
	// Alt çizgi, uzun sayılarda rakam gruplarını ayırmak için kullanılır. Bu, hata yapma riskini azaltır.
	a = 0x1F_C0               // Hexadecimal, alt çizgi kullanılarak okunabilirliği artırılmıştır.
	b = 0b0001_1111_1100_0000 // Binary, alt çizgi ile gruplandırılmıştır.
	c = 0o177_00              // Octal, alt çizgi kullanılarak daha okunabilir hale getirilmiştir.
	fmt.Printf("Hexadecimal a: %d\n", a)
	fmt.Printf("Binary b: %d\n", b)
	fmt.Printf("Octal c: %d\n", c)

	// ------------------------------
	// Bilimsel Gösterim (Scientific Notation)
	// ------------------------------
	// Avogadro sayısı gibi çok büyük değerlerin ifade edilmesi için üstel gösterim kullanılır.
	avogadro := 6.02e23
	// %e format belirteci, üstel (exponential) gösterimi kullanarak sayıyı formatlar.
	fmt.Printf("Avogadro: %e\n", avogadro)
	// %f format belirteci, sayıyı ondalık biçimde gösterir.
	fmt.Printf("Avogadro (normal format): %f\n", avogadro)
	// %T format belirteci, değişkenin veri tipini çıktı olarak verir.
	fmt.Printf("Type: %T\n", avogadro)

	// ------------------------------
	// Hexadecimal Floating Point Literal
	// ------------------------------
	// Hexadecimal tabanlı ondalık sayı literal örneği, farklı gösterim biçimlerinin kullanımını sağlar.
	hexFloat := 0xABp3
	fmt.Printf("Hexadecimal Floating Point: %f\n", hexFloat)
	fmt.Printf("Type: %T\n", hexFloat)

	// ------------------------------
	// Doğru Kaçış Dizisi Kullanımları (Escape Sequences)
	// ------------------------------
	// Dosya yolları gibi metinlerde ters eğik çizgi karakteri (backslash) kaçış dizileriyle doğru biçimde yazılmalıdır.
	// Aşağıdaki örnek, Windows dosya yolu yazımında ters eğik çizgilerin (backslash) nasıl kullanılacağını gösterir.
	fmt.Print("C:\\test\\numbers.txt\n")
	// Raw string literal kullanıldığında, escape dizileri etkisiz hale gelir.
	fmt.Print(`C:\test\numbers.txt`)

	// ------------------------------
	// Yanlış Kaçış Dizisi Kullanımı
	// ------------------------------
	// Aşağıdaki satır, yanlış kaçış dizisi kullanımı sebebiyle derleme hatası vereceğinden yorum satırı içerisine alınmıştır.
	// fmt.Print("C:\jest\numbers.txt")
}
