package GenericQueue

func NewGenericQueue[T any]() *Queue[T] {
	queue := Queue[T]{}
	return &queue
}
