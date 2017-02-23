package main

import (
	"fmt"
	"encoding/json"
)

func validateJson(jsonString string) {

	fmt.Println();

	var message string = "JSON validator: Invalid JSON";

	if isJSON(string(jsonString)) {
		message = "JSON validator: Valid JSON";
	}

	fmt.Println(message);
}

func isJSON(jsonString string) bool {

	var js map[string] interface{}

	return json.Unmarshal([]byte(jsonString), &js) == nil
}
