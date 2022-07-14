package figure

import (
	"strconv"
	"strings"
)

const signature = "flf2"
const reverseFlag = "1"

var charDelimiters = [3]string{"@", "#", "$"}
var hardBlanksBlacklist = [2]byte{'a', '2'}

func getHeight(metadata string) int {
	datum := strings.Fields(metadata)[1]
	height, _ := strconv.Atoi(datum)
	return height
}

func getBaseLine(metadata string) int {
	datum := strings.Fields(metadata)[2]
	baseLine, _ := strconv.Atoi(datum)
	return baseLine
}

func getHardBlank(metadata string) byte {
	datum := strings.Fields(metadata)[0]
	hardBlank := datum[len(datum)-1]
	if hardBlank == hardBlanksBlacklist[0] || hardBlank == hardBlanksBlacklist[1] {
		return ' '
	}
	return hardBlank
}

func getReverse(metadata string) bool {
	data := strings.Fields(metadata)
	return len(data) > 6 && data[6] == reverseFlag
}

func lastCharLine(text string, height int) bool {
	endOfLine, length := "  ", 2
	if height == 1 && len(text) > 0 {
		length = 1
	}
	if len(text) >= length {
		endOfLine = text[len(text)-length:]
	}
	return endOfLine == strings.Repeat(charDelimiters[0], length) ||
		endOfLine == strings.Repeat(charDelimiters[1], length) ||
		endOfLine == strings.Repeat(charDelimiters[2], length)
}
