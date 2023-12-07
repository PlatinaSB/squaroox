package main

import (
	"fmt"

	"github.com/inancgumus/screen"
)

func main() {
	segi := [4]string{"Persegi", "Lingkaran", "Segitiga", "Persegi Panjang"}
	for i, value := range segi {
		fmt.Printf("%d. %s\n", i+1, value)
	}
	fmt.Println("0. tutup program")

	var pil int
	fmt.Scan(&pil)
	switch pil {
	case 1:
		Persegi()
	case 2:
		Lingkaran()
	case 3:
		Segitiga()
	case 4:
		Persegi_Panjang()
	case 0:
		screen.Clear()
		fmt.Println("Program selesai")
		return

	default:
		screen.Clear()
		fmt.Println("Pilihan salah")
		main()
	}
}

func Persegi() {
	screen.Clear()
	fmt.Println("Masukan panjang")
	var panjang int
	fmt.Scan(&panjang)
	luas := panjang * panjang
	screen.Clear()
	fmt.Printf("luas Persegi adalah: %d\n", luas)
	lanjut()
}

func Lingkaran() {
	screen.Clear()
	fmt.Println("Masukan radius")
	var rads float64
	fmt.Scan(&rads)
	luas := 3.14 * rads * rads
	screen.Clear()
	fmt.Printf("luas Lingkaran adalah: %f\n", luas)
	lanjut()
}

func Segitiga() {
	screen.Clear()
	var alas, tinggi float64
	fmt.Print("Masukkan alas : ")
	fmt.Scan(&alas)
	fmt.Print("Masukkan tinggi : ")
	fmt.Scan(&tinggi)
	luas := alas * tinggi * 0.5
	screen.Clear()
	fmt.Printf("luas Segitiga adalah : %.2f\n", luas)
	lanjut()
}

func Persegi_Panjang() {
	screen.Clear()
	fmt.Println("Masukkan panjang:")
	var panjang float64
	fmt.Scan(&panjang)
	fmt.Println("Masukkan lebar:")
	var lebar float64
	fmt.Scan(&lebar)
	luas := panjang * lebar
	screen.Clear()
	fmt.Printf("Luas Persegi Panjang adalah: %f\n", luas)
	lanjut()
}

func lanjut() {
	println("lanjut Y/n")
	com := ""
	fmt.Scan(&com)

	if com == "Y" || com == "y" {
		screen.Clear()
		main()
	}
	if com == "N" || com == "n" {
		screen.Clear()
		fmt.Println("Program selesai")
		return
	}

}
