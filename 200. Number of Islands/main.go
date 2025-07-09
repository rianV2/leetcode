package main

import "fmt"

func numIslands(grid [][]byte) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}

	lRow, lCol := len(grid), len(grid[0])

	cIsland := 0
	for i := 0; i < lRow; i++ {
		for j := 0; j < lCol; j++ {

			if grid[i][j] == '1' {
				cIsland++
			}

			// visit all 1, stop if meets 0
			// exit the loop once finish visiting all 1
			// until we meet new 1
			visit(grid, i, j, &lRow, &lCol)
		}
	}

	return cIsland
}

func visit(grid [][]byte, row, col int, lRow, lCol *int) {
	if lRow == nil || lCol == nil {
		return
	}

	if row >= *lRow || row < 0 || col >= *lCol || col < 0 {
		return
	}

	// stop visit and increase island
	if grid[row][col] != '1' {
		return
	}

	// mark visit
	grid[row][col] = '2'
	printGrid(grid)

	// adjacent
	visit(grid, row-1, col, lRow, lCol) //up
	visit(grid, row, col+1, lRow, lCol) //right
	visit(grid, row+1, col, lRow, lCol) //bottom
	visit(grid, row, col-1, lRow, lCol) //left
}

func printGrid(grid [][]byte) {
	for _, row := range grid {
		for _, cell := range row {
			fmt.Printf("%c ", cell)
		}
		fmt.Println()
	}
	fmt.Println()
}

func main() {
	// grid := [][]byte{
	// 	{'1', '1', '1', '1', '0'},
	// 	{'1', '1', '0', '1', '0'},
	// 	{'1', '1', '0', '0', '0'},
	// 	{'0', '0', '0', '0', '0'},
	// } // 1
	grid := [][]byte{
		{'1', '1', '0', '0', '0'},
		{'1', '1', '0', '0', '0'},
		{'0', '0', '1', '0', '0'},
		{'0', '0', '0', '1', '1'},
	} //3

	// if ok := grid[0]; ok {
	// 	fmt.Println("ada")
	// }
	fmt.Println("result", numIslands(grid))
}
