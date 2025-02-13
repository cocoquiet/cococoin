package main

import (
	"github.com/cocoquiet/cococoin/explorer"
	"github.com/cocoquiet/cococoin/rest"
)

func main() {
	go explorer.Start(8000)
	rest.Start(4000)
}
