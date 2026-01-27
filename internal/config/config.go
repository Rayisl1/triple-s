package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	port := flag.Int("port", 8080, "Port number")
	dir := flag.String("dir", "./data", "Path to the directory")
	help := flag.Bool("help", false, "Show this screen")

	flag.Usage = func() {
		fmt.Printf("Simple Storage Service.\n\n")
		fmt.Printf("**Usage:**\n")
		fmt.Printf("    triple-s [-port <N>] [-dir <S>]\n")
		fmt.Printf("    triple-s --help\n\n")
		fmt.Printf("**Options:**\n")
		fmt.Printf("- --help     Show this screen.\n")
		fmt.Printf("- --port N   Port number\n")
		fmt.Printf("- --dir S    Path to the directory\n")
	}

	flag.Parse()

	if *help {
		flag.Usage()
		os.Exit(0)
	}

	fmt.Printf("Starting Triple-S on port %d, serving directory: %s\n", *port, *dir)
}
