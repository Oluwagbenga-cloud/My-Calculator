package main

import (
	"fmt"
	"github.com/oluwaengacloud/Calculator"
)

func main() {
	fmt.Println(Calculator.Calculate("2*38*4*1", "10/5/2/0", "2+3+4", "-10-5-6-1"))
	fmt.Println(Calculator.Multiplication([]float64{2, 38, 4, 1}))
	fmt.Println(Calculator.Division([]float64{25, 5, 2}))
	fmt.Println(Calculator.Addition([]float64{2, 3, 4, 1}))
	fmt.Println(Calculator.Subtraction([]float64{0, 8, 4, 2}))
}
