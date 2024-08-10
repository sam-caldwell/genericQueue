package GenericQueue

import "testing"

func TestNewGenericQueue(t *testing.T) {
	t.Run("basic test", func(t *testing.T) {
		queue := NewGenericQueue[int](nil)
		queue.lock.Lock()
		queue.lock.Unlock()
		if queue.head != nil {
			t.Fail()
		}
		if queue.tail != nil {
			t.Fail()
		}
		if queue.count != 0 {
			t.Fail()
		}
		if queue.compareFunc != nil {
			t.Fail()
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
			t.Fail()
		}
		if queue.tail != nil {
			t.Fail()
		}
		if queue.count != 0 {
			t.Fail()
		}
		if queue.compareFunc != nil {
			t.Fail()
		}
	})
}
