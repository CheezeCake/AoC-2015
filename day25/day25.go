package main

import "fmt"

func main() {
	var targetRow, targetColumn int
	fmt.Scanf("To continue, please consult the code grid in the manual.  Enter the code at row %d, column %d.\n", &targetRow, &targetColumn)

	code := 20151125
	row := 2
	col := 1
	for {
		r := row
		for r > 0 {
			code = (code * 252533) % 33554393
			if r == targetRow && col == targetColumn {
				fmt.Println("part 1:", code)
				return
			}
			col++
			r--
		}
		row++
		col = 1
	}
}
