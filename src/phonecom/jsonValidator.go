package main

import (
	"fmt"
	"encoding/json"
)

func validateJson(jsonString string) map[string] interface{} {

	fmt.Println();

	var message string = "JSON validator: Invalid JSON";

	var js map[string] interface{}
	json.Unmarshal([]byte(jsonString), &js)

	if (js != nil) {
		message = "JSON validator: Valid JSON";
		return js
	}

	fmt.Println(message);
	return nil
}
