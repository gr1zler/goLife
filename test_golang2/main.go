package main

import (
	"fmt"
	"time"

	"golang.org/x/exp/constraints"
)

func addd[T constraints.Ordered](a, b T) T {
	return a + b
}
func main() {
	p := people{"dff"}
	p = p.do("err")
	fmt.Println(p.first)
	fmt.Println(addd(31.2, 33.3))

}
func doMath(a int, b int, f func(int, int) int) int {
	return f(a, b)
}
func add(a int, b int) int {
	return a + b
}
func subtract(a int, b int) int {
	return a - b
}

func iste() {
	for i := 0; i <= 1e8; i++ {
		if 1e8 == i {
			fmt.Println(i)
		}
	}
}

type people struct {
	first string
}

func (p people) do(s string) people {
	p.first = s
	return p
}
func timefunc(f func()) {
	start := time.Now()
	//iste()
	duration := time.Since(start)
	fmt.Print(duration)
}
