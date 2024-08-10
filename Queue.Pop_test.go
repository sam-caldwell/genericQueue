package genericQueue

import "testing"

func TestQueue_Pop(t *testing.T) {
	t.Run("pop static setup", func(t *testing.T) {
		queue := Queue[uint]{
			count: 3,
			head: &(QueueNode[uint]{
				data: 1,
				nextPtr: &(QueueNode[uint]{
					data: 2,
					nextPtr: &(QueueNode[uint]{
						data:    3,
						nextPtr: nil,
					}),
				}),
			}),
		}
		if queue.head.data != 1 {
			t.Fatal("setup 1 is wrong")
		}
		if queue.count != 3 {
			t.Fatal("count is wrong")
		}
		if value, err := queue.Pop(); value != 1 || err != nil {
			t.Fatal("value expect 1, nil")
		}
		if queue.count != 2 {
			t.Fatal("count is wrong")
		}
		if value, err := queue.Pop(); value != 2 || err != nil {
			t.Fatal("value expect 2, nil")
		}
		if queue.count != 1 {
			t.Fatal("count is wrong")
		}
		if value, err := queue.Pop(); value != 3 || err != nil {
			t.Fatal("value expect 3, nil")
		}
		if queue.count != 0 {
			t.Fatal("count is wrong")
		}
		if value, err := queue.Pop(); value != 0 || err == nil {
			t.Fatal("value expect 0, 'queue empty' error")
		}
		if queue.count != 0 {
			t.Fatal("count is wrong")
		}
	})
}
