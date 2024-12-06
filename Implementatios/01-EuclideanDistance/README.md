
# Go Programı: İki Nokta Arasındaki Mesafeyi Hesaplama

## Açıklama
Bu program, 2D düzlemdeki iki nokta arasındaki Öklid mesafesini Go kullanarak hesaplar.

## Özellikler
- Kullanıcıdan iki noktanın koordinatlarını alır.
- Aşağıdaki formülü kullanarak noktalar arasındaki mesafeyi hesaplar:
  \[
  	ext{mesafe} = \sqrt{(x_2 - x_1)^2 + (y_2 - y_1)^2}
  \]
- Geçersiz girişler için hata yönetimi sağlar.

## Kullanım

### Gereksinimler
- Go programlama dili yüklü olmalıdır. [Go Yükleme Kılavuzu](https://go.dev/doc/install)

### Çalıştırma Adımları
1. Bu depoyu klonlayın veya indirin.
2. `main.go` dosyasının bulunduğu dizinde bir terminal açın.
3. Aşağıdaki komutu çalıştırın:
   ```bash
   go run main.go
   ```

### Giriş Formatı
- Program, aşağıdaki formatta iki noktanın koordinatlarını girmenizi isteyecektir:
  ```
  x1 y1
  x2 y2
  ```

### Örnek
#### Girdi:
```
1. noktanın koordinatlarını girin [x1 y1]: 3 4
2. noktanın koordinatlarını girin [x2 y2]: 6 8
```

#### Çıktı:
```
Noktalar arasındaki mesafe: 5.00
```

## Kod Yapısı
- **Point Yapısı (Struct)**: `x` ve `y` koordinatlarına sahip bir noktayı temsil eder.
- **main Fonksiyonu**: Giriş alır, doğrular ve mesafeyi hesaplar.

## Hata Yönetimi
- Beklenmeyen formatta girişler için uygun hata mesajları verir.
- Girişler doğru formatta değilse program sonlanır.

## Lisans
Bu proje MIT Lisansı ile lisanslanmıştır.
