# dice-game

sebuah script permainan dadu yang menerima input N jumlah pemain dan M jumlah
dadu, dengan peraturan sebagai berikut:
1. Pada awal permainan, setiap pemain mendapatkan dadu sejumlah M unit.
2. Semua pemain akan melemparkan dadu mereka masing-masing secara bersamaan
3. Setiap pemain akan mengecek hasil dadu lemparan mereka dan melakukan evaluasi
seperti berikut:
* Dadu angka 6 akan dikeluarkan dari permainan dan ditambahkan sebagai poin
bagi pemain tersebut.
* Dadu angka 1 akan diberikan kepada pemain yang duduk disampingnya.
Contoh, pemain pertama akan memberikan dadu angka 1 nya ke pemain kedua.
* Dadu angka 2,3,4 dan 5 akan tetap dimainkan oleh pemain.
4. Setelah evaluasi, pemain yang masih memiliki dadu akan mengulangi step yang ke-2
sampai tinggal 1 pemain yang tersisa.
a. Untuk pemain yang tidak memiliki dadu lagi dianggap telah selesai bermain.
5. Pemain yang memiliki poin terbanyak lah yang menang.

tambahan peraturan:
* Apabila terdapat lebih dari 1 pemain dengan poin tertinggi, maka permainan berakhir dengan imbang untuk para pemain tersebut
