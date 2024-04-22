package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	if err := run(os.Args[1:]); err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "timenow: error: %v\n", err)
		os.Exit(1)
	}
}

const help = `timenow is cross-platform command-line tool to print the current time in various formats.

Usage: timenow [flags] [args]

Flags:
`

func run(args []string) error {
	f := flag.NewFlagSet("timenow", flag.ExitOnError)
	f.Usage = func() {
		_, _ = fmt.Fprint(f.Output(), help)
		f.PrintDefaults()
	}

	var (
		format   string
		timezone string
	)

	f.StringVar(&format, "format", "epochs", "time format, one of: epochs, epochs-millis, epochs-nano, rfc3339, date-time, date-only (default: epochs)")
	f.StringVar(&timezone, "timezone", "UTC", "time zone (default: UTC)")

	if err := f.Parse(args); err != nil {
		if errors.Is(flag.ErrHelp, err) {
			f.Usage()
			return nil
		}
		return fmt.Errorf("parse flags: %w", err)
	}

	loc, err := time.LoadLocation(timezone)
	if err != nil {
		return fmt.Errorf("load location: %w", err)
	}

	now := time.Now().In(loc)

	switch v := strings.ToLower(format); v {
	case "epochs":
		fmt.Println(now.Unix())
	case "epochs-millis":
		fmt.Println(now.UnixMilli())
	case "epochs-nano":
		fmt.Println(now.UnixNano())
	case "rfc3339":
		fmt.Println(now.Format(time.RFC3339))
	case "date-time":
		fmt.Println(now.Format(time.DateTime))
	case "date-only":
		fmt.Println(now.Format(time.DateOnly))
	default:
		return fmt.Errorf("unknown format: %s", v)
	}

	return nil
}
