package iterator

type Collection interface {
	CreateIterator() Iterator
}

type ConcreteCollection struct {
	items []any
}

func NewConcreteCollection() *ConcreteCollection {
	return &ConcreteCollection{
		items: []any{},
	}
}

func (c *ConcreteCollection) Append(value ...any) {
	c.items = append(c.items, value...)
}

func (c *ConcreteCollection) CreateIterator() Iterator {
	return NewIteratorCollection(c.items)
}
