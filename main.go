package main

import (
	"log"
	"os"

	"github.com/benburkert/go-libgit2"
)

var (
	repo *libgit2.Repository

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

	wd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	repo, err = libgit2.OpenRepository(wd)
	if err != nil {
		log.Fatal(err)
	}

	switch os.Args[1] {
	case "log":
		Log()
	case "shortlog":
		ShortLog()
	default:
		help()
		os.Exit(1)
	}
}

func help() {}
