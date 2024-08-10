package genericQueue

import "testing"

func TestQueue_Push(t *testing.T) {
	t.Run("push one", func(t *testing.T) {
		queue := Queue[int]{}
		if queue.count != 0 {
			t.Fatalf("expect count 0")
		}
		if queue.head != nil {
			t.Fatalf("expect head nil")
		}
		if queue.tail != nil {
			t.Fatalf("expect tail nil")
		}
		queue.Push(1337)
		if queue.count != 1 {
			t.Fatalf("expect count 1")
		}
		if queue.head.data != 1337 {
			t.Fatalf("expect head.data 1337")
		}
		if queue.tail.data != 1337 {
			t.Fatalf("expect tail.data 1337")
		}
	})
	t.Run("push ten", func(t *testing.T) {
		queue := Queue[uint]{}
		for i := uint(0); i < 10; i++ {
			if queue.count != i {
				t.Fatalf("expect count %d", i)
			}
			queue.Push(i)
			if queue.head.data != 0 {
				t.Fatalf("expect head.data 0 (unchanged)")
			}
			if queue.tail.data != i {
				t.Fatalf("expect tail.data %d (unchanged)", i)
			}
		}
		if queue.count != 10 {
			t.Fatalf("expect count 10")
		}
		if queue.tail.data != 9 {
			t.Fatalf("expect tail.data 9 (unchanged)")
		}
	})
}
