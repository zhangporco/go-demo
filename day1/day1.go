package day1

import "fmt"

func AddConst() {
	const a = 212.0
	const b = 120
	fmt.Println("a + b = ", a + b)
}

func Var() {
	var b, f, s = true, 2.3, "four"
	fmt.Println(b, f, s)
}

func Var2() {
	a, b := 1, "2"
	fmt.Println(a, b)
}

func Point() {
	x := 1
	p := &x
	fmt.Println(p, *p)
	changePoint(p)
	fmt.Println(p, *p)
}
func changePoint(point *int) {
	*point = 5
}

// 最大公约数
func Gcd(x, y int) int {
	for y != 0 {
		x, y = y, x%y
	}
	return x
}

// 斐波那契数列 n 位数
func Fib(n int) int {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		x, y = y, x+y
	}
	return x
}

func Trans(a *int, b *int) {
	*a, *b = *b, *a
}
