package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

// This Task was to count the number of increases from the list of numbers provided
//

func main() {

	// This reads the file into the input var

	input, err := ioutil.ReadFile("./input.json")
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}

	var payload = []int{}

	err = json.Unmarshal(input, &payload)
	if err != nil {
		log.Fatal("Error during Unmarshal(): ", err)
	}

	fmt.Println(payload[0])

	IncreaseCount := 0
	prevNum := 0

	for i := range payload {
		fmt.Println(payload[i])
		if payload[i] > prevNum {
			IncreaseCount++
		}
		prevNum = payload[i]
	}
	fmt.Println(prevNum)

	// To get the correct answer the first comparison has to be discounted
	fmt.Println(IncreaseCount - 1)

}
