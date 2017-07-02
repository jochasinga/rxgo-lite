package types

import (
	"github.com/reactivex/rxgo-lite/interfaces"
)

type List struct {
	list []interface{}
	index int
}

func NewList(items []interface{}) *List {
	if items != nil {
		return &List{list: items}
	}
	return new(List)
}

func (li *List) Map(mappable interfaces.MappableFunc) interfaces.Iterable {
	_li := new(List)
	for _, val := range li.list {
		_li.list = append(_li.list, mappable(val))
	}
	return _li
}

func (li *List) Filter(filterable interfaces.FilterableFunc) interfaces.Iterable {
	_li := new(List)
	for _, val := range li.list {
		included := filterable(val)
		if included {
			_li.list = append(_li.list, included)
		}
	}
	return _li
}

func (li *List) Subscribe(ob interfaces.Observer) <-chan struct{} {
	done := make(chan struct{})

	go func() {
		for _, val := range li.list {
			ob.Handle(val)
		}
		done <- struct{}{}
		close(done)
	}()

	return done
}


