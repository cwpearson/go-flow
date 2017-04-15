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
	// Hook defines a function that is called on each task output
	t1 := flow.NewGoTask(incSource, []string{"x"}, []string{"y"})
	t2 := flow.NewGoTask(decSource, []string{"x"}, []string{"y"})
	t3 := flow.NewGoTask(mulSource, []string{"x", "y"}, []string{"z"})
	t3.Merge(t1, t2).Hook(Print)

	// dump the connections
	flow.Dump("graph.dot")

	// Inject inputs into the tasks
	t1.InputInt(a)
	t2.InputInt(b)

	// Wait for completion
	t3.Wait()
}
