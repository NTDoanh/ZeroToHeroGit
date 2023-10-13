package main

import (
	"fmt"
)

func add(a, b int) {
	fmt.Printf("%d + %d = %d\n", a, b, a + b)
}
func main() {
	var a, b int
	fmt.Print("Nhap so thu nhat: ")
	fmt.Scanln(&a)
	fmt.Print("Nhap so thu hai: ")
	fmt.Scanln(&b)
  add(a,b)
}
