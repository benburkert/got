package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/benburkert/go-libgit2"
)

var logOpts = struct {
	abbrevCommit bool
}{}

func Log() {
	fs := flag.NewFlagSet("git-log", flag.ExitOnError)
	fs.BoolVar(&logOpts.abbrevCommit, "abbrev-commit", false, "")
	fs.Parse(os.Args[2:])

	wd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	repo, err := libgit2.OpenRepository(wd)
	if err != nil {
		log.Fatal(err)
	}

	walker, err := repo.Walk(libgit2.Sorting(libgit2.SortTime))
	if err != nil {
		log.Fatal(err)
	}

	display(<-walker.C)
	for commit := range walker.C {
		fmt.Println()
		display(commit)
	}
}

func display(commit *libgit2.Commit) {
	sig, err := commit.Author()
	if err != nil {
		log.Fatal(err)
	}

	parents, err := commit.Parents()
	if err != nil {
		log.Fatal(err)
	}

	_, err = commit.ShortID()
	if err != nil {
		log.Fatal(err)
	}

	var cid string
	if logOpts.abbrevCommit {
		if cid, err = commit.ShortID(); err != nil {
			panic(err)
		}
	} else {
		cid = commit.String()
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
