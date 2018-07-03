package sentigo

import (
	"regexp"
	"strings"
)

func replace(text, rep, exp string) string {
	var re = regexp.MustCompile(exp)
	s := re.ReplaceAllString(text, rep)
	return s
}

func tokenize(text string) []string {
	text = strings.ToLower(text)
	text = replace(text, " ", `/\n/g`)
	text = replace(text, "", `/[.,\/#!$%\^&\*;:{}=_\"~()]/g`)
	return strings.Split(text, " ")
}

func score(text string) int {
	tokens := tokenize(text)
	sum := 0
	for _, i := range tokens {
		sum += wordlist[i]
	}
	return sum
}