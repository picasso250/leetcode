
type Island map[int]map[int]int

func NewIsland(x int, y int) Island {
	is := make(map[int]map[int]int)
	is[x] = make(map[int]int)
	is[x][y] = 1
	return is
}
func inIsland(x int, y int, i Island) bool {
	return i[x][y] == 1
}
func numIslands(grid [][]byte) int {
	islands := []Island{}
	for i, row := range grid {
		for j := range row {
			islands = addToIsland(i, j, islands)
		}
	}
	return len(islands)
}
func addToIsland(i int, j int, islands []Island) []Island {
	for id, is := range islands {
		if nearIsland(i, j, is) {
			if _, ok := islands[id][i]; !ok {
				islands[id][i] = make(map[int]int)
			}
			islands[id][i][j] = 1

			return islands
		}
	}
	return append(islands, NewIsland(i, j))
}
func nearIsland(i int, j int, is Island) bool {
	return inIsland(i-1, j, is) || inIsland(i+1, j, is) || inIsland(i, j-1, is) || inIsland(i, j+1, is)
}
