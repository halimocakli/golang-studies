
# Go Programı: İki Nokta Arasındaki Mesafeyi Hesaplama

## Açıklama
Bu program, 2D düzlemdeki iki nokta arasındaki Öklid mesafesini Go dilinde hesaplar.

## Özellikler
- Kullanıcıdan iki noktanın koordinatlarını alır.
- Aşağıdaki formülü kullanarak noktalar arasındaki mesafeyi hesaplar:
  \[
  	ext{distance} = \sqrt{(x_2 - x_1)^2 + (y_2 - y_1)^2}
  \]
- Geçersiz girişler için hata yönetimi içerir.

## Kullanım

### Gereksinimler
- Go programlama dili kurulu olmalıdır. [Go Yükleme Kılavuzu](https://go.dev/doc/install)

### Çalıştırma Adımları
1. Bu projeyi klonlayın veya indirin.
2. `main.go` dosyasının bulunduğu dizine gidin.
3. Terminalde aşağıdaki komutu çalıştırın:
   ```bash
   go run main.go
   ```

### Giriş Formatı
- Program sizden iki noktanın koordinatlarını şu formatta girmenizi isteyecek:
  ```
  x1 y1
  x2 y2
  ```

### Örnek
#### Girdi:
```
Enter coordinates for 1st point [x1 y1]: 3 4
Enter coordinates for 2nd point [x2 y2]: 6 8
```

#### Çıktı:
```
The distance between the points is: 5.00
```

## Kod Yapısı
- **Point Struct**: 2D uzayında bir noktayı temsil eder ve `x`, `y` koordinatlarını içerir.
- **main Function**: Kullanıcıdan giriş alır, doğrulama yapar ve mesafeyi hesaplar.

## Hata Yönetimi
- Beklenen formatta olmayan girişler için uygun hata mesajları görüntüler.
- Geçersiz bir giriş yapılırsa program sonlanır.

## Lisans
Bu proje MIT Lisansı ile lisanslanmıştır.
