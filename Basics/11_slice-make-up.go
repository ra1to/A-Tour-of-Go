package main

import "fmt"

func main() {
	n := make([]int, 3, 5)
	fmt.Println("len=%d cap=%d value=%v", len(n), cap(n), n)
	n = append(n, 0, 0)
	fmt.Println("len=%d cap=%d value=%v", len(n), cap(n), n)
	n = append(n, 1, 2, 3, 4, 5)
	fmt.Println("len=%d cap=%d value=%v", len(n), cap(n), n)

	a := make([]int, 3)
	fmt.Println("len=%d cap=%d value=%v", len(a), cap(a), a)
	b := make([]int, 0)
	fmt.Println("len=%d cap=%d value=%v", len(b), cap(b), b)
}
