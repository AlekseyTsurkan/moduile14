package main

import (
	"fmt"
	"math/rand"
	"time"
)

func miltiply(x int) int {

	x *= 10
	fmt.Println("Первое действие мы умножили и получили", x)
	return (x)
}

func plus(x int) int {

	x += 10
	fmt.Println("Второе действие прибавили и получили ", x)
	return (x)
}

func main() {

	rand.Seed(time.Now().UnixNano())
	x := rand.Intn(10)

	fmt.Println("Рандомное число равно", x)

	x = miltiply(x)
	x = plus(x)

}
