package main

import (
	"flag"
	"fmt"
	"os"
)

func usage() {
	fmt.Printf("\n")
	fmt.Printf("usage: slowloris [OPTIONS]\n")
	fmt.Printf("\n")
	fmt.Printf("OPTIONS\n")
	flag.PrintDefaults()
	fmt.Printf("\n")
	fmt.Printf("EXAMPLES\n")
	fmt.Printf("\t%s -target 127.0.0.1 -connections 500\n", os.Args[0])
	fmt.Printf("\t%s -target 127.0.0.1 -connections 500 -https\n", os.Args[0])
	fmt.Printf("\n")
	fmt.Printf("\n")
	fmt.Printf("\n")
}

func main() {
	flag.Usage = usage
}
