package vpinscoreparser

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

type Williams struct {
	StartOffset int
	EndOffset   int
}

func NewWilliamsS6Lmt5Parser() *Williams {
	return &Williams{
		StartOffset: 71,
		EndOffset:   77,
	}
}

func NewWilliamsS6Lmt6Parser() *Williams {
	return &Williams{
		StartOffset: 71,
		EndOffset:   78,
	}
}

func NewWilliamsS7Parser() *Williams {
	return &Williams{
		StartOffset: 103,
		EndOffset:   109,
	}
}

func (w *Williams) ParseScore(tableName string, bytes []byte) (int64, error) {
	scoreParts := []string{}
	for i := 0; i < len(bytes); i++ {
		if i > w.StartOffset && i <= w.EndOffset {
			byt := bytes[i]
			// fmt.Printf("%d  hex [%x] = number [%d]\n", i, byt, byt)

			bytAsStr := fmt.Sprintf("%x", byt)
			if len(bytAsStr) > 1 {
				trimmed := strings.TrimLeft(bytAsStr, "f")
				scoreParts = append(scoreParts, trimmed)
			} else {
				scoreParts = append(scoreParts, "0")
			}
		}
	}
	// fmt.Println("string parts: ", scoreParts)

	scoreStr := strings.Join(scoreParts, "")
	score, err := strconv.ParseInt(scoreStr, 10, 64)
	if err != nil {
		log.Fatal(err)
	}
	// fmt.Printf("score: %d\n", score)
	return score, nil
}
