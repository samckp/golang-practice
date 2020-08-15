package main

import "fmt"

func main() {

	// define MAP
	emps := make(map[string]string)

	emps["Ram"] = "IT"
	emps["Tom"] = "QA"
	emps["Raj"] = "Dev"

	fmt.Println(emps)

	delete(emps, "Raj")
	fmt.Println(emps)

	//Map declaration

	fmt.Println("-----------------------------------------------------------")
	employees := map[string]string{"Ravi": "IT", "Amit": "CS", "Satty": "Market", "Santy": "NGO"}

	fmt.Println(employees)
}
