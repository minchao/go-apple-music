package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/minchao/go-apple-music/token"
)

func usage() {
	fmt.Fprintln(os.Stderr, `Usage: token-generator [options]`)
	flag.PrintDefaults()
}

func main() {
	var (
		keyId  string
		teamId string
		ttl    int64
		pk     string
		pkFile string
		secret []byte
	)

	flag.StringVar(&keyId, "k", "", "MusicKit key")
	flag.StringVar(&teamId, "t", "", "Team ID")
	flag.Int64Var(&ttl, "l", 3600, "TTL (time-to-live), must not be greater than 15777000 (6 months in seconds)")
	flag.StringVar(&pk, "pk", "", "MusicKit private key, enter string without BEGIN and END annotations")
	flag.StringVar(&pkFile, "pf", "", "MusicKit private key, the path of private key file (.p8)")

	flag.Usage = usage
	flag.Parse()

	if len(os.Args) < 4 {
		flag.Usage()
		os.Exit(1)
	}

	if pk == "" && pkFile == "" {
		fmt.Fprintln(os.Stderr, "The -pk or -pf is required")
		os.Exit(1)
	}
	if pk != "" {
		pk = fmt.Sprintf("-----BEGIN PRIVATE KEY-----\n%s\n-----END PRIVATE KEY-----", pk)
		secret = []byte(pk)
	} else {
		var err error
		secret, err = ioutil.ReadFile(pkFile)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	}

	gen := token.Generator{
		KeyId:  keyId,
		TeamId: teamId,
		TTL:    ttl,
		Secret: secret,
	}

	t, err := gen.Generate()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	fmt.Fprintf(os.Stdout, "-----APPLE MUSIC API TOKEN-----\n%s\n", t)
	fmt.Fprintf(os.Stdout, "-----CURL EXAMPLE-----\ncurl -v -H 'Authorization: Bearer %s' https://api.music.apple.com/v1/storefronts/tw\n", t)
}
