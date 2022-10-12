package domain

import (
	"errors"
	"flag"
	"os"
)

type Flags struct {
	W     bool
	M     bool
	L     bool
	Files []string
}

const ErrIllegalFlags = "too many flags"

func NewFlags() (*Flags, error) {
	w := flag.Bool("w", false, "count words")
	m := flag.Bool("m", false, "count symbols")
	l := flag.Bool("l", false, "count lines")
	flag.Parse()
	return Valid(&Flags{
		L:     *l,
		M:     *m,
		W:     *w,
		Files: os.Args[2:],
	})
}
func Valid(flags *Flags) (*Flags, error) {
	if (flags.W && flags.M) || (flags.W && flags.L) || (flags.M && flags.L) {
		return nil, errors.New(ErrIllegalFlags)
	}
	if !(flags.W || flags.M || flags.L) {
		flags.W = true
		flags.Files = os.Args[1:]
	}
	return flags, nil
}
