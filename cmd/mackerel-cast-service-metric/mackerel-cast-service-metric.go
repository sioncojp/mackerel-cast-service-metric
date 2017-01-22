package main

import (
	"log"
	"os"

	"github.com/jessevdk/go-flags"
	m "github.com/sioncojp/mackerel-cast-service-metric"
)

const (
	app = "mackarel-cast-service-metric"
)

// CommandOpts ...Load option.
type CommandOpts struct {
	Config string `long:"config" short:"c" description:"load toml config path" required:"true"`
}

func main() {
	log.SetOutput(os.Stderr)
	log.SetPrefix(app + ": ")

	opts := CommandOpts{}
	if _, err := flags.ParseArgs(&opts, os.Args[1:]); err != nil {
		log.Fatalf("[Failed] %s", err)
	}

	if err := m.Run(opts.Config); err != nil {
		log.Fatalf("[Failed] %s", err)
	}

	log.Println("OK")
}
