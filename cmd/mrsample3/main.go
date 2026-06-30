// Command mrsample3 demonstrates the ds.QueueInt type by adding and removing
// integer values.
package main

import (
	"fmt"
	"mrmodule/internal/ds"
)

func main() {
	qi1 := ds.QueueInt{}
	qi1.Add(1)
	qi1.Add(2)
	qi1.Add(3)
	fmt.Println(qi1)
	qi1.Remove()
	fmt.Println(qi1)
}
