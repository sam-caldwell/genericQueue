package genericQueue

import "testing"

func TestQueue_Count(t *testing.T) {
	t.Run("Test values 0-100", func(t *testing.T) {
		for i := uint(0); i < 100; i++ {
			queue := Queue[int]{
				count: i,
			}
			if queue.count != i {
				t.Fatalf("expected count (%d) not found. Got %d", i, queue.count)
			}
		}
	})

}
