package main

import (
	"flag"
	"fmt"
	"log"
	"regexp"
	"strings"
)

type stringSlice []string

func (s *stringSlice) String() string {
	return "my string slice representation"
}

func (s *stringSlice) Set(value string) error {
	*s = append(*s, value)
	return nil
}

// flag variables
var (
	randFlag        bool
	shuffleFlag     stringSlice
	lengthFlag      int
	separatorFlag   string
	punctuationFlag bool
)

// flag definitions
const (
	randUsage        = "shuffles all words received from text file(s) after merging"
	shuffleUsage     = "shuffles word in each file received before merging"
	lengthUsage      = "set length of the text after merging"
	separatorUsage   = "set separator used to separate words in files provided"
	punctuationUsage = "setting flag to true removes punctuation marks from final text"
)

// initializeFlags is used to initialize all the flags to be used as command line arguments
func initializeFlags() {
	flag.BoolVar(&randFlag, "rand", false, randUsage)
	flag.BoolVar(&randFlag, "r", false, randUsage+" (shorthand)")

	flag.Var(&shuffleFlag, "shuffle", shuffleUsage)
	flag.Var(&shuffleFlag, "shf", shuffleUsage+" (shorthand)")

	flag.IntVar(&lengthFlag, "length", 0, lengthUsage)
	flag.IntVar(&lengthFlag, "l", 0, lengthUsage+" (shorthand)")

	flag.StringVar(&separatorFlag, "separator", " ", separatorUsage)
	flag.StringVar(&separatorFlag, "sep", " ", separatorUsage+" (shorthand)")

	flag.BoolVar(&punctuationFlag, "punctuation", false, punctuationUsage)
	flag.BoolVar(&punctuationFlag, "pun", false, punctuationUsage+" (shorthand)")

	flag.Parse()
}

// executeFlags is used to process command line arguments
// it returns a text string that contains processed text ready to be used in the program
func executeFlags() string {
	var text string

	if len(shuffleFlag) > 0 {
		for _, path := range shuffleFlag {
			text += shuffleWords(readFile(path))
		}
	}
	if len(flag.Args()) > 0 {
		for _, path := range flag.Args() {
			text += readFile(path)
		}
	} else {
		text += randomWords(words)
	}
	if randFlag {
		text = shuffleWords(text)
	}
	if lengthFlag > 0 {
		arr := strings.Split(text, separatorFlag)
		arr = arr[:lengthFlag]
		text = strings.Join(arr, separatorFlag)
	}
	if punctuationFlag {
		reg, err := regexp.Compile("[^a-zA-Z0-9 ]+")
		if err != nil {
			log.Fatal(err)
		}
		text = reg.ReplaceAllString(text, "")
	}

	return text
}

// printFlags prints out current flag values
func printFlags() {
	fmt.Println("rnd:", randFlag)
	fmt.Println("shf:", shuffleFlag)
	fmt.Println("len:", lengthFlag)
	fmt.Println("sep:", separatorFlag)
	fmt.Println("pun:", punctuationFlag)
	fmt.Println("tail:", flag.Args())
}
