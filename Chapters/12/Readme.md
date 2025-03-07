## 1. Yapılar (Structures – Yapılar) Nedir?

- **Tanım:**

  - Go dilinde, programcıların yerleşik (built-in – yerleşik) türler dışında kendi özel (user defined – kullanıcı tanımlı) türlerini tanımlamalarına olanak sağlayan yapılardır.
  - Bir grup verinin (field – alan) bir araya gelmesiyle oluşturulan bileşik (compound – bileşik) türlerdir.
  - Hafıza yönetimi açısından, yapılar sabit bir bellek düzeni sunar; bu da onları performans odaklı uygulamalarda tercih edilir hale getirir.

- **Kullanım Alanları:**

  - **Veri Modelleme:** Birbiriyle ilişkili verilerin mantıksal olarak gruplanması
  - **Veri Yapıları (Data Structures – Veri Yapıları):** Özel listeler, ağaçlar, hash tablolar gibi karmaşık yapılar oluşturulabilir
  - **Nesne Yönelimli Yaklaşım (Object-Oriented – Nesne Yönelimli):** Go, klasik sınıf (class – sınıf) yapısı sunmasa da, yapılar ve metotlar sayesinde nesne benzeri davranışlar elde edilmesini sağlar.

- **Genel Sözdizimi:**

  ```go
  type Isim struct {
      field1 Tip1
      field2 Tip2
      // ...
  }
  ```

  - **Not:** Yapı tanımlanırken alan isimlerini ve türlerini belirtmek zorunludur. Alan isimleri büyük harfle başlarsa (exported – dışa aktarılabilir) diğer paketlerden erişim sağlanabilir.

- **Ek Örnek – Anahtar-Değer İle Başlatma:**

  ```go
  package main

  import "fmt"

  type Person struct {
      Name string
      Age  int
  }

  func main() {
      // Alan isimlerine göre başlatma (key-value initialization)
      p := Person{
          Name: "Ahmet",
          Age:  30,
      }
      fmt.Println("Person:", p)
  }
  ```

---

## 2. Yapı Elemanları (Structure Members – Yapı Elemanları) ve Erişim

- **Yapı Elemanları:**

  - Yapı içinde tanımlanan değişkenlere, "field" veya "member" denir.
  - Her alan, yapının içindeki bir bileşeni temsil eder ve farklı türlerde olabilir.

- **Erişim Yöntemi:**
  - Yapı örneğine nokta operatörü (`.`) ile erişilir.
  - Eğer yapı bir pointer (işaretçi – pointer) ise, Go otomatik olarak dereference işlemi yapar.
  - Örneğin:
    ```go
    p := Point{100, 200}
    fmt.Println(p.x)  // x alanına erişim
    ```
- **Ek Açıklamalar:**
  - **Varsayılan Değerler:** Bir yapı örneği oluşturulduğunda, alanlar kendi türlerinin sıfır değerlerini alır (örneğin, int için 0, float için 0.0, string için boş).
  - **İç İçe Yapılar (Nested Structures – İç İçe Yapılar):** Yapı alanları başka yapılar da olabilir. Bu sayede hiyerarşik veri modelleri oluşturulabilir.
- **Ek Örnek – İç İçe Yapı Kullanımı:**

  ```go
  package main

  import "fmt"

  type Address struct {
      City    string
      ZipCode int
  }

  type Person struct {
      Name    string
      Age     int
      Address Address
  }

  func main() {
      p := Person{
          Name: "Ayşe",
          Age:  25,
          Address: Address{
              City:    "Istanbul",
              ZipCode: 34000,
          },
      }
      // İç içe yapıda alanlara erişim: p.Address.City
      fmt.Printf("Kişi: %s, Şehir: %s\n", p.Name, p.Address.City)
  }
  ```

- **Tablo: Yapı Elemanlarına Erişim**

  | **Erişim Durumu**      | **Sözdizimi**   | **Açıklama**                                                       |
  | ---------------------- | --------------- | ------------------------------------------------------------------ |
  | Yapı örneği üzerinden  | `nesne.field`   | Yapının alanına direkt erişim.                                     |
  | Yapı pointer üzerinden | `pointer.field` | Go otomatik dereference yapar, ekstra `(*pointer).field` gerekmez. |

---

## 3. Yapıların Fonksiyonları (Methods – Metotlar) ve Pointer ile Değer Gönderimi

### 3.1. Metot Tanımlama

- **Metot Nedir?**

  - Herhangi bir tipe (yapı gibi) ait fonksiyonlardır.
  - Metotlar, belirli bir alıcı (receiver – receiver) üzerinden çağrılır; bu alıcı, metodu hangi örnek üzerinde çalıştıracağımızı belirler.

- **Metot Sözdizimi:**

  ```go
  func (p Tip) MethodName(parametreler) (dönüşTipleri) {
      // Metot gövdesi
  }
  ```

  - Eğer alıcı bir kopya olarak gönderilirse (value receiver), orijinal yapı üzerinde değişiklik yapılmaz.
  - Eğer alıcı pointer olarak gönderilirse (pointer receiver), metot orijinal yapı üzerinde değişiklik yapabilir.

- **Örnek Kod – Değer Alıcısı Kullanımı:**

  ```go
  package main

  import "fmt"

  type Point struct {
      x, y int
  }

  // Value receiver kullanılarak sadece okuma işlemi yapılır.
  func (p Point) Print() {
      fmt.Printf("x = %d, y = %d\n", p.x, p.y)
  }

  func main() {
      p := Point{100, 200}
      p.Print()  // x = 100, y = 200
  }
  ```

### 3.2. Pointer Receiver (İşaretçi Alıcısı) Kullanımı

- **Açıklama:**

  - Pointer receiver kullanıldığında, metot yapı örneğinin orijinal belleği üzerinde çalışır.
  - Değişiklikler orijinal yapıya yansır. Bu, özellikle büyük veri yapılarında kopyalama maliyetini önlemek açısından önemlidir.

- **Örnek Kod – Değiştirme İşlemi:**

  ```go
  package main

  import "fmt"

  type Point struct {
      x, y int
  }

  // Pointer receiver kullanılarak orijinal yapı üzerinde değişiklik yapılır.
  func (p *Point) Set(x, y int) {
      p.x = x
      p.y = y
  }

  func (p *Point) Print() {
      fmt.Printf("x = %d, y = %d\n", p.x, p.y)
  }

  func main() {
      p := &Point{100, 100}
      p.Print()   // x = 100, y = 100

      p.Set(200, 300)
      p.Print()   // x = 200, y = 300
  }
  ```

- **Ek Detay – Değer Receiver ile Karşılaştırma:**

  - **Value Receiver:** Yapı kopyası üzerinde işlem yapar. Örneğin, bir `Set` metodu value receiver olarak yazılırsa, orijinal yapı değişmez.
  - **Pointer Receiver:** Orijinal yapı üzerinde işlem yapar. Bu sebeple, verinin güncellenmesi gereken durumlarda tercih edilir.

- **Tablo: Value Receiver vs. Pointer Receiver**

  | **Özellik**           | **Value Receiver**                                                   | **Pointer Receiver**                                          |
  | --------------------- | -------------------------------------------------------------------- | ------------------------------------------------------------- |
  | **Bellek Kopyalama**  | Yapının kopyası oluşturulur (memberwise copy – eleman bazlı kopya)   | Bellek kopyası oluşturulmaz; orijinal veri üzerinde çalışılır |
  | **Değişiklik Etkisi** | Değişiklikler yalnızca kopya üzerinde etkilidir                      | Orijinal yapı üzerinde doğrudan değişiklik yapılır            |
  | **Performans**        | Küçük yapılar için fark etmez, büyük yapılar için ek maliyet yaratır | Genellikle daha performanslı; kopyalama işlemi yoktur         |

---

## 4. Yapıların Fonksiyonlara Geçirilmesi (Passing Structures to Functions)

- **Memberwise Copy (Üye Bazlı Kopyalama):**

  - Yapı nesnesi fonksiyona değer olarak gönderildiğinde, her alanı ayrı ayrı kopyalanır.
  - Bu işlem küçük yapılar için genellikle sorun yaratmaz; ancak büyük yapılar veya sık kullanılan durumlarda performans etkisi göz önünde bulundurulmalıdır.

- **Pointer ile Geçirme:**

  - Yapının adresi (pointer) gönderildiğinde, kopyalama yapılmadan orijinal veri üzerinde işlem gerçekleştirilir.
  - Bu, özellikle yapının boyutu büyük olduğunda ve/veya değişiklik yapılması gerektiğinde tercih edilir.

- **Örnek Kod – Değer Gönderimi:**

  ```go
  package main

  import "fmt"

  type Point struct {
      x, y int
  }

  func printPoint(p Point) {
      fmt.Println("printPoint (değerle):", p)
  }

  func main() {
      p := Point{100, 100}
      printPoint(p) // p'nin kopyası gönderilir.
  }
  ```

- **Örnek Kod – Pointer Gönderimi:**

  ```go
  package main

  import "fmt"

  type Point struct {
      x, y int
  }

  func printPoint(p *Point) {
      fmt.Println("printPoint (pointer ile):", *p)
  }

  func main() {
      p := Point{100, 100}
      printPoint(&p) // p'nin adresi gönderilir.
  }
  ```

- **Ek Açıklama:**
  - Fonksiyon çağrılarında, yapının boyutu küçükse değer ile gönderim gayet kabul edilebilir.
  - Ancak, büyük veri yapılarında kopyalama maliyetini azaltmak ve fonksiyon içinde yapılan değişikliklerin orijinale yansımasını sağlamak için pointer kullanımı tercih edilmelidir.

---

## 5. Type Alias (Tip Takma Adı – Type Alias) ve Metot Eklenmesi

- **Tanım ve Kullanım:**

  - `type` anahtar sözcüğü ile temel türlere yeni bir isim verilebilir.
  - Bu yeni isim, orijinal türün tüm özelliklerini taşır ve bu tipe özgü metotlar tanımlanabilir.
  - Özellikle kodun okunabilirliğini artırmak ve belirli anlamlar yüklemek için kullanılır.

- **Örnek Kod – Basit Tip Alias:**

  ```go
  package main

  import "fmt"

  type CardValue int

  func main() {
      var cv CardValue = 10
      fmt.Println("CardValue:", cv)
  }
  ```

- **Örnek Kod – Tip Alias’a Metot Eklemek:**

  ```go
  package main

  import "fmt"

  type IntValue int

  // Metot ekleniyor (IsEven: Çift mi?)
  func (a IntValue) IsEven() bool {
      return a % 2 == 0
  }

  func main() {
      var a IntValue = 4
      fmt.Printf("%d -> %t\n", a, a.IsEven())
  }
  ```

- **Ek Örnek – String Tabanlı Tip Alias:**

  ```go
  package main

  import (
      "fmt"
      "strings"
  )

  type MyString string

  // MyString üzerinde bir metot tanımlayarak, tüm karakterleri büyük harfe çeviren bir fonksiyon ekleyelim.
  func (s MyString) ToUpper() MyString {
      return MyString(strings.ToUpper(string(s)))
  }

  func main() {
      var s MyString = "merhaba dünya"
      fmt.Println("Orijinal:", s)
      fmt.Println("Büyük Harf:", s.ToUpper())
  }
  ```

- **Not:**
  - Tip alias kullanılırken, orijinal türle uyumlu çalışılması nedeniyle ek dönüşümlere (casting) dikkat edilmelidir.

---

## 6. Uygulamalı Örnekler: Tarih (Date – Tarih) ve Kesir (Fraction – Kesir) Yapıları

### 6.1. Tarih Yapısı (CSDDate)

- **Amaç ve Gereksinimler:**

  - Tarih bilgilerini (gün, ay, yıl) saklayan bir yapı oluşturulacak.
  - Yeni bir tarih oluştururken geçerlilik kontrolü yapılacak; hatalı tarih girişlerinde `nil` dönecek.
  - Tarihle ilgili metotlar:
    - `IsLeap`: Yılın artık yıl olup olmadığını kontrol eder.
    - `DayOfYear`: Yılın kaçıncı günü olduğunu hesaplar.
    - `Set`, `Day`, `Month`, `Year`: Tarih bileşenlerini güncelleyen, geçerlilik kontrolü yapan metotlardır.
    - `DayOfWeek`: Tarihin haftanın hangi günü olduğunu belirler (ekstra fonksiyon olarak tanımlanmıştır).

- **Kod Açıklamaları ve Ek Detaylar:**

  - **Geçerlilik Kontrolü:**

    - `isValidDate` fonksiyonu, girilen gün, ay ve yıl bilgilerini kontrol ederek tarih geçerliliğini belirler.
    - `getDays` fonksiyonu, ilgili ayın gün sayısını döndürür; Şubat için artık yıl kontrolü yapılır.

  - **Hesaplamalar:**
    - `getDayOfYear` fonksiyonu, verilen tarih için yılın kaçıncı günü olduğunu hesaplamak üzere switch-case yapısıyla ayların günlerini toplar.
    - `dayOfWeek` fonksiyonu, yılın başlangıcından itibaren geçen toplam gün sayısını baz alarak haftanın gününü hesaplar.

- **Kod Örneği:**

  ```go
  package main

  import (
      "fmt"
      "math"
  )

  type CSDDate struct {
      day, month, year int
  }

  func NewDate(day, month, year int) *CSDDate {
      if isValidDate(day, month, year) {
          return &CSDDate{day, month, year}
      }
      return nil
  }

  func (d *CSDDate) IsLeap() bool {
      return isLeapYear(d.year)
  }

  func (d *CSDDate) DayOfYear() int {
      return getDayOfYear(d.day, d.month, d.year)
  }

  func (d *CSDDate) Set(day, month, year int) bool {
      if !isValidDate(day, month, year) {
          return false
      }
      d.day = day
      d.month = month
      d.year = year
      return true
  }

  func (d *CSDDate) Day(day int) bool {
      if isValidDate(day, d.month, d.year) {
          d.day = day
          return true
      }
      return false
  }

  func (d *CSDDate) Month(month int) bool {
      if isValidDate(d.day, month, d.year) {
          d.month = month
          return true
      }
      return false
  }

  func (d *CSDDate) Year(year int) bool {
      if isValidDate(d.day, d.month, year) {
          d.year = year
          return true
      }
      return false
  }

  // Ek yardımcı fonksiyonlar
  func getDayOfYear(day, month, year int) int {
      result := day
      switch month - 1 {
      case 11:
          result += 30
          fallthrough
      case 10:
          result += 31
          fallthrough
      case 9:
          result += 30
          fallthrough
      case 8:
          result += 31
          fallthrough
      case 7:
          result += 31
          fallthrough
      case 6:
          result += 30
          fallthrough
      case 5:
          result += 31
          fallthrough
      case 4:
          result += 30
          fallthrough
      case 3:
          result += 31
          fallthrough
      case 2:
          result += 28
          if isLeapYear(year) {
              result++
          }
          fallthrough
      case 1:
          result += 31
      }
      return result
  }

  func isValidDate(day, month, year int) bool {
      return (1 <= day && day <= 31) && (1 <= month && month <= 12) && (1900 <= year) && (day <= getDays(month, year))
  }

  func getDays(month, year int) int {
      days := 31
      switch month {
      case 4, 6, 9, 11:
          days = 30
      case 2:
          days = 28
          if isLeapYear(year) {
              days++
          }
      }
      return days
  }

  func isLeapYear(year int) bool {
      return year%4 == 0 && year%100 != 0 || year%400 == 0
  }

  func main() {
      // runApp fonksiyonu kullanıcı ile etkileşimi sağlar
      runApp()
  }

  func runApp() {
      for {
          var d, m, y int
          fmt.Print("Tarih bilgisini giriniz (gün ay yıl): ")
          fmt.Scan(&d, &m, &y)
          if d == 0 && m == 0 && y == 0 {
              break
          }
          printDateTR(d, m, y)
      }
  }

  func printDateTR(day, month, year int) {
      date := NewDate(day, month, year)
      if date == nil {
          fmt.Println("Geçersiz tarih!")
          return
      }
      // Haftanın gününü hesaplama (0: Pazar, 1: Pazartesi, …, 6: Cumartesi)
      switch dayOfWeek(day, month, year) {
      case 0:
          fmt.Printf("%02d.%02d.%04d Pazar\n", day, month, year)
      case 1:
          fmt.Printf("%02d.%02d.%04d Pazartesi\n", day, month, year)
      case 2:
          fmt.Printf("%02d.%02d.%04d Salı\n", day, month, year)
      case 3:
          fmt.Printf("%02d.%02d.%04d Çarşamba\n", day, month, year)
      case 4:
          fmt.Printf("%02d.%02d.%04d Perşembe\n", day, month, year)
      case 5:
          fmt.Printf("%02d.%02d.%04d Cuma\n", day, month, year)
      case 6:
          fmt.Printf("%02d.%02d.%04d Cumartesi\n", day, month, year)
      }
  }

  // Haftanın gününü hesaplayan yardımcı fonksiyonlar
  func dayOfWeek(day, month, year int) int {
      total := totalDays(getDayOfYear(day, month, year), year)
      return total % 7
  }

  func totalDays(dof, year int) int {
      result := dof
      for y := 1900; y < year; y++ {
          result += 365
          if isLeapYear(y) {
              result++
          }
      }
      return result
  }
  ```

- **Ek Notlar:**
  - Tarih hesaplamalarında kullanılan algoritmalar, özellikle artık yıl kontrolü ve ayların gün sayılarının toplanması kritik önem taşır.
  - Kullanıcı girişleri doğrulanarak, hatalı veri girişlerinde uygun uyarı mesajları verilmektedir.

---

### 6.2. Kesir Yapısı (Fraction)

- **Amaç ve Gereksinimler:**
  - Kesirleri (rasyonel sayılar – rational numbers) temsil etmek için bir yapı oluşturulacak.
  - Kesirlerin sadeleştirilmesi, işaret kontrolü ve aritmetik işlemleri (toplama, çıkarma, çarpma, bölme) için fonksiyonlar/metotlar tanımlanacaktır.
- **Özellikler ve İşlevler:**

  - **Sadeleştirme (Simplify):**

    - `simplify` metodu, pay ve paydayı ortak bölen ile sadeleştirir.
    - Bu, kesirin en sade formunu elde etmek için gereklidir.

  - **İşaret Ayarı (SetSign):**

    - Negatif işaretin yalnızca payda üzerinde olmamasını sağlar; böylece tutarlı bir gösterim elde edilir.

  - **Aritmetik İşlemler:**

    - `add`, `subtract`, `multiply`, `divide` gibi fonksiyonlar, iki kesirin aritmetik işlemlerini gerçekleştirir.
    - Fonksiyonlar hem kesirler arası hem de kesir ile tam sayı işlemlerine izin verir.

  - **Gerçek Değer Hesaplama:**
    - `GetRealValue` metodu, kesiri ondalık değere çevirir.

- **Kod Örneği:**

  ```go
  package main

  import (
      "fmt"
      "math"
      "math/rand"
  )

  type Fraction struct {
      a, b int
  }

  // Kesiri sadeleştirir.
  func (f *Fraction) simplify() {
      minValue := int(math.Min(math.Abs(float64(f.a)), math.Abs(float64(f.b))))
      for i := minValue; i >= 2; i-- {
          if f.a%i == 0 && f.b%i == 0 {
              f.a /= i
              f.b /= i
              break
          }
      }
  }

  // Kesirin pay ve paydasını ayarlar; sıfır pay durumunda payda 1 olarak ayarlanır.
  func (f *Fraction) set(a, b int) {
      f.a = a
      f.b = b
      if a != 0 {
          f.setSign()
          f.simplify()
      } else {
          f.b = 1
      }
  }

  // İşaret düzenlemesi: Eğer payda negatifse, işaretler değiştirilir.
  func (f *Fraction) setSign() {
      if f.b < 0 {
          f.a = -f.a
          f.b = -f.b
      }
  }

  func check(a, b int) int {
      if b == 0 {
          if a == 0 {
              return -1 // Belirsiz (indeterminate)
          }
          return 0 // Tanımsız (undefined)
      }
      return 1 // Geçerli
  }

  func add(a1, b1, a2, b2 int) Fraction {
     	p, _ := New(a1*b2+a2*b1, b1*b2)
      return *p
  }

  func subtract(a1, b1, a2, b2 int) Fraction {
      return add(a1, b1, -a2, b2)
  }

  func multiply(a1, b1, a2, b2 int) Fraction {
      p, _ := New(a1*a2, b1*b2)
      return *p
  }

  func divide(a1, b1, a2, b2 int) (*Fraction, int) {
      return New(a1*b2, a2*b1)
  }

  // Yeni bir Fraction oluşturur. Geçersiz durumlarda nil döner.
  func New(a, b int) (*Fraction, int) {
      status := check(a, b)
      if status != 1 {
          return nil, status
      }
      var f *Fraction = &Fraction{}
      f.set(a, b)
      return f, 1
  }

  func (f *Fraction) SetNumerator(val int) {
      f.set(val, f.b)
  }

  func (f *Fraction) SetDenominator(val int) bool {
      if check(f.a, val) != 1 {
          return false
      }
      f.set(f.a, val)
      return true
  }

  func (f *Fraction) GetRealValue() float64 {
      return float64(f.a) / float64(f.b)
  }

  func (f *Fraction) Add(other *Fraction) Fraction {
      return add(f.a, f.b, other.a, other.b)
  }

  func (f *Fraction) AddWith(val int) Fraction {
      return add(f.a, f.b, val, 1)
  }

  func (f *Fraction) Subtract(other Fraction) Fraction {
      return subtract(f.a, f.b, other.a, other.b)
  }

  func (f *Fraction) SubtractWith(val int) Fraction {
      return subtract(f.a, f.b, val, 1)
  }

  func (f *Fraction) Multiply(other Fraction) Fraction {
      return multiply(f.a, f.b, other.a, other.b)
  }

  func (f *Fraction) MultiplyWith(val int) Fraction {
      return multiply(f.a, f.b, val, 1)
  }

  func (f *Fraction) Divide(other Fraction) (*Fraction, int) {
      return divide(f.a, f.b, other.a, other.b)
  }

  func (f *Fraction) DivideWith(val int) (*Fraction, int) {
      return divide(f.a, f.b, val, 1)
  }

  func (f *Fraction) Inc() {
      f.a += f.b
  }

  func (f *Fraction) Dec() {
      f.a -= f.b
  }

  func (f *Fraction) Print() {
      fmt.Printf("%d / %d = %f\n", f.a, f.b, f.GetRealValue())
  }

  func main() {
      for {
          a := rand.Intn(20) - 10
          b := rand.Intn(20) - 10
          p, status := New(a, b)
          fmt.Println("------------------------------------")
          fmt.Printf("a = %d, b = %d\n", a, b)
          if status == 1 {
              p.Print()
          } else if status == 0 {
              fmt.Println("Undefined")
          } else {
              fmt.Println("Indeterminate")
          }
          fmt.Println("------------------------------------")
          if a == 0 && b == 0 {
              break
          }
      }
  }
  ```

- **Ek Açıklamalar:**
  - Kesirlerde, payda sıfır olmamalıdır. Bu durum kontrol edilerek, hata durumlarında uygun bir geri dönüş sağlanır.
  - `simplify` fonksiyonu, ortak bölenleri bularak kesiri en sade haline getirir; bu, matematiksel işlemlerin tutarlılığını artırır.

---

## 7. Go'da Yapıların Kullanım Alanları ve Önemli Noktalar

- **Kullanım Alanları:**

  - **Veri Modelleme:** İlişkili verilerin bir arada tutulması (örneğin, kullanıcı bilgileri, ürün detayları)
  - **Nesne Yönelimli Programlama Yaklaşımı:** Yapılar ve metotlar kullanılarak, sınıf benzeri davranışlar elde edilebilir.
  - **Veri Yapıları Oluşturma:** Listeler, ağaçlar, graf yapıları gibi karmaşık veri modelleri geliştirme.
  - **Performans ve Bellek Yönetimi:** Küçük yapılar için kopyalama maliyeti düşüktür; ancak büyük yapılar için pointer kullanımı, performansı artırır.

- **Önemli İpuçları:**

  - **Okunabilirlik:** Yapı tanımları ve metotlar, kodun okunabilirliğini artıracak şekilde sade ve anlaşılır tutulmalıdır.
  - **Doğru Receiver Seçimi:** Değişiklik yapılması gereken durumlarda pointer receiver; salt okuma işlemlerinde value receiver tercih edilmelidir.
  - **Fonksiyonlara Geçirme:** Gereksiz kopyalama işlemlerinden kaçınmak için, büyük yapılar pointer olarak fonksiyonlara geçirilmelidir.
  - **Tip Alias Kullanımı:** Temel türlere ek işlevsellik kazandırmak ve kodun anlamını güçlendirmek için tip alias (type alias) kullanımı oldukça yararlıdır.

- **Tablo: Yapıların Avantajları ve Dikkat Edilmesi Gerekenler**

  | **Avantaj / Önemli Nokta**     | **Açıklama**                                                                                                   |
  | ------------------------------ | -------------------------------------------------------------------------------------------------------------- |
  | Basit Sözdizimi                | C benzeri, ancak daha sade ve anlaşılır yapısı sayesinde öğrenilmesi kolaydır.                                 |
  | Hızlı Derleme ve Hata Yakalama | Derleyici, birçok hatayı derleme zamanında yakalayarak, güvenli kod yazımını destekler.                        |
  | Esnek Veri Modelleme           | Karmaşık verilerin mantıksal olarak gruplanması ve yönetilmesi mümkündür.                                      |
  | Performans                     | Küçük yapılar için kopyalama maliyeti düşük; büyük yapılar için pointer kullanımı ile verimlilik sağlanır.     |
  | Metot Eklenebilmesi            | Yapılar üzerine tanımlanan metotlar, nesneye yönelik programlama yaklaşımını destekler.                        |
  | Tip Alias Desteği              | Temel türlere yeni anlamlar yüklenebilir ve bu türlere özgü işlevler eklenerek kod modüler hale getirilebilir. |

---

## Sonuç

Gönderdiğiniz örnek metinler, Go dilinde yapıların tanımlanması, kullanılması, fonksiyonlara aktarılması, metot tanımlama ve type alias (tip takma adı) konularını kapsamlı ve doğru biçimde ele almıştır.  
Aşağıdaki detaylandırılmış açıklamalar ve ek örnekler sayesinde:

- **Yapıların tanımını ve temel özelliklerini,**
- **Yapı elemanlarına erişim yöntemlerini,**
- **Value (değer) ve pointer (işaretçi) receiver arasındaki farkları,**
- **Fonksiyonlara yapı gönderme yöntemlerini,**
- **Type alias ile ek işlevsellik kazandırma yöntemlerini,**
- **Tarih (Date – Tarih) ve kesir (Fraction – Kesir) yapılarının uygulamalı örneklerini**

daha derinlemesine ve kapsamlı bir şekilde kavrayabilirsiniz.
