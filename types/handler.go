package types

type HandlerFunc func(interface{})

func (handlerFunc HandlerFunc) Handle(item interface{}) {
	handlerFunc(item)
}

type Handler struct {
	nextFunc HandlerFunc
	errFunc  HandlerFunc
	doneFunc HandlerFunc
}

var DefaultHandler = &Handler{
	nextFunc: func(interface{}) {},
	errFunc: func(interface{}) {},
	doneFunc: func(interface{}) {},
}

func (handler Handler) processItem(item interface{}) {
	switch item := item.(type) {
	case error:
		handler.OnError(item)
	case bool:
		handler.OnDone()
	default:
		handler.OnNext(item)
	}
}

func NewHandler(nextFunc HandlerFunc, fs ...HandlerFunc) *Handler {
	handler := DefaultHandler
	handler.nextFunc = nextFunc
	if len(fs) > 0 {
		for i, _ := range fs {
			switch i {
			case 0:
				handler.errFunc = fs[i]
			case 1:
				handler.doneFunc = fs[i]
			default:
			}
		}
	}
	return handler
}

func (handler Handler) Handle(item interface{}) {
	handler.processItem(item)
}

func (handler Handler) OnNext(item interface{}) {
	handler.nextFunc(item)
}

func (handler Handler) OnError(err error) {
	handler.errFunc(err)
}

func (handler Handler) OnDone() {
	handler.doneFunc(nil)
}
