package cli

import (
	"flag"
	"fmt"
	"runtime"

	"github.com/cocoquiet/cococoin/explorer"
	"github.com/cocoquiet/cococoin/rest"
)

func usage() {
	fmt.Printf("Welcome to CoCo Coin\n\n")
	fmt.Printf("Please use the following flags:\n\n")
	fmt.Printf("-port:		Set port of the server\n")
	fmt.Printf("-mode:		Choose between 'html' and 'rest'\n\n")

	runtime.Goexit()
}

func Start() {
	port := flag.Int("port", 4000, "Set port of the server")
	mode := flag.String("mode", "rest", "Choose between 'html' and 'rest'")

	flag.Parse()

	switch *mode {
	case "rest":
		rest.Start(*port)
	case "html":
		explorer.Start(*port)
	default:
		usage()
	}
}
