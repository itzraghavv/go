package main

func getNameCounts(names []string) map[rune]map[string]int {
	result := map[rune]map[string]int{}

	for _, name := range names {
		nameRune := []rune(name)[0]
		if _, ok := result[nameRune]; !ok {
			result[nameRune] = make(map[string]int)
		}
		result[nameRune][name]++
	}
	return result
}
