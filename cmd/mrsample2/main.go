// Command mrsample2 demonstrates the ds.StackInt type by pushing, inspecting,
// and popping integer values.
package main

import (
	"fmt"
	"mrmodule/internal/ds"
)

func main() {
	si1 := ds.StackInt{}
	fmt.Println(si1.Count())
	si1.Push(1)
	si1.Push(2)
	si1.Push(3)
	fmt.Println(si1)
	fmt.Print(si1.Count())
	fmt.Print(si1.Total())
	fmt.Print(si1.Average())
	for i := 0; i < 3; i++ {
		v, err := si1.Pop()
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			continue
		}

		fmt.Println(v)
	}
}
