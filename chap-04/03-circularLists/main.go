package main

import (
	"container/ring"
	"fmt"
)

func main() {
	integers := []int{2, 3, 5, 7}

	circularList := ring.New(len(integers))

	for i := 0; i < circularList.Len(); i++ {
		circularList.Value = integers[i]
		circularList = circularList.Next()
	}

	fmt.Println("Printing values from circular linked lists (forward):")
	for i := 0; i < circularList.Len(); i++ {
		fmt.Println(i, circularList.Value)
		circularList = circularList.Next()
	}

	fmt.Println("Printing values from circular linked lists (backward):")
	for i := 0; i < circularList.Len(); i++ {
		fmt.Println(i, circularList.Value)
		circularList = circularList.Next()
	}

	circularList.Do(func(element interface{}) {
		fmt.Print(element, "\t")
	})
	fmt.Println()

	circularList = circularList.Move(2)
	circularList.Do(func(element interface{}) {
		fmt.Print(element, "\t")
	})
	fmt.Println()
}
