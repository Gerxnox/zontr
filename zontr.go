package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var domains []string

	// Check if there are command-line arguments
	if len(os.Args) > 1 {
		// Use the provided file path
		filePath := os.Args[1]

		// Open the file
		file, err := os.Open(filePath)
		if err != nil {
			fmt.Println("Error opening file:", err)
			os.Exit(1)
		}
		defer file.Close()

		// Read domains from the file
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			domains = append(domains, scanner.Text())
		}
		if err := scanner.Err(); err != nil {
			fmt.Println("Error reading file:", err)
			os.Exit(1)
		}
	} else {
		// Read domains from stdin
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			domains = append(domains, scanner.Text())
		}
		if err := scanner.Err(); err != nil {
			fmt.Println("Error reading stdin:", err)
			os.Exit(1)
		}
	}

	// Process domains
	for _, domain := range domains {
		// Add your zontr logic here
		fmt.Println("Processing domain:", domain)
	}
}
