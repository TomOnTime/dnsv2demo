package myrdata

import "strconv"

type YO struct {
	Priority uint8
	Yo       string `dns:"txt"`
}

// String returns the string representation of the YO RDATA.
// TODO(tlim): Escape the Yo string.
func (rr YO) String() string {
	return strconv.FormatUint(uint64(rr.Priority), 10) + " " + rr.Yo
}

func (rr YO) Len() int {
	return 1 + len(rr.Yo)
}
