package types

import (
	"errors"
	"testing"

	"github.com/reactivex/rxgo-lite/interfaces"
	"github.com/stretchr/testify/assert"
)

func TestHandlerFuncImplementsObserver(t *testing.T) {
	assert := assert.New(t)
	assert.Implements((*interfaces.Observer)(nil), (*HandlerFunc)(nil))
}

func TestHandlerImplementsObserver(t *testing.T) {
	assert := assert.New(t)
	assert.Implements((*interfaces.Observer)(nil), (*Handler)(nil))
	
}

func TestCreateHandlerWithFirstArg(t *testing.T) {
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
	handler := NewHandler(func(i interface{}) {
		nums = append(nums, i.(int) * 10)
	})
	
	for i := 0; i < len(tt); i++ {
		handler.Handle(i)
	}

	for i, t := range tt {
		assert.Equal(t.output, nums[i])
	}
}


func TestCreateHandlerWithFirstTwoArgs(t *testing.T) {
	assert := assert.New(t)

	tt := []struct{
		input int
		output int
	}{
		{1, 10},
		{2, 20},
		{3, 30},
		{4, 40},
		{5, 50},
	}
	
	nums := []int{}
	errs := []error{}
	
	handler := NewHandler(
		func(i interface{}) {
			nums = append(nums, i.(int) * 10)
		},
		func(err interface{}) {
			errs = append(errs, err.(error))
		},
	)

	items := []interface{}{1, 2, 3, 4, 5, errors.New("oh my"), 6, 7}

	for _, item := range items {
		handler.Handle(item)
	}

	for i, t := range tt {
		assert.Equal(t.output, nums[i])
	}

	assert.Equal("oh my", errs[0].Error())
}


