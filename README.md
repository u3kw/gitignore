# gitignore

`gitignore` is a command-line tool to generate language-specific gitignore file.

# Install

## Building from source
#### With Go 1.16 or higher:
```sh
$ go install github.com/kemorels/gitignore
```


# Usage
### Print gitignore template
To print the gitignore template to stdout, run the `gitignore` command followed by the language name:
```sh
$ gitignore go
```
### Save to file
Use the `-o` flag to save the gitignore template to the current directory, or use shell's redirection:
```sh
$ gitignore -o go
$ gitignore go > .gitignore
```
# Contributing 
Pull requests for new feautures, bug fixes, and suggestions are always welcome.


Templates source: [collection](https://github.com/github/gitignore)