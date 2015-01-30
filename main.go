package main

import (
	"os"

	"github.com/benburkert/go-libgit2"
)

var (
	colorDiffCommit = "yellow" // git config color.diff.commit

	fmtDiffCommit = colorFmt(colorDiffCommit)
)

func main() {
	if len(os.Args) < 2 {
		help()
		os.Exit(1)
	}

	libgit2.Init()
	defer libgit2.Shutdown()

	switch os.Args[1] {
	case "log":
		Log()
	default:
		help()
		os.Exit(1)
	}
}

func help() {}
