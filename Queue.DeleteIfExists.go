package GenericQueue

// DeleteIfExists - if a given claimedToken exists, return true and delete the item from the GenericQueue.
func (list *Queue[T]) DeleteIfExists(claimedToken T, compareFunc func(a, b T) int) bool {

	list.lock.Lock()
	defer list.lock.Unlock()

	var prevRecord *QueueNode[T] = nil
	for currentRecord := list.head; currentRecord.nextPtr != nil; currentRecord = currentRecord.nextPtr {
		if compareFunc(currentRecord.data, claimedToken) == 0 {
			prevRecord.nextPtr = currentRecord.nextPtr
			currentRecord.nextPtr = nil
			currentRecord = nil
			list.count--
			return true
		}
		prevRecord = currentRecord
	}

	return false

}
