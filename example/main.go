package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/erigontech/speedtest/speedtest"
)

var cloudflareHeaders = http.Header{
	"lsjdjwcush6jbnjj3jnjscoscisoc5s": []string{"I%OSJDNFKE783DDHHJD873EFSIVNI7384R78SSJBJBCCJBC32JABBJCBJK45"},
}

func main() {
	servurl := []string{
		"v1:https://caplin-snapshots-sepolia.erigon.network",
		"v1:https://erigon2-v1-snapshots-sepolia.erigon.network/",
		"v1:https://erigon2-v2-snapshots-sepolia.erigon.network/",
		"v1:https://erigon2-v3-snapshots-sepolia.erigon.network/",
		"v1:https://erigon3-v1-snapshots-sepolia.erigon.network/",
		"v1:https://erigon3-v3-snapshots-sepolia.erigon.network/",
		"v1:https://erigon3-v3-snapshots-sepolia.erigon.network/v2/",
	}

	urlstr, err := speedtest.SelectSegmentFromWebseeds(servurl, cloudflareHeaders)
	if err != nil {
		log.Fatal(err)
	}

	s, err := speedtest.CustomServer(urlstr)
	if err != nil {
		log.Fatal(err)
	}

	checkError(s.PingTest(nil))
	checkError(s.DownloadTest())

	// Note: The unit of s.DLSpeed, s.ULSpeed is bytes per second, this is a float64.
	fmt.Printf("Latency: %s, Download: %s\n", s.Latency, s.DLSpeed)
	s.Context.Reset()
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
