package main

import (
	"flag"
)

func main() {
	var org organizer
	flag.StringVar(&org.From, "from", "", "Dir you want to move your files from")
	flag.StringVar(&org.To, "to", "", "Dir you want to move your files to")
	flag.StringVar(&org.Regex, "regex", "", "Regex")
	flag.Parse()

	if org.Regex == "" {
		org.Regex = "2023(\\w\\w)"
	}

	CleanInput(&org.From)
	CleanInput(&org.To)

	org.DirExists()
	org.GetEntries()
	org.ParseFiles()
	org.MakeDirs()
	org.Move()
}
