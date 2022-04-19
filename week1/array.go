package main

import "fmt"

func ArrayTest() {

	arr := []string{
		"i",
		"am",
		"stupid",
		"and",
		"weak",
	}

	for idx, item := range arr {
		if item == "stupid" {
			arr[idx] = "smart"
		} else if item == "weak" {
			arr[idx] = "strong"
		}
	}

	for _, item := range arr {
		fmt.Print(item + " ")
	}

}
