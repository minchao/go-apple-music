# go-apple-music

[![GoDoc](https://godoc.org/github.com/minchao/go-apple-music?status.svg)](https://godoc.org/github.com/minchao/go-apple-music)
[![Build Status](https://travis-ci.org/minchao/go-apple-music.svg?branch=master)](https://travis-ci.org/minchao/go-apple-music)
[![Go Report Card](https://goreportcard.com/badge/github.com/minchao/go-apple-music)](https://goreportcard.com/report/github.com/minchao/go-apple-music)

A Go client library for accessing the [Apple Music API][].

This library is heavily inspired by [go-github][].

## Installation

Use go get.

```go
go get -u github.com/minchao/go-apple-music
```

## Usage

```go
import "github.com/minchao/go-apple-music"
```

Construct a new API client, then use to access the Apple Music API. For example:

```go
ctx := context.Background()
tp := applemusic.Transport{Token: "APPLE_MUSIC_API_TOKEN"}
client := applemusic.NewClient(tp.Client())

// Fetch all the storefronts in alphabetical order
storefronts, _, err := client.Storefront.GetAll(ctx, nil)
```

### Create a developer token

Use the [token generator](examples/token-generator) tool to quickly create a developer token.

    > cd examples/token-genrator
    > go build

Usage:

    > ./generate-toke
    Usage: generate-token [options]
      -k string
            MusicKit key
      -l int
            TTL (time-to-live), must not be greater than 15777000 (6 months in seconds) (default 3600)
      -pf string
            MusicKit private key, the path of private key file (.p8)
      -pk string
            MusicKit private key, enter string without BEGIN and END annotations
      -t string
            Team ID

Run:

    > ./generate-toke \
    > -k=MUSICKIT_KEY \
    > -t=TEAM_ID \
    > -pk=MUSICKIT_PRIVATE_KEY_FILE

### Create a Music User Token

Use the [requestUserToken(forDeveloperToken:completionHandler:)][] method in the StoreKit framework.

## License

See the [LICENSE](LICENSE) file for license rights and limitations (MIT).

[Apple Music API]: https://developer.apple.com/library/content/documentation/NetworkingInternetWeb/Conceptual/AppleMusicWebServicesReference/
[go-github]: https://github.com/google/go-github
[requestUserToken(forDeveloperToken:completionHandler:)]: https://developer.apple.com/documentation/storekit/skcloudservicecontroller/2909079-requestusertokenfordevelopertoke