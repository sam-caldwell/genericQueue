package GenericQueue

// Count - return the number of elements currently in the GenericQueue
func (list *Queue[T]) Count() uint {

	list.lock.Lock()
	defer list.lock.Unlock()

	return list.count

}
