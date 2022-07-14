package figure

import (
	"io"
	"log"
	"reflect"
	"strings"
)

const asciiOffset = 32
const firstASCII = ' '
const lastASCII = '~'

// Figure prints beautiful ASCII art from text
type Figure struct {
	phrase string
	font
	strict bool
	color  string
}

// NewFigure create an instance of Figure.
func NewFigure(phrase, fontName string, strict bool) Figure {
	font := newFont(fontName)
	if font.reverse {
		phrase = reverse(phrase)
	}
	return Figure{phrase: phrase, font: font, strict: strict}
}

// NewColorFigure create an instance with the console color of Figure.
func NewColorFigure(phrase, fontName string, color string, strict bool) Figure {
	color = strings.ToLower(color)
	if _, found := colors[color]; !found {
		log.Fatalf("invalid color. must be one of: %s", reflect.ValueOf(colors).MapKeys())
	}

	fig := NewFigure(phrase, fontName, strict)
	fig.color = color
	return fig
}

// NewFigureWithFont create an instance with a custom font
func NewFigureWithFont(phrase string, reader io.Reader, strict bool) Figure {
	font := newFontFromReader(reader)
	if font.reverse {
		phrase = reverse(phrase)
	}
	return Figure{phrase: phrase, font: font, strict: strict}
}

// Slicify return the slice of strings
func (fig Figure) Slicify() (rows []string) {
	for r := 0; r < fig.font.height; r++ {
		printRow := ""
		for _, char := range fig.phrase {
			if char < firstASCII || char > lastASCII {
				if fig.strict {
					log.Fatal("invalid input.")
				} else {
					char = '?'
				}
			}
			fontIndex := char - asciiOffset
			charRowText := scrub(fig.font.letters[fontIndex][r], fig.font.hardBlank)
			printRow += charRowText
		}
		if r < fig.font.baseLine || len(strings.TrimSpace(printRow)) > 0 {
			rows = append(rows, strings.TrimRight(printRow, " "))
		}
	}
	return rows
}

func scrub(text string, char byte) string {
	return strings.Replace(text, string(char), " ", -1)
}

func reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
