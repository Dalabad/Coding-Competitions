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
	for _, i := range d.Intersections {
		i.Schedule.Duration = append(i.Schedule.Duration, 1)
	}
	d.Simulate()
	d.WriteOutput(files[runDataset])

	fmt.Printf("Final Score: %d\n", d.Score)
}