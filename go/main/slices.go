package main

import "fmt"

func main() {
	s := make([]string, 3)
	s[0] = "H"
	s[1] = "e"
	s[2] = "l"
	// fmt.Print("this is :", s)
	s = append(s, "l")
	s = append(s, "o")
	fmt.Print("this is :", s)
}
