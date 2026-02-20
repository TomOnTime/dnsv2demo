package mytype

import "codeberg.org/miekg/dns"

const MyTypeYO = 65281
const NameYO = "YO"

const MyTypeCLOUDFLARESINGLEREDIRECT = 65283
const NameCLOUDFLARESINGLEDIRECT = "CLOUDFLARE_SINGLE_REDIRECT"

func Register() {

	dns.TypeToRR[MyTypeYO] = func() dns.RR { return new(YO) }
	dns.TypeToString[MyTypeYO] = NameYO
	dns.StringToType[NameYO] = MyTypeYO

	dns.TypeToRR[MyTypeCLOUDFLARESINGLEREDIRECT] = func() dns.RR { return new(CLOUDFLARESINGLEREDIRECT) }
	dns.TypeToString[MyTypeCLOUDFLARESINGLEREDIRECT] = NameCLOUDFLARESINGLEDIRECT
	dns.StringToType[NameCLOUDFLARESINGLEDIRECT] = MyTypeCLOUDFLARESINGLEREDIRECT

}
