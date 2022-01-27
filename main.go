package main

import (
	"flag"
	"fmt"
	"os"
)

const (
	usage = `Usage: gitignore [flags] [language]
Flags:
   -help         print help information
   -list         print list of available gitignore templates

Examples:
   gitignore go
   gitignore java
   gitignore -list`
)

var (
	list bool
)

func main() {
	flag.BoolVar(&list, "list", false, "print to stdout the available gitignore templates")
	flag.Usage = func() {
		fmt.Println(usage)
		os.Exit(1)
	}
	flag.Parse()

	run()
}

func run() {
	if flag.NArg() != 1 && !list {
		flag.Usage()
	}

	switch {
	case list:
		listTemplates()
	case len(os.Args[1:]) == 1:
		dumpTemplate(os.Args[1], false)
	default:
		flag.Usage()
	}
}
