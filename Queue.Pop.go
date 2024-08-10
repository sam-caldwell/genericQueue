package GenericQueue

// Pop - Remove and return the item at the head of the GenericQueue
func (list *Queue[T]) Pop() (data T) {

	list.lock.Lock()
	defer list.lock.Unlock()

	if list.head == nil {
		var nothing T
		return nothing
	}

	currentRecord := list.head
	if currentRecord.nextPtr == nil {
		list.head = nil
	}
	//move head to next node.
	list.head = currentRecord.nextPtr
	currentRecord.nextPtr = nil

	list.count-- // Decrement count

	return currentRecord.data

}
