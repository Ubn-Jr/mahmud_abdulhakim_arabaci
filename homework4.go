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

//Prints given array by int value,
//if given 0 prints from frist element to last
//if given 1 prints from last to first element
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

//Throw in a string array then array will have 5 fruits inserted into it
func insertFruitIntoArray(arr *[]string) {
	rnm := [5]string{
		"Elma",
		"Karpuz",
		"Çilek",
		"Muz",
		"Kayısı"}

	*arr = rnm[:]
}
//Todo delete add fruits, order fruits by name
