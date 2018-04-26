package pullword

import (
	//"fmt"
	"math"
	"strings"
)

// Entropy computes the Shannon entropy of a distribution or the distance between
// two distributions. The natural logarithm is used.
//  - sum_i (p_i * log_e(p_i))
func Entropy(p []float64) float64 {
	if len(p) == 0 {
		return 1 //math.MaxFloat64
	}
	var e float64
	for _, v := range p {
		if v != 0 { // Entropy needs 0 * log(0) == 0
			e -= v * math.Log(v)
		}
	}
	return e
}

func entropy(m map[string]float64) float64 {
	var p []float64
	var total float64
	for _, v := range m {
		total += v
	}
	for _, v := range m {
		p = append(p, v/total)
	}
	return Entropy(p)
}

func Process(m map[string]*Token, total float64) {
	if total < 1 {
		return
	}
	for k, v := range m {
		v.Freq = v.Freq / total * float64(len(strings.Fields(k)))
	}
	calc(m)
}

func calc(m map[string]*Token) {
	for k, v := range m {
		// calculate the degree of polymerization
		terms := strings.Fields(k)
		//fmt.Printf("%+v\n", terms)
		n := len(terms)
		if n <= 1 {
			v.Poly = 1
		} else {
			var max float64 = 0
			for i := 1; i < n; i++ {
				sub1 := strings.Join(terms[:i], " ")
				sub2 := strings.Join(terms[i:], " ")
				//fmt.Printf("sub1=[%s], sub2=[%s]\n", sub1, sub2)
				s := m[sub1].Freq * m[sub2].Freq
				if s > max {
					max = s
				}
			}
			//fmt.Printf("max=%f\n", max)
			if max > 0 {
				v.Poly = v.Freq / max
			} else {
				v.Poly = 1
			}
		}

		// calculate the degree of flexibility
		v.Flex = math.Min(entropy(v.Left), entropy(v.Right))

		// calculate score
		v.Score = v.Freq * v.Poly * v.Flex
	}
}
