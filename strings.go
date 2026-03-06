package main

func IsSubsequence(word string, arr []string) bool {
	if len(word) == 0 {
		return true
	}
	j := 0
	for i := 0; i < len(arr) && j < len(word); i++ {
		if string(word[j]) == arr[i] {
			j++
		}
	}
	return j == len(word)
}

func LongestSubstringWithoutRepeating(str string) int {
	lastSeen := make(map[rune]int)
	start := 0
	best := 0
	runes := []rune(str)
	for i, ch := range runes {
		if prev, ok := lastSeen[ch]; ok && prev >= start {
			start = prev + 1
		}
		lastSeen[ch] = i
		if curr := i - start + 1; curr > best {
			best = curr
		}
	}
	return best
}

func ReverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
