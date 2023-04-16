package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

func main() {
	// to take input from user
	scanner := bufio.NewScanner(os.Stdin)

	//printing out the format of the output
	fmt.Printf("domain, hasMX, hasSPF, spfRecords, hasDMARC, dmarcRecords\n")

	// actually taking the input and calling the checker function
	for scanner.Scan() {
		checkDomain(scanner.Text())
	}

	// error handling
	if err := scanner.Err(); err != nil {
		log.Printf("Error: Couldn't read from input: %v\n", err)
	}
}

func checkDomain(domain string) {
	// declaring variables for output
	var hasMX, hasSPF, hasDMARC bool
	var spfRecords, dmarcRecords string

	// using net.LookupMX to look for mx for the entered domain
	mxRecords, err := net.LookupMX(domain)
	if err != nil {
		log.Printf("Error: %v\n", err)
	}

	// if the mxRecords length is not 0, then there are mxRecords
	if len(mxRecords) > 0 {
		hasMX = true
	}

	// using net.LookupTXT to look for all records for the domain
	txtRecords, err := net.LookupTXT(domain)
	if err != nil {
		log.Printf("Error: %v\n", err)
	}

	// iterating through the records and if there is a record with 'spf1' in it then it is our spfRecords
	for _, record := range txtRecords {
		if strings.HasPrefix(record, "v=spf1") {
			hasSPF = true
			spfRecords = record
			break
		}
	}

	// using net.LookupTXT to look for records with the prefix 'dmarc'
	txtRecords, err = net.LookupTXT("_dmarc." + domain)
	if err != nil {
		log.Printf("Error: %v\n", err)
	}

	// iterating through the records to see if there is a record with 'dmarc1' in it hence making it out dmarcRecords
	for _, record := range txtRecords {
		if strings.HasPrefix(record, "v=dmarc1") {
			hasDMARC = true
			dmarcRecords = record
			break
		}
	}

	// printing the returned values
	fmt.Printf("%v ,%v, %v, %v, %v, %v", domain, hasMX, hasSPF, spfRecords, hasDMARC, dmarcRecords)
}
