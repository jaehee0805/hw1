package matrix

// If needed, you may define helper functions here.

// AreAdjacent returns true iff a and b are adjacent in lst.
func AreAdjacent(a, b int, lst []int) bool {
	adjacent := false

	if len(lst) < 2 {
		return adjacent
	}

	for i := 0; i < len(lst)-1; i++ {
		if (lst[i] == a && lst[i+1] == b) || (lst[i] == b && lst[i+1] == a) {
			adjacent = true
		}
	}

	return adjacent
}

// Transpose returns the transpose of the 2D matrix mat.
func Transpose(mat [][]int) [][]int {
	if mat == nil {
		return nil
	}

	rows := len(mat)
	if rows == 0 {
		return [][]int{}
	}

	cols := len(mat[0])

	transposed := [][]int{}

	for i := 0; i < cols; i++ {
		s := make([]int, 0)
		for j := 0; j < rows; j++ {
			s = append(s, mat[j][i])
		}
		transposed = append(transposed, s)
	}

	return transposed
}

// AreNeighbors returns true iff a and b are neighbors in the 2D matrix mat.
func AreNeighbors(mat [][]int, a, b int) bool {
	neighbors := false
	if mat == nil {
		return false
	}

	rows := len(mat)
	if rows == 0 {
		return false
	}

	cols := len(mat)

	if rows == 1 && cols == 1 {
		return false
	}

	if rows == 1 {
		i := 0
		for j := 0; j < cols-1; j++ {
			if (mat[i][j] == a && mat[i][j+1] == b) || (mat[i][j+1] == b && mat[i][j] == a) {
				neighbors = true
			}
		}
	}

	if cols == 1 {
		j := 0
		for i := 0; i < rows-1; i++ {
			if (mat[i][j] == a && mat[i+1][j] == b) || (mat[i+1][j] == b && mat[i][j] == a) {
				neighbors = true
			}
		}
	}

	for i := 0; i < rows; i++ {
		for j := 0; j < cols-1; j++ {
			if (mat[i][j] == a && mat[i][j+1] == b) || (mat[i][j+1] == b && mat[i][j] == a) {
				neighbors = true
			}
		}
	}

	for j := 0; j < cols; j++ {
		for i := 0; i < rows-1; i++ {
			if (mat[i][j] == a && mat[i+1][j] == b) || (mat[i+1][j] == b && mat[i][j] == a) {
				neighbors = true
			}
		}
	}

	return neighbors
}
