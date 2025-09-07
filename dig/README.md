# Package dig

Dig recursively retrieves a value of type T from a nested map[string]any structure.
It traverses the map using the provided fields as keys in order.
If any key in the path is missing or the value cannot be cast to type T, an error is returned.

```go
var data = map[string]any{}
json.Unmarshal(jsonBody, &data)
value, err := Dig[int](data, "user", "age") extracts the int value at data["user"]["age"].
```

Given a JSON body in []bytes

```go
jsonBody := []byte(`{"data":{"id":1,"username":"JSON Stathan","email":""},"error":null,"status":true}`)
digger := NewJsonDigger(jsonBody)
digger.String("data", "username") => "JSON Stathan"
```
