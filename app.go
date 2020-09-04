package main

import (
	"fmt"
	mapset "github.com/deckarep/golang-set"
)

func main() {
	basicColors := mapset.NewSet()
	basicColors.Add("Red")
	basicColors.Add("Blue")
	basicColors.Add("Green")

	if basicColors.Contains("Green") {
		fmt.Println("Yay! 'Green' is a basic color")
	} else {
		fmt.Println("What a disappointment! 'Green' is not a basic color")
	}
	fmt.Println("hello world")
}
