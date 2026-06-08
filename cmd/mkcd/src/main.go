package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"unicode"
)

func main() {
	show_help()
	dir := os.Args[1]
	unicode_only_checker(dir)
	directory_exist_check(dir)
	make_directory(dir)
	abs := read_abs_file_path(dir)
	fmt.Println(quote_path_check(abs))
}

func quote_path_check(p string) string {
	if !strings.ContainsAny(p, " \t\n\"'`") {
		return p
	}
	if !strings.Contains(p, "'") {
		return "'" + p + "'"
	}
	p = strings.ReplaceAll(p, "\"", "\\\"")
	return "\"" + p + "\""
}

func show_help() {
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "Usage: mkcd <directory>")
		os.Exit(1)
	}
}

func unicode_only_checker(dir string) {
	for _, r := range dir {
		if unicode.IsControl(r) {
			log.Fatal("control characters are not allowed")
		}
	}
}

func directory_exist_check(dir string) {
	info, err := os.Stat(dir)
	if err == nil && info.IsDir() {
		fmt.Fprintln(os.Stderr, "Directory already exists ✔")
	} else if err == nil {
		log.Fatalf("'%s' exists but is not a directory", dir)
	}
}

func make_directory(dir string) {
	err := os.MkdirAll(dir, 0755)
	if err != nil {
		log.Fatal("Error creating directory:", err)
	}
	fmt.Fprintln(os.Stderr, "Directory created ✔")
}

func read_abs_file_path(dir string) string {
	abs, err := filepath.Abs(dir)
	if err != nil {
		log.Fatal("Error resolving path:", err)
	}
	return abs
}