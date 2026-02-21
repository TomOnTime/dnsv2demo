package myrdata

import "fmt"

type CLOUDFLARESINGLEREDIRECT struct {
	Description string `dns:"string"` // human readable name
	Code        uint16 `dns:"uint16"` // 301, 302, 307, or 308
	When        string `dns:"string"` // matching regex
	Then        string `dns:"string"` // replacement regex
}

// String returns the string representation of the CLOUDFLARESINGLEREDIRECT RDATA.
// TODO(tlim): Escape the strings.  "%q" is not what a zonefile expects.
func (rr CLOUDFLARESINGLEREDIRECT) String() string {
	return fmt.Sprintf(`%q %03d %q %q`, rr.Description, rr.Code, rr.When, rr.Then)
}

func (rr CLOUDFLARESINGLEREDIRECT) Len() int {
	return 1 + len(rr.Description) + 2 + len(rr.When) + 2 + len(rr.Then)
}
