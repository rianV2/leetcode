package main

import "fmt"

/*
Notes:
- binary is backward
	binary = "0" + binary

*/

func findComplement(num int) int {
	binary := ""
	for num > 0 {
		if num%2 == 0 {
			binary = "0" + binary //prepand
		} else {
			binary = "1" + binary
		}
		num = num / 2
	}

	fmt.Println(binary)

	length := len(binary)
	total := 0
	for i := length - 1; i >= 0; i-- {
		fmt.Println("df ", i)
		fmt.Println("df2 ", length-1-i)
		//while using length - 1 -i
		// 3-1-2 = 0
		// 3-1-1 = 1
		// 3-1-0 = 2
		// why square(2, i) not working
		// its backward
		// 2*2, 2*1, 2*0 etc
		if string(binary[i]) == "0" {
			total += square(2, length-1-i)
		}
	}

	return total
}

// Always remember this
func square(v, n int) int {
	total := 1
	for i := 0; i < n; i++ {
		total = total * v
	}
	return total
}

func main() {
	// fmt.Println(findComplement(5))
	fmt.Println(findComplement(2))
	// fmt.Println(findComplement(25))
	// fmt.Println(findComplement(0))

	// fmt.Println(square(2, 1))
}
