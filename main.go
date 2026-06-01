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

	a = 0

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

//============ Manage Data ============
func initData(m *arrMenu, a *int) {
	var i, input, harga int
	var nama, kategori, komposisi string
	var ketersediaan bool

	if *a == 0 {
		fmt.Println("Range ID untuk makanan: 1-99")
		fmt.Println("Range ID untuk minuman non kopi: 100-199")
		fmt.Println("Range ID untuk minuman kopi: 200-299")
		fmt.Println("<ID> <NamaMenu> <Kategori> <Harga> <Komposisi> <Ketersediaan(true/false)>")
		fmt.Println("Ketik -1 untuk keluar")
		i = 0

		for {
			fmt.Scan(&input)

			if i >= NMAX || input == -1 {
				break
			}

			fmt.Scan(
				&nama,
				&kategori,
				&harga,
				&komposisi,
				&ketersediaan)

			if (kategori == "food" && (input > 0 && input < 100)) || (kategori == "noncoffee" && (input >= 100 && input < 200)) || (kategori == "coffee" && (input >= 200 && input < 300)) {
				m[i].id = input
				m[i].nama = nama
				m[i].kategori = kategori
				m[i].harga = harga
				m[i].komposisi = komposisi
				m[i].ketersediaan = ketersediaan

				i = i + 1
			} else {
				fmt.Println("Kategori tidak sesuai dengan range ID, data tidak disimpan!")
			}

		}

		*a = i
	} else {
		fmt.Println("Data sudah diinisialisasi!")
		return
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

	fmt.Println("Range ID untuk makanan: 1-99")
	fmt.Println("Range ID untuk minuman non kopi: 100-199")
	fmt.Println("Range ID untuk minuman kopi: 200-299")
	fmt.Println("Ketik -1 untuk keluar")

	for {
		fmt.Print("Masukkan ID menu: ")
		fmt.Println("")
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

	fmt.Println("<NamaMenu> <Kategori> <Harga> <Komposisi> <Ketersediaan(true/false)>")

	fmt.Scan(
		&nama,
		&kategori,
		&harga,
		&komposisi,
		&ketersediaan)
	fmt.Println("")
	if (kategori == "makanan" && (input > 0 && input < 100)) || (kategori == "minumanNC" && (input >= 100 && input < 200)) || (kategori == "minumanC" && (input >= 200 && input < 300)) {
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

	if *a != 0 {
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
	} else {
		fmt.Print("Data Kosong! Harap Masukan Data Terlebih Dahulu!")
		fmt.Println("")
	}
}

func ubahData(m *arrMenu, a *int) {
	var edit, i int

	fmt.Println("Range ID untuk makanan: 1-99")
	fmt.Println("Range ID untuk minuman non kopi: 100-199")
	fmt.Println("Range ID untuk minuman kopi: 200-299")
	fmt.Println("Ketik -1 untuk keluar")
	fmt.Print("Masukkan ID menu: ")
	fmt.Scan(&edit)

	if edit == -1 {
		fmt.Println("Keluar dari ubah data!")
		return
	}

	for i = 0; i < *a; i++ {
		if m[i].id == edit {
			fmt.Println("<NamaMenu> <Kategori> <Harga> <Komposisi> <Ketersediaan(true/false)>")
			fmt.Scan(
				&m[i].nama,
				&m[i].kategori,
				&m[i].harga,
				&m[i].komposisi,
				&m[i].ketersediaan)

			fmt.Println("Data Berhasil Diedit")
			return
		}
	}

	fmt.Println("ID tidak ditemukan")
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
			fmt.Println("Masukkan Kategori (food/noncoffee/coffee)")
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
// category by binary binggung bngt buset dah wkwkwkkw
func lihatDataByCategory(m *arrMenu, a *int) {
	var i, mode int
	var category string

	fmt.Println("Pilih metode search:\n1. Sequential\n2. Binary ")
	fmt.Scan(&mode)
	fmt.Println("")

	switch mode {
	case 1:
		fmt.Println("Masukkan Kategori (food/noncoffee/coffee)")
		fmt.Scan(&category)
		fmt.Println("")
		fmt.Println("Menu Kategori ", category)

		for i = 0; i < *a; i++ {
			if m[i].kategori == category {
				fmt.Printf(
					"%d %s %s %d %s %v\n",
					m[i].id,
					m[i].nama,
					m[i].kategori,
					m[i].harga,
					m[i].komposisi,
					m[i].ketersediaan)
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
	var i, j, l, r, mid, min int
	var kategori string
	var temp Menu
	//untuk binary search, data harus diurutkan dulu berdasarkan kategori range sesuai dengan kategorinya, jadi makanan 1-99, minuman non kopi 100-199, minuman kopi 200-299
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

	fmt.Println("Masukkan Kategori (food/noncoffee/coffee)")
	fmt.Scan(&kategori)
	fmt.Println("")
	fmt.Println("Menu Kategori ", kategori)
	l = 0
	r = *a - 1

	for l <= r {
		mid = (l + r) / 2
		if kategori == m[mid].kategori {
			fmt.Printf(
				"%d %s %s %d %s %v\n",
				m[mid].id,
				m[mid].nama,
				m[mid].kategori,
				m[mid].harga,
				m[mid].komposisi,
				m[mid].ketersediaan)
		}
		if kategori < m[mid].kategori {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}

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

	for k = 0; k < *a; k++ {
		fmt.Printf(
			"%d %s %s %d %s %v\n",
			m[k].id,
			m[k].nama,
			m[k].kategori,
			m[k].harga,
			m[k].komposisi,
			m[k].ketersediaan)
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

	for k = 0; k < *a; k++ {
		fmt.Printf(
			"%d %s %s %d %s %v\n",
			m[k].id,
			m[k].nama,
			m[k].kategori,
			m[k].harga,
			m[k].komposisi,
			m[k].ketersediaan)
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
		fmt.Println("0. Inisialisasi Data")
		fmt.Println("1. Tambah Data")
		fmt.Println("2. Edit Data")
		fmt.Println("3. Hapus Data")
		fmt.Println("4. Lihat Data")
		fmt.Println("5. Lihat Statistik Cafe")
		fmt.Println("6. Kembali")
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
