package main

import "fmt"

const NMAX int = 999

type Menu struct {
	id           int
	nama         string
	kategori     string
	harga        int
	komposisi    string
	ketersediaan bool
}

type arrMenu [NMAX]Menu

func main() {
	var mode, a int
	var Menu arrMenu

	for {
		fmt.Printf("%-4c %s\n", ' ', "Cafe Menu")
		fmt.Println("-------------------")
		fmt.Println("1. Admin")
		fmt.Println("2. Pelanggan")
		fmt.Println("3. Exit")
		fmt.Scan(&mode)

		switch mode {
		case 1:
			adminMode(&Menu, &a)

		case 2:
			userMode()

		case 3:
			exit()

		default:
			fmt.Println("Mode tidak valid!")
		}
	}

}

// initData -> mengisi data menu ke dalam array arrMenu secara sequential.
// Data dimasukkan oleh user melalui input terminal.
//
// Proses berhenti jika:
// - user memasukkan -1 (sentinel value untuk keluar)
// - jumlah data sudah mencapai NMAX (kapasitas maksimum array)
//
// Parameter:
// m : pointer ke array menu yang akan diisi
// a : pointer ke jumlah data yang berhasil dimasukkan
//
// Output:
// - array m akan terisi data menu
// - nilai a akan berisi jumlah data yang berhasil disimpan
//
// Mekanisme:
// - Input ID menu terlebih dahulu
// - Jika valid, lanjut input detail menu (nama, kategori, harga, dll)
// - Data disimpan secara berurutan berdasarkan index i
func initData(m *arrMenu, a *int) {
	var i, input, harga int
	var nama, kategori, komposisi string
	var ketersediaan bool

	fmt.Println("<ID> <NamaMenu> <Kategori> <Harga> <Komposisi> <Ketersediaan(true/false)>")
	fmt.Println("Ketik -1 untuk keluar")
	i = 0

	for {
		fmt.Scan(&input)

		if i >= NMAX || input == -1{
			break
		}

		m[i].id = input

		fmt.Scan(
			&nama, 
			&kategori, 
			&harga, 
			&komposisi, 
			&ketersediaan)

		m[i].nama = nama
		m[i].kategori = kategori
		m[i].harga = harga
		m[i].komposisi = komposisi
		m[i].ketersediaan = ketersediaan

		i = i + 1
	}

	*a = i
}


// tambahData -> menambahkan 1 data menu baru ke dalam array arrMenu.
//
// Proses penambahan dilakukan dengan langkah berikut:
// - Mengecek apakah data sudah penuh (mencapai NMAX)
// - Menampilkan format input yang diminta user
// - Membaca input ID menu dari user secara berulang sampai valid
// - Jika input = -1, proses penambahan dibatalkan (exit manual)
// - Melakukan pengecekan apakah ID sudah ada (menghindari duplikasi)
// - Jika ID sudah ada, user diminta memasukkan ID lain (loop sampai valid)
// - Jika ID belum ada, sistem lanjut ke input detail menu
// - Data menu (nama, kategori, harga, komposisi, ketersediaan) dimasukkan ke array
// - Jumlah data (*a) akan bertambah 1 setelah data berhasil ditambahkan
//
// Parameter:
// m : pointer ke array menu yang menyimpan seluruh data
// a : pointer ke jumlah data yang sedang tersimpan
//
// Output:
// - Menambahkan data baru ke array jika valid
// - Mengupdate nilai a jika penambahan berhasil
// - Menampilkan pesan sesuai kondisi (penuh, duplicate, sukses, atau exit)
func tambahData(m *arrMenu, a *int) {
	var i, input, harga int
	var nama, kategori, komposisi string
	var duplikat, ketersediaan bool

	if *a == NMAX {
		fmt.Println("Data Penuh!")
		return
	}

	fmt.Println("<ID> <NamaMenu> <Kategori> <Harga> <Komposisi> <Ketersediaan(true/false)>")
	fmt.Println("Ketik -1 untuk keluar")

	for {
		fmt.Scan(&input)

		duplikat = false

		if input == -1 {
			fmt.Println("Keluar dari tambah data!")
			return
		}

		for i = 0; i < *a; i++ {
			if m[i].id == input {
				duplikat = true
				break
			}
		}

		if duplikat == false {
			break
		} else {
			fmt.Println("ID sudah ada, silakan input ulang ID: ")
		}
	}

	m[*a].id = input

	fmt.Scan(
		&nama, 
		&kategori, 
		&harga, 
		&komposisi, 
		&ketersediaan)

	m[*a].nama = nama
	m[*a].kategori = kategori
	m[*a].harga = harga
	m[*a].komposisi = komposisi
	m[*a].ketersediaan = ketersediaan

	*a = *a + 1

	fmt.Println("Menu baru ditambahkan!")

}


// lihatData -> menampilkan seluruh data menu yang tersimpan di dalam array arrMenu.
//
// Fungsi ini melakukan traversal dari index 0 sampai (a - 1) untuk menampilkan
// setiap data menu yang sudah diinput.
//
// Data yang ditampilkan meliputi:
// - ID menu
// - Nama menu
// - Kategori menu
// - Harga
// - Komposisi
// - Status ketersediaan
//
// Parameter:
// m : pointer ke array menu yang berisi data
// a : pointer ke jumlah data yang tersimpan
//
// Output:
// - Menampilkan seluruh data menu ke terminal dalam format terstruktur
func lihatData(m *arrMenu, a *int) {
	var i int

	for i = 0; i < *a; i++ {
		fmt.Printf(
			"%d %s %s %d %s %v\n", 
			m[i].id, 
			m[i].nama, 
			m[i].kategori,
			m[i].harga, 
			m[i].komposisi, 
			m[i].ketersediaan)
	}

	fmt.Println("")
}


// ubahData -> melakukan update (perubahan) terhadap data menu berdasarkan ID yang dicari.
//
// Fungsi ini akan:
// - Menerima input ID menu yang ingin diubah
// - Jika input = -1, proses ubdah data dibatalkan (exit manual)
// - Melakukan pencarian data menu berdasarkan ID (sequential search)
// - Jika data ditemukan, seluruh field menu akan diperbarui dengan input baru
// - Jika ID tidak ditemukan, akan menampilkan pesan error
//
// Parameter:
// m : pointer ke array menu yang menyimpan data
// a : pointer ke jumlah data yang tersimpan
//
// Output:
// - Data menu akan diperbarui jika ID ditemukan
// - Menampilkan pesan "Data Berhasil Diedit" jika sukses
// - Menampilkan "ID tidak ditemukan" jika data tidak ada
func ubahData(m *arrMenu, a *int) {
	var edit, i, harga int
	var nama, kategori, komposisi string
	var ketersediaan bool

	fmt.Print("Masukkan ID menu: ")
	fmt.Println("Ketik -1 untuk keluar")
	fmt.Scan(&edit)

	if edit == -1 {
		fmt.Println("Keluar dari ubah data!")
		return
	}

	for i = 0; i < *a; i++ {
		if m[i].id == edit {
			fmt.Scan(
				&nama,
				&kategori,
				&harga,
				&komposisi,
				&ketersediaan)
			
			m[i].id = edit
			m[i].nama = nama
			m[i].kategori = kategori
			m[i].harga = harga
			m[i].komposisi = komposisi
			m[i].ketersediaan = ketersediaan

			fmt.Println("Data Berhasil Diedit")
			return
		}
	}

	fmt.Println("ID tidak ditemukan")
}

func hapusData() {

}

func adminMode(menu *arrMenu, a *int) {
	var mode int
	
	for {
		fmt.Printf("%-4c %s\n", ' ', "Mode Admin")
		fmt.Println("-------------------")
		fmt.Println("0. Inisialisasi Data")
		fmt.Println("1. Tambah Data")
		fmt.Println("2. Edit Data")
		fmt.Println("3. Hapus Data")
		fmt.Println("4. Lihat Data")
		fmt.Println("5. Kembali")
		fmt.Scan(&mode)
		fmt.Println("")

		switch mode {
		case 0:
			initData(menu, a)
		case 1:
			if *a == 0 {
				fmt.Println("Data kosong! Memasuki inisialisasi data...")
				initData(menu, a)
			} else {
				tambahData(menu, a)
			}
			
		case 2:
			ubahData(menu, a)

		case 3:
			// belum diimplementasikan

		case 4:
			lihatData(menu, a)

		case 5:
			return

		default:
			fmt.Println("Pilihan tidak valid!")
		}
	}
}

func userMode() {

}

func exit() {
	fmt.Println("Terima kasih telah menggunakan layanan kami!")
}
