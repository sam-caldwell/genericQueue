package GenericQueue

// NewGenericQueue - Generate and configure the initial state for a GenericQueue.Queue object
func NewGenericQueue[T any]() *Queue[T] {
	queue := Queue[T]{
		head:  nil,
		tail:  nil,
		count: 0,
	}
	return &queue
}
