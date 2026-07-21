package main

func findScore(scores map[string]int, name string) (int, bool) { 
	score, found := scores[name] 
	return score, found 
} 

func countWords(words []string) map[string]int { 
	wordCount := make(map[string]int) 
	
	for _, value := range words { 
		wordCount[value]++ 
		} 

	return wordCount 
} 

func removeInactive(users map[string]bool) { 
	for i, status := range users { 
		if status == false { 
			delete(users, i) 
		} 
	} 
}

func maps() {
	
}