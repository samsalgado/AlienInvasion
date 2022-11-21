package queue
import (
	"sort"
	"testing"
)
type testSortLine struct {
	priority int
}
func (test *testSortLine) Priority(other interface{}) bool {
	if t, ok:= other.(*testSortLine); ok {
		return test.priority > t.priority
	}
	return false 
}
func TestPriorityQueue(t *testing.T) {
	pq := NewPriorityQueue()
	l:= 10
	e := make([]int, 0, l)

	for i:= 0; i<l; i++ {
		p:= (i+1) * 5
		e = append(e, p)
		pq.Push(&testSortLine{priority: p})
	}
	sort.Ints(e)
	for i := l - 1; i>=0; i-- {
		r:= pq.Pop()
		if r.(*testSortLine).priority != e[i] {
			t.Errorf("INCORRECT: expected: %v, got:%v", e[i], r.(*testSortLine).priority)
		}
	}
	if pq.Size() != 0 {
		t.Errorf("INCORRECT: expected:%v, got:%v", 0, pq.Size())
	}
}
