package genericQueue

// Count - return the number of elements currently in the GenericQueue
func (queue *Queue[T]) Count() uint {

	queue.lock.Lock()
	defer queue.lock.Unlock()

	return queue.count

}
