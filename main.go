package main

import "fmt"

func Add(a, b int) int {
	return a + b
}
func Sub(a, b int) int {
	return a - b
}
func Mul(a, b int) int {
	return a * b
}
func main() {
	fmt.Println(`Addition: `, Add(2024, 132143))
	fmt.Println(`Subtraction: `, Sub(2, 2))
	fmt.Println(`Multiplication: `, Mul(10, 30))

}
