package main

import (
	"context"
	"log"
	"os"

	"github.com/Brahma-fi/brahma-builder/app/cli"
	_ "github.com/lib/pq"
)

func main() {
	err := cli.BuildCLI().Run(context.Background(), os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
