package iterator

type Iterator interface {
	HasNext() bool
	Next() any
}

type ConcreteIterator struct {
	Items []any
	index int
}

func NewIteratorCollection(data []any) *ConcreteIterator {
	return &ConcreteIterator{
		Items: data,
	}
}

func (i *ConcreteIterator) HasNext() bool {
	return i.index < len(i.Items)
}

func (i *ConcreteIterator) Next() any {
	item := i.Items[i.index]
	i.index++
	return item
}
