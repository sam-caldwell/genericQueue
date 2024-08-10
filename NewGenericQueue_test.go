package genericQueue

import "testing"

func TestNewGenericQueue(t *testing.T) {
	t.Run("basic test", func(t *testing.T) {
		queue := NewGenericQueue[int](nil)
		queue.lock.Lock()
		queue.lock.Unlock()
		if queue.head != nil {
			t.Fatal("head should be nil")
		}
		if queue.tail != nil {
			t.Fatal("tail should be nil")
		}
		if queue.count != 0 {
			t.Fatal("count should be 0")
		}
		if queue.compareFunc != nil {
			t.Fatal("compareFunc should be nil")
		}
	})
	t.Run("test with compareFunc", func(t *testing.T) {
		queue := NewGenericQueue[int](func(a, b int) int {
			if a == b {
				return EqualTo
			}
			if a < b {
				return LessThan
			}
			return GreaterThan
		})
		queue.lock.Lock()
		queue.lock.Unlock()
		if queue.head != nil {
			t.Fatal("head should be nil")
		}
		if queue.tail != nil {
			t.Fatal("tail should be nil")
		}
		if queue.count != 0 {
			t.Fatal("count should be 0")
		}
		if queue.compareFunc == nil {
			t.Fatal("compareFunc should be nil")
		}
	})
}
