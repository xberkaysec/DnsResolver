package main

import (
	"fmt"
	"log"
	"time"

	"github.com/miekg/dns"
)

func DNS_Resolver(domain string, queryType uint16) []dns.RR {
	msg := new(dns.Msg)
	msg.SetQuestion(dns.Fqdn(domain), queryType)
	msg.RecursionDesired = true

	client := &dns.Client{Timeout: 10 * time.Second}

	response, _, err := client.Exchange(msg, "1.1.1.1:53")

	if err != nil {
		log.Fatalf("[CRITICAL ERROR] : %v ", err)
		return nil
	}

	if response == nil {
		log.Fatalf("[CRITICAL ERROR] : no reply from the server\n")
		return nil
	}

	return response.Answer
}

func RRToString(rr dns.RR) string {
	switch rr := rr.(type) {
	case *dns.A:
		return fmt.Sprintf("Domain: %s\nTTL: %d\nClass: %s\nQuery Type: A\nIP Address: %s\n",
			rr.Hdr.Name, rr.Hdr.Ttl, dns.Class(rr.Hdr.Class).String(), rr.A.String())
	default:
		return "Unknown record type!"
	}
}

func main() {

	var domain string
	fmt.Printf("Enter a domain : ")
	fmt.Scanln(&domain)
	fmt.Println("")

	answers := DNS_Resolver(domain, dns.TypeA)

	for _, answer := range answers {
		fmt.Printf(RRToString(answer))
		fmt.Print("-------------------------\n")
	}
}

