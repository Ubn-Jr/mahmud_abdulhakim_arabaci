package main

import (
	"fmt"
	"math/rand"
	"time"
)

var rndName []string

func main() {
	rand.Seed(time.Now().UnixNano())
	fmt.Println("Meyveleri rasgele baştan sona yada sondan başa yazan programa hoşgeldiniz")
	insertFruitIntoArray(&rndName)
	for {
		arrPrinter(rndName, rand.Intn(2))
		fmt.Println("Devam etmek istiyor musunuz ? e/h")
		chc := "h"
		fmt.Scanln(&chc)
		if loopcheck(chc) {
			break
		}
	}

}

func loopcheck(chc string) bool {
	stp := false
	switch chc {
	case "e":
		fmt.Println("Program Devam Ediyor...")
		break
	case "h":
		fmt.Println("Program Sonlandırılıyor....")
		stp = true
		break
	default:
		fmt.Println("Yanlış Giriş Yaptınız Başa Dönülüyor...")
		break
	}
	return stp

}

func arrPrinter(arr []string, reverse int) {
	if len(arr) == 0 {
		fmt.Println("Dizi Boş")
	} else {
		if reverse == 1 {
			for i := len(arr) - 1; i >= 0; i-- {
				fmt.Println(arr[i])
			}
		} else {
			for i := 0; i < len(arr); i++ {
				fmt.Println(arr[i])
			}
		}

	}

}

func insertFruitIntoArray(arr *[]string) []string {
	rnm := [5]string{
		"Elma",
		"Karpuz",
		"Çilek",
		"Muz",
		"Kayısı"}

	*arr = rnm[:]
	return *arr
}
