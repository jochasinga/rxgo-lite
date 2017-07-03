package types

import (
	"github.com/reactivex/rxgo-lite/interfaces"
)

type List struct {
	list []interface{}
	index int
}

type SubscriptionResult struct {
	err interface{}
	done bool
}

func (re *SubscriptionResult) Result() (interface{}, bool) {
	return re.err, re.done
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

func (li *List) Subscribe(ob interfaces.Observer) <-chan interfaces.IteratorResult {
	done := make(chan interfaces.IteratorResult)
	sub := new(SubscriptionResult)
	count := 0

	go func() {
		for i, val := range li.list {
			count = i
			ob.Handle(val)
			if err, ok := val.(error); ok {
				sub.err = err
				break
			}
		}
		if count == len(li.list) - 1 {
			sub.done = !sub.done
		} 
		done <- sub
		close(done)
	}()
	return done
}


