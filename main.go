package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/ctrlrsf/logdna"
)

var (
	apiKey = flag.String("key", "", "LogDNA API key")
)

func main() {
	flag.Parse()

	if len(strings.TrimSpace(*apiKey)) == 0 {
		fmt.Println("empty api key")
		os.Exit(1)
	}

	client := logdna.NewClient(logdna.Config{
		APIKey:   *apiKey,
		LogFile:  "sample.log",
		Hostname: "labo.yukpiz.me",
	})
	defer client.Close()

	client.Log(time.Now(), "!!!!!! hello !!!!!!")
	if err := client.Flush(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
