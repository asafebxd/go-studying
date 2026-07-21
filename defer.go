package main

import "fmt"

func printSteps() {
	fmt.Println("Start")
	defer fmt.Println("End")
	fmt.Println("Middle")
}

func runSafely(shouldPanic bool) {

	defer func() {
		if recovered := recover(); recovered != nil {
			fmt.Println("Recovered:", recovered)
		}
	}()

	if shouldPanic {
		panic("operation failed")
	}

	fmt.Println("Operation completed")
}

func defferedValue() {
	value := 5
	defer fmt.Println(value)
	value = 10
	// Prints 5
}

func defferedValueCorrected() {
	value := 5

	defer func() {
		fmt.Println(value)	
	}()

	value = 10
}

func deferFunc() {

}