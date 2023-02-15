package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"sync"
)

func resolveSubdomain(subdomain string, ips chan<- map[string]string, wg *sync.WaitGroup) {
	defer wg.Done()

	ip, err := net.LookupIP(subdomain)
	if err != nil {
		fmt.Printf("Error resolving %s: %s\n", subdomain, err)
		return
	}
	ips <- map[string]string{subdomain: ip[0].String()}
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Please provide the path to the subdomains file as an argument")
		return
	}

	file, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Printf("Error opening file: %s\n", err)
		return
	}
	defer file.Close()
	fmt.Println("Resolving subdomains into IP ....")
	scanner := bufio.NewScanner(file)

	ips := make(map[string]string)
	ipsChan := make(chan map[string]string)

	var wg sync.WaitGroup

	for scanner.Scan() {
		subdomain := scanner.Text()
		wg.Add(1)
		go resolveSubdomain(subdomain, ipsChan, &wg)
	}

	go func() {
		wg.Wait()
		close(ipsChan)
	}()

	for ip := range ipsChan {
		for subdomain, subdomainIP := range ip {
			ips[subdomain] = subdomainIP
		}
	}

	if len(os.Args) >= 3 && os.Args[2] == "-o" {
		outputFile, err := os.Create("output.txt")
		if err != nil {
			fmt.Printf("Error creating output file: %s\n", err)
			return
		}
		defer outputFile.Close()

		for subdomain, ip := range ips {
			_, err := outputFile.WriteString(fmt.Sprintf("%s: %s\n", subdomain, ip))
			if err != nil {
				fmt.Printf("Error writing to output file: %s\n", err)
				return
			}
		}
		fmt.Println("Output saved to output.txt")
	} else {
		fmt.Println(ips)
	}
}
