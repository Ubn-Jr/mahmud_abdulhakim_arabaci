package main

import "fmt"

func main() {
	fmt.Println("Final Vize Hesaplamaya Hoşgeldiniz !")
	var GecisNotu int = 60
	var Vize, Final float64

	fmt.Println("Vize Notunuzu Giriniz")
	fmt.Scanln(&Vize)

	fmt.Println("Final Notunuzu Giriniz")
	fmt.Scanln(&Final)

	if Vize < 0 || Vize > 100 {
		fmt.Printf("Vize Notunuz 100'den büyük yada 0'dan küçük olamaz... Girdiginiz değer %d\n", int(Vize))
	} else if Final < 0 || Final > 100 {
		fmt.Printf("Final Notunuz 100'den büyük yada 0'dan küçük olamaz... Girdiginiz değer %d\n", int(Final))
	} else {
		if CalcIfPass(CalcMid(Vize, Final), GecisNotu) {
			fmt.Printf("Tebrikler ! Dersi Geçtiniz Notunuz %.1f", CalcMid(Vize, Final))
		} else {
			fmt.Printf("Malesef Dersten Kaldınız... Notunuz %.1f", CalcMid(Vize, Final))
		}
	}

}

func CalcMid(x, y float64) float64 {
	var VizeMult float64 = 0.4
	var FinalMult float64 = 0.6
	x = x * VizeMult
	y = y * FinalMult

	return x + y
}

func CalcIfPass(x float64, y int) bool {

	if int(x) >= y {
		return true
	} else {
		return false
	}

}
