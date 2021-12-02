package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
	Task1()
	Task2()
}

// This Task was to count the number of increases from the list of numbers provided
//

func Task1() {
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

	//fmt.Println(payload[0])

	IncreaseCount := 0
	prevNum := 0

	for i := range payload {
		//fmt.Println(payload[i])
		if payload[i] > prevNum {
			IncreaseCount++
		}
		prevNum = payload[i]
	}
	//fmt.Println(prevNum)

	// To get the correct answer the first comparison has to be discounted
	fmt.Println("The answer to task 1 is:")
	fmt.Println(IncreaseCount - 1)

}

func Task2() {
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

	//fmt.Println(payload[0])

	//IncreaseCount := 0
	prevNum1 := 0
	prevNum2 := 0

	var windows = []int{}

	for i := range payload {
		if prevNum1 == 0 {
			prevNum1 = payload[i]
			continue
		}
		if prevNum2 == 0 {
			prevNum2 = payload[i]
			continue
		}

		sum := payload[i] + prevNum1 + prevNum2

		windows = append(windows, sum)

		prevNum1 = prevNum2
		prevNum2 = payload[i]

	}

	IncreaseCount2 := 0
	previousNum := 0

	for i := range windows {
		if windows[i] > previousNum {
			IncreaseCount2++
		}

		if i < 20 {
			fmt.Println("Current Num: ", windows[i])
			fmt.Println("Previous Num: ", previousNum)
			fmt.Println(IncreaseCount2)
		}

		previousNum = windows[i]

	}

	fmt.Println("The answer to task 2 is:")
	fmt.Println(IncreaseCount2 - 1)

}
