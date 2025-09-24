package box

import (
	"bytes"
	"encoding/gob"
)

type BoxT[T any] struct {
	store StoreAdapter
}

func NewBoxT[T any](adapter StoreAdapter) *BoxT[T] {
	return &BoxT[T]{store: adapter}
}

func (box *BoxT[T]) Store(o T) (int64, error) {
	var data bytes.Buffer // Stand-in for the network.
	enc := gob.NewEncoder(&data)

	err := enc.Encode(&o)
	if err != nil {
		return 0, err
	}
	//TODO: consider not passing storage error directly to the user
	return box.store.Store(data.Bytes())

}

func (box *BoxT[T]) Fetch(id int64) (T, error) {
	var o T
	data, err := box.store.Fetch(id)
	if err != nil {
		return o, err
	}
	buff := bytes.NewBuffer(data)
	dec := gob.NewDecoder(buff)

	err = dec.Decode(&o)
	if err != nil {
		return o, err
	}
	return o, nil
}

func (box *BoxT[T]) Register(o any) {
	gob.Register(o)
}
