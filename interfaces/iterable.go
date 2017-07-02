package interfaces

type (
	MappableFunc func(interface{}) interface{}
	FilterableFunc func(interface{}) bool
)


type Iterable interface {
	Map(MappableFunc) Iterable
	Filter(FilterableFunc) Iterable
}


