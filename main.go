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
	//d.Simulate()
	d.WriteOutput(files[runDataset])

	fmt.Printf("Final Score: %d\n", d.Score)
}
