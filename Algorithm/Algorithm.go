package main

import (
	"fmt"
	"os"
)

func readMaze(filename string) [][]int {
	var row, col int
	file, err := os.Open(filename)
	defer file.Close()
	if err != nil {
		panic(err)
	}
	fmt.Fscanf(file, "%d %d", &row, &col)
	maze := make([][]int, row)
	for i := range maze {
		maze[i] = make([]int, col)
		for j := range maze[i] {
			fmt.Fscanf(file, "%d", &maze[i][j])
		}
	}
	return maze
}

type point struct {
	i, j int
}

//上左下右
var directions = [4]point{
	{-1, 0},
	{0, -1},
	{1, 0},
	{0, 1},
}

//点相加
func (p point) add(r point) point {
	return point{p.i + r.i, p.j + r.j}
}

//这个点是不是墙或者已经走过的或者越界
func (p point) at(grid [][]int) (int, bool) {
	if p.i < 0 || p.i >= len(grid) {
		return 0, false
	}
	if p.j < 0 || p.j >= len(grid[p.i]) {
		return 0, false
	}
	return grid[p.i][p.j], true
}
func walk(maze [][]int, start, end point) [][]int {

	//建一个maze 复件
	steps := make([][]int, len(maze))
	for i := range steps {
		steps[i] = make([]int, len(maze[i]))
		for j := range steps[i] {
			steps[i][j] = maze[i][j]
		}
	}
	queue := []point{start}
	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]
		if current == end {
			break
		}
		for _, dir := range directions {
			next := current.add(dir)
			//maze at next is 0   steps at next 0 next != start

			val, ok := next.at(steps)
			if !ok || val != 0 || val == 15 {
				continue
			}
			if next == start {
				continue
			}
			cursteps, _ := current.at(steps)
			steps[next.i][next.j] = cursteps + 1
			queue = append(queue, next)
		}
	}
	return steps
}

func way(maze [][]int, start, end point) [][]int {
	route := make([][]int, len(maze))
	for i := range route {
		route[i] = make([]int, len(maze[i]))

	}

	current := end
	route[current.i][current.j] = maze[end.i][end.j]
	for {
		val, _ := current.at(maze)
		if val == 0 {
			break
		}
		for _, dir := range directions {
			next := current.add(dir)
			val, ok := next.at(maze)
			number := maze[current.i][current.j]
			if !ok || (val+1 != number) {
				continue
			}
			route[next.i][next.j] = val
			current = next
		}
	}

	return route
}

func main() {
	maze := readMaze("Algorithm/maze.txt")
	steps := walk(maze, point{0, 0}, point{len(maze) - 1, len(maze[0]) - 1})

	s := way(steps, point{0, 0}, point{len(maze) - 1, len(maze[0]) - 1})
	for _, row := range s {
		for _, val := range row {
			fmt.Printf("%4d", val)
		}
		fmt.Println()
	}
}
