package main

import (
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

type organizer struct {
	Files   []fs.DirEntry
	NewDirs []string

	From      string
	To        string
	Recursive bool
	Regex     string
}

func CleanInput(input *string) {
	if strings.HasPrefix(*input, ".\\") {
		*input = strings.TrimPrefix(*input, ".")
		current, _ := os.Getwd()
		*input = current + *input

	} else if strings.HasSuffix(*input, "\\") || strings.HasSuffix(*input, "/") {
		*input = (*input)[:len(*input)-1]

	} else if *input == "." {
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
		fmt.Println("The program will continue with the value it could collect before the error.", err)
	}

	for _, entry := range entries {
		if !entry.IsDir() {
			o.Files = append(o.Files, entry)
		}
	}
}

func (o *organizer) recursive() {
	filepath.WalkDir(o.From, func(path string, d fs.DirEntry, err error) error {

		if d.IsDir() {
			o.From = path
			o.GetEntries()
			o.ParseFiles()
			o.MakeDirs()
			o.Move()
		}

		o.Files = nil
		o.NewDirs = nil
		return nil
	})
}

func (o *organizer) ParseFiles() {
	re := regexp.MustCompile(o.Regex)
	for _, f := range o.Files {
		match := re.FindStringSubmatch(f.Name())

		if len(match) >= 2 {
			value := match[1]
			// TODO: Check if there's repeated values in o.NewDirs

			o.NewDirs = append(o.NewDirs, value)
		}
	}
}

func (o *organizer) MakeDirs() {
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
				fmt.Println(err)
			}
		}
	}
}
