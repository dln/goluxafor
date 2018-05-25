package main

import (
	"github.com/dln/goluxafor"
	"math/rand"
	"time"
)

func randomColor() uint8 {
	return uint8(rand.Intn(0xff))
}

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	luxafor := goluxafor.NewLuxafor()
	// luxafor.Devices[0].Color(goluxafor.LedAll, randomColor(), randomColor(), randomColor(), 0)
	// luxafor.Devices[0].Strobe(goluxafor.LedAll, randomColor(), randomColor(), randomColor(), 10, 2)
	// luxafor.Devices[0].Wave(goluxafor.Wave4, randomColor(), randomColor(), randomColor(), 10, 2)
	luxafor.Devices[0].Pattern(goluxafor.Pattern3, 1)
	luxafor.Close()
}
