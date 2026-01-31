package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/goccy/go-yaml"
	"github.com/joshua-zingale/dasko/pkg/course"
)

func main() {
	fmt.Println("Welcome to Dasko!")

	type T struct {
		A []int `json:"a,omitempty"`
		V int   `json:"b"`
	}

	var t course.ActivitySpec

	err := yaml.UnmarshalWithOptions([]byte(`
title: Hello!
points: 100
overflow-scoring: false
exact-scoring: false
uploads:
  files:
  - name: bob.txt
  - name: main.c
autograder:
  exec: ./evaluate
  points: 40
`), &t, yaml.DisallowUnknownField())

	if err != nil {
		fmt.Printf("%s", yaml.FormatError(err, true, true))
	}

	fmt.Println(t)

	enc := json.NewEncoder(os.Stdout)

	enc.Encode(t)
}
