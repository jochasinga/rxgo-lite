package interfaces

type Observer interface {
	Handler
	OnNext(interface{})
	OnError(error)
	OnDone()
}


