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

		if mode == 1 {
			adminMode(&Menu, &a)
		} else if mode == 2 {
			userMode()
		} else if mode == 3 {
			exit()
			break
		}
	}

}

func tambahData(m *arrMenu, a *int) {
	var i, j, input int
	var duplikat bool

	i = 0
	if m[i].id == 0 {
		fmt.Println("ID NamaMenu Kategori Harga Komposisi Ketersediaan(true/false)")
		fmt.Scan(&input)
		for input != 0 {
			m[i].id = input
			fmt.Scan(&m[i].nama, &m[i].kategori, &m[i].harga, &m[i].komposisi, &m[i].ketersediaan)
			i++
			fmt.Scan(&input)
		}
		*a = i
		return
	} else {
		if *a == NMAX {
			fmt.Println("Data Penuh!")
			fmt.Println("")
			return
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
			var nama, kategori, komposisi string
			var harga int
			var tersedia bool
			fmt.Scan(&nama, &kategori, &harga, &komposisi, &tersedia)
			fmt.Println("ID sudah ada, data dilewati!")
			fmt.Println("")
			return
		} else {
			m[*a].id = input
			fmt.Scan(&m[*a].nama, &m[*a].kategori, &m[*a].harga, &m[*a].komposisi, &m[*a].ketersediaan)
			*a += 1
			fmt.Println("Menu baru ditambahkan!")
			fmt.Println("")
			return
		}
	}
}

func lihatData(m *arrMenu, a *int) {
	var i int

	for i = 0; i < *a; i++ {
		fmt.Printf("%d %s %s %d %s %v\n", m[i].id, m[i].nama, m[i].kategori, m[i].harga, m[i].komposisi, m[i].ketersediaan)
	}
	fmt.Println("")
	return
}

func ubahData() {

}

func hapusData() {

}

func adminMode(menu *arrMenu, a *int) {
	var mode int
	for {
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
			tambahData(menu, a)
		} else if mode == 2 {

		} else if mode == 3 {

		} else if mode == 4 {
			lihatData(menu, a)
		} else if mode == 5 {
			return
		} else {
			fmt.Println("Pilihan tidak valid!")
		}
	}
}

func userMode() {

}

func exit() {
	fmt.Println("Terima kasih telah menggunakan layanan kami!")
}
