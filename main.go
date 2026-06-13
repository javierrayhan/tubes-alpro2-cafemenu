package main

import "fmt"

const NMAX int = 300

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

	Menu[0].id = 1
	Menu[0].nama = "NasiGoreng"
	Menu[0].kategori = "food"
	Menu[0].harga = 15000
	Menu[0].komposisi = "Nasi_Telur_Ayam_Bumbu"
	Menu[0].ketersediaan = true

	Menu[1].id = 2
	Menu[1].nama = "MieGoreng"
	Menu[1].kategori = "food"
	Menu[1].harga = 12000
	Menu[1].komposisi = "Mie_Telur_Ayam_Bumbu"
	Menu[1].ketersediaan = false

	Menu[2].id = 100
	Menu[2].nama = "EsTehManis"
	Menu[2].kategori = "noncoffe"
	Menu[2].harga = 7000
	Menu[2].komposisi = "Air_Gula_Es_Teh"
	Menu[2].ketersediaan = true

	Menu[3].id = 101
	Menu[3].nama = "EsJeruk"
	Menu[3].kategori = "noncoffe"
	Menu[3].harga = 10000
	Menu[3].komposisi = "Air_Gula_Es_Jeruk"
	Menu[3].ketersediaan = false

	Menu[4].id = 202
	Menu[4].nama = "KopiHitam"
	Menu[4].kategori = "coffe"
	Menu[4].harga = 15000
	Menu[4].komposisi = "Kopi_Air"
	Menu[4].ketersediaan = true

	Menu[5].id = 203
	Menu[5].nama = "Cappuccino"
	Menu[5].kategori = "coffe"
	Menu[5].harga = 25000
	Menu[5].komposisi = "Air_Kopi_Susu"
	Menu[5].ketersediaan = false

	a = 6

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
			userMode(&Menu, &a)

		case 3:
			fmt.Println("Terima kasih telah menggunakan layanan kami!")
			return

		default:
			fmt.Println("Mode tidak valid!")
		}
	}

}

func tambahData(m *arrMenu, a *int) {
	var i, input, harga int
	var nama, kategori, komposisi string
	var ketersediaan bool
	var duplikat bool

	if *a == NMAX {
		fmt.Println("Data Penuh!")
		return
	}

	fmt.Println("Ketik -1 untuk keluar")

	for {
		fmt.Print("Masukkan ID menu: ")
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

	fmt.Println("<NamaMenu> <Kategori> <Harga> <Ketersediaan(true/false)> <Komposisi>")
	fmt.Scan(
		&nama,
		&kategori,
		&harga,
		&ketersediaan,
		&komposisi)

	fmt.Println("")
	if (kategori == "food" && (input > 0 && input < 100)) || (kategori == "noncoffe" && (input >= 100 && input < 200)) || (kategori == "coffe" && (input >= 200 && input < 300)) {
		m[*a].id = input
		m[*a].nama = nama
		m[*a].kategori = kategori
		m[*a].harga = harga
		m[*a].komposisi = komposisi
		m[*a].ketersediaan = ketersediaan

		*a = *a + 1
		fmt.Println("Menu baru ditambahkan!")
		fmt.Println("")
		return
	} else {
		fmt.Println("Kategori tidak sesuai dengan range ID, data tidak disimpan!")
		fmt.Println("")
		return
	}

}

func lihatData(m *arrMenu, a *int) {
	var i int

	fmt.Printf(
		"%-7s | %-20s | %-15s | %-10s | %-15v | %s\n",
		"ID",
		"Nama",
		"Kategori",
		"Harga",
		"Ketersediaan",
		"Komposisi")

	if *a != 0 {
		for i = 0; i < *a; i++ {
			fmt.Printf(
				"%-7d | %-20s | %-15s | %-10d | %-15v | %s\n",
				m[i].id,
				m[i].nama,
				m[i].kategori,
				m[i].harga,
				m[i].ketersediaan,
				m[i].komposisi,
			)
		}
		fmt.Println("")
	} else {
		fmt.Print("Data Kosong! Harap Masukan Data Terlebih Dahulu!")
		fmt.Println("")
	}
}

func ubahData(m *arrMenu, a *int) {
	var edit, i int
	var pilih string

	for {
		fmt.Println("Ketik -1 untuk keluar")
		fmt.Print("Masukkan ID menu: ")
		fmt.Scan(&edit)

		if edit == -1 {
			fmt.Println("Keluar dari ubah data!")
			return
		}
		for i = 0; i < *a; i++ {
			if m[i].id == edit {
				break
			}
		}
		fmt.Println("Pilih yang mau diubah: ")
		fmt.Println("<NamaMenu> <Kategori> <Harga> <Ketersediaan(true/false)> <Komposisi>")
		fmt.Scan(&pilih)

		if pilih == "NamaMenu" || pilih == "namamenu" {
			fmt.Print("Masukkan Nama Menu Baru: ")
			fmt.Scan(&m[i].nama)
			fmt.Println("Nama Menu berhasil diubah!")
		} else if pilih == "Kategori" || pilih == "kategorimenu" {
			fmt.Print("Masukkan Kategori Baru: ")
			fmt.Scan(&m[i].kategori)
			fmt.Println("Kategori berhasil diubah!")
		} else if pilih == "Harga" || pilih == "harga" {
			fmt.Print("Masukkan Harga Baru: ")
			fmt.Scan(&m[i].harga)
			fmt.Println("Harga berhasil diubah!")
		} else if pilih == "Komposisi" || pilih == "komposisi" {
			fmt.Print("Masukkan Komposisi Baru: ")
			fmt.Scan(&m[i].komposisi)
			fmt.Println("Komposisi berhasil diubah!")
		} else if pilih == "Ketersediaan" || pilih == "ketersediaan" {
			fmt.Print("Masukkan Ketersediaan Baru: ")
			fmt.Scan(&m[i].ketersediaan)
			fmt.Println("Ketersediaan berhasil diubah!")
		} else {
			fmt.Println("Pilihan Tidak Valid!")
			fmt.Println("")
		}
	}

}

func hapusData(m *arrMenu, a *int) {
	var i, j, idTarget int
	var found bool

	fmt.Println("Masukan id yang akan dihapus: ")
	fmt.Scan(&idTarget)

	for i = 0; i < *a; i++ {
		if m[i].id == idTarget {
			found = true

			*a = *a - 1

			for j = i; j < *a; j++ {
				m[j] = m[j+1]
			}

			fmt.Printf("\n Data dengan id %d berhasil dihapus!\n", idTarget)
			return
		}
	}

	if found == false {
		fmt.Println("ID tidak ditemukan!")
	}
}

//============ Statistik ============
func StatistikCafe(m *arrMenu, a *int) {
	var category string
	var i, x, mode, allCategory int
	var rataAllMenu, rataMenuCategory float64

	for {
		fmt.Printf("%-4c %s\n", ' ', "Statistik Cafe")
		fmt.Println("-------------------")
		fmt.Println("1. Statistik Semua Menu")
		fmt.Println("2. Statistik Menu By Kategori")
		fmt.Println("3. Kembali")
		fmt.Scan(&mode)
		fmt.Println("")

		if mode == 1 {
			for i = 0; i < *a; i++ {
				allCategory++
				rataAllMenu += float64(m[i].harga)
			}

			fmt.Printf("Jumlah Semua Menu: %d\n", allCategory)
			fmt.Printf("Jumlah Rata Rata Harga Semua Menu: %.2f\n", rataAllMenu/float64(allCategory))
			return
		} else if mode == 2 {
			fmt.Println("Masukkan Kategori (food/coffe/noncoffe)")
			fmt.Scan(&category)
			fmt.Println("")

			for i = 0; i < *a; i++ {
				if m[i].kategori == category {
					x++
					rataMenuCategory += float64(m[i].harga)
				}
			}

			fmt.Printf("Jumlah Menu kategori %s: %d\n", category, x)
			fmt.Printf("Jumlah Rata Rata Harga Menu kategori %s: %.2f\n", category, rataMenuCategory/float64(x))
			return
		} else if mode == 3 {
			return
		} else {
			fmt.Println("Pilihan Tidak Valid")
		}
	}
}

//============ Lihat menu by Category ============
// category by binary dan Sequential
func lihatDataByCategory(m *arrMenu, a *int) {
	var i, mode int
	var category string

	fmt.Println("Pilih metode search:\n1. Sequential\n2. Binary ")
	fmt.Scan(&mode)
	fmt.Println("")

	switch mode {
	case 1:
		fmt.Println("Masukkan Kategori (food/coffe/noncoffe)")
		fmt.Scan(&category)
		fmt.Println("")
		fmt.Println("Menu Kategori ", category)

		fmt.Printf(
			"%-7s | %-20s | %-15s | %-10s | %-15v | %s\n",
			"ID",
			"Nama",
			"Kategori",
			"Harga",
			"Ketersediaan",
			"Komposisi",
		)

		for i = 0; i < *a; i++ {
			if m[i].kategori == category {
				fmt.Printf(
					"%-7d | %-20s | %-15s | %-10d | %-15v | %s\n",
					m[i].id,
					m[i].nama,
					m[i].kategori,
					m[i].harga,
					m[i].ketersediaan,
					m[i].komposisi)
			}

		}

	case 2:
		categoryBinarySearch(m, a)

	default:
		fmt.Println("Pilihan tidak valid!")
	}

	fmt.Println("")
}

func categoryBinarySearch(m *arrMenu, a *int) {
	var l, r, mid, i int
	var kategori string

	sortBinarySearch(m, a)
	fmt.Println("Masukkan Kategori (food/coffe/noncoffe)")
	fmt.Scan(&kategori)
	fmt.Println("")

	l = 0
	r = *a - 1

	for l <= r {
		mid = (l + r) / 2

		if m[mid].kategori == kategori {
			i = mid
			for i > 0 && m[i-1].kategori == kategori {
				i--
			}

			fmt.Printf("%-7s | %-20s | %-15s | %-10s | %-15v | %s\n", "ID", "Nama", "Kategori", "Harga", "Ketersediaan", "Komposisi")
			for i < *a && m[i].kategori == kategori {
				fmt.Printf("%-7d | %-20s | %-15s | %-10d | %-15v | %s\n", m[i].id, m[i].nama, m[i].kategori, m[i].harga, m[i].ketersediaan, m[i].komposisi)
				i++
			}
			return
		}

		if kategori < m[mid].kategori {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}

	fmt.Println("Kategori tidak ditemukan!")
}

//============ Lihat menu by price ============
func hargaSelectionSort(m *arrMenu, a *int) {
	var i, j, k, min int
	var temp Menu

	for i = 0; i < *a-1; i++ {
		min = i

		for j = i + 1; j < *a; j++ {
			if m[j].harga < m[min].harga {
				min = j
			}
		}

		temp = m[i]
		m[i] = m[min]
		m[min] = temp
	}

	fmt.Printf(
		"%-7s | %-20s | %-15s | %-10s | %-15v | %s\n",
		"ID",
		"Nama",
		"Kategori",
		"Harga",
		"Ketersediaan",
		"Komposisi")

	for k = 0; k < *a; k++ {
		fmt.Printf(
			"%-7d | %-20s | %-15s | %-10d | %-15v | %s\n",
			m[k].id,
			m[k].nama,
			m[k].kategori,
			m[k].harga,
			m[k].ketersediaan,
			m[k].komposisi)
	}

}

//============ Sort Binary Search ============
func sortBinarySearch(m *arrMenu, a *int) {
	var i, j, min int
	var temp Menu

	for i = 0; i < *a-1; i++ {
		min = i

		for j = i + 1; j < *a; j++ {
			if m[j].kategori < m[min].kategori {
				min = j
			}
		}

		temp = m[i]
		m[i] = m[min]
		m[min] = temp
	}

}

func hargaInsertionSort(m *arrMenu, a *int) {
	var i, j, k int
	var temp Menu

	for i = 1; i < *a; i++ {
		temp = m[i]
		j = i - 1

		for j >= 0 && m[j].harga > temp.harga {
			m[j+1] = m[j]
			j = j - 1
		}

		m[j+1] = temp
	}

	fmt.Printf(
		"%-7s | %-20s | %-15s | %-10s | %-15v | %s\n",
		"ID",
		"Nama",
		"Kategori",
		"Harga",
		"Ketersediaan",
		"Komposisi")

	for k = 0; k < *a; k++ {
		fmt.Printf(
			"%-7d | %-20s | %-15s | %-10d | %-15v | %s\n",
			m[k].id,
			m[k].nama,
			m[k].kategori,
			m[k].harga,
			m[k].ketersediaan,
			m[k].komposisi)
	}
}

func lihatDataByPrice(m *arrMenu, a *int) {
	var mode int

	fmt.Println("Pilih metode sort:\n1. Selection\n2. Insertion ")
	fmt.Scan(&mode)
	fmt.Println("")

	switch mode {
	case 1:
		hargaSelectionSort(m, a)

	case 2:
		hargaInsertionSort(m, a)

	default:
		fmt.Println("Pilihan tidak valid!")
	}
}

//============ Admin Mode ============
func adminMode(menu *arrMenu, a *int) {
	var mode int

	for {
		fmt.Printf("%-4c %s\n", ' ', "Mode Admin")
		fmt.Println("-------------------")
		fmt.Println("1. Tambah Data")
		fmt.Println("2. Edit Data")
		fmt.Println("3. Hapus Data")
		fmt.Println("4. Lihat Data")
		fmt.Println("5. Lihat Statistik Cafe")
		fmt.Println("6. Kembali")
		fmt.Scan(&mode)
		fmt.Println("")

		switch mode {
		case 1:
			tambahData(menu, a)

		case 2:
			ubahData(menu, a)

		case 3:
			hapusData(menu, a)

		case 4:
			lihatData(menu, a)

		case 5:
			StatistikCafe(menu, a)

		case 6:
			return

		default:
			fmt.Println("Pilihan tidak valid!")
		}
	}
}

//============ User Mode ============
func userMode(menu *arrMenu, a *int) {
	var mode int

	for {
		fmt.Printf("%-4c %s\n", ' ', "Mode Pelanggan")
		fmt.Println("-------------------")
		fmt.Println("1. Lihat Menu")
		fmt.Println("2. Search Menu By Category")
		fmt.Println("3. Search menu By Descending Price")
		fmt.Println("4. Kembali")
		fmt.Scan(&mode)
		fmt.Println("")

		switch mode {
		case 1:
			lihatData(menu, a)

		case 2:
			lihatDataByCategory(menu, a)

		case 3:
			lihatDataByPrice(menu, a)

		case 4:
			return

		default:
			fmt.Println("Pilihan tidak valid!")
		}
	}
}
