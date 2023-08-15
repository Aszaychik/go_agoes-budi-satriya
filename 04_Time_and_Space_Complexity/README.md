# Time Complexity & Space Complexity

## Time Complexity

Penggunaan Time Complexity memudahkan dalam memperkirakan waktu eksekusi program. Kompleksitas dapat dilihat sebagai jumlah maksimum operasi primitif yang dapat dilakukan oleh program. Operasi-reguler adalah penambahan tunggal, perkalian, penugasan, dll. Beberapa operasi mungkin tidak dihitung dan fokus diberikan pada operasi yang dilakukan sebanyak mungkin. Operasi seperti ini disebut sebagai _Dominan_.

### Contoh Operasi Dominan :

```
func dominant (n int) int {
  var result int = 0
  for i:=0 ; i<n; i++{
  result +=1
  }
  result = result +10
  return result
}
```

Pada kode diatas operasi pada baris ke-4 adalah dominan dan akan dieksekusi sebanyak n kali. Kompleksitas dijelaskan dalam notasi Big-O dalam hal ini O(n) - Kompleksitas Linear

### Berikut contoh beberapa Time Compexities yang berbeda :

- Constant Time - O(1)
- Linear Time - O(n)
- Linear Time - O(n+m)
- Logarithmic Time - O(log n)
- Quadratic Time - O(n^2)

## Batas Waktu

Saat ini, rata rata komputer dapat melakukan 10^8 operasi dalam waktu kurang dari satu detik. Batas waktu yang ditetapkan untuk tes online biasanya dari 1 hingga 10 detik

- n<=1.000.000, Kompleksitas waktu yang diharapkan adalah O(n) atau 0(n log n)
- n<=10.000, kompleksitas waktu yang diharapkan adalah 0(n^2)
- n<=500, kompleksitas waktu yang diharapkan adalah 0(n^3)

Batasan ini belum tentu tepat, ini hanya perkiraan dan akan bervariasi tergantung pada tugas tertentu

## Space Complexity

Memori limits memberikan informasi tentang Space Complexity yang diharapkan. Anda dapat memperkirakan jumlah variabel yang dapat dideklarasikan dalam program Anda. Secara singkat, jika Anda memiliki jumlah variabel yang tetap, Anda juga memiliki Space Complexity yang tetap: dalam notasi Big-O, ini adalah O(1). Jika Anda perlu mendeklarasikan sebuah array dengan n elemen, Anda memiliki Space Complexity linear — O(n).
