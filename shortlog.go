package main

import (
	"fmt"
	"log"
	"sort"
	"strings"

	"github.com/benburkert/go-libgit2"
)

func ShortLog() {
	names := sort.StringSlice{}
	subjectsByName := map[string][]string{}

	commits, err := repo.Walk(libgit2.Sorting(libgit2.SortTime))
	if err != nil {
		log.Fatal(err)
	}

	for commit := range commits.C {
		author, err := commit.Author()
		if err != nil {
			log.Fatal(err)
		}

		if cs, ok := subjectsByName[author.Name]; ok {
			subjectsByName[author.Name] = append(cs, commit.Subject())
		} else {
			names = append(names, author.Name)
			subjectsByName[author.Name] = []string{commit.Subject()}
		}
	}

	names.Sort()

	for _, name := range names {
		subjects := subjectsByName[name]
		fmt.Printf("%s (%d):\n", name, len(subjects))

		for i := len(subjects) - 1; i >= 0; i-- {
			fmt.Printf("      %s\n", shortlogSubject(subjects[i]))
		}
		fmt.Println()
	}
}

// git-shortlog subs 's/\\n/\s{5}/g' on multi-line subjects
func shortlogSubject(s string) string {
	return strings.TrimSpace(strings.Replace(s, "\n", "     ", -1))
}
