package main

import "fmt"

// 無名関数・クロージャ
// 状態を持つ関数を作る
func store() func(int) int {
	var x int
	return func(n int) int {
		x += n
		return x
	}
}

func main() {
	s1 := store()
	s2 := store()
	fmt.Printf("s1: %d, s2 : %d\n", s1(1), s2(2))
	fmt.Printf("s1: %d, s2 : %d\n", s1(1), s2(2))
	fmt.Printf("s1: %d, s2 : %d\n", s1(1), s2(2))
}
