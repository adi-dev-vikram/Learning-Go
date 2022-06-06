package main

import (
	"fmt"
	"log"
	"net"
	"strings"
)

/*func checkForValidTLD(str string) bool {
	etld, im := publicsuffix.PublicSuffix(str)
	var validtld = false
	if im { // ICANN managed
		validtld = true
	} else if strings.IndexByte(etld, '.') >= 0 { // privately managed
		validtld = true
	}
	return validtld
}
*/
func checkDomain(domain string) {
	var hasMX, hasSPF, hasdmarc bool
	var spfRecord, dmarcRecord string

	mxRecords, err := net.LookupMX(domain)

	if err != nil {
		log.Printf("Error: %v\n", err)
	}
	if len(mxRecords) > 0 {
		hasMX = true

	}

	txtRecords, err := net.LookupTXT(domain)

	if err != nil {
		log.Printf("Error: %v\n", err)
	}

	for _, record := range txtRecords {
		if strings.HasPrefix(record, "v=spf1") {
			hasSPF = true
			spfRecord = record
			break
		}

	}

}

func main() {
	res := strings.SplitN("abc@goole.com", "@", 2)

	fmt.Printf("\nSlice 1: %s", res[1])

	/*
		scanner := bufio.NewScanner(os.Stdin)

		fmt.Printf("domain,hasMX,hasSPF,hasDMARC,dmarcRecord\n")

		for scanner.Scan() {
			checkDomain(scanner.Text())
		}

		err := scanner.Err()
		if err != nil {
			log.Fatal("issue in scanning")
		}
		arguments := os.Args

		if len(arguments) == 1 {
			fmt.Println("Please provide an IP address!")
			os.Exit(100)
		}

		IP := arguments[1]
		addr := net.ParseIP(IP)
		if addr == nil {
			fmt.Println("Not a valid IP address!")
			os.Exit(100)
		}

		hosts, err := net.LookupAddr(IP)
		if err != nil {
			fmt.Println(err)
			os.Exit(100)
		}
		for _, hostname := range hosts {
			fmt.Println(hostname)
		}

		hostname := arguments[2]
		IPs, err := net.LookupHost(hostname)
		if err != nil {
			fmt.Println(err)
			os.Exit(100)
		}
		for _, IP := range IPs {
			fmt.Println(IP)
		}

		//DNS-related problems and finding out the status of a domain
		//using NS lookup

		domain := arguments[3]
		NSs, err := net.LookupNS(domain)

		if err != nil {
			fmt.Println(err)
			os.Exit(100)
		}
		for _, NS := range NSs {
			fmt.Println(NS.Host)
		}
	*/
}
