package main

import "fmt"

func main() {
	var a [5]int
	fmt.Println("a = ", a)
	a[4] = 100
	fmt.Println("a = ", a)
	fmt.Println("a = ", a[4])
	fmt.Println("a len = ", len(a))
	fmt.Println("============================================================================")
	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("b = ", b)
	fmt.Println("============================================================================")
	var c [2][3]int
	fmt.Println("len c = ", len(c))
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			c[i][j] = i + j
		}
	}
	fmt.Println("c = ", c)
}
