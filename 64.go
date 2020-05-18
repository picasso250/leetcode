package main

import "fmt"

func sumI(row []int) int {
	s := 0
	for _, v := range row {
		s += v
	}
	return s
}
func sumJ(grid [][]int, j int) int {
	s := 0
	for i := range grid {
		s += grid[i][j]
	}
	return s
}
func minPathSumIter(grid [][]int, i0, j0, i1, j1 int) int {
	fmt.Printf("%d %d - %d %d\n", i0, j0, i1, j1)
	if i1-i0 == 1 {
		return sumI(grid[i0])
	}
	if j1-j0 == 1 {
		return sumJ(grid, j0)
	}
	v0 := grid[i0][j0]
	a := v0 + minPathSumIter(grid, i0, j0+1, i1, j1)
	b := v0 + minPathSumIter(grid, i0+1, j0, i1, j1)
	fmt.Printf("%d %d\n", a, b)
	if a < b {
		return a
	}
	return b
}
func deepStructure(grid [][]int) [][]int {
	s := make([][]int, len(grid))
	for i := range grid {
		s[i] = make([]int, len(grid[0]))
	}
	return s
}
func evalue(s [][]int, grid [][]int, i, j int) {
	si, sj := s[i][j+1], s[i+1][j]
	m := sj
	if si < sj {
		m = si
	}
	s[i][j] = grid[i][j] + m
}
func minPathSum(grid [][]int) int {
	m := len(grid)
	if m == 0 {
		return 0
	}
	n := len(grid[0])
	s := deepStructure(grid)

	i, j := m-1, n-1
	s[i][j] = grid[i][j]

	// __|
	for i, j = m-1, n-1; i-1 >= 0; i-- {
		s[i-1][j] = grid[i-1][j] + s[i][j]
	}

	for i, j = m-1, n-1; j-1 >= 0; j-- {
		s[i][j-1] = grid[i][j-1] + s[i][j]
	}
	// inner
	for i = m-2; i >= 0; i-- {
		for j = n - 2; j >= 0; j-- {
			evalue(s, grid, i, j)
		}
	}
	return s[0][0]
}
func main() {
	grid := [][]int{
		{1, 2, 5},
		{3, 2, 1},
	}
	fmt.Printf("%d\n", minPathSum(grid))
}
