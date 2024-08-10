package GenericQueue

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
		if value := queue.Pop(); value != 1 {
			t.Fatal("value expect 1")
		}
		if queue.count != 2 {
			t.Fatal("count is wrong")
		}
		if value := queue.Pop(); value != 2 {
			t.Fatal("value expect 2")
		}
		if queue.count != 1 {
			t.Fatal("count is wrong")
		}
		if value := queue.Pop(); value != 3 {
			t.Fatal("value expect 3")
		}
		if queue.count != 0 {
			t.Fatal("count is wrong")
		}
		if value := queue.Pop(); value != 0 {
			t.Fatalf("value expect 0 got: %d", value)
		}
		if queue.count != 0 {
			t.Fatal("count is wrong")
		}
	})
}
