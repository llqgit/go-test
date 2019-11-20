package main

import (
	"encoding/json"
	"fmt"
	"log"
)

func main() {
	m := make(map[string]string)
	data := `{"foo": "bar"}`
	err := json.Unmarshal([]byte(data), &m)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(m)
}
