package main

import (
	"encoding/binary"
	"fmt"
	"net"
	"os"
	"strings"
)

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error:%s", err.Error())
		os.Exit(1)
	}
}

func main() {
	service := "223.5.5.5:53"
	udpAddr, err := net.ResolveUDPAddr("udp", service)
	checkError(err)
	conn, err := net.DialUDP("udp", nil, udpAddr)
	checkError(err)

	question := dnsQuestion{"www.cyeam.com", dnsTypeA, dnsClassINET}
	// question1 := dnsQuestion{"blog.cyeam.com", dnsTypeALL, dnsClassINET}
	out := DnsMsg{}
	out.Id = 2015
	out.Bits |= _RD
	out.Questions = append(out.Questions, question)
	// out.Questions = append(out.Questions, question1)
	// fmt.Println(append(out.Pack(), []byte{1, 0, 0, 1, 0, 0, 0, 0, 0, 0, 3, 119, 119, 119, 5, 99, 121, 101, 97, 109, 3, 99, 111, 109, 0, 0, 1, 0, 1}...))
	fmt.Println(out.Pack())
	_, err = conn.Write(out.Pack())
	// _, err = conn.Write([]byte{11, 117, 1, 0, 0, 1, 0, 0, 0, 0, 0, 0, 3, 119, 119, 119, 5, 99, 121, 101, 97, 109, 3, 99, 111, 109, 0, 0, 1, 0, 1})
	checkError(err)

	buf := []byte{}
	buf = make([]byte, 512)
	n, err := conn.Read(buf[0:])
	checkError(err)
	fmt.Println(buf[0:n])
	fmt.Println(out.Unpack(buf[0:n]))
	os.Exit(0)
}

type DnsMsg struct {
	Id                                 uint16
	Bits                               uint16
	Qdcount, Ancount, Nscount, Arcount uint16
	Questions                          []dnsQuestion
	Answers                            []dnsAnswer
}

type dnsQuestion struct {
	Name   string `net:"domain-name"` // `net:"domain-name"` specifies encoding; see packers below
	Qtype  uint16
	Qclass uint16
}

type dnsAnswer struct {
	Name   uint16
	Qtype  uint16
	Qclass uint16
	QLive  uint32
	QLen   uint16
	CName  string `net:"domain-name"` // `net:"domain-name"` specifies encoding; see packers below
}

func (this *DnsMsg) Pack() []byte {
	bs := make([]byte, 12)
	binary.BigEndian.PutUint16(bs[0:2], this.Id)
	binary.BigEndian.PutUint16(bs[2:4], this.Bits)
	binary.BigEndian.PutUint16(bs[4:6], uint16(len(this.Questions)))
	binary.BigEndian.PutUint16(bs[6:8], this.Ancount)
	binary.BigEndian.PutUint16(bs[8:10], this.Nscount)
	binary.BigEndian.PutUint16(bs[10:12], this.Arcount)

	ds := strings.Split(this.Questions[0].Name, ".")
	for _, d := range ds {
		bs = append(bs, byte(len(d)))
		bs = append(bs, []byte(d)...)
	}
	bs = append(bs, 0)

	temp := make([]byte, 2)
	binary.BigEndian.PutUint16(temp, this.Questions[0].Qtype)
	bs = append(bs, temp...)
	binary.BigEndian.PutUint16(temp, this.Questions[0].Qclass)
	bs = append(bs, temp...)
	return bs
}

func (this *DnsMsg) Unpack(buf []byte) *DnsMsg {
	res := new(DnsMsg)
	res.Id = binary.BigEndian.Uint16(buf[0:2])
	res.Bits = binary.BigEndian.Uint16(buf[2:4])
	res.Qdcount = binary.BigEndian.Uint16(buf[4:6])
	res.Questions = make([]dnsQuestion, int(res.Qdcount))
	res.Ancount = binary.BigEndian.Uint16(buf[6:8])
	res.Answers = make([]dnsAnswer, int(res.Ancount))
	res.Nscount = binary.BigEndian.Uint16(buf[8:10])
	res.Arcount = binary.BigEndian.Uint16(buf[10:12])
	i := 13
	j := 0
	for ; j < int(res.Qdcount); j++ {
		domain_count := int(buf[i-1])
		question := dnsQuestion{}
		for buf[i] != 0 {
			if domain_count > 0 {
				question.Name += string(buf[i:i+domain_count]) + "."
				i += domain_count
				domain_count = 0
			} else {
				domain_count = int(buf[i])
				i++
			}
		}
		i++
		question.Name = strings.TrimRight(question.Name, ".")
		question.Qtype = binary.BigEndian.Uint16(buf[i : i+2])
		question.Qclass = binary.BigEndian.Uint16(buf[i+2 : i+4])
		i += 4
		res.Questions[j] = question
	}
	for j = 0; j < int(res.Ancount); j++ {
		answer := dnsAnswer{}
		answer.Name = binary.BigEndian.Uint16(buf[i : i+2])
		i += 2
		answer.Qtype = binary.BigEndian.Uint16(buf[i : i+2])
		i += 2
		answer.Qclass = binary.BigEndian.Uint16(buf[i : i+2])
		i += 2
		answer.QLive = binary.BigEndian.Uint32(buf[i : i+4])
		i += 4
		answer.QLen = binary.BigEndian.Uint16(buf[i : i+2])
		i += 2

		if answer.Qtype == dnsTypeCNAME {
			domain_count := int(buf[i])
			i++
			for buf[i] != 0 {
				if domain_count > 0 {
					answer.CName += string(buf[i:i+domain_count]) + "."
					i += domain_count
					domain_count = 0
				} else {
					domain_count = int(buf[i])
					i++
				}
			}
			i++
			answer.CName = strings.TrimRight(answer.CName, ".")
		} else if answer.Qtype == dnsTypeA {
			for m := 0; m < int(answer.QLen); m++ {
				answer.CName += fmt.Sprintf("%d.", buf[i+m])
			}
			answer.CName = strings.TrimRight(answer.CName, ".")
		}
		res.Answers[j] = answer
	}

	return res
}

const (
	// dnsHeader.Bits
	_QR = 1 << 15 // query/response (response=1)
	_AA = 1 << 10 // authoritative
	_TC = 1 << 9  // truncated
	_RD = 1 << 8  // recursion desired
	_RA = 1 << 7  // recursion available
)

const (
	// valid dnsRR_Header.Rrtype and dnsQuestion.qtype
	dnsTypeA     = 1
	dnsTypeNS    = 2
	dnsTypeMD    = 3
	dnsTypeMF    = 4
	dnsTypeCNAME = 5
	dnsTypeSOA   = 6
	dnsTypeMB    = 7
	dnsTypeMG    = 8
	dnsTypeMR    = 9
	dnsTypeNULL  = 10
	dnsTypeWKS   = 11
	dnsTypePTR   = 12
	dnsTypeHINFO = 13
	dnsTypeMINFO = 14
	dnsTypeMX    = 15
	dnsTypeTXT   = 16
	dnsTypeAAAA  = 28
	dnsTypeSRV   = 33

	// valid dnsQuestion.qtype only
	dnsTypeAXFR  = 252
	dnsTypeMAILB = 253
	dnsTypeMAILA = 254
	dnsTypeALL   = 255

	// valid dnsQuestion.qclass
	dnsClassINET   = 1
	dnsClassCSNET  = 2
	dnsClassCHAOS  = 3
	dnsClassHESIOD = 4
	dnsClassANY    = 255

	// dnsMsg.rcode
	dnsRcodeSuccess        = 0
	dnsRcodeFormatError    = 1
	dnsRcodeServerFailure  = 2
	dnsRcodeNameError      = 3
	dnsRcodeNotImplemented = 4
	dnsRcodeRefused        = 5
)
