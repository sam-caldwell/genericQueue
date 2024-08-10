package GenericQueue

import "testing"

func TestNewGenericQueue(t *testing.T) {
	queue := NewGenericQueue[int]()
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
}
