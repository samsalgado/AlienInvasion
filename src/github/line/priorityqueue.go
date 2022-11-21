package queue
import ("container/heap")
type (
	Heapable interface {
		Priority(other interface{}) bool
	}
	//internal container for priority queue 
	items []Heapable
	PriorityQueue struct {
		queue *items
	}
)
//new priority queue referencing new Priority Queue
func NewPriorityQueue() *PriorityQueue {
	pq := &PriorityQueue{queue: &items{}}
	//Referencing items by values
	heap.Init(pq.queue)
	return 	pq 
}
func (pq *PriorityQueue) Push(s Heapable) {
	heap.Push(pq.queue, s)
}
func (pq *PriorityQueue) Pop() Heapable {
	return heap.Pop(pq.queue).(Heapable)
}

// Size returns the size of the priority queue.
func (pq *PriorityQueue) Size() int {
	return pq.queue.Len()
}
func (pq items) Len() int {
	return len(pq)
}

// Less implements the sort.Interface.
func (pq items) Less(i, j int) bool {
	return pq[i].Priority(pq[j])
}

// Swap implements the sort.Interface.
func (pq items) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

// Push implements the heap interface.
func (pq *items) Push(x interface{}) {
	item := x.(Heapable)
	*pq = append(*pq, item)
}

// Pop implements the heap interface.
func (pq *items) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}
