package main

import (
	"fmt"
	"io/fs"
	"log"
	"os"
	"regexp"
	"strings"
)

type organizer struct {
	/*
		Maybe
			Files   map[string]fs.DirEntry
		could be better?
	*/

	Files   []fs.DirEntry
	Dirs    []fs.DirEntry
	NewDirs []string

	From      string
	To        string
	Recursive bool

	OldPath string
	NewPath string

	Regex string
}

func CleanInput(input *string) {
	switch {
	case strings.HasSuffix(*input, "\\") || strings.HasSuffix(*input, "/"):
		*input = (*input)[:len(*input)-1]

	case strings.HasPrefix(*input, ".\\"):
		*input = strings.TrimPrefix(*input, ".")
		current, _ := os.Getwd()
		*input = current + *input

	case *input == ".":
		*input, _ = os.Getwd()
	}
}

func (o *organizer) DirExists() {
	err := os.Chdir(o.To) //Return err?

	if err != nil && os.IsNotExist(err) {
		log.Fatalf("Folder %s does not exist", o.To)
	}
}

func (o *organizer) GetEntries() {
	entries, err := os.ReadDir(o.From)
	if err != nil {
		fmt.Println("ERROR: The program will continue with the value it could collect before the error.", err)
	}

	for _, entry := range entries {
		if !entry.IsDir() {
			o.Files = append(o.Files, entry)
		}
	}
}

func (o *organizer) GetAllEntries() {
	// TODO (Recursive)
}

func (o *organizer) ParseFiles() {
	re := regexp.MustCompile(o.Regex)
	for _, f := range o.Files {
		match := re.FindStringSubmatch(f.Name())

		if len(match) >= 2 {
			value := match[1]
			// TODO: Check if there's repeated values
			o.NewDirs = append(o.NewDirs, value)
		}
	}
}

func (o *organizer) MakeDir() {
	err := os.Chdir(o.To)
	if err != nil {
		log.Fatal(err)
	}

	for _, dir := range o.NewDirs {
		err := os.Mkdir(dir, os.ModeDir)
		if err != nil && !os.IsExist(err) {
			log.Fatal(err)
		}
	}

}

func (o *organizer) Move() {
	re := regexp.MustCompile(o.Regex)
	for _, f := range o.Files {
		match := re.FindStringSubmatch(f.Name())
		if len(match) >= 2 {
			// TODO: Rewrite
			value := match[1]
			old := o.From + "\\" + f.Name()
			new := o.To + "\\" + value + "\\" + f.Name()

			err := os.Rename(old, new)
			if err != nil {
				log.Fatal(err)
			}
		}
	}
}
