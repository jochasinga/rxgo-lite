package types

import (
	"testing"

	"github.com/reactivex/rxgo-lite/interfaces"
	"github.com/stretchr/testify/assert"
)

func TestListImplementsObservable(t *testing.T) {
	assert := assert.New(t)
	assert.Implements((*interfaces.Observable)(nil), (*List)(nil))
}

func TestListImplementsIterable(t *testing.T) {
	assert := assert.New(t)
	assert.Implements((*interfaces.Iterable)(nil), (*List)(nil))
}


