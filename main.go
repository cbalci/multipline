package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) < 2 {
		printArgumentErrorMessage()
	}

	times, err := strconv.Atoi(os.Args[1])
	if err != nil {
		printArgumentErrorMessage()
	}

	err = multipline(os.Stdin, os.Stdout, times)
	if err != nil {
		fmt.Fprintf(os.Stderr, err.Error())
		os.Exit(1)
	}
}

func multipline(r io.Reader, w io.Writer, times int) error {
	stdinReader := bufio.NewReader(r)
	for {
		lineBytes, _, err := stdinReader.ReadLine()
		if err != nil {
			if err == io.EOF {
				break
			}
			return fmt.Errorf("Unable to read line from stdin: %s", err.Error())
		}
		for i := 0; i < times; i++ {
			w.Write(lineBytes)
			w.Write([]byte("\n"))
		}
	}
	return nil
}

func printArgumentErrorMessage() {
	fmt.Fprintln(os.Stderr, "Usage error. multiplne expects integer as first (and only)")
	fmt.Fprintln(os.Stderr, "argument for the number of repetitions. Please see usage below.")
	printUsage()
	os.Exit(1)
}

func printUsage() {
	fmt.Print("\nSummary:\n")
	fmt.Print("\tmultipline is a trivial cmdline utility which reads stdin\n")
	fmt.Print("\tline by line and writes each line to stdout 'n' times.\n\n")
	fmt.Print("Usage:\n\t$>echo \"turok help!\" | multipline 3\n")
	fmt.Print("\tturok help!\n\tturok help!\n\tturok help!\n\n")
}
