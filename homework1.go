package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

func main() {
	fmt.Print("1 : Merhaba Dünya Yaz\n2 : Rasgele 2 sayının karesini al\n3 : Rasgele 2 sayıyı Topla\n4 : Rasgele 2 sayıyı Çarp\n5 : Rasgele 2 sayıyı Böl")
	fmt.Println("Yapmak istediğiniz işlem numarasını giriniz....")
	var choice int
	fmt.Scanln(&choice)
	rand.Seed(time.Now().UnixNano())
	choiceAct(choice)

}

//TODO
//golang için isimlerin büyük harfle başlamasının özel bir anlamı vardır bu konu işlenecektir ++
//go dilinde en ideal kullanım biçimi camelCase yöntemidir ++
//bunu araştırıp tüm fonksiyon ve değişken isimlerini camelCase olarak yeniden yazalım ++
//proje dosyası isimlerinin sonuna .go ifadesini koyalım ++
//BAŞARILI bir çalışma

func choiceAct(cha int) {
	if cha < 1 || cha > 5 {
		fmt.Println("Yanlış Giriş Yaptınız")
	} else {
		if cha == 1 {
			hw(true)
		} else if cha == 2 {
			fmt.Printf("Karesi : %d\n", sq(rand.Intn(100)))
		} else if cha == 3 {
			fmt.Printf("Toplamı %d\n", add(rand.Intn(200), rand.Intn(200)))
		} else if cha == 4 {
			fmt.Printf("Çarpımı %d\n", multiply(rand.Intn(100), rand.Intn(100)))
		} else if cha == 5 {
			fmt.Printf("Bölümü %.2f\n", newdivide(float64(rand.Intn(200)), float64(rand.Intn(200))))
		}
	}
}

func hw(x bool) {
	xSwitch := rand.Intn(2)
	if xSwitch == 0 {
		x = !x
	}
	if !x {
		println("Merhaba Dünya")
	} else {
		println("Merhabayı kestim...")
	}

}
func sq(x int) int {
	fmt.Printf("Rasgele Sayınız %d\n", x)
	return x * x
}
func add(x, y int) int {
	fmt.Printf("1.Rasgele Sayınız %d\n", x)
	fmt.Printf("2.Rasgele Sayınız %d\n", y)
	return x + y
}
func multiply(x, y int) int {
	fmt.Printf("1.Rasgele Sayınız %d\n", x)
	fmt.Printf("2.Rasgele Sayınız %d\n", y)
	return x * y
}
func newdivide(x, y float64) float64 {
	x = math.Round((x * 100) / 100)
	y = math.Round((y * 100) / 100)
	fmt.Printf("1.Rasgele Sayınız %.0f\n", x)
	fmt.Printf("2.Rasgele Sayınız %.0f\n", y)
	if x > y {
		return x / y
	} else if y >= x {
		return y / x
	} else {
		return -1
	}
}
