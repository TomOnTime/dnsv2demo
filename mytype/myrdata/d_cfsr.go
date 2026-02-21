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
	return fmt.Sprintf(`%s %03d %s %s`,
		ZoneEscapeString(rr.Description),
		rr.Code,
		ZoneEscapeString(rr.When),
		ZoneEscapeString(rr.Then))
}

func (rr CLOUDFLARESINGLEREDIRECT) Len() int {
	return len(rr.String())
	//return 1 + len(rr.Description) + 2 + len(rr.When) + 2 + len(rr.Then)
}
