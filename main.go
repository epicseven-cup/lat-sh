package main

import (
	"flag"
	"fmt"
	"github.com/epicseven-cup/lat-sh/internal/latsh"
	"os"
)

var source string
var destination string
var overview string

func init() {
	flag.StringVar(&source, "source", "", "The name of the shorten url e.g. dandandondon for https://lat.sh/dandandondon")
	flag.StringVar(&source, "s", "", "The name of the shorten url e.g. dandandondon for https://lat.sh/dandandondon (shorthand)")
	flag.StringVar(&destination, "destination", "", "The destination url e.g https://google.com")
	flag.StringVar(&destination, "d", "", "The destination url e.g https://google.com (shorthand)")
	flag.StringVar(&overview, "overview", "", "Overview status of the url e.g. https://lat.sh/dandandondon")
	flag.StringVar(&overview, "o", "", "Overview status of the url e.g. https://lat.sh/dandandondon (shorthand)")
}

func main() {
	flag.Parse()

	if overview != "" {
		client := latsh.NewDefaultLatShClient()
		err := client.Overview(overview)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		os.Exit(0)
	}

	if source == "" {
		fmt.Println("Please provide a source URL")
		flag.Usage()
		os.Exit(1)
	}
	if destination == "" {
		fmt.Println("Please provide a destination URL")
		flag.Usage()
		os.Exit(1)
	}

	if source != "" && destination != "" {
		client := latsh.NewDefaultLatShClient()
		err := client.CreateUrl(source, destination)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		os.Exit(0)
	}

	os.Exit(0)

}
