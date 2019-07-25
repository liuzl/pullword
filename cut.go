package pullword

import (
	"strings"

	"github.com/liuzl/ling"
)

var nlp = ling.MustNLP(ling.Norm)

func Cut(line string) []string {
	d := ling.NewDocument(strings.TrimSpace(line))
	if err := nlp.Annotate(d); err != nil {
		return nil
	}
	return d.XRealTokens(ling.Norm)
}
