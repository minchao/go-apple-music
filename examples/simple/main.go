package main

import (
	"context"
	"flag"
	"fmt"
	"os"

	"github.com/minchao/go-apple-music"
)

func usage() {
	fmt.Fprintln(os.Stderr, `Usage: simple [options]`)
	flag.PrintDefaults()
}

func main() {
	var token string

	flag.StringVar(&token, "t", "", "Apple Music API Token")
	flag.Usage = usage
	flag.Parse()

	if len(os.Args) < 2 {
		flag.Usage()
		os.Exit(1)
	}

	ctx := context.Background()
	tp := applemusic.Transport{Token: token}
	client := applemusic.NewClient(tp.Client())

	// Fetch all the storefronts in alphabetical order
	storefronts, _, err := client.Storefront.GetAll(ctx, nil)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	fmt.Fprintf(os.Stdout, "result: %+v\n", storefronts)
}
