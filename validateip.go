package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

//check ip address valid or not
func ip4or6(s string) string {

	for i := 0; i < len(s); i++ {

		switch s[i] {
		case '.':
			return "IPv4"
		case ':':
			return "IPv6"

		}
	}
	return "Invalid ip"
}

//file create and take user input
func CreateFile(filename, text string) {

	file, err := os.Create(filename)

	if err != nil {
		log.Fatalf("failed creating file: %s", err)
	}

	defer file.Close()

	len, err := file.WriteString(text)
	if err != nil {
		log.Fatalf("failed writing to file: %s", err)
	}

	fmt.Printf("\nFile Name: %s", file.Name())
	fmt.Printf("\nLength: %d bytes", len)
}

//read user input from file
func ReadFile(filename string) {

	f, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {

		address := scanner.Text()
		testInput := net.ParseIP(address)

		if testInput.To4() == nil {
			fmt.Printf("%s : is  not a valid IPv4 address\n", testInput)
		}
		if testInput.To16() == nil {
			fmt.Printf("%s : is  not a valid IPv6 address\n", testInput)

		}

		fmt.Printf("%s is : %s \n", address, ip4or6(address))

	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

}

// main function
func main() {

	// user input for filename
	fmt.Println("Enter filename: ")
	var filename string
	fmt.Scanln(&filename)

	// user input for file content
	fmt.Println("Enter text: ")
	inputReader := bufio.NewReader(os.Stdin)
	input, _ := inputReader.ReadString('\n')

	// file is created and read
	CreateFile(filename, input)
	fmt.Println("")

	ReadFile(filename)
}
