package main

import (
	"github.com/cocoquiet/cococoin/cli"
	"github.com/cocoquiet/cococoin/db"
)

func main() {
	defer db.Close()

	cli.Start()
}
