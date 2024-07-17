### Description
Bu proje, `Go` programlama dili kullanılarak oluşturulmuş bir `web uygulamasıdır`. Kullanıcıların metin girişi yaparak `ASCII sanatı` oluşturmasına olanak tanır. Web tarayıcısı üzerinde çalışan bir localhost sunucusu olarak tasarlanmıştır.

### installation steps
1. DockerHub'da `coderboyjxp/ascii-art-docker` yazarak konteyneri bulun.
2. Konteynerin sayfasına gidip `Docker Pull Command`metnini kopyalayın.
3. `Terminal` ekranında bu komutu yapıştırarak uygulamayı indirin.
4. komutu kullanarak uygulamayı çalıştırın = `docker run --rm -i -t -d -p 8045:8045 coderboyjxp/ascii-art-docker` 
5. Websitesini kullanmak için ` localhost:8045/index.html` adresine gidin.

## installation steps Terminal
1. Bu kodu `klonlayın` veya `indirin`.
2. Gerekli bağımlılıkları yüklemek için `go mod tidy` komutunu çalıştırın.
3. Kodu çalıştırmak için `go run main.go` komutunu kullanın.
4. Tarayıcınızda `http://localhost:8045` adresine giderek uygulamayı kullanmaya başlayabilirsiniz.


### Usage: how to run
Projeyi çalıştırmak için kullanıcı, web arayüzünde bir metin girişi yapmalıdır. Varsayılan olarak, metin girişi ve `ASCII` sanatı oluşturma işlemi `standard` yazı tipiyle gerçekleşir. Kullanıcı istediği kelimeyi girdikten sonra, yazı tipini değiştirmek istiyorsa `standard`, `shadow` veya `thinkertoy` seçeneklerinden birini seçebilir. Ardından, `Generate ASCII Art` butonuna tıklayarak ASCII sanatını görebilir.

### Implementation details
Bu projenin temel amacı, kullanıcıların web arayüzü üzerinden metin girişi yaparak ASCII sanatını oluşturmasına olanak tanımaktır. Proje, Go dilinde yazılmıştır ve aşağıdaki ana bileşenleri içerir:

- **HTTP Sunucusu**: Proje, standart HTTP paketi kullanılarak bir HTTP sunucusu olarak çalışır. HTTP sunucusu, kullanıcıların tarayıcılarından istekleri dinler ve web sayfalarını sunar.
- **HTML Şablonları**: Proje, `html/template` paketi kullanılarak HTML şablonlarıyla çalışır. Şablonlar, sunucu tarafında işlenir ve istemciye gönderilmeden önce dinamik içerik eklemek için kullanılır.
- **ASCII Sanatı Oluşturma**: Kullanıcıların girdiği metinler, belirli bir ASCII sanatı stili seçilerek ASCII sanatına dönüştürülür. Bu dönüştürme işlemi, projenin temel algoritması olan `getAsciiArt` fonksiyonu tarafından gerçekleştirilir.
- **Dosya Okuma**: Proje, belirli bir ASCII sanat stili dosyasından ASCII karakterlerini okumak için dosya okuma işlemi gerçekleştirir. Bu işlem, `linesRead` fonksiyonu tarafından gerçekleştirilir.

Projenin temel algoritması, kullanıcının girdiği metinleri ASCII sanatına dönüştürmek için adım adım aşağıdaki şekilde çalışır:

1. Kullanıcıdan gelen metin, satır satır ayrıştırılır ve her bir karakterin ASCII değeri hesaplanır.
2. ASCII değerleri kullanılarak, belirli bir ASCII sanat stili dosyasından ilgili ASCII karakterleri okunur.
3. Okunan ASCII karakterleri birleştirilerek, metnin ASCII sanatı oluşturulur.
