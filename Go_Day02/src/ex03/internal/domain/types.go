package domain

import (
	"flag"
	"os"
)

type Flags struct {
	A       bool
	Archive string
	Files   []string
}

func NewFlags() *Flags {
	a := flag.Bool("a", false, "Add to archive")
	flag.Parse()

	var arch string
	var files []string
	if *a == true && len(os.Args) > 2 {
		arch = os.Args[2]
		files = os.Args[3:]
	} else {
		arch = ""
		files = os.Args[1:]
	}
	return &Flags{
		A:       *a,
		Archive: arch,
		Files:   files,
	}
}
