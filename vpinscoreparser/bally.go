package vpinscoreparser

import (
	"fmt"
	"strconv"
	"strings"
)

type Bally struct {
	StartOffset int
	EndOffset   int
	Tables      []string
}

func NewBallyParser() *Bally {
	return &Bally{
		StartOffset: 172,
		EndOffset:   166,
	}
}

func (b *Bally) ParseScore(tableName string, bytes []byte) (int64, error) {
	scoreParts := []string{}
	for i := len(bytes) - 1; i >= 0; i-- {
		byt := bytes[i]

		if i <= b.StartOffset && i >= b.EndOffset {
			// fmt.Printf("%d  hex [%x] = number [%d]\n", i, byt, byt)
			bytAsStr := fmt.Sprintf("%x", byt)
			if len(bytAsStr) > 1 {
				trimmed := strings.TrimRight(bytAsStr, "f")
				scoreParts = append(scoreParts, trimmed)
			} else {
				scoreParts = append(scoreParts, "0")
			}
		}
	}
	scoreStr := strings.Join(scoreParts, "")
	// fmt.Println(scoreStr)

	score, err := strconv.ParseInt(scoreStr, 10, 64)
	if err != nil {
		return -1, err
	}
	return score, nil
}
