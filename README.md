## Pengerjaan

### Pembenahan Sistem
Misi utama benahi sistem:
Sesuai dengan misi utama, hal pertama yang saya lakukan adalah mencari cara untuk membuat prediksi image lebih baik. 
Setelah membaca mengenaik hasing image, saya menemukan sebuah library dan metode yang menarik yakni membandingkan icon dari kedua image. 
<br />
Hal ini berdasrkan metode hashing yang saya pelajari. Metode hashing average dan preceptual hash di awali dengan pemampatan ukuran dari image. Berdasarkan paper yang saya baca, Perceptual menghasilkan hasil yang lebih optimal, karena nilai collisonnya jauh di bawah average dengan nilai simalarity yang mirip. Untuk menghitung jaraka, saya menggunakan levithense distance antara dua buah string hash dari image.
<br />
Similarity merupakan nilai persentase yang sama antara dua buah string. Dengan rumus (len(hash)-distance)/len(has).

### Pembenahan Tampilan
Karena saya melihat tampilannya tidak menarik jadi saya lakukan sedikit perubahan. Dengan demikian, diharapkan pengunjung terkesan dan mau mengunjungi web ini lagi.
