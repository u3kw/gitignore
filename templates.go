package main

import _ "embed"

// CGitignoreTemplate is the template for C .gitignore.
//go:embed .templates/C.gitignore
var CGitignoreTemplate string

// CppGitignoreTemplate is the template for C++ .gitignore.
//go:embed .templates/C++.gitignore
var CppGitignoreTemplate string

// GoGitignoreTemplate is the template for Go .gitignore.
//go:embed .templates/Go.gitignore
var GoGitignoreTemplate string

// JavaGitIgnoretTemplate is the template for Java .gitignore.
//go:embed .templates/Java.gitignore
var JavaGitignoreTemplate string

// NodeGitignoreTemplate is the template for Node .gitignore.
//go:embed .templates/Node.gitignore
var NodeGitignoreTemplate string

// PythonGitignoreTemplate is the template for Python .gitignore.
//go:embed .templates/Python.gitignore
var PythonGitignoreTemplate string

// RubyGitignoreTemplate is the template for Ruby .gitignore.
//go:embed .templates/Ruby.gitignore
var RubyGitignoreTemplate string

// RustGitignoreTemplate is the template for Rust .gitignore.
//go:embed .templates/Rust.gitignore
var RustGitignoreTemplate string
