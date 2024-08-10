package GenericQueue

// QueueNode - Generic Linked list node
type QueueNode[T any] struct {
	data    T
	nextPtr *QueueNode[T]
}
