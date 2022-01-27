package main

import (
	"flag"
	"fmt"
	"os"
	"sort"
)

var gitignoreTemplates = map[string]struct {
	language string
	template string
}{
	"c":      {"The C programming language gitignore", CGitignoreTemplate},
	"cpp":    {"The C++ programming language gitignore", CppGitignoreTemplate},
	"go":     {"The Go programming language gitignore", GoGitignoreTemplate},
	"java":   {"The Java programming languge gitignore", JavaGitignoreTemplate},
	"node":   {"Node.js gitignore", NodeGitignoreTemplate},
	"python": {"The Python programming language gitignore", PythonGitignoreTemplate},
	"ruby":   {"The Ruby programming languague gitignore", RubyGitignoreTemplate},
	"rust":   {"The Rust programming language gitignore", RustGitignoreTemplate},
}

// listTemplates prints all the available gitignore templates
func listTemplates() {
	keys := make([]string, 0, len(gitignoreTemplates))
	for key := range gitignoreTemplates {
		keys = append(keys, key)
	}

	sort.Strings(keys)
	for _, key := range keys {
		fmt.Printf("%-14s(%s)\n", key, gitignoreTemplates[key].language)
	}
}

// dumpTemplate prints to stdout the language gitignore template defined by `key`
func dumpTemplate(key string, save bool) {
	if save {
		err := os.WriteFile(".gitignore", []byte(gitignoreTemplates[key].template), 0666)
		if err != nil {
			fmt.Printf("write failed: %v\n", err)
			flag.Usage()
		}
		os.Exit(0)
	}
	fmt.Println(gitignoreTemplates[key].template)
}
