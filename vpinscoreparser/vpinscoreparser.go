package vpinscoreparser

import (
	"fmt"
	"os"
)

type VPinScoreParser struct{}

func NewVPinScoreParser() *VPinScoreParser {
	return &VPinScoreParser{}
}

func (vp *VPinScoreParser) ReadNVRamFile(filePath string) ([]byte, error) {
	if _, err := os.Stat(filePath); err != nil {
		return nil, fmt.Errorf("could not find filePath '%s': %s\n", filePath, err)
	}

	bytes, err := os.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("could not read filePath '%s': %s\n", filePath, err)
	}
	return bytes, nil
}

type Parser interface {
	ParseScore(tableName string, bytes []byte) (int64, error)
}

func (vp *VPinScoreParser) Parse(tableName string, bytes []byte) (int64, error) {
	if tableName == "" {
		return -1, fmt.Errorf("tableName cannot be empty")
	}

	var parser Parser
	switch TableMapping[tableName] {
	case SystemBally:
		parser = NewBallyParser()
	case SystemStern:
		parser = NewSternParser()
	case SystemWilliamS6Lmt5:
		parser = NewWilliamsS6Lmt5Parser()
	case SystemWilliamS6Lmt6:
		parser = NewWilliamsS6Lmt6Parser()
	case SystemWilliamS7:
		parser = NewWilliamsS7Parser()

	default:
		return -1, fmt.Errorf("tableName is not supported")
	}

	return parser.ParseScore(tableName, bytes)
}

// Supported NVRam systems
const (
	SystemBally         = iota // Bally MPU AS-2518-35
	SystemStern                // Stern M-200 MPU
	SystemWilliamS6Lmt5        // Williams System 6, 5 chars
	SystemWilliamS6Lmt6        // Williams System 6, 6 chars
	SystemWilliamS7            // Williams System 7
)

var TableMapping = map[string]int{
	"elektra.nv": SystemBally,
	"fathom.nv":  SystemBally,

	"seawitch.nv": SystemStern,
	"viper.nv":    SystemStern,

	"scrpn_l1.nv": SystemWilliamS6Lmt5,
	"frpwr_b7.nv": SystemWilliamS6Lmt6,

	"barra_l1.nv": SystemWilliamS7,
	"bk_l4.nv":    SystemWilliamS7,
}
