package cherry

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCheck(t *testing.T) {
	t.Run("no error", func(t *testing.T) {
		assert.NotPanics(t, func() { Check(nil) })
	})

	t.Run("with error", func(t *testing.T) {
		assert.Panics(t, func() { Check(errors.New("something went wrong")) })
	})
}

func TestCheck2(t *testing.T) {
	t.Run("no error", func(t *testing.T) {
		val := Check2(42, nil)
		assert.Equal(t, 42, val)
	})

	t.Run("with error", func(t *testing.T) {
		assert.Panics(t, func() { Check2(42, errors.New("something went wrong")) })
	})

	t.Run("return value", func(t *testing.T) {
		val := Check2("hello", nil)
		assert.Equal(t, "hello", val)
	})

	t.Run("return value with error", func(t *testing.T) {
		assert.Panics(t, func() { Check2("hello", errors.New("something went wrong")) })
	})
}
