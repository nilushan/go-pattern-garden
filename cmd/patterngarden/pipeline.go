package main

import (
	"fmt"
	"patterngarden/patterns/pipeline"
)

func demoPipeline() {
	p := pipeline.NewPipeline(
		pipeline.FilterState(func(i interface{}) bool {
			val, ok := i.(int)
			return ok && val > 0
		}),
		pipeline.TransformStage(func(i interface{}) interface{} {
			return i.(int) * 2
		}),
		pipeline.BatchStage(10),
	)

	input := make(chan interface{})
	output := p.Execute(input)

	go func() {
		// Let's use a smaller range for a cleaner example
		for i := -5; i <= 100; i++ {
			input <- i
		}
		close(input)
	}()

	for batch := range output {
		fmt.Printf("Batch: %v\n", batch)
	}
}
