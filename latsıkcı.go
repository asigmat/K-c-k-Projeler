package main

import (
	"fmt"
)

func main() {
	var alers int
	fmt.Println("selamlar ne tür bir araç için lastik değişimi istemiştiniz araba=1 uçak=2 bisiklet=3 ?")
	_, err := fmt.Scanln(&alers) // & işareti ile değişkenin adresini veriyoruz
	if err != nil {
		fmt.Println("Geçersiz giriş! Lütfen sayı girin.")
		return
	}
	switch alers {
	case 1:
		Car()
	case 2:
		Bike()
	case 3:
		Fly()
	}

}
func Car() {
	var lastik int
	var mevsim string
	fmt.Println("Selam Aracınızda kaç tane lastik değiştiriceksiniz 1-2-3-4 ve kış lastiği mi yaz mı")
	_, err := fmt.Scanln(&lastik, &mevsim)
	if err != nil {
		fmt.Println("geçersiz bir seçim yaptınız")
	}
	var onelast int = lastik * 20

	if mevsim == "yaz" {
		fmt.Println("seçiminiz değelendirilid ve ek ücret talep etmiyoruz")
		println(onelast)
	}
	if mevsim == "kış" {

		var kıslast int = 20
		var sonuclsat int = (lastik * kıslast) + onelast
		fmt.Println(sonuclsat)
	}

}
func Bike() {
	var last int
	fmt.Println("Selam Aracınızda kaç tane lastik değiştiriceksiniz 1-2")
	_, err := fmt.Scanln(&last)
	if err != nil {
		fmt.Println("geçersiz bir seçim yaptınız")
	}
	var sonuc int = last * 10
	fmt.Println("Ücretiniz", sonuc)
}
func Fly() {
	var lastİ int
	fmt.Println("Selam Uçağınızda  kaç tane lastik değiştiriceksiniz")
	_, err := fmt.Scanln(&lastİ)
	if err != nil {
		fmt.Println("geçersiz bir seçim yaptınız")
	}
	var sonuc int = lastİ * 10
	fmt.Println("Ücretiniz", sonuc)

}
