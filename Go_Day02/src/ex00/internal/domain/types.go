package domain

import (
	"flag"
	"os"
)

type Flags struct {
	SymLink  bool
	Dir      bool
	File     bool
	Ext      string
	RootName string
}

func NewFlags() *Flags {
	symLink := flag.Bool("sl", false, "symbolic links only")
	dir := flag.Bool("d", false, "directories only")
	file := flag.Bool("f", false, "files only")
	ext := flag.String("ext", "", "files extension")
	flag.Parse()
	return &Flags{
		SymLink:  *symLink,
		Dir:      *dir,
		File:     *file,
		Ext:      *ext,
		RootName: os.Args[len(os.Args)-1],
	}
}
