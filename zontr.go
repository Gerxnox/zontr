package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func checkZoneTransfer(domain string) {
	fmt.Printf("Checking zone transfer for: %s\n", domain)

	// Perform a DNS zone transfer
	cmd := exec.Command("dig", "axfr", fmt.Sprintf("@%s", domain))
	output, err := cmd.CombinedOutput()

	if err != nil {
		fmt.Printf("Error checking zone transfer for %s: %v\n", domain, err)
		return
	}

	result := string(output)

	// Check if the result contains a transfer failed message
	if strings.Contains(result, "Transfer failed") {
		fmt.Printf("Zone transfer failed for %s\n", domain)
	} else {
		fmt.Printf("Zone transfer may be possible for %s\n", domain)

		// Save the result to a file
		fileName := fmt.Sprintf("zone_transfer_%s.txt", domain)
		err := os.WriteFile(fileName, []byte(result), 0644)
		if err != nil {
			fmt.Printf("Error saving zone transfer result for %s: %v\n", domain, err)
		} else {
			fmt.Printf("Zone transfer result saved to %s\n", fileName)
		}
	}

	fmt.Println("----------------------------------------------")
}

func main() {
	// Replace 'domains.txt' with the path to your domains list file
	domainsFile := "domains.txt"

	file, err := os.Open(domainsFile)
	if err != nil {
		fmt.Printf("Error opening file %s: %v\n", domainsFile, err)
		os.Exit(1)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		domain := scanner.Text()
		checkZoneTransfer(domain)
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading file %s: %v\n", domainsFile, err)
		os.Exit(1)
	}
}
