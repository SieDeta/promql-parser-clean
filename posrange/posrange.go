package posrange

import "fmt"

// Pos is the position in a string.
// Negative numbers indicate undefined positions.
type Pos int

// PositionRange describes a position in the input string of the parser.
type PositionRange struct {
	Start Pos
	End   Pos
}

// StartPosInput uses the query string to convert the PositionRange into a
// line:col string, indicating when this is not possible if the query is empty
// or the position is invalid. When this is used to convert ParseErr to a string,
// lineOffset is an additional line offset to be added, and is only used inside
// unit tests.
func (p PositionRange) StartPosInput(query string, lineOffset int) string {
	if query == "" {
		return "unknown position"
	}
	pos := int(p.Start)
	if pos < 0 || pos > len(query) {
		return "invalid position"
	}

	lastLineBreak := -1
	line := lineOffset + 1
	for i, c := range query[:pos] {
		if c == '\n' {
			lastLineBreak = i
			line++
		}
	}
	col := pos - lastLineBreak
	return fmt.Sprintf("%d:%d", line, col)
}
