package main

import "time"

func main() {
	score := 30
	if score >= 30 {
		println("Score=", score)
	} else if score >= 20 {
		println("Score=", score)
	}
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		println("周末来了")
	default:
		println("搬砖搬砖")
	}

	t := time.Now()

	switch {
	case t.Hour() < 12:
		println("上午")
	default:
		println("下午")
	}
}
