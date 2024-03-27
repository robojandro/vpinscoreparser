package vpinscoreparser_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"vpinscoreparser"
)

func Test_VPinScoreParser(t *testing.T) {
	vpinScoreParser := vpinscoreparser.NewVPinScoreParser("nvram")

	tests := []struct {
		name  string
		score int64
	}{
		{
			name:  "seawitch.nv",
			score: 180340,
		},
		{
			name:  "viper.nv",
			score: 98010,
		},
		{
			name:  "elektra.nv",
			score: 885050,
		},
		{
			name:  "fathom.nv",
			score: 135650,
		},
		{
			name:  "scrpn_l1.nv",
			score: 555555,
		},
		{
			name:  "frpwr_b7.nv",
			score: 87850,
		},
		{
			name:  "barra_l1.nv",
			score: 121360,
		},
		{
			name:  "bk_l4.nv",
			score: 111111,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hiScore, err := getScore(vpinScoreParser, tt.name)
			if err != nil {
				t.Fatal(err)
			}
			assert.Equal(t, hiScore, tt.score)
		})
	}
}

func getScore(vpinScoreParser *vpinscoreparser.VPinScoreParser, table string) (int64, error) {
	bytes, err := vpinScoreParser.ReadNVRamFile(table)
	if err != nil {
		return -1, err
	}

	hiScore, err := vpinScoreParser.Parse(table, bytes)
	if err != nil {
		return -1, err
	}
	return hiScore, nil
}
