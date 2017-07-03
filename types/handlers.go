package types

import "github.com/reactivex/rxgo-lite/interfaces"

type (
	NextFunc func(interface{})
	ErrFunc  func(error)
	DoneFunc func()
)

func (nextFunc NextFunc) Handle(item interface{}) {
	if _, ok := item.(error); ok {
		return
	}
	nextFunc(item)
}

func (errFunc ErrFunc) Handle(item interface{}) {
	if v, ok := item.(error); ok {
		errFunc(v)
	}
}

func (doneFunc DoneFunc) Handle(item interface{}) {
	doneFunc()
}

type Watcher struct {
	nextFunc interfaces.Handler
	errFunc  interfaces.Handler
	doneFunc interfaces.Handler
}

var DefaultWatcher = Watcher{
	nextFunc: NextFunc(func(i interface{}) {}),
	errFunc: ErrFunc(func(err error) {}),
	doneFunc: DoneFunc(func() {}),
}

func (w Watcher) processItem(item interface{}) {
	switch item := item.(type) {
	case error:
		w.OnError(item)
	case struct{}:
		w.OnDone()
	default:
		w.OnNext(item)
	}
}

func NewWatcher(nextFunc interfaces.Handler, fs ...interfaces.Handler) Watcher {
	w := DefaultWatcher
	w.nextFunc = nextFunc
	if len(fs) > 0 {
		for i, _ := range fs {
			switch i {
			case 0:
				w.errFunc = fs[i]
			case 1:
				w.doneFunc = fs[i]
			default:
			}
		}
	}
	return w
}

func (w Watcher) Handle(item interface{}) {
	w.processItem(item)
}

func (w Watcher) OnNext(item interface{}) {
	if nf, ok := w.nextFunc.(NextFunc); ok {
		nf(item)
	}
}

func (w Watcher) OnError(err error) {
	if ef, ok := w.errFunc.(ErrFunc); ok {
		ef(err)
	}
}

func (w Watcher) OnDone() {
	if df, ok := w.doneFunc.(DoneFunc); ok {
		df()
	}
}
