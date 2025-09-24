package dig

import (
	"fmt"
	"strings"
)

func dig[T any](data map[string]any, fields ...string) (T, error) {
	var defT T
	currentField := fields[0]
	if _, ok := data[currentField]; !ok {
		return defT, fmt.Errorf("key '%s' not found", currentField)
	}
	switch t := data[currentField].(type) {
	case map[string]any:
		return dig[T](t, fields[1:]...)
	default:
		v, ok := t.(T)
		if ok {
			return v, nil
		} else {
			return defT, fmt.Errorf("can't convert %T to %T", t, defT)
			// return defT, fmt.Errorf("failed to convert %+v of type %T to %T", t, t, defT)
		}
	}
}

// Dig recursively retrieves a value of type T from a nested map[string]any structure.
// It traverses the map using the provided fields as keys in order.
// If any key in the path is missing or the value cannot be cast to type T, an error is returned.
//
// Example: Dig[int](data, "user", "age") extracts the int value at data["user"]["age"].
func Dig[T any](data map[string]any, fields ...string) (T, error) {
	v, err := dig[T](data, fields...)
	if err != nil {
		return v, fmt.Errorf("%w in $.%+v", err, strings.Join(fields, "."))
	}
	return v, nil
}

func DigNil(data map[string]any, fields ...string) (any, error) {
	v, _ := Dig[any](data, fields...)
	switch v.(type) {
	case nil:
		return nil, nil
	default:
		return nil, fmt.Errorf("expected null value at %+v", fields)
	}
}
