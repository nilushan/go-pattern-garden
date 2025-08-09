package factory

import (
	"encoding/json"
	"fmt"
	"sync"
)

type DataProcessor interface {
	Process(data []byte) ([]byte, error)
	Type() string
}

type ProcessorFactory struct {
	mu         sync.RWMutex
	processors map[string]func() DataProcessor
}

func NewProcessorFactory() *ProcessorFactory {
	return &ProcessorFactory{
		processors: make(map[string]func() DataProcessor),
	}
}

func (f *ProcessorFactory) Register(name string, constructor func() DataProcessor) {

	f.mu.Lock()
	defer f.mu.Unlock()
	f.processors[name] = constructor
}

func (f *ProcessorFactory) Create(name string) (DataProcessor, error) {
	f.mu.Lock()
	defer f.mu.Unlock()
	constructor, ok := f.processors[name]
	if !ok {
		return nil, fmt.Errorf("processor %s not registered", name)
	}
	return constructor(), nil
}

type JSONProcessor struct{}

func (j *JSONProcessor) Process(data []byte) ([]byte, error) {
	var v interface{}
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return json.MarshalIndent(v, "", "")
}

func (j *JSONProcessor) Type() string {
	return "json"
}

type XMLProcessor struct{}

func (x *XMLProcessor) Process(data []byte) ([]byte, error) {

	return data, nil
}

func (x *XMLProcessor) Type() string {
	return "xml"
}

var defaultFactory = NewProcessorFactory()

// This function is automatically called when the package is first used.
func init() {
	defaultFactory.Register("json", func() DataProcessor { return &JSONProcessor{} })
	defaultFactory.Register("xml", func() DataProcessor { return &XMLProcessor{} })

	// jsonProcessor, err := factory.Create("json")
	// if err != nil {
	// 	panic(err)
	// }
	// jsonProcessor.Process([]byte(`{"key": "value"}`))
}

func CreateProcessor(name string) (DataProcessor, error) {
	return defaultFactory.Create(name)
}
