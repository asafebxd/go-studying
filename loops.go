package main

import "fmt"

func sumUntil(limit int) int { 
	count := 0 
		
	for i := 1; i <= limit; i++ { 
		count += i 
	} 
	return count 
	} 

func countEven(limit int) int { 
	even := 0 
	
	for i := 1; i <= limit; i++ { 
		if i % 2 == 0 { 
			even += 1 
			} 
		} 
	
	return even 
	} 

func fizzBuzz(limit int) { 
	for i := 1; i <= limit; i++ { 
		if i % 5 == 0 && i % 3 == 0 { 
			fmt.Println("FizzBuzz") 
		} else if i % 3 == 0 { 
			fmt.Println("Fizz") 
		} else if i % 5 == 0 { 
			fmt.Println("Buzz") 
		} else { 
			fmt.Println(i) 
			} 
		} 
	}

func loops() {

}