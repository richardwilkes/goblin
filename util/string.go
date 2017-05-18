package util

import "bytes"

// QuotedString returns a quoted string suitable for a script statement.
func QuotedString(str string) string {
	var buffer bytes.Buffer
	buffer.WriteString(`"`)
	for _, ch := range ([]rune)(str) {
		switch ch {
		case '"':
			buffer.WriteString(`\"`)
		case '\n':
			buffer.WriteString(`\n`)
		case '\t':
			buffer.WriteString(`\t`)
		default:
			buffer.WriteRune(ch)
		}
	}
	buffer.WriteString(`"`)
	return buffer.String()
}
