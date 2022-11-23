package main

import (
	"fmt"
	"math/rand"
	"time"
)

func getnumber(x, y int) {
	x = 2*x + 10
	y = -3*y - 5
	fmt.Println("Ответ x:", x, "Ответ y:", y)
}

func main() {

	rand.Seed(time.Now().UnixNano())
	x := rand.Intn(10)
	y := rand.Intn(11)

	fmt.Println("Рандомные числа x и y равны", x, y)
	getnumber(x, y)
}
