package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"

	"m3ujson/internal/converter"
)

var (
	version string
)

func main() {
	showVersion := flag.Bool("version", false, "Show version and exit")
	flag.Parse()

	if *showVersion {
		fmt.Println(version)
		os.Exit(0)
	}

	args := flag.Args()

	if len(args) < 1 {
		flag.Usage()
		log.Fatalln("target should be specifically")
	}

	c := converter.New(args[0])

	if err := c.Convert(context.Background(), os.Stdout); err != nil {
		log.Fatalf("convert: %v", err)
	}
}
