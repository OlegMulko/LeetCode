package pascalstriangle2

// GetRow ...
func GetRow(rowIndex int) []int {
	if rowIndex < 0 {
		return nil
	}
	pascalTriangle := generate(rowIndex + 1)
	return pascalTriangle[rowIndex]
}

func generate(numRows int) [][]int {
	if numRows <= 0 {
		return nil
	}
	pascalTriangle := make([][]int, numRows)
	generateRow(pascalTriangle, numRows)
	return pascalTriangle
}

func generateRow(pascalTriangle [][]int, numRows int) {
	var lenRow int
	var previosRow []int

	if numRows == 0 {
		return
	}
	generateRow(pascalTriangle, numRows-1)
	row := make([]int, numRows)
	i := 0
	row[0] = 1
	if numRows != 1 {
		row[numRows-1] = 1
		previosRow = pascalTriangle[numRows-2]
		lenRow = len(previosRow)
	}
	for {
		if i+1 >= lenRow {
			break
		}
		row[i+1] = previosRow[i] + previosRow[i+1]
		i++
	}
	pascalTriangle[numRows-1] = row
}
