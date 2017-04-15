package flow

import (
	"io/ioutil"
	"os"
	"strconv"
)

var std *Graph

type Graph struct {
	graph [][][][]int
}

func NewGraph() *Graph {
	g := &Graph{}
	g.Enlarge(1, 1, 1, 1)
	return g
}

func (g *Graph) Enlarge(n1, a1, n2, a2 int) {
	for n1 > len(g.graph) {
		g.graph = append(g.graph, [][][]int{})
	}

	for i := range g.graph {
		for a1 > len(g.graph[i]) {
			g.graph[i] = append(g.graph[i], [][]int{})
		}
	}
	for i := range g.graph {
		for j := range g.graph[i] {
			for n2 > len(g.graph[i][j]) {
				g.graph[i][j] = append(g.graph[i][j], []int{})
			}
		}
	}
	for i := range g.graph {
		for j := range g.graph[i] {
			for k := range g.graph[i][j] {
				for a2 > len(g.graph[i][j][k]) {
					g.graph[i][j][k] = append(g.graph[i][j][k], 0)
				}
			}
		}
	}
}

func (g *Graph) Connect(n1, a1, n2, a2 int) {
	g.Enlarge(n1+1, a1+1, n2+1, a2+1)
	g.graph[n1][a1][n2][a2] = 1
}

func Connect(n1, a1, n2, a2 int) {
	std.Connect(n1, a1, n2, a2)
}

func (g *Graph) Dims(slices ...[]interface{}) []int {
	return []int{len(g.graph), len(g.graph[0]), len(g.graph[0][0]), len(g.graph[0][0][0])}
}

func (g *Graph) Dump(path string) {
	os.Open(path)

	str := ""

	// Edges
	for n1 := range g.graph {
		for a1 := range g.graph[n1] {
			for n2 := range g.graph[n1][a1] {
				for a2 := range g.graph[n1][a1][n2] {
					if g.graph[n1][a1][n2][a2] != 0 {
						str += strconv.Itoa(n1) + "x" + strconv.Itoa(a1) + "->" + strconv.Itoa(n2) + "x" + strconv.Itoa(a2) + ";\n"
					}
				}
			}
		}
	}

	ioutil.WriteFile(path, []byte(str), 0666)
}

func Dump(path string) {
	std.Dump(path)
}

func Dims() []int {
	return std.Dims()
}

func init() {
	std = NewGraph()
}
