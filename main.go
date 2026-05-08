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

type arrMenuMinumanCoffe [NMAX]Menu
type arrMenuMinumanNonCoffe [NMAX]Menu
type arrMenuMakanan [NMAX]Menu

func main() {
	var mode, a, b, c int
	var MenuMakanan arrMenuMakanan
	var MenuMinumanNonCoffe arrMenuMinumanNonCoffe
	var MenuMinumanCoffe arrMenuMinumanCoffe

	fmt.Printf("%-4c %s\n", ' ', "Cafe Menu")
	fmt.Println("-------------------")
	fmt.Println("1. Admin")
	fmt.Println("2. Pelanggan")
	fmt.Println("3. Exit")
	fmt.Scan(&mode)

	if mode == 1 {
		adminMode(&MenuMakanan, &MenuMinumanCoffe, &MenuMinumanNonCoffe, &a, &b, &c)
	} else if mode == 2 {
		userMode()
	} else if mode == 3 {
		exit()
	}
}

func tambahData(m *arrMenuMakanan, mc *arrMenuMinumanCoffe, mnc *arrMenuMinumanNonCoffe, a, b, c *int) {
	var i, j, input, mode int
	var duplikat bool

	fmt.Printf("%-4c %s\n", ' ', "Cafe Menu")
	fmt.Println("-------------------")
	fmt.Println("1. Menu Makanan")
	fmt.Println("2. Menu Minuman Coffe")
	fmt.Println("3. Menu Minuman Non Coffe")
	fmt.Println("4. Kembali")
	fmt.Scan(&mode)
	fmt.Println("")

	if mode == 1 {
		i = 0
		if m[i].id == 0 {
			fmt.Println("Format: <ID> <Nama Menu> <Kategori> <Harga> <Komposisi> <Ketersediaan (true/false)> ")
			fmt.Scan(&input)
			for input != 0 {
				m[i].id = input
				fmt.Scan(&m[i].nama, &m[i].kategori, &m[i].harga, &m[i].komposisi, &m[i].ketersediaan)
				i++
				fmt.Scan(&input)
			}
			*a = i
			adminMode(m, mc, mnc, a, b, c)
		} else {
			if *a == NMAX {
				fmt.Println("Data Penuh!")
				fmt.Println("")
				adminMode(m, mc, mnc, a, b, c)
			}
			fmt.Scan(&input)
			duplikat = false
			for j = 0; j < *a; j++ {
				if m[j].id == input {
					duplikat = true
					break
				}
			}
			if duplikat {
				fmt.Println("ID sudah ada, data dilewati!")
				fmt.Println("")
				adminMode(m, mc, mnc, a, b, c)
			} else {
				m[*a].id = input
				fmt.Scan(&m[*a].nama, &m[*a].kategori, &m[*a].harga, &m[*a].komposisi, &m[*a].ketersediaan)
				*a += 1
				fmt.Println("Menu baru ditambahkan!")
				fmt.Println("")
				adminMode(m, mc, mnc, a, b, c)
			}
		}
	} else if mode == 2 {
		i = 0
		if mc[i].id == 0 {
			fmt.Println("Format: <ID> <Nama Menu> <Kategori> <Harga> <Komposisi> <Ketersediaan (true/false)> ")
			fmt.Scan(&input)
			for input != 0 {
				mc[i].id = input
				fmt.Scan(&mc[i].nama, &mc[i].kategori, &mc[i].harga, &mc[i].komposisi, &mc[i].ketersediaan)
				i++
				fmt.Scan(&input)
			}
			*b = i
			adminMode(m, mc, mnc, a, b, c)
		} else {
			if *b == NMAX {
				fmt.Println("Data Penuh!")
				fmt.Println("")
				adminMode(m, mc, mnc, a, b, c)
			}
			fmt.Scan(&input)
			duplikat = false
			for j = 0; j < *b; j++ {
				if mc[j].id == input {
					duplikat = true
					break
				}
			}
			if duplikat {
				fmt.Println("ID sudah ada, data dilewati!")
				fmt.Println("")
				adminMode(m, mc, mnc, a, b, c)
			} else {
				mc[*b].id = input
				fmt.Scan(&mc[*b].nama, &mc[*b].kategori, &mc[*b].harga, &mc[*b].komposisi, &mc[*b].ketersediaan)
				*b += 1
				fmt.Println("Menu baru ditambahkan!")
				fmt.Println("")
				adminMode(m, mc, mnc, a, b, c)
			}
		}
	} else if mode == 3 {
		i = 0
		if mnc[i].id == 0 {
			fmt.Println("Format: <ID> <Nama Menu> <Kategori> <Harga> <Komposisi> <Ketersediaan (true/false)> ")
			fmt.Scan(&input)
			for input != 0 {
				mnc[i].id = input
				fmt.Scan(&mnc[i].nama, &mnc[i].kategori, &mnc[i].harga, &mnc[i].komposisi, &mnc[i].ketersediaan)
				i++
				fmt.Scan(&input)
			}
			*c = i
			adminMode(m, mc, mnc, a, b, c)
		} else {
			if *c == NMAX {
				fmt.Println("Data Penuh!")
				fmt.Println("")
				adminMode(m, mc, mnc, a, b, c)
			}
			fmt.Scan(&input)
			duplikat = false
			for j = 0; j < *c; j++ {
				if mnc[j].id == input {
					duplikat = true
					break
				}
			}
			if duplikat {
				fmt.Println("ID sudah ada, data dilewati!")
				fmt.Println("")
				adminMode(m, mc, mnc, a, b, c)
			} else {
				mnc[*c].id = input
				fmt.Scan(&mnc[*c].nama, &mnc[*c].kategori, &mnc[*c].harga, &mnc[*c].komposisi, &mnc[*c].ketersediaan)
				*c += 1
				fmt.Println("Menu baru ditambahkan!")
				fmt.Println("")
				adminMode(m, mc, mnc, a, b, c)
			}
		}
	}
}

func lihatData(m *arrMenuMakanan, mc *arrMenuMinumanCoffe, mnc *arrMenuMinumanNonCoffe, a, b, c *int) {
	var i, mode int
	fmt.Printf("%-4c %s\n", ' ', "Cafe Menu")
	fmt.Println("-------------------")
	fmt.Println("1. Menu Makanan")
	fmt.Println("2. Menu Minuman Coffe")
	fmt.Println("3. Menu Minuman Non Coffe")
	fmt.Println("4. Kembali")
	fmt.Scan(&mode)
	fmt.Println("")

	if mode == 1 {
		for i = 0; i < *a; i++ {
			fmt.Printf("%d %s %s %d %s %v\n", m[i].id, m[i].nama, m[i].kategori, m[i].harga, m[i].komposisi, m[i].ketersediaan)
		}
		fmt.Println("")
		adminMode(m, mc, mnc, a, b, c)
	} else if mode == 2 {
		for i = 0; i < *b; i++ {
			fmt.Printf("%d %s %s %d %s %v\n", mc[i].id, mc[i].nama, mc[i].kategori, mc[i].harga, mc[i].komposisi, mc[i].ketersediaan)
		}
		fmt.Println("")
		adminMode(m, mc, mnc, a, b, c)
	} else if mode == 3 {
		for i = 0; i < *c; i++ {
			fmt.Printf("%d %s %s %d %s %v\n", mnc[i].id, mnc[i].nama, mnc[i].kategori, mnc[i].harga, mnc[i].komposisi, mnc[i].ketersediaan)
		}
		fmt.Println("")
		adminMode(m, mc, mnc, a, b, c)
	} else if mode == 4 {
		fmt.Println("")
		adminMode(m, mc, mnc, a, b, c)
	}
}

func ubahData() {

}

func hapusData() {

}

func adminMode(makan *arrMenuMakanan, minumC *arrMenuMinumanCoffe, minumNC *arrMenuMinumanNonCoffe, a *int, b *int, c *int) {
	var mode int

	fmt.Printf("%-4c %s\n", ' ', "Mode Admin")
	fmt.Println("-------------------")
	fmt.Println("1. Tambah Data")
	fmt.Println("2. Edit Data")
	fmt.Println("3. Hapus Data")
	fmt.Println("4. lihat Data")
	fmt.Println("5. Kembali")
	fmt.Scan(&mode)
	fmt.Println("")

	if mode == 1 {
		tambahData(makan, minumC, minumNC, a, b, c)
	} else if mode == 2 {

	} else if mode == 3 {

	} else if mode == 4 {
		lihatData(makan, minumC, minumNC, a, b, c)
	} else if mode == 5 {
		adminMode(makan, minumC, minumNC, a, b, c)
	}
}

func userMode() {

}

func exit() {
	fmt.Println("Terima kasih telah menggunakan layanan kami!")
}
