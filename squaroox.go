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
		Persegi_Panjang()
	case 4:

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
	fmt.Print("later")
}
func Persegi_Panjang() {
	fmt.Print("later")
}
