# Golang Operatörleri İçin Notlar

Bu not defteri, Golang operatörleri ile ilgili temel kavramları, operatörlerin sınıflandırılmasını, öncelik ve ilişki kurallarını ve yazılan kodla ilgili tüm önerileri içerir. Notlar, hem öğretici hem de pratik bilgileri kapsar.

## 1. **Operatörlerin Tanımı ve Sınıflandırılması**

- **Operatör (Operator)**: Bir işleme yol açan ve işlem sonucunda bir değer üreten sembollerdir.
- **Operand**: Bu operatörlerle işlem yapan ifadelerdir.
- **Sınıflandırma:**
  - **Aritmetik Operatörler (Arithmetic Operators)**: `+`, `-`, `*`, `/`, `%`
  - **Karşılaştırma Operatörleri (Comparison Operators)**: `==`, `!=`, `<`, `>`, `<=`, `>=`
  - **Mantıksal Operatörler (Logical Operators)**: `&&`, `||`, `!`
  - **Bitsel Operatörler (Bitwise Operators)**: `&`, `|`, `^`, `<<`, `>>`
  - **Özel Amaçlı Operatörler (Special Purpose Operators)**: `&`, `*`, `=`, `:=`

## 2. **Operatör Önceliği ve İşlem Sırası**

- Operatörler, öncelik düzeylerine göre işlenmektedir.
- **Aritmetik operatörlerin Öncelik Sıralaması**: `*`, `/`, `%` > `+`, `-`
- **Sağdan sola (right associative)** veya **soldan sağa (left associative)** olabilir.
  - Aritmetik operatörler genellikle soldan sağa (left associative) çalışır.
  - Atama operatörleri (`=`) sağdan sola (right associative) çalışır.

## 3. **Golang Operatörleri Hakkında Detaylı Bilgiler**

- **Modulo Operatörü (%)**:
  - Yalnızca `int` türleri ile kullanılabilir.
  - `float64` veya `complex` türleriyle kullanılamaz, aksi halde derleme hatası alınır.
  - Matematiksel mod işleminden farklı olarak, Go'da negatif sayıların mod işlemi, operandın işaretiyle ilişkilidir.
- **Aritmetik Operatörler** (`+`, `-`, `*`, `/`):

  - Hem `int` hem de `float64` değerleriyle kullanılabilir.
  - `complex` değerleriyle de kullanılabilir.
  - `/` operatörü bölme işleminde, eğer bölen 0 ise hata oluşturur.

- **Mantıksal Operatörler** (`&&`, `||`, `!`):

  - Mantıksal işlemler için kullanılır ve yalnızca `bool` türleriyle çalışır.
  - Kısa devre mantığı (short-circuiting) kullanılır.

- **Karşılaştırma Operatörleri** (`==`, `!=`, `<`, `>`, `<=`, `>=`):

  - Bütün sayısal türlerle ve `string` ile kullanılabilir.

- **Bitsel Operatörler** (`&`, `|`, `^`, `<<`, `>>`):

  - Yalnızca `int` türleriyle kullanılabilir.

- **Atama Operatörleri** (`=`, `+=`, `-=`, `*=`, `/=`, `%=`):
  - Değişkenlere değer atamak veya mevcut değeri güncellemek için kullanılır.


## 4. **Sonuç**

Bu not defteri, Golang operatörleri hakkında önemli bilgiler ve kod için iyileştirme önerileri içermektedir. Tüm bu öneriler, kodun okunabilirliğini, modülerleştirilmesini ve hata kontrolünü geliştirir.
