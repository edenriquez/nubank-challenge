package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	textWellFormed = []byte(`line1
line2
line3`)

	textWithLinebreak = []byte(`line1
line2
line3
`)

	textWithSpacesBefore = []byte(`
		line1
    line2
		line3`)

	textWithLinebreakInBetween = []byte(`
		line1
		
		line2
		
		line3
		
		`)
)

func TestParseLines(t *testing.T) {
	res := ParseByteToStringLines(textWellFormed)
	assert.Len(t, res, 3)
}

func TestParseLinesWithBreakLine(t *testing.T) {
	res := ParseByteToStringLines(textWithLinebreak)
	assert.Equal(t, `line1`, res[0])
	assert.Equal(t, `line2`, res[1])
	assert.Equal(t, `line3`, res[2])
}

func TestParseLinesWithSpacesBefore(t *testing.T) {
	res := ParseByteToStringLines(textWithSpacesBefore)
	assert.Equal(t, `line1`, res[0])
	assert.Equal(t, `line2`, res[1])
	assert.Equal(t, `line3`, res[2])
}

func TestParseLinesWithBreaksInBetween(t *testing.T) {
	res := ParseByteToStringLines(textWithLinebreakInBetween)
	assert.Equal(t, `line1`, res[0])
	assert.Equal(t, `line2`, res[1])
	assert.Equal(t, `line3`, res[2])
}
