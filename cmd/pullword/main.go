package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"sort"
	"strings"

	"github.com/golang/glog"
	"github.com/liuzl/pullword"
	"zliu.org/goutil"
)

var (
	input = flag.String("i", "input.txt", "input file")
)

func main() {
	flag.Parse()
	file, err := os.Open(*input)
	if err != nil {
		glog.Fatal(err)
	}
	defer file.Close()
	br := bufio.NewReader(file)
	m := make(map[string]*pullword.Token)
	var total float64
	for {
		line, c := br.ReadString('\n')
		if c == io.EOF {
			break
		}
		ret, cnt := pullword.GetNGram(1, 5, line)
		total += float64(cnt)
		for k, v := range ret {
			if m[k] == nil {
				m[k] = v
			} else {
				m[k].Freq += v.Freq
				for w, c := range v.Left {
					m[k].Left[w] += c
				}
				for w, c := range v.Right {
					m[k].Right[w] += c
				}
			}
		}
	}
	pullword.Process(m, total)
	var l pullword.WordList
	for k, v := range m {
		if v.Score > 0.5 && v.Flex > 0.5 {
			l = append(l, pullword.Word{goutil.Join(strings.Fields(k)), v})
		}
	}
	sort.Sort(l)
	for _, v := range l {
		fmt.Printf("%s = %+v\n", v.Str, v.Info)
	}
}
