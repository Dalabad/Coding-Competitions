package main

import (
	"fmt"
	"hashcode/src"
)

func main() {

	files := []string{"a", "b", "c", "d", "e", "f"}

	runDataset := 0

	d := src.Dataset{}
	d.ReadInput(files[runDataset])
	d.Simulate()
	d.WriteOutput(files[runDataset])

	fmt.Printf("Final Score: %d\n", d.Score)
}

func (d *Dataset) simulate() {
	for simulationTimestamp := 0; simulationTimestamp < d.Time; simulationTimestamp++ {
		for _, street := range d.Streets {
			car := street.Cars[0]

			if street.EndIntersection.isGreen(street, simulationTimestamp) {
				// Set car to next street
				if len(car.Path) > 1 {
					car.Path = car.Path[1:]
				} else {
					// Car has completed its path, remove
					car.Delete()
					d.UpdateScore(simulationTimestamp)
					continue
				}

				// Remove car from street
				if len(street.Cars) > 1 {
					street.Cars = street.Cars[0:]
				} else {
					street.Cars = []*Car{}
				}

				// Move car to next street
				nextStreet := car.Path[0]
				nextStreet.Cars = append(nextStreet.Cars, car)
			}
		}
	}
}