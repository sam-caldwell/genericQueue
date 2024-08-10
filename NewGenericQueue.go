package genericQueue

// NewGenericQueue - Generate and configure the initial state for a GenericQueue.Queue object
func NewGenericQueue[T any](compareFunc func(T, T) int) *Queue[T] {
	queue := Queue[T]{
		head:        nil,
		tail:        nil,
		count:       0,
		compareFunc: compareFunc,
	}
	return &queue
}
