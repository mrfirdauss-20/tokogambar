## Pengerjaan

### Pembenahan Sistem
Misi utama benahi sistem:
Sesuai dengan misi utama, hal pertama yang saya lakukan adalah mencari cara untuk membuat prediksi image lebih baik. 
Setelah membaca mengenaik hasing image, saya menemukan sebuah library dan metode yang menarik yakni membandingkan icon dari kedua image. 
<br />
Hal ini berdasrkan metode hashing yang saya pelajari. Metode hashing average dan preceptual hash di awali dengan pemampatan ukuran dari image. Seperti kita ketahui bersama, komputer membaca image bagaikan sebuah matriks. Jika ukurannya berbeda, color grading berbeda pasti lah hash dari image juga akan berbeda.
<br />
Kita melihat kode awal menggunakan hash standar sha256, di mana jika beda 1 huruf saja (1 komponen matriks) hasil hashnya bisa jauh berbeda. Oleh karena itu, perlu digunakan metode lainnya.<br />
Saya mengubah sedikit model struktr data agar menyimpan byte dari tiap iamge, kemudia byte dijadikan image dan dicompress menajdi ikon, agar skala pengecekan lebih kecil dan program lebih efisien. Kemudian untuk tiap foto dilakukan iterasi dan dibandingkan iconnya dengan foto yang diunggah.

### Pembenahan Tampilan
Karena saya melihat tampilannya tidak menarik jadi saya lakukan sedikit perubahan. Dengan demikian, diharapkan pengunjung terkesan dan mau mengunjungi web ini lagi.
