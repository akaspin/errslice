package errslice_test

import (
	"errors"
	"github.com/akaspin/errslice"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestError_Error(t *testing.T) {
	t.Run("nil", func(t *testing.T) {
		var err errslice.Error
		assert.Empty(t, err.Error())
	})
	t.Run("error", func(t *testing.T) {
		err := errslice.Append(errors.New("1"), errors.New("2"))
		assert.EqualError(t, err, "1,2")
	})
}
