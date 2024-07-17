## ascii-art-web

### Hedefler

Ascii-art-web, son projeniz olan [ascii-art](../ascii-) web **GUI** (grafiksel kullanıcı arayüzü) sürümünü kullanmanın mümkün olacağı bir sunucunun oluşturulması ve çalıştırılmasından oluşur. sanat).

Web sayfanız farklı bannerların kullanımına izin vermelidir:

- [shadow](../ascii-art/shadow.txt)
- [standart](../ascii-art/standard.txt)
- [thinkertoy](../ascii-art/thinkertoy.txt)

Aşağıdaki HTTP uç noktalarını uygulayın:

1. GET `/`: Ana sayfaya HTML yanıtı gönderir.\
   1.1. GET İpucu: Sunucudan veri almak ve görüntülemek için [şablonlara gidin](https://golang.org/pkg/html/template/).

2. POST `/ascii-art`: Go sunucusuna veri gönderir (metin ve banner)\
   2.1. POST İpucu: Gönderi isteğinde bulunmak için formu ve diğer [etiket türlerini](https://developer.mozilla.org/en-US/docs/Web/HTML/Element) kullanın.\

POST sonucunu görüntüleme şekliniz size bağlıdır. Tavsiye ettiğimiz şeyler aşağıdakilerden biridir:

- POST tamamlandıktan sonra sonucu `/ascii-art` yolunda görüntüleyin. Yani ana sayfadan başka bir sayfaya geçiyoruz.
- Veya POST sonucunu ana sayfada görüntüleyin. Bu şekilde sonuçlar ana sayfaya eklenir.

Ana sayfada şunlar bulunmalıdır:

- metin girişi
- radyo düğmeleri, banner'lar arasında geçiş yapmak için nesneyi veya başka bir şeyi seçin
- '/ascii-art' dosyasına POST isteği gönderen ve sonucu sayfaya çıkaran düğme.

### HTTP durum kodu

Uç noktalarınız uygun HTTP durum kodlarını döndürmelidir.

- Tamam (200), eğer her şey hatasız giderse.
- Hiçbir şey bulunamazsa (örneğin şablonlar veya afişler) Bulunamadı.
- Yanlış istekler için Kötü İstek.
- İşlenmeyen hatalar için Dahili Sunucu Hatası.

