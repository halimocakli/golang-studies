package main

import (
	"fmt"
	"math"
)

func main() {
	// Negatif bir sayının karekökünü hesaplama.
	// Gerçek sayı sisteminde tanımsız olduğundan IEEE 754'e uygun olarak NaN üretir.
	sqrtNeg := math.Sqrt(-1)

	fmt.Printf("math.Sqrt(-1) = %v\n", sqrtNeg) // Beklenen çıktı: NaN

	// Üretilen değerin NaN olup olmadığını kontrol etme (hata yönetimi açısından kritik)
	if math.IsNaN(sqrtNeg) {
		fmt.Println("sqrtNeg is NaN (Not-a-Number)!")
	}

	fmt.Print("\n")

	// Pozitif sonsuzluk (+Infinity) oluşturma;
	// IEEE 754 standardına göre sonsuzluk kavramı uygulanır.
	inf := math.Inf(1)
	// Pozitif sonsuzluktan pozitif sonsuzluk çıkarılması belirsizlik içerdiğinden NaN üretir.
	infMinusInf := inf - inf

	fmt.Printf("(+Infinity) - (+Infinity) = %v\n", infMinusInf) // Beklenen çıktı: NaN

	// İşlem sonucunun NaN olup olmadığını kontrol etme
	if math.IsNaN(infMinusInf) {
		fmt.Println("infMinusInf is NaN (Not-a-Number)!")
	}

	fmt.Print("\n")

	// Negatif bir sayının logaritmasını hesaplama.
	// Matematiksel olarak tanımsız olduğundan NaN sonucu verir.
	logNeg := math.Log(-10)

	fmt.Printf("math.Log(-10) = %v\n", logNeg) // Beklenen çıktı: NaN

	// Hesaplamadan elde edilen sonucun NaN olup olmadığını kontrol etme
	if math.IsNaN(logNeg) {
		fmt.Println("logNeg is NaN (Not-a-Number)!")
	}

	fmt.Print("\n")

	// Sıfırın logaritmasının hesaplanması.
	// Matematiksel tanım gereği negatif sonsuzluk (negative Infinity) sonucu üretir.
	logZero := math.Log(0)

	fmt.Printf("math.Log(0) = %v\n", logZero)

	// logZero değerinin sonsuzluk (Infinity) olup olmadığını kontrol etme.
	// 0 parametresi, her iki yönlü sonsuzluğu kapsar.
	if math.IsInf(logZero, 0) {
		fmt.Println("logZero is negative Infinity!")
	}
}
