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
   -o	         save gitignore template in the current directory
Examples:
   gitignore go
   gitignore java
   gitignore -list`
)

var (
	list bool
	out  bool
)

func main() {
	flag.BoolVar(&list, "list", false, "print to stdout the available gitignore templates")
	flag.BoolVar(&out, "o", false, "save gitignore template in the current directory")
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
	case len(os.Args[1:]) == 2:
		dumpTemplate(os.Args[2], out)
	default:
		flag.Usage()
	}
}
