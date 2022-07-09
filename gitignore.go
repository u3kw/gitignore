package main

import (
	"flag"
	"fmt"
	"io"
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

// dumpTemplate write to w the language gitignore template defined by `key`
func dumpTemplate(key string, w io.Writer) {
	_, err := io.WriteString(w, gitignoreTemplates[key].template)
	if err != nil {
		fmt.Printf("write failed: %v\n", err)
		flag.Usage()
	}

}
