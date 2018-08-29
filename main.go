package main

import (
	"bufio"
	"cloud.google.com/go/logging"
	"encoding/json"
	"github.com/jessevdk/go-flags"
	"golang.org/x/net/context"
	"log"
	"os"
)

var opts struct {
	LogID     string `short:"l" long:"log-id" description:"LOG_ID" required:"true"`
	ProjectID string `short:"p" long:"project-id" description:"PROJECT_ID" required:"true"`
}

func main() {
	_, err := flags.ParseArgs(&opts, os.Args)
	if err != nil {
		if flagsErr, ok := err.(*flags.Error); ok && flagsErr.Type == flags.ErrHelp {
			os.Exit(0)
		} else {
			log.Fatalf("%v", err)
		}
	}

	ctx := context.Background()
	client, err := logging.NewClient(ctx, opts.ProjectID)
	if err != nil {
		os.Exit(1)
	}
	defer client.Close()

	lg := client.Logger(opts.LogID)
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		str := scanner.Text()
		j := []byte(str)

		var decode_data interface{}
		if err := json.Unmarshal(j, &decode_data); err != nil {
			lg.Log(logging.Entry{Payload: str})
		} else {
			lg.Log(logging.Entry{Payload: json.RawMessage(j)})
		}
	}
}
