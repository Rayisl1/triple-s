package main

import "fmt"

func main() {
	var a int
	var array [][]int
	for i := 0; i < a; i++ {
		for j := 0; j < a; j++ {
			fmt.Scan(&(array[i][j]))
		}
	}
	count := 0
	for i := 0; i < a; i++ {
		for j := 0; j < a; j++ {
			if array[i][j] == 1 {
				count++
			}
		}
		if count >= 2 {
			fmt.Print(i + 1)
			break
		}
	}
}
