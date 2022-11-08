package main

import "fmt"

func main() {
	arr := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			num := arr[i][j]
			if j > 1 {
				goto EXIT
			}
			fmt.Println(num)
		}
	}
EXIT:
	fmt.Println("Exit.")
}
