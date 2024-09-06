package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	data := map[string]interface{}{
		"key1": "value1",
		"key2": 42,
		"key3": true,
	}
	data["huangqi"] = "123"
	jsonData, err := json.Marshal(data)
	if err != nil {
		fmt.Println("Error marshalling map to JSON:", err)
		return
	}

	jsonString := string(jsonData)
	fmt.Println("JSON data as string:", jsonString)
}
