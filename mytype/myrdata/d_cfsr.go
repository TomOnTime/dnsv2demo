package myrdata

import "fmt"

type CLOUDFLARESINGLEREDIRECT struct {
	Description string `dns:"string"` // human readable name
	Code        uint16 `dns:"uint16"` // 301, 302, 307, or 308
	When        string `dns:"string"` // matching regex
	Then        string `dns:"string"` // replacement regex
}

func (rr CLOUDFLARESINGLEREDIRECT) String() string {
	// TODO(tlim): Escape the strings.
	return fmt.Sprintf(`%s %03d %s %s`, rr.Description, rr.Code, rr.When, rr.Then)
}

func (rr CLOUDFLARESINGLEREDIRECT) Len() int {
	return 1 + len(rr.Description) + 2 + len(rr.When) + 2 + len(rr.Then)
}
