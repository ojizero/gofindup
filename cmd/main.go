package main

import (
	"fmt"

	flag "github.com/spf13/pflag"

	"github.com/ojizero/gofindup"
)

func main() {
	var (
		startdir string
		nametofind string
	)

	flag.StringVarP(&startdir, "startdir", "s", "", "Directory to start searching from, if not given would search from current working directory")
	flag.Parse()
	nametofind = flag.Arg(0)

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
