package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"time"
)

func transform(input string, output time.Duration) (string, error) {
	dur, err := time.ParseDuration(input)
	if err != nil {
		return "", err
	}
	out := float64(dur) / float64(output)
	return fmt.Sprintf("%f", out), nil
}

func main() {
	var timeOutput time.Duration
	var toMs bool
	var toS bool
	var input string
	flag.BoolVar(&toMs, "ms", true, "output in time milliseconds")
	flag.BoolVar(&toS, "s", false, "output in time seconds")
	flag.StringVar(&input, "i", "", "string to transform")
	flag.Parse()

	timeOutput = time.Millisecond
	if toS {
		timeOutput = time.Second
	}

	if input != "" {
		out, err := transform(input, timeOutput)
		if err != nil {
			panic(err)
		}
		fmt.Println(out)
		os.Exit(0)
	}

	scanner := bufio.NewScanner(os.Stdin)
	for {
		scanner.Scan()
		text := scanner.Text()
		if len(text) != 0 {
			out, err := transform(text, timeOutput)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println(out)
			}
		} else {
			// exit if user entered an empty string
			break
		}
	}
}
