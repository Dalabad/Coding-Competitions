package main

import "fmt"

func main() {

	files := []string{"a", "b", "c", "d", "e", "f"}

	d := Dataset{}
	d.readInput(files[0])
	d.simulate()
	d.writeOutput()

	fmt.Printf("Final Score: %d\n", d.Score)
}

func (d *Dataset) simulate() {
	for simulationTimestamp := 0; simulationTimestamp < d.Time; simulationTimestamp++ {
		for _, car := range d.Cars {
			currentStreet := car.Path[0]
			if currentStreet.EndIntersection.isGreen(currentStreet, simulationTimestamp) {
				// Set car to next street
				if len(car.Path) > 1 {
					car.Path = car.Path[1:]
				} else {
					// Car has completed its path, remove
					car.Delete()
					d.UpdateScore(simulationTimestamp)
				}
			}
		}
	}
}

func (i *Intersection) isGreen(street Street, timestamp int) bool {
	if len(i.Schedule.Duration) == 0 {
		return false
	}

	overallDuration := 0
	for _, v := range i.Schedule.Duration {
		overallDuration += v
	}

	remaining := timestamp % overallDuration

	for streetIndex, duration := range i.Schedule.Duration {
		if remaining > duration {
			remaining -= duration
		} else {
			return i.Schedule.Streets[streetIndex].Name == street.Name
		}
	}

	return false
}

func (c *Car) Delete() {
	// TODO: Remove car
}

func (d *Dataset) UpdateScore(timestamp int) {
	addScore := 1000 + (d.Time - timestamp)
	d.Score += addScore
	fmt.Printf("Increase Score by %d\n", addScore)
}
