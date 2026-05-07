package main

import "fmt"

type menu struct {
	id           int
	nama         string
	kategori     string
	harga        int
	komposisi    string
	ketersediaan bool
}

type arrMenu [999]menu

func main() {
	var mode int
	fmt.Printf("%-4c %s\n", ' ', "Cafe Menu")
	fmt.Println("-------------------")
	fmt.Println("1. Admin")
	fmt.Println("2. Pelanggan")
	fmt.Println("3. Exit")
	fmt.Scan(&mode)
	if mode == 1 {
		adminMode()
	} else if mode == 2 {
		UserMode()
	} else if mode == 3 {
		Exit()
	}
}

func adminMode() {

}

func UserMode() {

}

func Exit() {
	fmt.Println("")
	fmt.Println("Terima kasih telah menggunakan layanan kami!")
}
