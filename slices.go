package main

func sumNumbers(numbers []int) int { 
	sum := 0 
	for _, value := range numbers { 
		sum += value 
		} 
	return sum 
}

func findLargest(numbers []int) int { 
	largest := 0 
	for _, value := range numbers { 
		if value > largest { 
			largest = value 
			} 
		} 
	return largest 
}

func doubleValues(numbers []int) { 
	for i := range numbers { 
		numbers[i] = numbers[i] * 2 
		} 
}

func slices() {

}