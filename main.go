package main

import (
	"flag"
)

var source string
var destination string

func init() {
	flag.StringVar(&source, "source", "", "source file")
	flag.StringVar(&destination, "destination", "", "destination file")
}

func main() {

}
