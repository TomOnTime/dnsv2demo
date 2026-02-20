package myrdata

import "fmt"

type CLOUDFLARESINGLEREDIRECT struct {
	Description string `dns:"comment"` // human readable name
	Code        uint16 `dns:"uint16"`  // 301, 302, 307, or 308
	When        string `dns:"txt"`     // matching regex
	Then        string `dns:"txt"`     // replacement regex
}

func (rr CLOUDFLARESINGLEREDIRECT) String() string {
	// TODO(tlim): Escape the strings.
	return fmt.Sprintf("%q %03d %q %q", rr.Description, rr.Code, rr.When, rr.Then)
}

func (rr CLOUDFLARESINGLEREDIRECT) Len() int {
	return 1 + len(rr.Description) + 1 + len(rr.When) + 1 + len(rr.Then)
}
