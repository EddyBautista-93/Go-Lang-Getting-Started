package main

import "fmt"

func main() {
	//arrays have to be fixed length
	//var ages [3]int = [3]int{20, 25, 30}

	//shorthand
	var ages = [3]int{20, 25, 30}
	fmt.Println(ages, len(ages))

	greetings := [4]string{"hello", "Howdy", "Hi", "Hey"}
	greetings[2] = "Bye"
	fmt.Println(greetings, len(greetings))

	// slices(use arrays under the hood)
	var scores = []int{100, 20, 201, 433, 32, 34}
	scores[2] = 0
	fmt.Println(scores)
	// you can only append to slices
	scores = append(scores, 292)

	// slice ranges

	rangeOne := greetings[1:3]
	fmt.Println(rangeOne)

}
