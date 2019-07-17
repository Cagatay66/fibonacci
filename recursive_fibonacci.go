package main

import "fmt"

func fibo(elemanlar []int, limit int) []int {
	if len(elemanlar) > limit {
		return elemanlar
	} else if len(elemanlar) < 2 {
		return fibo(append(elemanlar, 0+1), limit)
	} else {
		return fibo(append(elemanlar,
			elemanlar[len(elemanlar)-2]+elemanlar[len(elemanlar)-1]),
			limit)
	}
}

func main() {
	var elemanlar []int

	fmt.Println(fibo(elemanlar, 500))
}
