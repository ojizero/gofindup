package main

import (
	"fmt"
	"os"

	flag "github.com/spf13/pflag"

	"github.com/ojizero/gofindup"
)

var (
	version = "dev"
	commit = "HEAD"
	date = "unknown"
)

func main() {
	var (
		startdir string
		nametofind string
		printversion bool
	)

	flag.BoolVarP(&printversion, "version", "v", false, "Display version info and exit program")
	flag.StringVarP(&startdir, "startdir", "s", "", "Directory to start searching from, if not given would search from current working directory")
	flag.Parse()
	nametofind = flag.Arg(0)

	if printversion {
		fmt.Printf("Version: %v, Build commit: %v, Released on: %v\n", version, commit, date)
		os.Exit(0)
	}

	findup := gofindup.Findup
	if startdir != "" {
		findup = func(name string) (string, error) {
			return gofindup.FindupFrom(name, startdir)
		}
	}

	found, err := findup(nametofind)
	if err != nil {
		panic(err)
	}

	fmt.Print(found)
}
