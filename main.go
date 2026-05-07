package main

import "fmt"

const NMAX int = 999

type Menu struct {
	id           int			// id untuk hapus dan mengubah data
	nama         string			// nama menu
	kategori     string			// coffe atau noncoffe
	harga        int			// harga dalam Rp
	komposisi    string			// komposisi 
	ketersediaan bool 			// true = tersedia, false = tidak tersedia
}

type arrMenu [NMAX]Menu

func inputData(m *arrMenu, n int) {
	var i int

	fmt.Println("Format: <ID> <Nama Menu> <Kategori> <Harga> <Komposisi> <Ketersediaan (true/false)> ")

	// loop sebanyak n kali
	for i = 0; i < n; i++ {
		// scan array
		fmt.Scan(&m[i].id, &m[i].nama, &m[i].kategori, &m[i].harga, &m[i].komposisi, &m[i].ketersediaan)
	}

	fmt.Println("Menu telah diinput ke dalam Array")
}

func tambahData(m *arrMenu, n *int, add int) {
	var menuNew Menu
	var i, j int
	var duplikat bool

	fmt.Println("Format: <ID> <Nama Menu> <Kategori> <Harga> <Komposisi> <Ketersediaan (true/false)> ")

	// loop sebanyak berapa menu yang mau ditambah
	for i = 0; i < add; i++ {
		// cek n lebih dari NMAX
		if *n >= NMAX {
			fmt.Println("Data Penuh!")
			return
		}

		// scan array
		fmt.Scan(&menuNew.id, &menuNew.nama, &menuNew.kategori, &menuNew.harga, &menuNew.komposisi, &menuNew.ketersediaan)
		
		// cek duplikat
		duplikat = false

		for j = 0; j < *n; j++ {
			if m[j].id == menuNew.id {
				duplikat = true
				break
			}
		}

		// kalo ada duplikat id, data lain ga akan diinput
		if duplikat == true {
			fmt.Println("ID sudah ada, data dilewati!")
		} else {
			m[*n] = menuNew
			*n = *n + 1
		}
	}

	fmt.Println("Menu baru ditambahkan!")
}

func ubahData(){

}

func hapusData() {

}

func main() {
	var mode int
	var n int
	var dataMenu Menu

	fmt.Printf("%-4c %s\n", ' ', "Cafe Menu")
	fmt.Println("-------------------")
	fmt.Println("1. Admin")
	fmt.Println("2. Pelanggan")
	fmt.Println("3. Exit")
	fmt.Scan(&mode)

	if mode == 1 {
		adminMode(&dataMenu, &n)
	} else if mode == 2 {
		userMode()
	} else if mode == 3 {
		exit()
	}
}

func adminMode(m *Menu, n* int) {

}

func userMode() {

}

func exit() {
	fmt.Println("\nTerima kasih telah menggunakan layanan kami!")
}
