package app

import "fmt"

func WorkingWithAnArray(arr [5]int) string {
	minNum, maxNum := arr[0], arr[0]
	sum := 0

	//find the minimum element in the loop
	for _, value := range arr {
		if value < minNum {
			minNum = value
		}
	}

	//find the maximum element in the loop
	for _, value := range arr {
		if value > maxNum {
			maxNum = value
		}
	}

	//in a loop find the sum of the elements of the array
	for i := 0; i < len(arr); i++ {
		sum = sum + arr[i]
	}

	final := fmt.Sprintf("minimum: %d, maximum: %d", sum-maxNum, sum-minNum)
	return final
}
