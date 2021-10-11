package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	const ptnNumber int = 48
	const hand int = 5 // 手札
	const player int = 6

	var ptnIndex [ptnNumber]int

	for i := 0; i < ptnNumber; i++ {
		ptnIndex[i] = i + 1
	}

	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(ptnIndex), func(i, j int) { ptnIndex[i], ptnIndex[j] = ptnIndex[j], ptnIndex[i] })
	for i := 0; i < player; i++ {
		fmt.Printf("Player%v\n", i+1)
		for j := 0; j < hand; j++ {
			fmt.Printf("　パターン%v\n", ptnIndex[5*i+j])
		}
	}
	fmt.Println(ptnIndex)
}
