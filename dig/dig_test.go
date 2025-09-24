package dig

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/jmrepetti/kdr/cherry"

	"github.com/stretchr/testify/assert"
)

var jsonBody = []byte(`{"data":{"id": 1,
	"first_name": "John",
	"last_name": "Doe",
	"age": 30,
	"nickname": "",
	"address": {
		"street": "123 Main St",
		"ext": null,
		"city": "Madrid",
		"zip": "28013",
		"country": "Spain"}
	},
"error":null,
"status":true}`)

func data() map[string]any {
	var data = map[string]any{}
	json.Unmarshal(jsonBody, &data)
	return data
}

func TestDigTypes(t *testing.T) {
	data := data()
	assert.Equal(t, "123 Main St", cherry.Check2(Dig[string](data, "data", "address", "street")))
	assert.Equal(t, float64(1), cherry.Check2(Dig[float64](data, "data", "id")))
	assert.Equal(t, true, cherry.Check2(Dig[bool](data, "status")))
}

func TestDigNotFound(t *testing.T) {
	data := data()
	_, err := Dig[string](data, "data", "addr", "street")
	assert.Error(t, err)
	assert.ErrorContains(t, err, "key 'addr' not found")
}

func TestDigFailToConvert(t *testing.T) {
	data := data()
	_, err := Dig[string](data, "data", "age")
	assert.Error(t, err)
	assert.ErrorContains(t, err, "can't convert float64 to string in $.data.age")
	_, err = Dig[string](data, "data", "address", "ext")
	assert.Error(t, err)
	fmt.Println(err)
	assert.ErrorContains(t, err, "can't convert <nil> to string in $.data.address.ext")
}

func TestDigEmptyStringIsString(t *testing.T) {
	data := data()
	_, err := Dig[string](data, "data", "nickname")
	assert.NoError(t, err)
}

func TestDigNull(t *testing.T) {
	data := data()
	_, err := DigNil(data, "data", "address", "zip")
	assert.Error(t, err)
	assert.ErrorContains(t, err, "expected null value at [data address zip]")
	v, err := DigNil(data, "data", "address", "ext")
	assert.NoError(t, err)
	assert.Equal(t, nil, v)
}
