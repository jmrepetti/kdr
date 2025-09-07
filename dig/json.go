package dig

import (
	"encoding/json"
	"fmt"
)

// Giveb a JSON body in []bytes
// jsonBody := []byte(`{"data":{"id":1,"username":"JSON Stathan","email":""},"error":null,"status":true}`)
// digger := NewDigger(jsonBody)
// digger.String("data", "username") => "JSON Stathan"
type JsonDigger struct {
	data map[string]any
}

func NewJsonDigger(bs []byte) (JsonDigger, error) {
	var data = map[string]any{}
	err := json.Unmarshal(bs, &data)
	return JsonDigger{data}, err
}

func (d *JsonDigger) Any(fields ...string) (any, error) {
	return Dig[any](d.data, fields...)
}

func (d *JsonDigger) Int(fields ...string) (int, error) {
	return Dig[int](d.data, fields...)
}

func (d *JsonDigger) Bool(fields ...string) (bool, error) {
	return Dig[bool](d.data, fields...)
}

func (d *JsonDigger) String(fields ...string) (string, error) {
	return Dig[string](d.data, fields...)
}

func (d *JsonDigger) Null(fields ...string) (any, error) {
	v, _ := d.Any(fields...)
	switch v.(type) {
	case nil:
		return nil, nil
	default:
		return nil, fmt.Errorf("expected null value at %+v", fields)
	}
}

func (d *JsonDigger) Float64(fields ...string) (float64, error) {
	return Dig[float64](d.data, fields...)
}

func (d *JsonDigger) Float32(fields ...string) (float32, error) {
	return Dig[float32](d.data, fields...)
}
