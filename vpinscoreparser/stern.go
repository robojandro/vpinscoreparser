package vpinscoreparser

import (
	"fmt"
	"strconv"
)

type Stern struct {
	MillionsOffset int
	TenKsOffset    int
	HundredsOffset int
	OnesOffset     int
}

func NewSternParser() *Stern {
	return &Stern{
		MillionsOffset: 202,
		TenKsOffset:    203,
		HundredsOffset: 204,
		OnesOffset:     205,
	}
}

func (s *Stern) ParseScore(tableName string, bytes []byte) (int64, error) {
	var millionsVal, tenKsVal, hundredsVal, onesVal int64

	for i := 0; i < len(bytes)-1; i++ {
		if i >= s.MillionsOffset && i <= s.OnesOffset {
			byt := bytes[i]

			// fmt.Printf("%d  hex [%x] = number [%d] = string [%s]\n", i, byt, byt, string(byt))

			hexStr := fmt.Sprintf("%x", byt)
			var err error
			if i == s.MillionsOffset {
				millionsVal, err = strconv.ParseInt(hexStr, 10, 16)
				if err != nil {
					return -1, err
				}
			}
			if i == s.TenKsOffset {
				tenKsVal, err = strconv.ParseInt(hexStr, 10, 16)
				if err != nil {
					return -1, err
				}
			}
			if i == s.HundredsOffset {
				hundredsVal, err = strconv.ParseInt(hexStr, 10, 16)
				if err != nil {
					return -1, err
				}
			}
			if i == s.OnesOffset {
				onesVal, err = strconv.ParseInt(hexStr, 10, 16)
				if err != nil {
					return -1, err
				}
			}
		}
	}
	/*
		fmt.Printf("millions: %d\n", millionsVal)
		fmt.Printf("100 Ks: %d\n", tenKsVal)
		fmt.Printf("100s: %d\n", hundredsVal)
		fmt.Printf("1s: %d\n", onesVal)
	*/

	score := (millionsVal * 1_000_000) +
		(tenKsVal * 10_000) +
		(hundredsVal * 100) +
		(onesVal)

	// fmt.Printf("score: %d\n", score)
	return score, nil
}
