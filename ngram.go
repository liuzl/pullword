package pullword

import (
	"strings"
)

func GetNGramFromArray(min, max int, words []string) map[string]*Token {
	if min <= 0 || max <= 0 || min > max {
		return nil
	}
	dict := make(map[string]*Token)
	n := len(words)
	for i := 0; i < n; i++ {
		for j := min; j <= max; j++ {
			if i+j > n {
				break
			}
			k := strings.Join(words[i:i+j], " ")
			if dict[k] == nil {
				dict[k] = &Token{Left: make(map[string]float64),
					Right: make(map[string]float64)}
			}
			dict[k].Freq++
			if i > 0 {
				dict[k].Left[words[i-1]]++
			}
			if i+j < n {
				dict[k].Right[words[i+j]]++
			}
		}
	}
	return dict
}

func GetNGram(min, max int, input string) (map[string]*Token, int) {
	words := Cut(input)
	return GetNGramFromArray(min, max, words), len(words)
}
