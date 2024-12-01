package main

import (
	"log"
	"math"
)

const (
	MAP_SIZE       float64 = 300
	TIME_START     float64 = 0
	TIME_STOP      float64 = 100
	TIME_INCREMENT float64 = 1
)

func main() {
	dw, err := newDataWriter()
	if err != nil {
		log.Fatal("error creating data writer:", err)
	}
	defer func() {
		if err := dw.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	dw.WriteInfo(
		TIME_START, TIME_STOP, TIME_INCREMENT,
		int(MAP_SIZE), int(MAP_SIZE),
	)
	for t := TIME_START; t <= TIME_STOP; t = t + TIME_INCREMENT {
		for i := 0.0; i < MAP_SIZE; i++ {
			for j := 0.0; j < MAP_SIZE; j++ {
				rx := i - MAP_SIZE/2
				ry := j - MAP_SIZE/2
				r := math.Sqrt(rx*rx + ry*ry)
				dw.WriteData(t, r*math.Exp(-0.05*r)*math.Sin(r-math.Pi*t/20))
			}
		}
	}
}
