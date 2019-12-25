package isValidSudoku

//Determine if a 9x9 Sudoku board is valid. Only the filled cells need to be validated according to the following rules:
//
//Each row must contain the digits 1-9 without repetition.
//Each column must contain the digits 1-9 without repetition.
//Each of the 9 3x3 sub-boxes of the grid must contain the digits 1-9 without repetition.
//The Sudoku board could be partially filled, where empty cells are filled with the character '.'.
//
//Example 1:
//
//Input:
//[
//["5","3",".",".","7",".",".",".","."],
//["6",".",".","1","9","5",".",".","."],
//[".","9","8",".",".",".",".","6","."],
//["8",".",".",".","6",".",".",".","3"],
//["4",".",".","8",".","3",".",".","1"],
//["7",".",".",".","2",".",".",".","6"],
//[".","6",".",".",".",".","2","8","."],
//[".",".",".","4","1","9",".",".","5"],
//[".",".",".",".","8",".",".","7","9"]
//]
//Output: true
//Example 2:
//
//Input:
//[
//["8","3",".",".","7",".",".",".","."],
//["6",".",".","1","9","5",".",".","."],
//[".","9","8",".",".",".",".","6","."],
//["8",".",".",".","6",".",".",".","3"],
//["4",".",".","8",".","3",".",".","1"],
//["7",".",".",".","2",".",".",".","6"],
//[".","6",".",".",".",".","2","8","."],
//[".",".",".","4","1","9",".",".","5"],
//[".",".",".",".","8",".",".","7","9"]
//]
//Output: false
//Explanation: Same as Example 1, except with the 5 in the top left corner being
//modified to 8. Since there are two 8's in the top left 3x3 sub-box, it is invalid.
//Note:
//
//A Sudoku board (partially filled) could be valid but is not necessarily solvable.
//Only the filled cells need to be validated according to the mentioned rules.
//The given board contain only digits 1-9 and the character '.'.
//The given board size is always 9x9.

func isValidSudoku(board [][]byte) bool {
	//Each row must contain the digits 1-9 without repetition
	if !isValidRows(board) {
		return false
	}

	//Each column must contain the digits 1-9 without repetition
	boardTmp := make([][]byte, len(board))
	for i := range boardTmp {
		boardTmp[i] = make([]byte, len(board[i]))
	}
	for i := range board {
		for j := i; j < len(board[i]); j++ {
			boardTmp[i][j], boardTmp[j][i] = board[j][i], board[i][j]
		}
	}
	if !isValidRows(boardTmp) {
		return false
	}
	//Each of the 9 3x3 sub-boxes of the grid must contain the digits 1-9 without repetition
	for i := 0; i < len(board); i += 3 {
		for j := 0; j < len(board[i]); j += 3 {
			var tmpArray []byte
			tmpArray = append(append(append(tmpArray,
				board[i][j:j+3]...),
				board[i+1][j:j+3]...),
				board[i+2][j:j+3]...)
			if containsDuplicate(tmpArray) {
				return false
			}
		}
	}
	return true
}

func isValidRows(board [][]byte) bool {
	for _, row := range board {
		if containsDuplicate(row) {
			return false
		}
	}
	return true
}

func containsDuplicate(bytes []byte) bool {
	bytesMap := make(map[byte]bool)
	for _, v := range bytes {
		if v == '.' {
			continue
		}
		_, ok := bytesMap[v]
		if !ok {
			bytesMap[v] = true
		} else {
			return true
		}
	}
	return false
}
