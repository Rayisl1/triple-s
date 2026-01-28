package config

import (
	"errors"
	"flag"
)

type Config struct {
	Port int
	Dir  string
	Help bool
}

func UsageText() string {
	return `Simple Storage Service.

Usage:
  triple-s [-port <N>] [-dir <S>]
  triple-s --help

Options:
  --help     Show this screen.
  --port N   Port number
  --dir S    Path to the directory
`
}

func Parse(args []string) (Config, error) {
	cfg := Config{}

	flagSet := flag.NewFlagSet("triple-s", flag.ContinueOnError)
	flagSet.SetOutput(nil)

	flagSet.IntVar(&cfg.Port, "port", 8080, "Port number")
	flagSet.StringVar(&cfg.Dir, "dir", "data", "Path to the directory")
	flagSet.BoolVar(&cfg.Help, "help", false, "Show help")

	err := flagSet.Parse(args)
	if err != nil {
		return Config{}, err
	}

	if cfg.Help {
		return cfg, nil
	}

	if cfg.Port < 1 || cfg.Port > 65535 {
		return Config{}, errors.New("wrong port number")
	}

	if cfg.Dir == "" {
		return Config{}, errors.New("directory is empty")
	}

	return cfg, nil
}
