# Basic Programming

## Bahasa Go (Golang)

Merupakan bahasa pemrograman open source yang dibangun oleh Google dengan tujuan membuat proses pengembangan menjadi lebih sederhana, praktis, dan efisien.

Bahasa Go sangat cocok untuk menulis program dengan tingkat kedalaman yang lebih rendah (lower level programming), termasuk pembuatan aplikasi dan program sistem.

### Contoh Aplikasi:

1. **E-commerce**
2. **Music Player**
3. **Social Media Apps**

### Contoh System Program :

1. **APIs**
2. **Game Engine**
3. **CLI apps**

## Kenapa Harus Golang ?

Terdapat beberapa alasan mengapa kita sebaiknya mempelajari Go :

- Golang adalah bahasa yang sederhana dan membuat proses pengkodean menjadi lebih menyenangkan.
- Golang memiliki sintaks yang ringan dan sudah memiliki dukungan concurrency (pemrosesan bersamaan) bawaan.
- Golang adalah proyek open source yang banyak digunakan oleh perusahaan besar sebagai bahasa pemrograman andalan mereka.

## Variabel dan Tipe Data

Variabel adalah tempat untuk menyimpan data. data biasa memiliki nama dan nama digunakan saat kita ingin memanggil data yang ada di dalamnya.

### Tipe data yang ada di bahasa golang :

- **Boolean :** True or False
- **Numeric :** integer, float, dan complex
- **String**

### Contoh pembuatan Variabel :

```
var age int = 21
```

### Ada dua cara deklarasi dalam bahasa golang

- **Long :**

```
var age int
```

- **short :**

```
age :=
```

#### Kedua pendekatan di atas valid tergantung pada preferensi pribadi.

## Nilai Default

### Lalu dalam golang juga ada nilai Default value untuk variabel yaitu :

- **booleans** : false
- **floats** : 0.0
- **integers :** 0
- **strings :** “”

## Const

Const atau Konstanta merupakan variabel konstan yang tidak dapat diubah, sehingga nilai hanya dimasukan diawal dan tidak bisa diubah

### const bisa diinisialisasikan dengan menggunakan 2 cara yaitu :

- **single constant :** `const pi float = 3.14`
- **multiple constant :**

```
const (
  pi float64 = 3.14
  pi2
  umur = 10
)

```

## Operator

Seperti bahasa pemrograman lainnya, Go juga memiliki berbagai operator yang dapat digunakan:

### Aritmatika :

- (+) (Penjumlahan)
- (-) (Pengurangan)
- (/) (Pembagian)
- (\*) (Perkalian)
- (%) (Modulus)
- (++) (Increment)
- (—) (Decrement)

### Comparision :

- (==)
- (!=)
- (>)
- (<)
- (>=)
- (<=)

### Logical :

- (&&)
- (||)
- (!)

### Bitwise :

- (&)
- (|)
- (^)
- (<<)
- (>>)

### Assignment :

- (=)
- (+=)
- (-=)
- (\*=)
- (/=)
- (%=)
- (<<=)
- (>>=)
- (&=)
- (^=)
- (!=)

### Miscellaneous :

- (&)
- (\* (Pointer))

## Control Structure

Control Structure atau Struktur kontrol adalah konsep dalam pemrograman yang digunakan untuk mengontrol aliran eksekusi program. Ini memungkinkan kita membuat keputusan, melakukan pengulangan kode, dan mengorganisir jalannya program sesuai kebutuhan.

## Conditional Statements

Conditional Statements atau Percabangan adalah salah satu bentuk dari struktur kontrol yang memungkinkan kita untuk menjalankan kode tertentu berdasarkan kondisi atau kriteria yang diberikan. Dengan menggunakan percabangan, kita dapat membuat program yang mampu beradaptasi terhadap berbagai situasi yang mungkin terjadi.

### Terdapat beberapa metode utama yang digunakan untuk melakukan percabangan:

- **Logika If-Else :** Percabangan ini memungkinkan kita untuk menjalankan satu blok kode jika suatu kondisi tertentu terpenuhi, dan menjalankan blok kode lainnya jika kondisi tersebut tidak terpenuhi.
- **Logika If-Else :** Percabangan ini memungkinkan kita untuk menjalankan satu blok kode jika suatu kondisi tertentu terpenuhi, dan menjalankan blok kode lainnya jika kondisi tersebut tidak terpenuhi.

## Perulangan

Pengulangan digunakan untuk menjalankan kode berulang kali selama kondisinya terpenuhi. Pengulangan dapat dicapai menggunakan struktur pengulangan **for**.

### Contohnya, untuk mencetak angka yang kurang dari 5:

```
for i := 0; i < 5; i++ {
  fmt.Println(i)
}
```
