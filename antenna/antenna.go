package antenna

import (
	"os"
	"os/signal"
)

type Antenna struct {
	Sigs chan os.Signal
	// Done chan bool
}

func NewAntenna(sigs ...os.Signal) Antenna {
	antenna := Antenna{
		Sigs: make(chan os.Signal, 1),
		// Done: make(chan bool, 1),
	}
	signal.Notify(antenna.Sigs, sigs...)
	// go func() {
	// 	sig := <-antenna.Sigs
	// 	fmt.Println(sig)
	// 	antenna.Done <- true
	// }()
	return antenna
}

func (a *Antenna) Wait() <-chan os.Signal {
	return a.Sigs
}
