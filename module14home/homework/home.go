package main

import (
	"fmt"
	"math/rand"
	"time"
)

func result(x int) {

	if x%2 == 0 {
		fmt.Println("Число четное ", true)
	} else {
		fmt.Println("число не четное", false)
		return
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())
	x := rand.Intn(999999)

	result(x)
}
