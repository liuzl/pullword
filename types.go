package pullword

import "fmt"

type Token struct {
	Freq  float64
	Poly  float64
	Flex  float64
	Score float64
	Left  map[string]float64
	Right map[string]float64
}

func (t *Token) String() string {
	return fmt.Sprintf("freq:%f,poly:%f,flex:%f,score:%f",
		t.Freq, t.Poly, t.Flex, t.Score)
}

type Word struct {
	Str  string
	Info *Token
}

type WordList []Word

func (l WordList) Swap(i, j int)      { l[i], l[j] = l[j], l[i] }
func (l WordList) Len() int           { return len(l) }
func (l WordList) Less(i, j int) bool { return l[i].Info.Score > l[j].Info.Score }
