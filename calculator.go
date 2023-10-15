package main

import (
	"fmt"
)

func add(a, b int) {
	fmt.Printf("%d + %d = %d\n", a, b, a+b)
}
func subtract(a, b int) {
	fmt.Printf("%d - %d = %d\n", a, b, a-b)
}

func main() {
	var a, b int
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
	default:
		fmt.Printf("Chua co phep nay hoac nhap sai")
	}
}
