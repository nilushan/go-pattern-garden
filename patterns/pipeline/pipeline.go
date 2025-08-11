package pipeline

type PipelineStage func(<-chan interface{}) <-chan interface{}

type Pipeline struct {
	stages []PipelineStage
}

func NewPipeline(stages ...PipelineStage) *Pipeline {
	return &Pipeline{
		stages: stages,
	}
}

func (p *Pipeline) Execute(input <-chan interface{}) <-chan interface{} {
	var output <-chan interface{} = input

	for _, stage := range p.stages {
		output = stage(output)
	}
	return output
}

func FilterState(predicate func(interface{}) bool) PipelineStage {

	return func(input <-chan interface{}) <-chan interface{} {
		output := make(chan interface{})

		go func() {
			defer close(output)
			for item := range input {
				if predicate(item) {
					output <- item
				}
			}
		}()
		return output
	}
}

func TransformStage(transform func(interface{}) interface{}) PipelineStage {
	return func(input <-chan interface{}) <-chan interface{} {
		output := make(chan interface{})

		go func() {
			defer close(output)
			for item := range input {
				output <- transform(item)
			}
		}()
		return output
	}
}

func BatchStage(size int) PipelineStage {
	return func(input <-chan interface{}) <-chan interface{} {
		output := make(chan interface{})

		go func() {
			defer close(output)
			batch := make([]interface{}, 0, size)

			for item := range input {
				batch = append(batch, item)
				if len(batch) == size {
					output <- batch
					batch = nil
				}
			}
			if len(batch) > 0 {
				output <- batch
			}
		}()
		return output
	}
}
