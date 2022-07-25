# Algorithms and Iterators

An experiment in using Go Generics to create iterators and build algorithms around the use of these iterators. The typical functional algorithms such as map, filter and reduce.

Here is an example of combining different algorithms in a code example. Notice that because Go does not have function overloading we need different functions for creating each iterator type. `NewSliceIterator` is used to create an iterator over a slice. A different function would be needed to iterate over a dictionary or binary tree.


	var numbers []int = []int{1, 2, 3, 4, 5}
	
    // Create iterator over a slice of integers
    iter := NewSliceIterator[int](numbers)

	// Pick values larger than 3
	filtered := Filter[int](iter, func(x int) bool {
		return x > 3
	})

	// Square all values
	mapped := Map[int](filtered, func(x int) int {
		return x * x
	})

	// Collect result from collection
	result := Collect[int](mapped)

	for _, x := range result {
		fmt.Println(x)
	}
    
Unlike what you may be used to from other code examples the Map and Filter functions return new iterators rather than the final result. To get get final result you call Collect which returns a slice with all the result data. An alternative is to simply use a for-loop to get each value in the iterator.

	for mapped.Next() {
		fmt.Println(mapped.Value())
	}
    
Each iterator must adhere to the following interface:

    type Iterator[T any] interface {
    	Next() bool
    	Value() T
    }
    
This is not actually what Google uses internally for their iterator interface. You can find the Google iterator API described here: [Google iterator package](https://pkg.go.dev/google.golang.org/api/iterator).

The Google iterator interface signals end of iteration with an error, while I am using a boolean. I choose to model the iterator interface on the pratices found in the Go standard library. The Scanner type uses an iterator like interface for instance.

# Reflections on the Use of Iterators in Go

Generics have made using iterators more practical in Go. However I am not sure iterators combined with map, filter and reduce are always the most obvious solution to Go programming problems.

I use this style of programming a lot in Julia, but I don't think Go is well suited for it. Julia has a REPL environment which makes this kind of coding easy because you can process data in a functional pipeline and see the output right away. It is easy to gradually build up code which runs data through these kinds of functional pipelines while observing results in a REPL.

That is not how you do Go programming. You are more likely rely more on using a debugger or tests. In such cases functional programming is not always ideal. It is much harder to debug data piped through lots of maps, filters and reduce statements than to just step through a simple for-loop.

The real value in iterators is for complex data structures such as binary trees, but even then you may want to process data using a simple for-loop.
