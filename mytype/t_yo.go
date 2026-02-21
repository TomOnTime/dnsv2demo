package mytype

import (
	"fmt"
	"strconv"

	"codeberg.org/miekg/dns"
	"codeberg.org/miekg/dns/dnsutil"
	"github.com/TomOnTime/dnsv2demo/mytype/myrdata"
)

// YO is a private RR: www.example.org. IN YO 10 Yo!
type YO struct {
	Hdr dns.Header
	Yo  myrdata.YO
}

// Typer interface.
func (rr *YO) Type() uint16 { return MyTypeYO }

// RR interface.
func (rr *YO) Header() *dns.Header { return &rr.Hdr }
func (rr *YO) Len() int            { return rr.Hdr.Len() + 1 + rr.Yo.Len() }
func (rr *YO) Data() dns.RDATA     { return rr.Yo }
func (rr *YO) Clone() dns.RR       { return &YO{Hdr: rr.Hdr, Yo: rr.Yo} }
func (rr *YO) String() string {
	return rr.Header().Name + "\t" +
		strconv.FormatInt(int64(rr.Header().TTL), 10) + "\t" +
		dnsutil.ClassToString(rr.Header().Class) + "\tYO\t" +
		rr.Yo.String()
}

// Parser interface.
func (rr *YO) Parse(tokens []string, _ string) error {
	fields := myrdata.TokensToFields(tokens)
	// for i, t := range fields {
	// 	fmt.Printf("DEBUG: YO.Fields[%d]: %q\n", i, t)
	// }

	if len(fields) < 2 { // no rdata
		return nil
	}

	i, err := strconv.ParseUint(fields[0], 10, 32)
	if err != nil || i > 255 {
		return fmt.Errorf("bad YO Priority")
	}

	rr.Yo = myrdata.YO{Priority: uint8(i), Yo: fields[1]}
	return nil
}
