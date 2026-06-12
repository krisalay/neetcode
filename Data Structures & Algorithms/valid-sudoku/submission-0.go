type Set struct {
	store map[byte]struct{}
}

func NewSet() *Set {
	return &Set{
		store: make(map[byte]struct{}),
	}
}

func (s *Set) AddVal(val byte) bool {
	if _, ok := s.store[val]; !ok {
		s.store[val] = struct{}{}
		return true
	}
	return false
}

func isValidSudoku(board [][]byte) bool {
	// verify rows
	for row, _ := range board {
		s := NewSet()
		for col, _ := range board[row] {
			if board[row][col] == '.' {
				continue
			}
			if ok := s.AddVal(board[row][col]); !ok {
				return false
			}
		}
	}
	// verify columns
	for col := 0; col < len(board[0]); col++ {
		s := NewSet()
		for row := 0; row < len(board); row++ {
			if board[row][col] == '.' {
				continue
			}
			if ok := s.AddVal(board[row][col]); !ok {
				return false
			}
		}
	}
	// verify sub-boxes
	subBoxes := make(map[int]*Set)
	for row := 0; row < len(board); row++ {
		for col := 0; col < len(board[0]); col++ {
			if board[row][col] == '.' {
				continue
			}
			idx := (row / 3) * 3 + (col / 3)
			existingSet, subBoxExists := subBoxes[idx]
			if !subBoxExists {
				existingSet = NewSet()
				subBoxes[idx] = existingSet
			}
			if ok := existingSet.AddVal(board[row][col]); !ok {
				return false
			}
		}
	}
	return true
}
