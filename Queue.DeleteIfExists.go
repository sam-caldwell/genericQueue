package genericQueue

// DeleteIfExists - if a given claimedToken exists, return true and delete the item from the GenericQueue.
func (queue *Queue[T]) DeleteIfExists(claimedToken T) bool {
	queue.lock.Lock()
	defer queue.lock.Unlock()

	if queue.compareFunc == nil {
		panic("nil comparison function in GenericQueue")
	}
	if queue.head == nil {
		return false
	}

	var prevRecord *QueueNode[T] = nil
	currentRecord := queue.head
	for currentRecord != nil {
		if queue.compareFunc(currentRecord.data, claimedToken) == EqualTo {
			if prevRecord == nil { // Deleting the head node
				queue.head = currentRecord.nextPtr
				if queue.head == nil { // Queue is now empty
					queue.tail = nil
				}
			} else {
				prevRecord.nextPtr = currentRecord.nextPtr
				if currentRecord.nextPtr == nil { // Deleting the tail node
					queue.tail = prevRecord
				}
			}
			currentRecord.nextPtr = nil
			queue.count--
			return true
		}
		prevRecord = currentRecord
		currentRecord = currentRecord.nextPtr
	}

	return false
}
