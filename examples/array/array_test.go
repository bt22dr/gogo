package array

import (
	"fmt"

	"github.com/bt22dr/gogo/hangul"
)

func Example() {
	fruits := [3]string{"사과", "바나나", "토마토"}
	for _, fruit := range fruits {
		fmt.Printf("%s는 맛있다.\n", fruit)
	}
	// Output:
	// 사과는 맛있다.
	// 바나나는 맛있다.
	// 토마토는 맛있다.
}

func Example_array() {
	fruits := [4]string{"사과", "바나나", "토마토", "멜론"}
	for _, fruit := range fruits {
		if hangul.HasConsonantSuffix(fruit) == true {
			fmt.Printf("%s은", fruit)
		} else {
			fmt.Printf("%s는", fruit)
		}
		fmt.Printf(" 맛있다.\n")
	}
	// Output:
	// 사과는 맛있다.
	// 바나나는 맛있다.
	// 토마토는 맛있다.
	// 멜론은 맛있다.
}
