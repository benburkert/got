package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	"unicode"

	"github.com/benburkert/go-libgit2"
)

var logOpts = struct {
	abbrevCommit bool
	format       string
}{}

func Log() {
	fs := flag.NewFlagSet("git-log", flag.ExitOnError)
	fs.BoolVar(&logOpts.abbrevCommit, "abbrev-commit", false, "")
	fs.StringVar(&logOpts.format, "pretty", "", "")
	fs.StringVar(&logOpts.format, "format", "", "")
	fs.Parse(os.Args[2:])

	walker, err := repo.Walk(libgit2.Sorting(libgit2.SortTime))
	if err != nil {
		log.Fatal(err)
	}

	display(<-walker.C, true)
	for commit := range walker.C {
		display(commit, false)
	}
}

func display(commit *libgit2.Commit, firstLine bool) {
	switch logOpts.format {
	case "oneline":
		displayOneLine(commit)
	case "short":
	case "medium", "":
		displayMedium(commit, firstLine)
	case "full":
	case "fuller":
	case "email":
	case "raw":
	}
}

func displayOneLine(commit *libgit2.Commit) {
	var (
		cid string
		err error
	)

	if logOpts.abbrevCommit {
		if cid, err = commit.ShortID(); err != nil {
			log.Fatal(err)
		}
	} else {
		cid = commit.String()
	}

	fmtDiffCommit.Printf(cid)
	message := strings.Replace(commit.Subject(), "\n", " ", -1)
	message = strings.TrimRightFunc(message, unicode.IsSpace)
	fmt.Printf(" %s\n", message)
}

func displayMedium(commit *libgit2.Commit, firstLine bool) {
	sig, err := commit.Author()
	if err != nil {
		log.Fatal(err)
	}

	parents, err := commit.Parents()
	if err != nil {
		log.Fatal(err)
	}

	var cid string
	if logOpts.abbrevCommit {
		if cid, err = commit.ShortID(); err != nil {
			log.Fatal(err)
		}
	} else {
		cid = commit.String()
	}

	if !firstLine {
		fmt.Println()
	}

	fmtDiffCommit.Printf("commit %s\n", cid)
	if len(parents) > 1 {
		fmt.Print("Merge:")
		for _, cmt := range parents {
			id, err := cmt.ShortID()
			if err != nil {
				log.Fatal(err)
			}
			fmt.Printf(" %s", id)
		}
		fmt.Println()
	}

	fmt.Printf("Author: %s <%s>\n", sig.Name, sig.Email)
	fmt.Printf("Date:   %s\n", sig.When.Format("Mon Jan 2 15:04:05 2006 -0700"))
	fmt.Print(prettify(commit.Message()))
}

func prettify(msg string) string {
	msg = strings.TrimRight(msg, " \n")
	if len(msg) == 0 {
		return msg
	}

	lines := strings.Split(msg, "\n")
	for i, line := range lines {
		lines[i] = "    " + strings.TrimRight(line, " \r\t\n")
	}

	msg = strings.Join(lines, "\n")
	return "\n" + msg + "\n"
}
