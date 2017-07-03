package interfaces

type Observable interface {
	Iterable
	Subscribe(Observer) <-chan IteratorResult
}


