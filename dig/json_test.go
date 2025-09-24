package dig

import (
	"testing"

	"github.com/jmrepetti/kdr/cherry"

	"github.com/stretchr/testify/assert"
)

func TestNewJsonDigger(t *testing.T) {
	_, err := NewJsonDigger(jsonBody)
	assert.NoError(t, err)
}

func TestNewJsonDiggerInvalidInput(t *testing.T) {
	_, err := NewJsonDigger([]byte(``))
	assert.Error(t, err)
}

func TestJsonDiggerTypes(t *testing.T) {
	d, err := NewJsonDigger(jsonBody)
	assert.NoError(t, err)
	assert.Equal(t, "123 Main St", cherry.Check2(d.String("data", "address", "street")))
	assert.Equal(t, float64(1), cherry.Check2(d.Float64("data", "id")))
	assert.Equal(t, true, cherry.Check2(d.Bool("status")))
}

func TestJsonDiggerNotFound(t *testing.T) {
	d, err := NewJsonDigger(jsonBody)
	assert.NoError(t, err)
	_, err = d.String("data", "addr", "street")
	assert.Error(t, err)
	assert.ErrorContains(t, err, "key 'addr' not found")
}

func TestJsonDiggerFailToConvert(t *testing.T) {
	d, err := NewJsonDigger(jsonBody)
	assert.NoError(t, err)
	_, err = d.String("data", "age")
	assert.Error(t, err)
	assert.ErrorContains(t, err, "can't convert float64 to string in $.data.age")
	_, err = d.String("data", "address", "ext")
	assert.Error(t, err)
	assert.ErrorContains(t, err, "can't convert <nil> to string in $.data.address.ext")
}

func TestJsonDiggerEmptyStringIsString(t *testing.T) {
	d, err := NewJsonDigger(jsonBody)
	assert.NoError(t, err)
	_, err = d.String("data", "nickname")
	assert.NoError(t, err)
}

func TestJsonDiggerNull(t *testing.T) {
	d, err := NewJsonDigger(jsonBody)
	assert.NoError(t, err)
	_, err = d.Null("data", "address", "zip")
	assert.Error(t, err)
	assert.ErrorContains(t, err, "expected null value at [data address zip]")
}
