package clay

type Arena[T any] struct {
	pool []T
	next int
}

func NewArena[T any](capacity int32) Arena[T] {
	return Arena[T]{pool: make([]T, capacity)}
}

func (a *Arena[T]) Allocate() *T {
	if a.next >= len(a.pool) {
		panic("arena: out of memory")
	}
	obj := &a.pool[a.next]
	a.next++
	return obj
}

func (a *Arena[T]) Reset() {
	a.next = 0
}
