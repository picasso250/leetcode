package main

type solveNQueensOccupy struct {
	cols map[int]bool
	rows map[int]bool
	x1   map[int]bool
	x2   map[int]bool
}

func solveNQueensOccupyNew() *solveNQueensOccupy {
	return &solveNQueensOccupy{
		make(map[int]bool),
		make(map[int]bool),
		make(map[int]bool),
		make(map[int]bool),
	}
}
func (o *solveNQueensOccupy) Add(i, j int) {
	o.rows[i] = true
	o.cols[j] = true
	o.x1[i-j] = true
	o.x2[i+j] = true
}
func (o *solveNQueensOccupy) Del(i, j int) {
	o.rows[i] = false
	o.cols[j] = false
	o.x1[i-j] = false
	o.x2[i+j] = false
}
func (o *solveNQueensOccupy) CanPlace(i, j int) bool {
	return o.rows[i] == false &&
		o.cols[j] == false &&
		o.x1[i-j] == false &&
		o.x2[i+j] == false
}

func solveNQueens(n int) [][]string {
	// if n == 1 {
	// 	return [][]string{{"Q"}}
	// }
	board := solveNQueensGetEmptyBoard(n)
	boardOccupy := solveNQueensOccupyNew()
	solution := make([][]string, 0)
	solution = solveNQueensIter(board, 0, 0, boardOccupy, solution)
	return solution
}
func solveNQueensIter(board [][]byte, i, j int, occupy *solveNQueensOccupy, solution [][]string) [][]string {
	n := len(board)
	if j >= n {
		// solution = solveNQueensIter(board, i+1, 0, occupy, solution)
		return solution
	}
	if i >= n {
		return solution
	}
	if occupy.CanPlace(i, j) {
		if i == n-1 {
			board[i][j] = 'Q'
			solution = append(solution, solveNQueensToStringSlice(solveNQueensCopyBoard(board)))
			board[i][j] = '.'
			return solution
		} else {
			board[i][j] = 'Q'
			occupy.Add(i, j)
			solution = solveNQueensIter(board, i+1, 0, occupy, solution)
			occupy.Del(i, j)
			board[i][j] = '.'
		}
	}
	solution = solveNQueensIter(board, i, j+1, occupy, solution)
	return solution
}
func solveNQueensToStringSlice(board [][]byte) []string {
	s := make([]string, len(board))
	for i, row := range board {
		s[i] = string(row)
	}
	return s
}
func solveNQueensCopyBoard(board [][]byte) [][]byte {
	ret := make([][]byte, len(board))
	for i, row := range board {
		ret[i] = make([]byte, len(row))
		copy(ret[i], row)
	}
	return ret
}

func solveNQueensSet(board [][]byte, i, j int) {
	n := len(board)
	if 0 <= i && i < n && 0 <= j && j < n {
		board[i][i] = 'x'
	}
}
func solveNQueensGetPos(board [][]byte, want byte) (ret [][2]int) {
	n := len(board)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] == want {
				ret = append(ret, [2]int{i, j})
			}
		}
	}
	return
}
func solveNQueensGetEmptyBoard(n int) (board [][]byte) {
	board = make([][]byte, n)
	for i := 0; i < n; i++ {
		board[i] = make([]byte, n)
		for j := 0; j < n; j++ {
			board[i][j] = '.'
		}
	}
	return
}
