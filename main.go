package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"flag"
)

func main() {
	// open file
	f, err := os.Open("textFile.txt")
	if err != nil {
		log.Fatal(err)
	}
	// remember to close the file at the end of the program
	defer f.Close()
	// read the file line by line using scanner
	scanner := bufio.NewScanner(f)
	var lines []string
	for scanner.Scan() {
		// do something with a line
		lines=append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	n := flag.Int("n", 1, "an int")
	flag.Parse()
	if *n>20 {
		*n=20
	} else if *n<1 {
		fmt.Println("You put in an invalid line number")
		os.Exit(4)
	}
	lines = lines[len(lines)-*n:]
	for _,l :=range lines{
		fmt.Println(l)
	}

}
