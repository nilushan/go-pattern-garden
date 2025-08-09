package main

import (
	"fmt"
	"log"
	"patterngarden/patterns/factory"
)

func demoFactory() {
	// Sample data
	jsonData := []byte(`{"message":"hello world","id":123}`)
	xmlData := []byte(`<message><greeting>hello world</greeting><id>456</id></message>`)

	// --- Use the factory to create a JSON processor ---
	fmt.Println("--- Creating JSON Processor ---")
	jsonProc, err := factory.CreateProcessor("json")
	if err != nil {
		log.Fatalf("Failed to create processor: %v", err)
	}

	// Use the processor via its interface
	processedJSON, err := jsonProc.Process(jsonData)
	if err != nil {
		log.Fatalf("JSON processing failed: %v", err)
	}
	fmt.Printf("Processor Type: %s\n", jsonProc.Type())
	fmt.Printf("Processed JSON Output:\n%s\n\n", string(processedJSON))

	// --- Use the factory to create an XML processor ---
	fmt.Println("--- Creating XML Processor ---")
	xmlProc, err := factory.CreateProcessor("xml")
	if err != nil {
		log.Fatalf("Failed to create processor: %v", err)
	}

	processedXML, err := xmlProc.Process(xmlData)
	if err != nil {
		log.Fatalf("XML processing failed: %v", err)
	}
	fmt.Printf("Processor Type: %s\n", xmlProc.Type())
	fmt.Printf("Processed XML Output:\n%s\n\n", string(processedXML))

	// --- Handle a processor that doesn't exist ---
	fmt.Println("--- Attempting to create a non-existent processor ---")
	_, err = factory.CreateProcessor("yaml")
	if err != nil {
		fmt.Printf("Received expected error: %v\n", err)
	}
}
