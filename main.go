package main

import "fmt"

type menu struct {
	id           int
	nama         string
	kategori     string
	harga        int
	komposisi    string
	ketersediaan bool //true = sedia,  false = tidak tersedia
}

type arrMenu [999]menu

func main() {
	var n int
	fmt.Scan(&n)
}
