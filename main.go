package main

import "fmt"

const NMAX int = 999
type menu struct {
	id           int
	nama         string
	kategori     string
	harga        int
	komposisi    string
	ketersediaan bool //true = sedia,  false = tidak tersedia
}

type arrMenu [NMAX]menu

func inputData(m *arrMenu, n int) {
	var i int

	fmt.Println("Format: <ID> <Nama Menu> <Kategori> <Harga> <Komposisi> <Ketersediaan (true/false)> ")
	for i = 0; i < n; i++ {
		fmt.Scan(&m[i].id, &m[i].nama, &m[i].kategori, &m[i].harga, &m[i].komposisi, &m[i].ketersediaan)
	}
}

func tambahData(m *arrMenu, n *int, add int) {
	var menuNew menu
	var i int

	fmt.Println("Format: <ID> <Nama Menu> <Kategori> <Harga> <Komposisi> <Ketersediaan (true/false)> ")
	for i = 0; i < add; i++ {
		fmt.Scan(&menuNew.id, &menuNew.nama, &menuNew.kategori, &menuNew.harga, &menuNew.komposisi, &menuNew.ketersediaan)

		m[*n] = menuNew
		*n = *n + 1

		if *n >= NMAX {
			fmt.Println("Data Penuh!")
			return
		}
	}

	fmt.Println("Menu baru ditambahkan!")
}

func main() {
	var n int
	var menuData arrMenu

	fmt.Scan(&n)

	fmt.Println(menuData)
}
