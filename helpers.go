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

func deferCase() {
	// defer is really hande when it comes to close a database for avoid memory leaks
	// dbOpen := true
	// defer closedb()
}

func getArray() {
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1])

	integers := [6]int{1, 2, 3, 4, 5, 6}
	fmt.Println(integers)
}

func getSlice() {
	var a []int
	c := make([]int, 5)

	b := []any{1, 2, 3, "helloo"}

	b = append(b, "world")
	fmt.Println(a, b[0:2], c)

	myslice := make([]int, 5)
	myslice[0] = 1
	myslice[1] = 1
	myslice[2] = 1

	for i, v := range myslice {
		fmt.Printf("index %v , value %v\n", i, v)
	}
}

func getMap() {
	m := make(map[string]int)

	m1 := map[string]int{}

	m["key1"] = 42
	m1["key1"] = 32

	delete(m, "key1")

	v, ok := m["key1"]
	fmt.Println("the value:", v, "Present?", ok)

	for key, value := range m1 {
		fmt.Println(key, value)
	}
}

func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 2)
}

func annonymusfn() {
	hole := func(x, y float64) float64 {
		return x + y
	}

	fmt.Println(hole(3, 4))
	fmt.Println(compute(hole))
}
