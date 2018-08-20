package queue

import "fmt"

func Example_queue() {
	QueueSample([]string{"1", "2", "pop", "3", "4", "pop"})
	QueueSample([]string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "11", "12"})
	// Output:
	// [1]
	// [1 2]
	// [2]
	// [2 3]
	// [2 3 4]
	// [3 4]

	//	[1]
	//	[1 2]
	//	[1 2 3]
	//	[1 2 3 4]
	//	[1 2 3 4 5]
	//	[1 2 3 4 5 6]
	//	[1 2 3 4 5 6 7]
	//	[1 2 3 4 5 6 7 8]
	//	[1 2 3 4 5 6 7 8 9]
	//	[1 2 3 4 5 6 7 8 9 10]
	//	[1 2 3 4 5 6 7 8 9 10 11]
	//	[1 2 3 4 5 6 7 8 9 10 11 12]
}

func ExampleMyQueue() {
	q := MakeMyQueue(5)
	q = Insert(q, 1)
	q = Insert(q, 2)
	q = Pop(q)
	q = Insert(q, 3)
	q = Insert(q, 4)
	q = Pop(q)
	fmt.Println(q)
	// Output:
	// [3 4]
}
