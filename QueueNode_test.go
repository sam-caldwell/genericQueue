package GenericQueue

import "testing"

func TestQueueNode_struct(t *testing.T) {
	_ = QueueNode[int]{
		data:    0,
		nextPtr: &(QueueNode[int]{}),
	}
}
