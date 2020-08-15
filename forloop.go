package main

import "fmt"

func main() {

	fmt.Println("For Loop scenarios ")

	// for loop
	for i := 0; i <= 10; i++ {

		fmt.Println(i)
	}

	//for loop with condition and continue
	for j := 0; j < 10; j++ {
		if j%2 == 0 {
			continue
		}
		fmt.Println(j)
	}
}
