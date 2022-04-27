package chip8

import (
	_ "embed"
	"log"
	"strconv"
	"strings"
)

const totalFontSize = 5 * 16 // 16 chars each 5 bytes

//go:embed font.bin
var fontFile string

func getFontData() []byte {
	font := make([]byte, 0, totalFontSize)
	chars := strings.Split(fontFile, ",")

	for _, char := range chars {
		number := strings.Replace(char, "0x", "", -1)
		number = strings.Trim(number, " ")
		if len(number) != 2 {
			continue
		}

		realNumber, err := strconv.ParseUint(number, 16, 8)

		if err != nil {
			log.Fatal(err)
		}

		font = append(font, byte(realNumber))
	}

	return font
}
