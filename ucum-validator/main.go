package main

import (
	"github.com/iimos/ucum"
	"syscall/js"
)

// validateUCUM function that checks the input text
func validateUCUM(this js.Value, inputs []js.Value) interface{} {
	if len(inputs) == 0 {
		return js.ValueOf(map[string]interface{}{
			"valid":  false,
			"reason": "No input provided",
		})
	}

	inputText := inputs[0].String()

	_, err := ucum.Parse([]byte(inputText))

	if err != nil {
		return js.ValueOf(map[string]interface{}{
			"valid":  false,
			"reason": err.Error(),
		})
	}

	return js.ValueOf(map[string]interface{}{
		"valid":  true,
		"reason": "UCUM is valid",
	})
}

func main() {
	// Expose validateUCUM function to JavaScript
	js.Global().Set("validateUCUM", js.FuncOf(validateUCUM))

	// Keep the Go program running
	select {}
}
