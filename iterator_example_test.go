package algos

import (
	"fmt"
	//"github.com/stretchr/testify/assert"
)

func ExampleSliceIterator_Next_int() {
	var numbers []int = []int{2, 4, 8, 10, 12}

	iter := SliceIterator[int]{Elements: numbers}

	for iter.Next() {
		fmt.Println(iter.Value())
	}

	// Output:
	// 2
	// 4
	// 8
	// 10
	// 12
}

func ExampleSliceIterator_Next_string() {
	strs := []string{"foo", "bar", "qux"}

	iter := SliceIterator[string]{Elements: strs}

	for iter.Next() {
		fmt.Println(iter.Value())
	}

	// Output:
	// foo
	// bar
	// qux
}

func ExampleMap_squared() {
	var numbers []int = []int{1, 2, 3, 4, 5}

	iter := SliceIterator[int]{Elements: numbers}

	mapped := Map[int](&iter, func(x int) int {
		return x * x
	})

	for mapped.Next() {
		fmt.Println(mapped.Value())
	}

	// Output:
	// 1
	// 4
	// 9
	// 16
	// 25
}

func ExampleCollect_identity() {
	var xs []int = []int{2, 4, 8, 10, 12}
	iter := SliceIterator[int]{Elements: xs}

	ys := Collect[int](&iter)

	for _, y := range ys {
		fmt.Println(y)
	}

	// Output:
	// 2
	// 4
	// 8
	// 10
	// 12
}

func ExampleCollect_mapped() {
	var xs []int = []int{2, 4, 8, 10, 12}
	iter := SliceIterator[int]{Elements: xs}
	mapped := Map[int](&iter, func(x int) int {
		return 10 + x
	})

	ys := Collect[int](mapped)

	for _, y := range ys {
		fmt.Println(y)
	}

	// Output:
	// 12
	// 14
	// 18
	// 20
	// 22
}

func ExampleFilter() {
	var numbers []int = []int{1, 2, 3, 4, 5}

	// Create iterator over a slice of integers
	iter := NewSliceIterator(numbers)

	// Pick values larger than 3
	filtered := Filter(iter, func(x int) bool {
		return x > 3
	})

	// Square all values
	mapped := Map(filtered, func(x int) int {
		return x * x
	})

	// Collect result from collection
	result := Collect(mapped)

	for _, x := range result {
		fmt.Println(x)
	}

	// Output:
	// 16
	// 25
}

func ExampleReduce() {
	var numbers []int = []int{1, 2, 3, 4, 5}

	iter := NewSliceIterator(numbers)

	result := Reduce(iter, func(accum, x int) int {
		return accum + x
	})

	fmt.Println(result)

	// Output:
	// 15
}
