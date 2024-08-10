package GenericQueue

// Push - Add a new element to the tail-end of the GenericQueue
func (list *Queue[T]) Push(data T) {

	list.lock.Lock()
	defer list.lock.Unlock()

	newNode := &QueueNode[T]{data: data}

	if list.tail == nil { // List is empty
		list.head = newNode
		list.tail = newNode
	} else {
		list.tail.nextPtr = newNode
		list.tail = newNode
	}

	list.count++ // Increment count
	list.sorted = false

}
