package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
	Task1()
	fmt.Println("Task2")
	Task2()
}

// This Task was to count the number of increases from the list of numbers provided
//

type Directions struct {
	Direction string
	Distance  int
}

func Task1() {
	// This reads the file into the input var

	input, err := ioutil.ReadFile("./input.json")
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}

	var payload = []Directions{}

	err = json.Unmarshal(input, &payload)
	if err != nil {
		log.Fatal("Error during Unmarshal(): ", err)
	}

	horizontal := 0
	depth := 0

	for i := range payload {
		if payload[i].Direction == "forward" {
			horizontal = horizontal + payload[i].Distance
		}
		if payload[i].Direction == "up" {
			depth = depth - payload[i].Distance
		}

		if payload[i].Direction == "down" {

			depth = depth + payload[i].Distance
		}

	}

	fmt.Println("Horizontal: ", horizontal)
	fmt.Println("Depth: ", depth)

	fmt.Println("Final Answer: ", horizontal*depth)
}

func Task2() {
	// This reads the file into the input var

	input, err := ioutil.ReadFile("./input.json")
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}

	var payload = []Directions{}

	err = json.Unmarshal(input, &payload)
	if err != nil {
		log.Fatal("Error during Unmarshal(): ", err)
	}

	horizontal := 0
	depth := 0
	aim := 0

	for i := range payload {
		if payload[i].Direction == "forward" {
			horizontal = horizontal + payload[i].Distance
			depth = depth + (aim * payload[i].Distance)
		}
		if payload[i].Direction == "up" {
			aim = aim - payload[i].Distance
		}

		if payload[i].Direction == "down" {
			aim = aim + payload[i].Distance
		}

	}

	fmt.Println("Horizontal: ", horizontal)
	fmt.Println("Depth: ", depth)

	fmt.Println("Final Answer: ", horizontal*depth)

}
