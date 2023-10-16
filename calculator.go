package main

import (
	"fmt"
)

func add(a, b float64) {
	fmt.Printf("%.2f + %.2f = %.2f\n", a, b, a+b)
}
func multiply(a, b int) {
	fmt.Printf("%d + %d = %d\n", a, b, a*b)
}
func subtract(a, b float64) {

	fmt.Printf("%.2f - %.2f = %.2f\n", a, b, a-b)
}
func divide(a, b float64) {
	fmt.Printf("%.2f / %.2f = %.2f\n", a, b, a/b)
}
func main() {
	var a, b float64
	var c string
	fmt.Print("Nhap so thu nhat: ")
	fmt.Scanln(&a)
	fmt.Print("Nhap so thu hai: ")
	fmt.Scanln(&b)
	fmt.Print("Nhap phep tinh: ")
	fmt.Scanln(&c)
	switch c {
	case "+":
		add(a, b)
	case "-":
		subtract(a, b)
	case "*":
		multiply(a, b)
	case "/":
		divide(a, b)
	default:
		fmt.Printf("Chua co phep nay hoac nhap sai")
	}
}
