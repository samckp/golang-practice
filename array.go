package main

import "fmt"

func main() {

	var arr [5]int
	fmt.Println(arr)

	fmt.Println("array length : ", len(arr))

	b := [5]int{10, 20, 30, 40, 50}

	fmt.Println("array elements : ", b)

	//two dimensional array

	var twoDArray [3][3]int

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {

			twoDArray[i][j] = i + j
		}
	}

	fmt.Println("2D Array \n", twoDArray)
}
