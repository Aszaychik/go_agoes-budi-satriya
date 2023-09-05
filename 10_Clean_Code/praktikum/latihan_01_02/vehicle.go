package main

import (
	"fmt"
)

// Struct Kendaraan mendefinisikan atribut yang umum untuk semua kendaraan.
type Kendaraan struct {
	TotalRoda       int // Jumlah roda kendaraan.
	KecepatanPerJam int // Kecepatan kendaraan dalam kilometer per jam.
}

// Struct Mobil adalah jenis kendaraan khusus yang menyertakan Kendaraan sebagai embedded struct.
type Mobil struct {
	Kendaraan // Mobil memiliki atribut yang sama dengan Kendaraan.
}

// Berjalan adalah metode untuk meningkatkan kecepatan Mobil sebesar kecepatanBaru.
func (m *Mobil) Berjalan() {
	m.TambahKecepatan(10)
}

// TambahKecepatan adalah metode untuk menambahkan kecepatan saat ini dengan kecepatanBaru.
func (m *Mobil) TambahKecepatan(kecepatanBaru int) {
	m.KecepatanPerJam += kecepatanBaru
}

func main() {
	// Membuat objek mobilCepat dari jenis Mobil.
	mobilCepat := Mobil{}

	// Memanggil metode Berjalan pada mobilCepat sebanyak 3 kali.
	mobilCepat.Berjalan()
	mobilCepat.Berjalan()
	mobilCepat.Berjalan()

	// Membuat objek mobillamban dari jenis Mobil.
	mobilLamban := Mobil{}

	// Memanggil metode Berjalan pada mobilLamban.
	mobilLamban.Berjalan()

	// Menampilkan kecepatan mobilLamban.
	fmt.Println("Kecepatan mobil lamban:", mobilLamban.KecepatanPerJam)
}
