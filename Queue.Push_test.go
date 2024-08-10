package GenericQueue

import "testing"

func TestQueue_Push(t *testing.T) {
	t.Run("push one", func(t *testing.T) {
		queue := Queue[int]{}
		if queue.count != 0 {
			t.Fatalf("expect count 0")
		}
		queue.Push(1)
		if queue.count != 1 {
			t.Fatalf("expect count 1")
		}
	})
	t.Run("push ten", func(t *testing.T) {
		queue := Queue[uint]{}
		for i := uint(0); i < 10; i++ {
			if queue.count != i {
				t.Fatalf("expect count %d", i)
			}
			queue.Push(i)
		}
		if queue.count != 10 {
			t.Fatalf("expect count 10")
		}
	})
}
