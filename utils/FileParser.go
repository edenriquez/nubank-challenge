package utils

import "strings"

// ParseByteToStringLines should parse line by line from file into an array of strings
// cases that this function should prevent are:
// - EOF
// - Break lines
func ParseByteToStringLines(data []byte) []string {
	linesByRow := strings.Split(string(data), "\n")
	var linesWithoutEOF []string
	for _, row := range linesByRow {
		cleanRow := strings.TrimSpace(row)
		if len(cleanRow) > 0 {
			linesWithoutEOF = append(linesWithoutEOF, cleanRow)
		}
	}
	return linesWithoutEOF
}

// AppendError should append a given error in a list of errors by pointer
//  in that way we can reuse list from current package that is calling this func
func AppendError(e error, listOfErrors *[]error) {
	if e != nil {
		*listOfErrors = append(*listOfErrors, e)
	}
}
