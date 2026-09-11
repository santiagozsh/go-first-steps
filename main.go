package main

import "fmt"

// this is the only way to initialize outside a func
// var (
// 	name string = "santiago"
// 	py   bool
// )

// struct arregment collects fields with different data types
type Vertex struct {
	X int
	Y int
	// common way to use is for parsing data from json into our struct, capital letter at first for given permission to be used outside the pacakge
	// Post string `json:"post_id"`
	// this cannot be used outside the pacakge1
	// getX int
}

type Money string

// always use the same patter with * or without it for prevent bugs
func (v *Vertex) Abs() int {
	return v.X + v.Y
}

func (m Money) ItsUsd() bool {
	return m == "USD"
}

func main() {
	v := Vertex{2, 3}
	fmt.Println(v.Abs())
}
