package main

import (
	"fmt"
	"os"

	"triple-s/internal/config"
	"triple-s/internal/server"
)

func main() {
	cfg, err := config.Parse(os.Args[1:])
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	if cfg.Help {
		fmt.Print(config.UsageText())
		return
	}

	if err := server.Run(cfg); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
