package main

import "fmt"

const z = 100
const g = 43
const d = 125

func f1(a *int) {
	*a += z
	fmt.Println("Первое сложили получили ответ", *a)

}

func f2(a *int) {
	*a += g
	fmt.Println("Второе сложили с первым получили ответ", *a)

}
func f3(a *int) {
	*a += d
	fmt.Println("Третье сложили со вторым получили ответ", *a)

}

func main() {
	var a int
	f1(&a)
	f2(&a)
	f3(&a)

}
