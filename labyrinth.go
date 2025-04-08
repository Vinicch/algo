package main

var dir = [][]int{
	{-1, 0},
	{0, -1},
	{1, 0},
	{0, 1},
}

type Point struct {
	x int
	y int
}

func SolveMaze(maze []string, wall byte, start, end Point) []Point {
	seen := [][]bool{}
	path := []Point{}

	for range maze {
		seen = append(seen, make([]bool, len(maze)))
	}

	walk(maze, wall, start, end, seen, path)
	return path
}

func walk(maze []string, wall byte, curr, end Point, seen [][]bool, path []Point) bool {
	if curr.x < 0 || curr.y < 0 || curr.x >= len(maze) || curr.y >= len(maze[0]) {
		return false
	}

	if maze[curr.x][curr.y] == wall {
		return false
	}

	if seen[curr.x][curr.y] {
		return false
	}

	seen[curr.x][curr.y] = true
	path = append(path, curr)

	if curr.x == end.x && curr.y == end.y {
		return true
	}

	for i := range dir {
		x, y := dir[i][0], dir[i][1]
		x += curr.x
		y += curr.y
		if walk(maze, wall, Point{x, y}, end, seen, path) {
			return true
		}
	}

	path = path[:len(path)-1]
	return false
}
