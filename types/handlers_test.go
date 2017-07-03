package types

import (
	"errors"
	"testing"

	"github.com/reactivex/rxgo-lite/interfaces"
	"github.com/stretchr/testify/assert"
)

func TestNextFuncImplementsHandler(t *testing.T) {
	assert := assert.New(t)
	assert.Implements((*interfaces.Handler)(nil), (*NextFunc)(nil))
}

func TestErrFuncImplementsHandler(t *testing.T) {
	assert := assert.New(t)
	assert.Implements((*interfaces.Handler)(nil), (*ErrFunc)(nil))
}

func TestDoneFuncImplementsHandler(t *testing.T) {
	assert := assert.New(t)
	assert.Implements((*interfaces.Handler)(nil), (*DoneFunc)(nil))
}

func TestWatcherImplementsHandler(t *testing.T) {
	assert := assert.New(t)
	assert.Implements((*interfaces.Handler)(nil), (*Watcher)(nil))
}

func TestWatcherImplementsObserver(t *testing.T) {
	assert := assert.New(t)
	assert.Implements((*interfaces.Observer)(nil), (*Watcher)(nil))
	
}

func TestCreateWatcherWithFirstArg(t *testing.T) {
	assert := assert.New(t)

	tt := []struct{
		input int
		output int
	}{
		{0, 0},
		{1, 10},
		{2, 20},
		{3, 30},
		{4, 40},
		{5, 50},
	}

	nums := []int{}
	nf := NextFunc(func(i interface{}) {
		nums = append(nums, i.(int) * 10)
	})
	w := NewWatcher(nf)
	
	for _, t := range tt {
		w.Handle(t.input)
	}

	for i, t := range tt {
		assert.Equal(t.output, nums[i])
	}
}


func TestCreateHandlerWithFirstTwoArgs(t *testing.T) {
	assert := assert.New(t)

	tt := []int{10, 20, 30, 40, 50, 60, 70}
	values := []interface{}{
		1, 2, 3, 4, 5, errors.New("oh my"), 6, 7,
	}
	nums := []int{}	


	nf := NextFunc(func(i interface{}) {
		nums = append(nums, i.(int) * 10)
	})
	ef := ErrFunc(func(err error) {})
	
	handler := NewWatcher(nf, ef)
	for _, val := range values {
		handler.Handle(val)
	}

	assert.EqualValues(tt, nums)
}



