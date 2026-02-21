package myrdata

import "strings"

// TokensToFields transforms a slice of tokens into fields, combining quoted strings.
// TODO(tlim): There may be other tokens we need to be handle. It's probably better
// for the dns library to provide this functionality.
func TokensToFields(tokens []string) []string {
	var fields []string
	inQuotes := false
	var current []string

	for i := 0; i < len(tokens); i++ {
		token := tokens[i]
		if token == `"` {
			if inQuotes {
				// End of quoted field
				fields = append(fields, strings.Join(current, " "))
				current = nil
				inQuotes = false
			} else {
				// Start of quoted field
				inQuotes = true
				current = nil
			}
		} else if inQuotes {
			current = append(current, token)
		} else {
			fields = append(fields, token)
		}
	}
	// Handle unclosed quote
	if inQuotes && len(current) > 0 {
		fields = append(fields, strings.Join(current, " "))
	}

	// de-escape quotes in fields
	for i, f := range fields {
		fields[i] = strings.ReplaceAll(f, `\"`, `"`)
	}
	return fields
}
