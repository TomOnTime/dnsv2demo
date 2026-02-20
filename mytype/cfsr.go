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
func (rr *CLOUDFLARESINGLEREDIRECT) Parse(tokens []string, _ string) error {
	if len(tokens) < 4 { // no rdata
		return nil
	}
	c, err := strconv.ParseUint(tokens[0], 10, 16)
	if err != nil || c > 999 {
		return fmt.Errorf("bad CLOUDFLARESINGLEREDIRECT Code")
	}

	rr.SingleRedirect = myrdata.CLOUDFLARESINGLEREDIRECT{Code: uint16(c), Description: tokens[1], When: tokens[2], Then: tokens[3]}
	return nil
}
