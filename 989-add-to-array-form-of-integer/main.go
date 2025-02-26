package main

import (
	"fmt"
	"reflect"
)

// using number manipulation
// Modulus: every number modules 10, it will result to last digit number
// Division: every number division 10, it will omit last element cause it's float, will result to comma
// num := 234
// fmt.Println(num % 10) // Output: 4
// num /= 10             // Remove last digit (now num = 23)

func addToArrayForm(arr []int, addition int) []int {
	var newArr []int
	idx := len(arr) - 1
	mem := 0
	for {
		// add := addition % 10
		add := addition%10 + mem
		if idx >= 0 {
			add = add + arr[idx]
		}
		mem = add / 10 //retrieve first digit
		if add > 9 {
			add = add - 10 //use only value under 10
		}
		newArr = append([]int{add}, newArr...)

		if addition/10 == 0 { //break if no more addition
			if idx > 0 || mem != 0 {
				if idx < 0 { //need to pass [0] for left over addition
					idx = 0
				}
				newArr = append(addToArrayForm(arr[0:idx], mem), newArr...) //add left over addition
			}
			break
		}
		idx -= 1
		addition /= 10

	}

	return newArr
}

func addToArrayForm1(arr []int, addition int) []int {
	mem := 0
	for i := len(arr) - 1; i >= 0; i-- {
		add := addition % 10
		total := arr[i] + add + mem
		mem = total / 10 //retrieve first digit

		if total > 9 {
			total = total - 10 //only use last digit
		}

		arr[i] = total

		addition /= 10
	}

	if addition != 0 {
		mem += addition
	}

	if mem != 0 {
		arr = append([]int{mem}, arr...)
	}

	return arr
}

func main() {
	type Input struct {
		arr      []int
		addition int
	}

	type Expect = []int

	tests := []struct {
		name     string
		input    Input
		expected Expect
	}{
		{"Case 1", Input{[]int{2, 7, 4}, 181}, Expect{4, 5, 5}},
		{"Case 2", Input{[]int{0}, 23}, Expect{2, 3}},
		{"Case 3", Input{[]int{0}, 10000}, Expect{1, 0, 0, 0, 0}},
		{"Case 4", Input{[]int{1}, 99}, Expect{1, 0, 0}},
		{"Case 4", Input{[]int{1, 2, 0, 0}, 34}, Expect{1, 2, 3, 4}},
		{"Case 4", Input{[]int{1, 2, 3}, 9}, Expect{1, 3, 2}},
	}

	for _, tt := range tests {
		result := addToArrayForm(tt.input.arr, tt.input.addition)
		if !reflect.DeepEqual(result, tt.expected) {
			fmt.Printf("❌ Test Failed: %s | Input %v, Expected %v, got %v\n", tt.name, tt.input, tt.expected, result)
		} else {
			fmt.Printf("✅ Test Passed: %s | Input %v, Expected %v\n", tt.name, tt.input, tt.expected)
		}
	}
}
