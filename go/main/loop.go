package main

func main() {
	i := 1
	for i <= 3 {
		println(i)
		i += 1
	}
	println("----------------------------------------------------------------")
	for j := 0; j <= 5; j++ {
		println(j)
	}
	println("----------------------------------------------------------------")
	for {
		println("for")
		break
	}
}
