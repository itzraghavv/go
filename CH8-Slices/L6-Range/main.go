package main

func indexOfFirstBadWord(msg []string, badWords []string) int {
	for i, word := range msg {
		for _, badWord := range badWords {
			if word == badWord {
				return i
			}
		}
	}
	return -1
}

func main() {
	indexOfFirstBadWord([]string{"hey", "there", "john"}, []string{"crap", "shoot", "frick", "dang"})
}
