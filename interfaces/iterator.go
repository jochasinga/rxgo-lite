package interfaces

type Iterator interface {
	Next() (interface{}, bool)
}

type IteratorResult interface {
	Result() (interface{}, bool)
}