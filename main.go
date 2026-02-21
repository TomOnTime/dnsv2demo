package main

import (
	"fmt"

	"codeberg.org/miekg/dns"
	"codeberg.org/miekg/dns/dnsutil"
	"codeberg.org/miekg/dns/rdata"
	"github.com/TomOnTime/dnsv2demo/mytype"
	"github.com/TomOnTime/dnsv2demo/mytype/myrdata"
)

/*

	The purpose of this program is to demonstrate the functionality that
	dnscontrol needs to adopt dnsv2.RR / dnsv2.RDATA as the native data
	types, replacing the many fields of models.RecordConfig{}.

	1. create an RDATA from the fields:
		rr1 := rdata.MX{ Preference: 10, Mx: "mx.plts.org."}

    2. create an RDATA only knowing the type and a zonefile-like string
		rr2 := parserdata(TypeMX, "10 mx.plts.org.")

    3. Round-trip an RDATA .String() to parserdata() and get back the same string.

    4. All of the above, using private record types.

	5. FUTURE: Test that ZoneParser works with custom types.
*/

func main() {

	mx1 := rdata.MX{Preference: 10, Mx: "mx.plts.org."}
	roundtrip(dns.TypeMX, mx1, parserdataBuiltin)

	mytype.Register()

	yo1 := myrdata.YO{Priority: 10, Yo: "yo!"}
	roundtrip(mytype.MyTypeYO, yo1, parserdataExternal)

	//sr1 := myrdata.CLOUDFLARESINGLEREDIRECT{Code: 301, Description: "Moved Permanently", When: "^http://example.com/(.*)$", Then: "https://example.com/$1"}
	sr1 := myrdata.CLOUDFLARESINGLEREDIRECT{Code: 301, Description: "simple", When: "when", Then: "then"}
	roundtrip(mytype.MyTypeCLOUDFLARESINGLEREDIRECT, sr1, parserdataExternal)

	sr2 := myrdata.CLOUDFLARESINGLEREDIRECT{Code: 301, Description: "Moved Permanently", When: "^http://example.com/(.*)$", Then: "https://example.com/$1"}
	roundtrip(mytype.MyTypeCLOUDFLARESINGLEREDIRECT, sr2, parserdataExternal)

	// TODO(tlim): Test that ZoneParser works with custom types.

}

func roundtrip(typ uint16, r dns.RDATA, parseFn func(uint16, string) (dns.RDATA, error)) {
	defer fmt.Println() // Always print a blank line after the output.

	// Step 1: String() the RDATA
	// Step 2: parserdata() the string back to an RDATA
	// Step 3: String() the new RDATA and compare to the original string

	typStr := dnsutil.TypeToString(typ)

	// Step 1:
	s1 := r.String()
	println("String:", s1)

	// Step 2:
	r2, err := parseFn(typ, s1)
	if err != nil {
		fmt.Printf("Step 2: Error parsing %v %q: %s\n", typStr, s1, err.Error())
		return
	}

	// Step 3:
	s3 := r2.String()
	println("Round-trip String:", s3)
	if s1 != s3 {
		fmt.Println("Step 3: Round-trip mismatch!")
	}
}

// parserdataBuiltin parses an RDATA string based on the type and returns the corresponding RDATA object.
func parserdataBuiltin(typ uint16, s string) (dns.RDATA, error) {
	return dns.NewData(typ, s)
}

// parserdataExternal uses ZoneParser().
func parserdataExternal(typ uint16, s string) (dns.RDATA, error) {
	rr, err := dns.New(fmt.Sprintf(". 0 IN %s %s", dns.TypeToString[typ], s))
	if err != nil {
		fmt.Printf("DEBUG: dns.New failed: %v\n", err)
	}

	return rr.Data(), err
}
