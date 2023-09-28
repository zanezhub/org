package main

import (
	"flag"
)

func main() {
	frm := flag.String("frm", "", "Folder that you want to read")
	to := flag.String("to", "", "Folder that you want to move the files to")
	flag.Parse()

	clean_input(frm)
	clean_input(to)

	app(frm, to)

}
