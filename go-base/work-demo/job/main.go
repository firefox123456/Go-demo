package main

import (
	"encoding/json"
	"fmt"
)

type Property struct {
	Key   string
	Value string
}

func main() {
	s := "[{\"Key\":\"pipeline.max-parallelism\",\"Value\":\"2048\"},{\"Key\":\"restart-strategy.fixed-delay.attempts\",\"Value\":\"10\"},{\"Key\":\"flink.kubernetes.diagnosis-volume-mounted\",\"Value\":\"true\"}]"
	dynamicProperties := make([]Property, 0)
	json.Unmarshal([]byte(s), &dynamicProperties)
	fmt.Println(dynamicProperties)
}
