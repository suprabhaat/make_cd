package main

import (
	"fmt"
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

  fmt.Println("Absolute path:")
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
	fmt.Println("Usage: mkcd <directory>")
	os.Exit(1)
  } 
}

func unicode_only_checker(dir string) {
  for _, r := range dir {
	if unicode.IsControl(r) {
		fmt.Println("control characters are not allowed")
		os.Exit(1)
	}
  }
}

func directory_exist_check(dir string) {
  info, err := os.Stat(dir)
  
  if err == nil && info.IsDir() {
	fmt.Println("Directory already exists ✔")
  } else {
	fmt.Println("Directory created ✔")
  } 
}

func make_directory(dir string){
  err := os.MkdirAll(dir, 0755)
  
  if err != nil {
	fmt.Println("Error creating directory:", err)
	  os.Exit(1)
  }
}

// abs = absolute
func read_abs_file_path(dir string) string{
  abs, err := filepath.Abs(dir)
  
  if err != nil {
	fmt.Println("Error resolving path:", err)
	os.Exit(1)
  }
  return abs
}