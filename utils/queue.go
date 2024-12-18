package utils

type Element[T any] struct {
	item     T
	priority int
}

type PriorityQueue[T any] struct {
	elements []Element[T]
}

func (q *PriorityQueue[T]) Enqueue(el T, priority int) {
	newElement := Element[T]{
		item:     el,
		priority: priority,
	}
	if len(q.elements) == 0 {
		q.elements = append(q.elements, newElement)
		return
	}

	index := 0
	for i, el := range q.elements {
		if el.priority <= priority {
			index = i + 1
		} else {
			break
		}
	}

	q.elements = append(q.elements[:index], append([]Element[T]{newElement}, q.elements[index:]...)...)
}

func (q *PriorityQueue[T]) Dequeue() T {
	deQueued := q.elements[0]
	q.elements = q.elements[1:]
	return deQueued.item
}

func (q *PriorityQueue[T]) Length() int {
	return len(q.elements)
}
