package search

import "fmt"

func Example_search() {
	data := []string{"aa", "bb", "cc", "dd"}
	var result bool

	result = Search(data, "aa")
	fmt.Println(result)

	result = Search(data, "ee")
	fmt.Println(result)

	// Output:
	// true
	// false
}
