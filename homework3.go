package main

import (
	"fmt"
	"math/rand"
	"time"
)

var stdName []string
var stdSurname []string
var stdVize []float64
var stdFinal []float64
var stdVizeMult float64 = 0.4
var stdFinalMult float64 = 0.6
var stdOrt []float64
var stdDrm []bool
var gecisNotu int = 60
var maxOgrenci = 10

var rngNames []string
var rngSurnames []string

func main() {
	rand.Seed(time.Now().UnixNano())
	fmt.Println("Ögrenci Durumu Hesaplamaya ve Yazdırmaya Hoşgeldiniz !")
	fmt.Printf("Ögrenci Mevcudunu Giriniz... (Max %d)\n", maxOgrenci)
	var std int
	fmt.Scanln(&std)

	if checkStdNumb(std) {

		fmt.Println("1 : Ögrenciler için değerleri kendiniz girmek istiyorsanız...")
		fmt.Println("2 : Ögrencileri Random olarak otomatik oluşturmak istiyorsanız..")
		var chc1 int
		fmt.Scanln(&chc1)

		if chc1 == 1 {
			stdCreator(std, false)
		} else if chc1 == 2 {
			stdCreator(std, true)
		} else {
			fmt.Println("Yanlış tercih yaptınız...")
			goto FINISH
		}
	Loops:
		var chc2 int
		fmt.Println("1 : Ögrencilerin Vize Final Hesaplaması yapıp kalan ve geçenleri listelemek için \n2 : Ögrencileri Ad Soyad Listelemek için\n3 : Girilen İndeksdeki Ögrencinin Durumunu Yazdır\n4 : Girilen İndeksdeki Harici Ögrencileri listele\n5 : Sonlandırmak için")
		// TODO Ad'a soyad'a göre Arama, Ad,Soyad, Vize, Final ve Ort Notlara göre sıralama(A-Z, Z-A), push to multi dim array
		fmt.Scanln(&chc2)
		if chc2 == 5 {
			//TODO
			//burada FINISH label i yerine break veya return yada os.Exit(-1) kullanılabilir
			//goto ile label kullanımı yazılımlar için kullanışlı değildir
			goto FINISH
		}
		listStd(chc2)
		goto Loops
	} else {
		fmt.Println("Değer yanlış girildiğinden dolayı program sonlanıyor....")
	}
FINISH:
}

func listStd(chpc int) {
	if chpc <= 0 || chpc > 5 {
		fmt.Println("Yanlış Giriş Yaptınız....")
	} else {
		if chpc == 1 {
			for i := 0; i < len(stdName); i++ {
				drm := "Kaldı"
				if stdDrm[i] {
					drm = "Geçti"
				}
				fmt.Printf("%d. Ögrencinin Adı: %s, Soyadı : %s, Vize: %.1f, Final: %.1f Ortalama: %.1f Durumu: %s...\n", i+1, stdName[i], stdSurname[i], stdVize[i], stdFinal[i], stdOrt[i], drm)
			}
		}
		if chpc == 2 {
			for i := 0; i < len(stdName); i++ {
				fmt.Printf("%d. Ögrencinin Adı: %s, Soyadı : %s \n", i+1, stdName[i], stdSurname[i])
			}
		}

		if chpc == 3 {
			fmt.Println("İndeks Giriniz")
			var pckStd int
			fmt.Scanln(&pckStd)
			if !checkIfinRange(pckStd-1, 0, len(stdName)) {
				fmt.Println("Yanlış değer girdiniz..")
				return
			}
			drm := "Kaldı"
			if stdDrm[pckStd-1] {
				drm = "Geçti"
			}
			fmt.Printf("%d. Ögrencinin Adı: %s, Soyadı : %s, Vize: %.1f, Final: %.1f Ortalama: %.1f Durumu: %s...\n", pckStd, stdName[pckStd-1], stdSurname[pckStd-1], stdVize[pckStd-1], stdFinal[pckStd-1], stdOrt[pckStd-1], drm)
		}
		if chpc == 4 {
			fmt.Println("İndeks Giriniz")
			var pckStd int
			fmt.Scanln(&pckStd)
			if !checkIfinRange(pckStd-1, 0, len(stdName)) {
				fmt.Println("Yanlış değer girdiniz..")
				return
			}
			for i := 0; i < len(stdName); i++ {
				if i == pckStd-1 {
					continue
				}
				drm := "Kaldı"
				if stdDrm[i] {
					drm = "Geçti"
				}
				fmt.Printf("%d. Ögrencinin Adı: %s, Soyadı : %s, Vize: %.1f, Final: %.1f Ortalama: %.1f Durumu: %s...\n", i+1, stdName[i], stdSurname[i], stdVize[i], stdFinal[i], stdOrt[i], drm)
			}
		}

	}
}

func checkIfinRange(x, min, max int) bool {
	if x >= min && x < max {
		return true
	} else {
		return false
	}

}

func setRngNameSurname(stdSize int) {
	rngname := [10]string{
		"Ahmet",
		"Zeynep",
		"Mehmet",
		"Cansu",
		"Gürcan",
		"Süleyman",
		"Gökhan",
		"Cağdaş",
		"Meltem",
		"Nur"}
	rngsurname := [10]string{
		"Çay",
		"Çorbacı",
		"Kızılay",
		"Yıldız",
		"Güney",
		"Güneş",
		"Karakapı",
		"Kaya",
		"Temiz",
		"Yazar"}

	rngNames = rngname[:]
	rngSurnames = rngsurname[:]

}

func stdCreator(stdSize int, rngfill bool) {

	stdName = make([]string, stdSize)
	stdSurname = make([]string, stdSize)
	stdVize = make([]float64, stdSize)
	stdFinal = make([]float64, stdSize)
	stdOrt = make([]float64, stdSize)
	stdDrm = make([]bool, stdSize)

	if rngfill {
		setRngNameSurname(stdSize)
	}

	for i := 0; i < stdSize; i++ {

		if rngfill {

			stdName[i] = rngNames[rand.Intn(stdSize)]
			stdSurname[i] = rngSurnames[rand.Intn(stdSize)]
			stdVize[i] = float64(rand.Intn(100))
			stdFinal[i] = float64(rand.Intn(100))
			stdOrt[i] = stdVize[i]*stdVizeMult + stdFinal[i]*stdFinalMult
			if stdOrt[i] >= float64(gecisNotu) {
				stdDrm[i] = true
			}
			fmt.Printf("%d. Ögrencinin Adı: %s, Soyadı : %s, Vize: %.1f, Final: %.1f Şeklinde Rasgele Atandı...\n", i+1, stdName[i], stdSurname[i], stdVize[i], stdFinal[i])

		} else {

			fmt.Printf("%d. Ögrencinin Adını giriniz\n", i+1)
			fmt.Scanln(&stdName[i])

			fmt.Printf("%d. Ögrencinin Soyadını giriniz\n", i+1)
			fmt.Scanln(&stdSurname[i])

			fmt.Printf("%d. Ögrencinin Vize Notunu giriniz\n", i+1)
			fmt.Scanln(&stdVize[i])

			fmt.Printf("%d. Ögrencinin Vize Notunu giriniz\n", i+1)
			fmt.Scanln(&stdFinal[i])

			stdOrt[i] = calcMid(stdVize[i], stdFinal[i])
			stdDrm[i] = calcIfPass(stdOrt[i], gecisNotu)
			fmt.Printf("%d. Ögrencinin Adı: %s, Soyadı : %s, Vize: %.1f, Final: %.1f Şeklinde Atandı...\n", i+1, stdName[i], stdSurname[i], stdVize[i], stdFinal[i])

		}
	}
}

func checkStdNumb(x int) bool {
	if x > 0 && x <= maxOgrenci {
		return true
	} else {
		if x < 0 || x == 0 {
			fmt.Println("Öğrenci sayısı 0 yada 0'dan küçük olamaz")
		}
		if x > maxOgrenci {
			fmt.Printf("Öğrenci sayısı Max Ögrenci sayısından fazla olamaz.. Max %d\n", maxOgrenci)
		}
		return false
	}

}

func calcMid(x, y float64) float64 {
	x = x * stdVizeMult
	y = y * stdFinalMult

	return x + y
}

func calcIfPass(x float64, y int) bool {

	if int(x) >= y {
		return true
	} else {
		return false
	}

}
