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


	var outfile = os.Stdout
	var err error


	switch {
	case list:
		listTemplates()
	case len(os.Args[1:]) == 2:
		if out {
			outfile, err = os.Create(".gitignore")
			if err != nil {
				fmt.Printf("fail creating .gitignore file: %v",
						   err)
				os.Exit(1)
			}
			defer outfile.Close()
		}

		dumpTemplate(os.Args[2], outfile)
	case len(os.Args[1:]) == 1:
		dumpTemplate(os.Args[1], outfile)
	default:
		flag.Usage()
	}
}
