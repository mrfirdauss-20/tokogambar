# TokoGambar

`TokoGambar` adalah sebuah platform dimana orang bisa menjual gambar yang mereka buat dengan harga tinggi. Gambar itu bisa berupa gambar apapun: wallpaper, fanart, komik, dan lain sebagainya.

Ketika seorang penjual ingin menjual gambarnya, pihak `TokoGambar` akan melakukan beberapa verifikasi untuk menguji bahwa gambar tersebut adalah otentik milik sang penjual. Salah satu langkahnya adalah mengecek apakah gambar tersebut sudah pernah dijual oleh orang lain di situs `TokoGambar` atau tidak.

Jika gambar yang ingin dijual sudah pernah dijual oleh orang lain, maka gambar tersebut akan ditolak oleh pihak `TokoGambar`.

Keunikan dari suatu gambar ditentukan berdasarkan kesamaan visualnya. Jika admin menganggap bahwa gambar tersebut sama secara visual dengan gambar lain yang berada di situs `TokoGambar`, itu artinya gambar tersebut tidak unik.

Pihak `TokoGambar` menggunakan sistem yang ada di repositori ini untuk mengecek otentisitas dari suatu gambar. Namun sayangnya sistem ini masih berupa prototipe, sehingga masih banyak perbaikan yang perlu dilakukan disana-sini.

## Misi Kamu

Perbaiki sistem yang ada di repositori ini.

Coba pikirkan hal-hal apa saja yang perlu diperbaiki dari sudut pandang bisnis `TokoGambar`. Prioritaskan perubahanmu berdasarkan apa yang memang menurut kamu layak untuk diprioritaskan.

## Pengerjaan Tantangan

1. Fork repositori ini.
2. Analisis masalah dengan menggunakan sudut pandang bisnis `TokoGambar`.
3. Tulis masalah-masalah yang kamu temukan dan apa solusinya di `README.md`.
4. Implementasikan solusi dari masalah yang kamu temukan.
5. Tuliskan juga di `README.md` hal apa saja yang akan kamu lakukan jika kamu punya waktu lebih banyak untuk mengerjakan tantangan ini.
6. Bagikan link repositori hasil fork beserta CV kamu kepada kami.

## Menjalankan Program

Untuk menjalankan program, cukup ketik perintah berikut di terminal:

```
make run
```

Setelah program berjalan, kamu bisa akses programnya di `http://localhost:7124`. Kamu bisa menggunakan gambar-gambar yang ada di direktori [input](./input) untuk mencoba programnya.

Kalau implementasimu program-mu sudah benar, seharusnya `input_1.jpg`, `input_2.jpg`, `input_3.jpg` semuanya akan memiliki hasil.

> **Note:**
>
> Pastikan setidaknya Go versi 1.15 sudah terinstall di komputer kamu.

## Pertanyaan

Jika ada hal-hal yang ingin ditanyakan, silakan buat issue di repositori ini.

## Pustaka Menarik

- [Perceptual Hash](http://www.hackerfactor.com/blog/index.php?/archives/432-Looks-Like-It.html)
- [Time Complexity Analysis](https://www.youtube.com/watch?v=fZc3ijGM0aM)
- [Levenshtein Distance](https://en.wikipedia.org/wiki/Levenshtein_distance)