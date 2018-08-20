package sort

import "fmt"

func Example_sort() {
	var result []string
	result = Sort([]string{"aa", "bb", "cc", "dd"})
	for i := 0; i < len(result); i++ {
		fmt.Println(result[i])
	}
	result = Sort([]string{"cc", "aa", "dd", "bb"})
	for i := 0; i < len(result); i++ {
		fmt.Println(result[i])
	}
	result = Sort([]string{"dd", "cc", "bb", "aa"})
	for i := 0; i < len(result); i++ {
		fmt.Println(result[i])
	}
	// Output:
	// aa
	// bb
	// cc
	// dd
	// aa
	// bb
	// cc
	// dd
	// aa
	// bb
	// cc
	// dd
}
