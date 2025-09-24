package box

import (
	"bytes"
	"encoding/gob"
)

type Box struct {
	store StoreAdapter
}

func NewBox(adapter StoreAdapter) *Box {
	return &Box{store: adapter}
}

func (box *Box) Store(o any) (int64, error) {
	// return 0, nil
	// gob.Register(Point{})
	var data bytes.Buffer // Stand-in for the network.
	enc := gob.NewEncoder(&data)

	err := enc.Encode(&o)
	if err != nil {
		return 0, err
	}
	return box.store.Store(data.Bytes())

}

func (box *Box) Fetch(id int64) (any, error) {
	var o any
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

func (box *Box) Register(o any) {
	gob.Register(o)
}
