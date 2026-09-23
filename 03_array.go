package main

import "fmt"

func main() {
	scores := []int{70, 80, 60, 90, 50}

	fmt.Println("Scores:", scores)
	fmt.Println("first Scores:", scores[0])
	fmt.Println("third Scores:", scores[2])

	for i := 0; i < len(scores); i++ {
		fmt.Println("index", i, ":", scores[i])
	}
}
