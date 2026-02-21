package mytype

import (
	"fmt"
	"strconv"

	"codeberg.org/miekg/dns"
	"codeberg.org/miekg/dns/dnsutil"
	"github.com/TomOnTime/dnsv2demo/mytype/myrdata"
)

type CLOUDFLARESINGLEREDIRECT struct {
	Hdr            dns.Header
	SingleRedirect myrdata.CLOUDFLARESINGLEREDIRECT
}

// Typer interface.
func (rr *CLOUDFLARESINGLEREDIRECT) Type() uint16 { return MyTypeCLOUDFLARESINGLEREDIRECT }

// RR interface.
func (rr *CLOUDFLARESINGLEREDIRECT) Header() *dns.Header { return &rr.Hdr }
func (rr *CLOUDFLARESINGLEREDIRECT) Len() int {
	return rr.Hdr.Len() + 1 + rr.SingleRedirect.Len()
}
func (rr *CLOUDFLARESINGLEREDIRECT) Data() dns.RDATA { return rr.SingleRedirect }
func (rr *CLOUDFLARESINGLEREDIRECT) Clone() dns.RR {
	return &CLOUDFLARESINGLEREDIRECT{Hdr: rr.Hdr, SingleRedirect: rr.SingleRedirect}
}
func (rr *CLOUDFLARESINGLEREDIRECT) String() string {
	return rr.Header().Name + "\t" +
		strconv.FormatInt(int64(rr.Header().TTL), 10) + "\t" +
		dnsutil.ClassToString(rr.Header().Class) + "\tCLOUDFLARESINGLEREDIRECT\t" +
		rr.SingleRedirect.String()
}

// Parser interface.
// TODO(tlim): Implement unescaping of description, when, then fields.
func (rr *CLOUDFLARESINGLEREDIRECT) Parse(tokens []string, _ string) error {
	fields := myrdata.TokensToFields(tokens)
	// for i, t := range tokens {
	// 	fmt.Printf("DEBUG: CLOUDFLARESINGLEREDIRECT.Token[%d]: %q\n", i, t)
	// }

	if len(fields) < 4 { // no rdata
		return nil
	}

	desc := fields[0]
	code, err := strconv.ParseUint(fields[1], 10, 16)
	if err != nil || code > 999 {
		return fmt.Errorf("bad CLOUDFLARESINGLEREDIRECT Code")
	}
	when := fields[2]
	then := fields[3]

	//fmt.Printf("CLOUDFLARESINGLEREDIRECT.Fields: %q %03d %q %q\n", desc, code, when, then)

	rr.SingleRedirect = myrdata.CLOUDFLARESINGLEREDIRECT{
		Description: desc,
		Code:        uint16(code),
		When:        when,
		Then:        then,
	}
	return nil
}
