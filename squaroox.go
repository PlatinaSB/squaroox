package main

import (
	"fmt"
	"time"
)

func main() {
	segi := [4]string{"Persegi", "Lingkaran", "Segitiga", "Persegi Panjang"}
	for i, value := range segi {
		fmt.Printf("%d. %s\n", i+1, value)
	}
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

	default:
		fmt.Println("Pilihan salah")
		time.Sleep(5 * time.Second)
		main()
	}
}

func Persegi() {
	fmt.Println("Masukan panjang")
	var panjang int
	fmt.Scan(&panjang)
	luas := panjang * panjang
	fmt.Printf("luas Persegi adalah: %d\n", luas)
	main()
}

func Lingkaran() {
	fmt.Println("Masukan radius")
	var rads float64
	fmt.Scan(&rads)
	luas := 3.14 * rads * rads
	fmt.Printf("luas Lingkaran adalah: %f\n", luas)
	main()
}

func Segitiga() {
	var alas, tinggi float64
	fmt.Print("Masukkan alas : ")
	fmt.Scan(&alas)
	fmt.Print("Masukkan tinggi : ")
	fmt.Scan(&tinggi)
	luas := alas * tinggi * 0.5
	fmt.Printf("luas Segitiga adalah : %.2f\n", luas)
	main()
}

func Persegi_Panjang() {
	fmt.Println("Masukkan panjang:")
	var panjang float64
	fmt.Scan(&panjang)
	fmt.Println("Masukkan lebar:")
	var lebar float64
	fmt.Scan(&lebar)
	luas := panjang * lebar
	fmt.Printf("Luas Persegi Panjang adalah: %f\n", luas)
	main()
}
