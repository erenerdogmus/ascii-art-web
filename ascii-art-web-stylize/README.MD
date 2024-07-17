### Description
Bu proje, `Go` programlama dili kullanılarak oluşturulmuş bir `web uygulamasıdır`. Kullanıcıların metin girişi yaparak `ASCII sanatı` oluşturmasına olanak tanır. Web tarayıcısı üzerinde çalışan bir localhost sunucusu olarak tasarlanmıştır

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

Bu bölüm, projenin iç işleyişi hakkında daha derin bir anlayış sağlar ve kodunuzun nasıl çalıştığını daha açık bir şekilde açıklar.
