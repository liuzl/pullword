package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"sort"
	"strings"

	"github.com/cheggaaa/pb"
	"github.com/golang/glog"
	"github.com/liuzl/pullword"
	"zliu.org/goutil"
)

var (
	input = flag.String("i", "input.txt", "input file")
	o     = flag.String("o", "output.txt", "output file")
)

func main() {
	flag.Parse()
	count, err := goutil.FileLineCount(*input)
	if err != nil {
		glog.Fatal(err)
	}
	file, err := os.Open(*input)
	if err != nil {
		glog.Fatal(err)
	}
	defer file.Close()
	br := bufio.NewReader(file)
	m := make(map[string]*pullword.Token)
	var total float64
	bar := pb.StartNew(count)
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
		bar.Increment()
	}
	bar.FinishPrint("calculate ngram done!")
	pullword.Process(m, total)
	var l pullword.WordList
	for k, v := range m {
		if v.Score > 0.5 && v.Flex > 0.5 {
			l = append(l, pullword.Word{goutil.Join(strings.Fields(k)), v})
		}
	}
	fmt.Printf("found %d words\n", len(l))
	if len(l) > 0 {
		out, err := os.Create(*o)
		if err != nil {
			glog.Fatal(err)
		}
		defer out.Close()
		sort.Sort(l)
		for _, v := range l {
			fmt.Fprintf(out, "%s = %+v\n", v.Str, v.Info)
		}
	}
}
