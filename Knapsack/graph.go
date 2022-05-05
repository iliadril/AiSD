package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

type Graph interface {
	Print()
	AddVertex(int)
	FindVertex(int) *Vertex
	AddEdge(int, int)
	DFS(*Vertex, []*Vertex) []*Vertex
	BFS(int) []*Vertex
	ColourFull() []*Vertex
	BFSTop() []*Vertex
}

// ELGraph Graph representation which uses Table of edges
type ELGraph struct {
	vertices []*Vertex
	edges    [][]*Vertex
}

func (g *ELGraph) Print()  {
	for i, edge := range g.edges {
		fmt.Printf("Edge index: %d:  %v -> %v\n", i, edge[0], edge[1])
	}
}

func (g *ELGraph) FindVertex(key int) *Vertex {
	for _, vertex := range g.vertices {
		if vertex.Key == key {
			return vertex
		}
	}
	return nil
}

func (g *ELGraph) AddVertex(key int)  {
	g.vertices = append(g.vertices, &Vertex{Key: key, Colour: 'w'})
}

func (g *ELGraph) AddEdge(from, to int)  {
	fromPointer := g.FindVertex(from)
	toPointer   := g.FindVertex(to)

	if fromPointer == nil {
		g.AddVertex(from)
		fromPointer = g.FindVertex(from)
	}
	if toPointer == nil {
		g.AddVertex(to)
		toPointer = g.FindVertex(to)
	}
	g.edges = append(g.edges, []*Vertex{fromPointer, toPointer})
}

func (g ELGraph) DFS(vertex *Vertex, path []*Vertex) []*Vertex {
	path = append(path, vertex)
	for _, edge := range g.edges {
		if edge[0] == vertex && !Contains(edge[1], path) {
			path = g.DFS(edge[1], path)
		}
	}
	return path
}

func (g ELGraph) BFS(key int) []*Vertex {
	vertex := g.FindVertex(key)
	path := []*Vertex{vertex}
	queue := []*Vertex{vertex}
	for len(queue) > 0 {
		vertex = queue[0]
		queue = queue[1:]
		for _, edge := range g.edges {
			if edge[0] == vertex && !Contains(edge[1], path) {
				neighbourPointer := edge[1]
				path = append(path, neighbourPointer)
				queue = append(queue, neighbourPointer)
			}
		}
	}
	return path
}

func (g ELGraph) colourTop(key int, sorted []*Vertex) []*Vertex {
	vertex := g.FindVertex(key)
	vertex.Colour = 'g'
	for _, edge := range g.edges {
		if edge[0] == vertex {
			if edge[1].Colour == 'g' && !Contains(edge[1], sorted) {
				log.Fatal("graf zawiera cykl")
			} else if edge[1].Colour == 'w' {
				sorted = g.colourTop(edge[1].Key, sorted)
			}
		}
	}
	vertex.Colour = 'b'
	sorted = prependVertex(sorted, vertex)
	return sorted
}

func (g ELGraph) ColourFull() []*Vertex {
	var result []*Vertex
	for _, vertex := range g.vertices {
		vertex.Colour = 'w'
	}
	for _, vertex := range g.vertices {
		if vertex.Colour == 'w' {
			result = g.colourTop(vertex.Key, result)
		}
	}
	return result
}

func (g ELGraph) CountInbounding(key int) int {
	sum := 0
	for _, edge := range g.edges {
		if edge[1].Key == key {
			sum++
		}
	}
	return sum
}

func (g ELGraph) BFSTop() []*Vertex {
	inDegree := make([]int, len(g.vertices))
	for i := range inDegree {
		inDegree[i] = g.CountInbounding(i)
	}
	//fmt.Println(inDegree)
	var sorted []*Vertex
	i := 0
	for i < len(g.vertices) {
		//fmt.Println(inDegree)
		if inDegree[i] == 0 {
			inDegree[i]--
			vertex := g.FindVertex(i)
			sorted = append(sorted, vertex)

			for _, edge := range g.edges {
				if edge[0] == vertex {
					inDegree[edge[1].Key]--
				}
			}
			i = 0
			continue
		}
		i++
	}

	for _, vertex := range g.vertices {
		if !Contains(vertex, sorted) {
			log.Fatal("graf zawiera cykl")
		}
	}
	return sorted
}


// ALGraph Graph representation which uses Adjacency List to store edges
type ALGraph struct {
	vertices []*Vertex
	adjList  [][]*Vertex
}

func (g *ALGraph) Print() {
	for i, neighbours := range g.adjList {
		fmt.Printf("%d: %v\n", i, neighbours)
	}
}

func (g *ALGraph) AddVertex(key int)  {
	g.vertices = append(g.vertices, &Vertex{Key: key, Colour: 'w'})
}

func (g *ALGraph) FindVertex(key int) *Vertex {
	for _, vertex := range g.vertices {
		if vertex.Key == key {
			return vertex
		}
	}
	return nil
}

func (g *ALGraph) AddEdge(from, to int) {
	toVertex := g.FindVertex(to)
	g.adjList[from] = append(g.adjList[from], toVertex)
}

func (g ALGraph) DFS(vertex *Vertex, path []*Vertex) []*Vertex {
	path = append(path, vertex)
	for _, neighbour := range g.adjList[vertex.Key] {
		if !Contains(neighbour, path) {
			path = g.DFS(neighbour, path)
		}
	}
	return path
}

func (g ALGraph) BFS(key int) []*Vertex {
	vertex := g.FindVertex(key)
	path := []*Vertex{vertex}
	queue := []*Vertex{vertex}
	for len(queue) > 0 {
		vertex = queue[0]
		queue = queue[1:]
		for _, neighbour := range g.adjList[vertex.Key] {
			if !Contains(neighbour, path) {
				path = append(path, neighbour)
				queue = append(queue, neighbour)
			}
		}
	}
	return path
}

func (g ALGraph) colourTop(vertex *Vertex, sorted []*Vertex) []*Vertex {
	vertex.Colour = 'g'
	for _, neighbour := range g.adjList[vertex.Key] {
		if neighbour.Colour == 'g' && !Contains(neighbour, sorted) {
			log.Fatal("graf zawiera cykl")
		} else if neighbour.Colour == 'w' {
			sorted = g.colourTop(neighbour, sorted)
		}
	}
	vertex.Colour = 'b'
	sorted = prependVertex(sorted, vertex)
	return sorted
}

func (g ALGraph) ColourFull() []*Vertex {
	var result []*Vertex
	for _, vertex := range g.vertices {
		vertex.Colour = 'w'
	}
	for _, vertex := range g.vertices {
		if vertex.Colour == 'w' {
			result = g.colourTop(vertex, result)
		}
	}
	return result
}

func (g ALGraph) CountInbounding(key int) int {
	sum := 0
	for _, vertices := range g.adjList {
		for _, vertex := range vertices {
			if vertex.Key == key {
				sum++
			}
		}
	}
	return sum
}

func (g ALGraph) BFSTop() []*Vertex {
	inDegree := make([]int, len(g.vertices))
	for i := range inDegree {
		inDegree[i] = g.CountInbounding(i)
	}
	var sorted []*Vertex
	i := 0
	for i < len(g.vertices) {
		if inDegree[i] == 0 {
			inDegree[i]--
			vertex := g.FindVertex(i)
			sorted = append(sorted, vertex)
			for _, neighbour := range g.adjList[i] {
				inDegree[neighbour.Key]--
			}
			i = 0
			continue
		}
		i++
	}
	for _, vertex := range g.vertices {
		if !Contains(vertex, sorted) {
			log.Fatal("graf zawiera cykl")
		}
	}
	return sorted
}

// AMGraph Graph representation which uses Adjacency Matrix to store edges
type AMGraph struct {
	vertices  []*Vertex
	adjMatrix [][]int
}

func (g *AMGraph) Print() {
	fmt.Print(" |")
	for i := range g.adjMatrix {
		fmt.Print(i)
	}
	fmt.Println("\n-+"+strings.Repeat("-", len(g.adjMatrix)))
	for i, matrix := range g.adjMatrix {
		fmt.Print(i, "|")
		for _, val := range matrix {
			fmt.Print(val)
		}
		fmt.Println()
	}
}

func (g *AMGraph) AddVertex(key int)  {
	g.vertices = append(g.vertices, &Vertex{Key: key, Colour: 'w'})
}

func (g *AMGraph) FindVertex(key int) *Vertex {
	for _, vertex := range g.vertices {
		if vertex.Key == key {
			return vertex
		}
	}
	return nil
}

func (g *AMGraph) AddEdge(from, to int)  {
	g.adjMatrix[from][to] = 1
}

func (g AMGraph) DFS(vertex *Vertex, path []*Vertex) []*Vertex {
	path = append(path, vertex)
	for i, neighbour := range g.adjMatrix[vertex.Key] {
		if neighbour == 1 && !Contains(g.FindVertex(i), path) {
			path = g.DFS(g.FindVertex(i), path)
		}
	}
	return path
}

func (g AMGraph) BFS(key int) []*Vertex {
	vertex := g.FindVertex(key)
	path := []*Vertex{vertex}
	queue := []*Vertex{vertex}
	for len(queue) > 0 {
		vertex = queue[0]
		queue = queue[1:]
		for i, neighbour := range g.adjMatrix[vertex.Key] {
			if neighbour == 1 && !Contains(g.FindVertex(i), path) {
				neighbourPointer := g.FindVertex(i)
				path = append(path, neighbourPointer)
				queue = append(queue, neighbourPointer)
			}
		}
	}
	return path
}

func (g AMGraph) colourTop(key int, sorted []*Vertex) []*Vertex {
	vertex := g.FindVertex(key)
	vertex.Colour = 'g'
	for i, neighbour := range g.adjMatrix[vertex.Key] {
		if neighbour == 1 {
			successor := g.FindVertex(i)
			if successor.Colour == 'g' && !Contains(successor, sorted) {
				log.Fatal("graf zawiera cykl")
			} else if successor.Colour == 'w' {
				sorted = g.colourTop(successor.Key, sorted)
			}
		}
	}
	vertex.Colour = 'b'
	sorted = prependVertex(sorted, vertex)
	return sorted
}

func (g AMGraph) ColourFull() []*Vertex {
	var result []*Vertex
	for _, vertex := range g.vertices {
		vertex.Colour = 'w'
	}
	for i, vertex := range g.vertices {
		if vertex.Colour == 'w' {
			result = g.colourTop(i, result)
		}
	}

	return result
}

func (g AMGraph) CountInbounding(key int) int {
	sum := 0
	for _, neighbours := range g.adjMatrix {
		if neighbours[key] == 1 {
			sum++
		}
	}
	return sum
}

func (g AMGraph) BFSTop() []*Vertex {
	inDegree := make([]int, len(g.vertices))
	for i := range inDegree {
		inDegree[i] = g.CountInbounding(i)
	}
	var sorted []*Vertex
	i := 0
	for  i < len(g.vertices) {
		if inDegree[i] == 0 {
			inDegree[i]--
			vertex := g.FindVertex(i)
			sorted = append(sorted, vertex)

			for i2, neighbour := range g.adjMatrix[i] {
				if neighbour == 1 {
					inDegree[i2]--
				}
			}
			i = 0
			continue
		}
		i++
	}
	for _, vertex := range g.vertices {
		if !Contains(vertex, sorted) {
			log.Fatal("graf zawiera cykl")
		}
	}
	return sorted
}


type Vertex struct {
	Key    int
	Colour rune
}

func (v *Vertex) String() string {
	return fmt.Sprintf("%d", v.Key)
}

// Functions connected with graphs

func MakeAMGraph(verticesCount int) *AMGraph {
	graph := AMGraph{adjMatrix: make([][]int, verticesCount)}
	for i := 0; i < verticesCount; i++ {
		graph.adjMatrix[i] = make([]int, verticesCount)
		graph.AddVertex(i)
	}
	return &graph
}

func MakeALGraph(verticesCount int) *ALGraph {
	graph := ALGraph{adjList: make([][]*Vertex, verticesCount)}
	for i := 0; i < verticesCount; i++ {
		graph.AddVertex(i)
	}
	return &graph
}

func MakeElGraph() *ELGraph {
	return &ELGraph{edges: nil}
}

func Contains(a *Vertex, list []*Vertex) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func prependVertex(list []*Vertex, vertex *Vertex) []*Vertex {
	list = append(list, &Vertex{})
	copy(list[1:], list)
	list[0] = vertex
	return list
}

// Functions connected with testing

// ReadInts reads integers in file and outputs slice of them as well as error
func ReadInts(r io.Reader) ([]int, error) {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanWords)
	var result []int

	for scanner.Scan() {
		x, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return result, err
		}
		result = append(result, x)
	}
	return result, scanner.Err()
}

func ParseFile(filename string) (int, int, [][]int, error) {
	file, err := os.Open(filename)
	if err != nil {
		return 0, 0, nil, err
	}
	l, err := ReadInts(file)
	if err != nil {
		return 0, 0, nil, err
	}
	verticesCount := l[0]
	edgesCount := l[1]
	var edges [][]int
	l = l[2:]
	for len(l) > 0 {
		from := l[0] - 1
		to := l[1] - 1
		edges = append(edges, []int{from, to})
		l = l[2:]
	}
	return verticesCount, edgesCount, edges, nil
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func AddEdges(graph Graph, verticesCount int, edges [][]int) {
	for i := 0; i < verticesCount; i++ {
		if graph.FindVertex(i) == nil {
			graph.AddVertex(i)
		}

	}
	for _, edge := range edges {
		graph.AddEdge(edge[0], edge[1])
	}
}

func Test(graph Graph)  {
	start := time.Now()
	graph.DFS(graph.FindVertex(0), []*Vertex{})
	elapsed := time.Since(start)
	fmt.Printf("DFS took = %s \n", elapsed)

	start = time.Now()
	graph.BFS(0)
	elapsed = time.Since(start)
	fmt.Printf("BFS took = %s \n", elapsed)

	start = time.Now()
	graph.ColourFull()
	elapsed = time.Since(start)
	fmt.Printf("DFS topology sort took = %s \n", elapsed)

	start = time.Now()
	graph.BFSTop()
	elapsed = time.Since(start)
	fmt.Printf("BFS topology sort took = %s \n", elapsed)
}

func main()  {
	verticesCount, _, edgesList, err := ParseFile("data.txt")
	check(err)

	var amGraph Graph = MakeAMGraph(verticesCount)
	var alGraph Graph = MakeALGraph(verticesCount)
	var elGraph Graph = MakeElGraph()

	AddEdges(amGraph, verticesCount, edgesList)
	fmt.Println("Adjency Matrix Graph created")
	AddEdges(elGraph, verticesCount, edgesList)
	fmt.Println("Adjency List Graph created")
	AddEdges(alGraph, verticesCount, edgesList)
	fmt.Println("Edges List Graph created")

	fmt.Println("\nAdjacency Matrix Graph Testing:")
	Test(amGraph)

	fmt.Println("\nEdge List Graph Testing:")
	Test(elGraph)

	fmt.Println("\nAdjacency List Graph Testing:")
	Test(alGraph)

}
