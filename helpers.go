package main

import (
	"fmt"
	"runtime"
	"time"
)

func sum(x int, y int) int {
	return x + y
}

// you can return 2 values from a func
func swap(x, y string) (string, string) {
	return y, x
}

// dont use this type of implicitr return, difficult to read, track
func split(num int) (x, y int) {
	x = num * 4 / 3
	y = num - 3
	return
}

func loops() {
	sum := 0

	// for loop, breaks when the condition is false
	for i := 0; i < 10; i++ {
		sum += i
	}

	// two ways of infinite loops
	for sum < 1000 {
		sum += sum
	}

	for {
		if sum < 100 {
			break
		}
		sum += sum
	}

	fmt.Println(sum)
}

// On if, switch, conditions you can declared vars that you can access within that scope, not outside
func switchcase() {
	os := runtime.GOOS
	switch os {
	case "linux":
		fmt.Println("Linux sistem")
	case "darwin":
		fmt.Println("mac os x")
	default:
		fmt.Println(os)
	}
	// you can do a if else with switch too
	t := time.Now()
	switch {
	// if
	case t.Hour() < 12:
		fmt.Println("Good morning")
	// else if
	case t.Hour() < 17:
		fmt.Println("Good afternoon")
	default:
		fmt.Println("Good evening.")
	}
}
