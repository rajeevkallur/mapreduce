// Command mrsample1 demonstrates the mapreduce package by doubling a slice of
// integers with MapInt and summing the result with ReduceInt.
package main

import (
	"fmt"
	"mrmodule/internal/mapreduce"
)

func main() {
	input := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(input)
	doublefunc := multiplefunc(2)
	//quadruplefunc := multiplefunc(4)
	output := mapreduce.Map[int, int](
		input[:],
		doublefunc,
	)
	fmt.Println(output)

	sum := sum(output)
	fmt.Println(sum)
}

func multiplefunc(factor int) func(int) int {
	return func(x int) int {
		return x * factor
	}
}

func sum(input []int) int {
	return mapreduce.Reduce(input, func(x, y int) int {
		return x + y
	}, 0)
}
