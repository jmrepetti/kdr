package antenna

import (
	"os"
	"os/signal"
)

type Antenna struct {
	Sigs chan os.Signal
}

func NewAntenna(sigs ...os.Signal) Antenna {
	antenna := Antenna{
		Sigs: make(chan os.Signal, 1),
	}
	signal.Notify(antenna.Sigs, sigs...)
	return antenna
}

func (a *Antenna) Wait() <-chan os.Signal {
	return a.Sigs
}
