package genericQueue

import (
	"testing"
)

func TestQueue_DeleteIfExists(t *testing.T) {

	t.Run("expect panic on nil comparison function", func(t *testing.T) {
		queue := NewGenericQueue[int](nil)
		defer func() {
			if r := recover(); r == nil {
				t.Errorf("The code did not panic")
			}
		}()
		queue.DeleteIfExists(1024)
	})

	t.Run("Test if exists.  Delete non head or tail node", func(t *testing.T) {
		queue := NewGenericQueue[uint](func(a, b uint) int {
			if a == b {
				return EqualTo
			}
			if a < b {
				return LessThan
			}
			return GreaterThan
		})
		queue.Push(1)
		queue.Push(2)
		queue.Push(3)
		queue.Push(4)

		if queue.tail.data != 4 {
			t.Fatalf("expect tail to be 4")
		}
		if ok := queue.DeleteIfExists(3); !ok {
			t.Fatalf("expected 3 to be found.")
		}
		if queue.count != 3 {
			t.Fatalf("expect count to be 3, got %d", queue.count)
		}
		if queue.tail.data != 4 {
			t.Fatalf("expect tail to still be 4")
		}
	})

	t.Run("Test if exists.  Delete tail node", func(t *testing.T) {
		queue := NewGenericQueue[uint](func(a, b uint) int {
			if a == b {
				return EqualTo
			}
			if a < b {
				return LessThan
			}
			return GreaterThan
		})
		queue.Push(1)
		queue.Push(2)
		queue.Push(3)
		queue.Push(4)

		if queue.tail.data != 4 {
			t.Fatalf("expect tail to be 4")
		}
		if queue.count != 4 {
			t.Fatalf("expect count to be 4, got %d", queue.count)
		}
		if ok := queue.DeleteIfExists(4); !ok {
			t.Fatalf("expected 4 to be found.")
		}
		if queue.count != 3 {
			t.Fatalf("expect count to be 3, got %d", queue.count)
		}
		if i := queue.tail.data; i == 4 {
			t.Fatalf("expect tail to no longer be 4. got %d", i)
		}
	})

	t.Run("Test if not exists", func(t *testing.T) {
		queue := NewGenericQueue[uint](func(a, b uint) int {
			if a == b {
				return EqualTo
			}
			if a < b {
				return LessThan
			}
			return GreaterThan
		})
		queue.Push(1)
		queue.Push(2)
		queue.Push(3)

		if ok := queue.DeleteIfExists(4); ok {
			t.Fatalf("expected 4 to not be found.")
		}
		if queue.count != 3 {
			t.Fatalf("expect count to be 3, got %d", queue.count)
		}
	})
}
