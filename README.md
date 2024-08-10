GenericQueue
============

## Description
Create a generic queue using a linked list so we don't need a lot of reallocations,
contiguous blocks of memory, etc.

## Usage

```go
package main

import (
	"fmt"
	"github.com/sam-caldwell/genericQueue"
)

func main() {

	queue := NewGenericQueue[uint](func(a, b uint) int {
		if a == b {
			return EqualTo
		}
		if a < b {
			return LessThan
		}
		return GreaterThan
	})
	queue.Push(1)
	queue.Push(2)
	queue.Push(3)
	if queue.Count() != 3 {
		panic("count is wrong")
    }
	v, err := queue.Pop()
	if err != nil {
		panic(err)
	}
	fmt.Println(v)
	if ok:=queue.DeleteIfExists(4); ok {
		panic("we can't delete an item that doesn't exist")
    }
}

```