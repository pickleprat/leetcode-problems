package main 

import "fmt" 

func DetermineBoard(row, col int) int {
    if row >= 0 || row <= 2 {
        if col >= 0 || col <= 2 {
            return 0
        } else if col >= 3 || col <= 5 {
            return 1 
        } else if col >= 6 || col <= 8 {
            return 2
        }
    } else if row >= 3 || row <= 5 {
        if col >= 0 || col <= 2 {
            return 3
        } else if col >= 3 || col <= 5 {
            return 4 
        } else if col >= 6 || col <= 8 {
            return 5
        }
    } else if row >= 6 || row <= 8 {
        if col >= 0 || col <= 2 {
            return 6
        } else if col >= 3 || col <= 5 {
            return 7
        } else if col >= 6 || col <= 8 {
            return 8
        }
    }
    return -1; 
}

func isValidSudoku(board [][]byte) bool {
    // the row in these matrices determine the row, col or mini board object 
    // and the column determines the number that was used repeatedly. 
    rowBoard := [9][9] int{}
    colBoard := [9][9] int{}
    miniBoard := [9][9] int{}

    numberMap := make(map[string] int, 0) 
    for i := 1; i <= 9; i++ {
        numberMap[fmt.Sprintf("%d", i)] = i-1
    }

    for row, rowObject := range board {
        for col, numberStr := range rowObject {
            if string(numberStr) == "." {
                continue
            }

            number := numberMap[string(numberStr)]

            if rowBoard[row][number] >= 1 {
                return false 
            } else {
                rowBoard[row][number]++; 
            }

            if colBoard[col][number] >= 1 {
                return false 
            } else {
                colBoard[col][number]++; 
            }


            bNo := DetermineBoard(row, col); 
	    fmt.Printf("row : %d, col : %d, boardSelected: %d, number: %d\n", row, col, bNo, number); 

            if miniBoard[bNo][number] >= 1 {
                return false 
            } else {
                miniBoard[bNo][number]++; 
            }
        }
    }

    return true 
    
}


func main() {
	board := [][]byte{
		{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
		{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
		{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
		{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
		{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
		{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
		{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
		{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
		{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
    	}

	fmt.Printf("The vector is : %+v\n", isValidSudoku(board)); 
} 
