package main

import (
	"flag"
	"log"
)

func main() {
	var org organizer
	flag.StringVar(&org.From, "from", "", "Dir you want to move your files from")
	flag.StringVar(&org.To, "to", "", "Dir you want to move your files to")
	flag.StringVar(&org.Regex, "regex", "2023(\\w\\w)", "Regex")
	flag.BoolVar(&org.Recursive, "recursive", false, "Regex")
	flag.Parse()

	if org.From == "" || org.To == "" {
		log.Fatal("The flags can't be empty strings")
	}

	CleanInput(&org.From)
	CleanInput(&org.To)
	org.DirExists()

	switch org.Recursive {
	case false:
		org.GetEntries()
		org.ParseFiles()
		org.MakeDirs()
		org.Move()
	case true:
		org.recursive()
	}

}
