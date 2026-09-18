package main

func predictPartyVictory(senate string) string {
	var r, d queue[int]
	for i, s := range senate {
		switch s {
		case 'R':
			r.enqueue(i)
		case 'D':
			d.enqueue(i)
		}
	}

	for r.len() > 0 && d.len() > 0 {
		ri, rok := r.dequeue()
		di, dok := d.dequeue()
		if !rok || !dok {
			break
		}
		if ri < di {
			r.enqueue(ri + len(senate))
		} else {
			d.enqueue(di + len(senate))
		}
	}

	if r.len() > 0 {
		return "Radiant"
	}

	return "Dire"
}

type queue[T any] struct {
	items []T
}

func (q *queue[T]) enqueue(v T) {
	q.items = append(q.items, v)
}

func (q *queue[T]) dequeue() (T, bool) {
	if len(q.items) == 0 {
		var zero T
		return zero, false
	}

	val := q.items[0]
	q.items = q.items[1:]
	return val, true
}

func (q *queue[T]) len() int {
	return len(q.items)
}
