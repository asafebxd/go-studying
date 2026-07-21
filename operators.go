package main

func classifyAge(age int) string { 
	switch { 
		case age < 13: 
			return "Child" 
		case age >= 13 && age <= 17: 
			return "Teenager" 
		case age >= 18 && age <= 64: 
			return "Adult" 
		default: 
			return "Senior" 
			} 
		} 
	
func canAccess(age int, isAuthorized bool, isBlocked bool) bool { 
	if age >= 18 && isAuthorized && !isBlocked { 
		return true 
		} else { 
			return false 
			} 
	} 
	
func statusMessage(status string) string { 
	switch status { 
		case "pending": 
			return "The task is waiting" 
		case "running": 
			return "The task is in progress" 
		case "completed": 
			return "The task is complete" 
		case "failed": 
			return "The task failed" 
		default: 
			return "Unknown status" 
			} 
	}

func operators() {

}