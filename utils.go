package main

import (
	"fmt"
	"io/fs"
	"os"
	"regexp"
	"slices"
)

func grepMonths(files *[]fs.DirEntry, to *string, from *string) ([]string, *regexp.Regexp) {
	pattern := "2023(\\w\\w)"
	re := regexp.MustCompile(pattern)

	var months []string
	for _, e := range *files {
		match := re.FindStringSubmatch(e.Name())

		if len(match) >= 2 {
			value := match[1] // The value "xx" is in the first capturing group
			if !slices.Contains(months, value) {
				months = append(months, value)
			}
		}

	}

	return months, re
}

func makeDir(to *string, months *[]string) {
	err := os.Chdir(*to)
	if err != nil {
		fmt.Printf("Can't find or change folder ro %s\n %e", *to, err)
	}

	for _, m := range *months {
		// [X] TODO: Revisar si ya se crearon las carpetas y saltar esto
		err := os.Mkdir(m, os.ModeDir)
		if err != nil && !os.IsExist(err) {
			fmt.Println(err)
		}
	}

}

func move(files *[]fs.DirEntry, to *string, from *string, re *regexp.Regexp) {
	for _, e := range *files {
		match := (*re).FindStringSubmatch(e.Name())

		if len(match) >= 2 {
			value := match[1]
			old := *from + "\\" + e.Name()
			new := *to + "\\" + value + "\\" + e.Name()

			// [?] TODO: handle error. Tal vez no aplica
			os.Rename(old, new)
		}

	}

}

// [?] TODO: Reescribir
// [X] TODO: Revisar los args por "\" al final de la cadena y quitarlos
func check(frm *string, to *string) {
	{
		char := (*frm)[len(*frm)-1:]
		if char == "\\" || char == "/" {
			*frm = (*frm)[:len(*frm)-1]
		}
	}

	char := (*to)[len(*to)-1:]
	if char == "\\" || char == "/" {
		*to = (*to)[:len(*to)-1]
	}

}
