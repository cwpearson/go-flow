package main

import (
	"fmt"

	"github.com/cwpearson/go-flow"
)

const incSource = `
func Inc(x int) int {
	y := x + 1
	return y
}`

const decSource = `
func Dec(x int) int {
	y := x - 1
	return y
}`

const mulSource = `
func Mul(x, y int) int {
	z := x * y
	return z
}`

func Inc(x int) int {
	y := x + 1
	return y
}

func Dec(x int) int {
	y := x - 1
	return y
}

func Mul(x, y int) int {
	z := x * y
	return z
}

func Print(i interface{}) {
	fmt.Println(i)
}

func main() {
	a := 1
	b := 2

	// Define tasks
	t1 := flow.NewGoTask(incSource, []string{"x"}, []string{"y"})
	t2 := flow.NewGoTask(decSource, []string{"x"}, []string{"y"})
	t3 := flow.NewGoTask(mulSource, []string{"x", "y"}, []string{"z"})

	n1 := t1

	t3.Merge(n1, t2).Hook(Print)

	t1.InputInt(a)
	t2.InputInt(b)

	t3.Wait()

}
