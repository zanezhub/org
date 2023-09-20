/*
	Parse args
		* From
		* To
	Revisar si existen las carpetas


	* Revisar qué mes tiene

*/

package main

import (
	"flag"
	"fmt"
	"io/fs"
	"os"
	"regexp"
	"slices"
)

func makedirs(arp *[]fs.DirEntry, to *string) {
	pattern := "2023(\\w\\w)"
	re := regexp.MustCompile(pattern)

	var months []string
	for _, e := range *arp {
		match := re.FindStringSubmatch(e.Name())

		if len(match) >= 2 {
			value := match[1] // The value "xx" is in the first capturing group
			if !slices.Contains(months, value) {
				months = append(months, value)
			}
		}

	}

	for _, m := range months {
		os.Mkdir(m, os.ModeDir)
	}

}

func main() {
	frm := flag.String("frm", "folder to read", "string")
	to := flag.String("to", "folder to move", "string")
	flag.Parse()

	// Revisar si existe el folder
	err := os.Chdir(*frm)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Leer contenidos
	entries, err := os.ReadDir(*frm)
	if err != nil {
		fmt.Println(err)
	}

	arp := &entries
	makedirs(arp, to)

}
