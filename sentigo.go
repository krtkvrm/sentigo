package sentigo

//Score will return the Sentiment Score
func Score(text string) int {
	tokens := tokenize(text)
	sum := 0
	for _, i := range tokens {
		sum += wordlist[i]
	}
	return sum
}
