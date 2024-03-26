package main

import (
	"fmt"
	"log"

	"parser/vpinscoreparser"
)

func main() {
	vpinScoreParser := vpinscoreparser.NewVPinScoreParser()

	for _, table := range []string{
		"seawitch.nv",
		"viper.nv",
		"elektra.nv",
		"fathom.nv",
		"scrpn_l1.nv",
		"frpwr_b7.nv",
		"barra_l1.nv",
		"bk_l4.nv",
	} {
		hiScore, err := getScore(vpinScoreParser, table)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%s = %d\n", table, hiScore)
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
