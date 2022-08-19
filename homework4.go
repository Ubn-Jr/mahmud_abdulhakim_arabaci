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
		rng := rand.Intn(1)
		rbg := false
		if rng == 1 {
			rbg = true
		} else {
			rbg = false
		}
		arrPrinter(rndName, rbg)
		fmt.Println("Devam etmek istiyor musunuz ? e/h")
		chc := "h"
		fmt.Scanf(chc)
		if chc == "e" || chc == "h" {
			if chc == "e" {
				fmt.Println("Program Devam Ediyor...")
			} else {
				fmt.Println("Program Sonlandırılıyor....")
				break
			}
		} else {
			fmt.Println("Yanlış Giriş Yaptınız Başa Dönülüyor...")
		}

	}

}

func arrPrinter(arr []string, reverse bool) {
	if len(arr) == 0 {
		fmt.Println("Dizi Boş")
	} else {
		if reverse {
			for i := len(arr) - 1; i == 0; i-- {
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
