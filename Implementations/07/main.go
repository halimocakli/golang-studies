package main

import (
	"fmt" // fmt paketi: Formatlı giriş/çıkış (formatted I/O) işlemleri için kullanılan temel pakettir.
)

func main() {
	// Tamsayılar (Integers):
	// Aşağıdaki satırlarda, farklı format belirteçleri kullanılarak tamsayı çıktıları elde ediliyor.
	// Bu, sayısal verilerin farklı sistemlerde (decimal, binary, octal, hexadecimal) nasıl gösterileceğini anlamak açısından önemlidir.
	fmt.Printf("%%d: %d\n", 42)    // %d: Tamsayıyı ondalık (decimal) sistemde gösterir.
	fmt.Printf("%%b: %b\n", 42)    // %b: Tamsayıyı ikilik (binary) sistemde gösterir.
	fmt.Printf("%%b: %016b\n", 42) // %016b: 16 karakter genişliğinde, eksik basamakları sıfır (0) ile doldurulmuş ikilik gösterim sağlar.
	fmt.Printf("%%o: %o\n", 42)    // %o: Tamsayıyı sekizlik (octal) sistemde gösterir.
	fmt.Printf("%%x: %x\n", 42)    // %x: Tamsayıyı onaltılık (hexadecimal) sistemde, küçük harflerle gösterir.
	fmt.Printf("%%x: %08x\n", 42)  // %08x: 8 haneli, eksik basamakları sıfır (0) ile tamamlanmış onaltılık gösterim.
	fmt.Printf("%%X: %X\n", 42)    // %X: Tamsayıyı onaltılık (hexadecimal) sistemde, büyük harflerle gösterir.
	fmt.Printf("%%c: %c\n", 65)    // %c: Tamsayıyı karşılık geldiği Unicode karakter (character) olarak gösterir (65 -> 'A').
	fmt.Printf("%%q: %q\n", 65)    // %q: Tamsayıyı tek tırnak içine alınmış karakter çıktısı şeklinde sunar.

	fmt.Print("\n")

	// Gerçek Sayılar (Floating-Point Numbers):
	// Bu bölümde, gerçek sayılar için farklı biçimlendirme seçenekleri kullanılarak, sabit nokta, bilimsel gösterim ve en uygun gösterim örnekleri sunulmaktadır.
	fmt.Printf("%%f: %f\n", 3.14159)   // %f: Sabit nokta (fixed-point) gösterimi sağlar; ondalıklı sayının düz metin olarak ifade edilmesidir.
	fmt.Printf("%%e: %e\n", 1234.5678) // %e: Bilimsel (exponential) gösterim; sayı çok büyük veya çok küçük olduğunda okunabilirliği artırır.
	fmt.Printf("%%g: %g\n", 1234.5678) // %g: Değerin en uygun biçimde gösterimini sağlar; otomatik olarak %f veya %e formatını seçer.

	fmt.Print("\n")

	// Dizgeler (Strings):
	// Dizge (string) çıktıları için çeşitli format belirteçleri kullanılarak farklı gösterim yöntemleri örneklenmiştir.
	fmt.Printf("%%s: %s\n", "hello") // %s: Dizgenin ham (raw) halini, yani orijinal metni olduğu gibi gösterir.
	fmt.Printf("%%q: %q\n", "hello") // %q: Dizgeyi çift tırnak içine alarak çıktı verir; bu, dizgenin sınırlarının belirginleşmesini sağlar.
	fmt.Printf("%%x: %x\n", "hello") // %x: Dizgedeki her karakterin hexadecimal (onaltılık) değerini gösterir; karakter kodlarının (ASCII/Unicode) analizi için faydalıdır.

	fmt.Print("\n")

	// Pointer (İşaretçi):
	// Değişkenin bellekteki adresini gösterir. Bu, bellek yönetimi (memory management) ve hata ayıklama (debugging) işlemleri için kritik öneme sahiptir.
	x := 42
	fmt.Printf("%%p: %p\n", &x) // %p: Değişken 'x'in bellekteki adresini hexadecimal biçimde gösterir.

	fmt.Print("\n")

	// Genel Yer Tutucular (General Placeholders):
	// Farklı veri tiplerinin varsayılan gösterimlerini ve veri tip bilgilerini görüntülemek için kullanılır.
	fmt.Printf("%%v: %v\n", []int{1, 2, 3})                // %v: Slice'in (dizi) varsayılan gösterimini sağlar.
	fmt.Printf("%%+v: %+v\n", struct{ Name string }{"Go"}) // %+v: Struct (yapı) verisini, alan adlarıyla birlikte detaylı gösterir.
	fmt.Printf("%%T: %T\n", x)                             // %T: Değişken 'x'in veri tipini (type) belirtir; çalışma zamanında veri tiplerinin kontrolü için yararlıdır.
}
