package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	frm := flag.String("frm", "", "Folder that you want to read")
	to := flag.String("to", "", "Folder that you want to move the files to")
	flag.Parse()

	clean_input(frm)
	clean_input(to)

	if dir_exists(frm) && dir_exists(to) {
		entries, err := os.ReadDir(*frm)
		if err != nil {
			fmt.Println("Couldn't get all the entries")
			fmt.Printf("The program will continue with the following number entries: %d\n", len(entries))
		}

		months, re := get_months(&entries, to, frm)
		err = make_dir(to, &months)
		if err != nil {
			fmt.Println(err)
		}

		err = move(&entries, to, frm, re)
		if err != nil {
			fmt.Println("Couldn't move/rename, the error is a type *LinkError")
			fmt.Println(err)
			return
		}
	} else {
		return
	}

}
