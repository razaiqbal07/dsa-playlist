
func isValidSudoku(board [][]byte) bool {
	checked := make(map[byte]bool)

	// checking each row
	for i := 0; i < 9; i++ {
		row := board[i]
		for j := 0; j < 9; j++ {
			_, ok := checked[row[j]]
			if ok {
				return false
			}
			if row[j] != '.' {
				checked[row[j]] = true
			}
		}
		checked = make(map[byte]bool)
	}

	checkedColumn := make(map[byte]bool)
	j := 0
	for j < 9 {
		for i := 0; i < 9; i++ {
			num := board[i][j]
			if num != '.' {
				_, ok := checkedColumn[num]
				if ok {
					return false
				}
				checkedColumn[num] = true
			}
		}
		checkedColumn = make(map[byte]bool)
		j++
	}

	check := make(map[byte]bool)
	columnPointer := 0
	rowPointer := 0
	for rowPointer < 9 {
		for columnPointer < 9 {
			for i := rowPointer; i < rowPointer+3; i++ {
				for j := columnPointer; j < columnPointer+3; j++ {
					num := board[i][j]
					if num != '.' {
						_, ok := check[num]
						if ok {
							return false
						}
						check[num] = true
					}
				}
			}
			columnPointer = columnPointer + 3
			check = make(map[byte]bool)
		}
		rowPointer = rowPointer + 3
		columnPointer = 0
	}
	return true
}