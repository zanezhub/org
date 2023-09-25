package main

import (
	"fmt"
	"io/fs"
	"os"
	"regexp"
	"strings"
)

func is_in_array(arr *[]string, str *string) bool {
	for _, value := range *arr {
		if *str == value {
			return true
		}
	}
	return false
}

func get_months(files *[]fs.DirEntry, to *string, from *string) ([]string, *regexp.Regexp) {
	pattern := "2023(\\w\\w)"
	re := regexp.MustCompile(pattern)

	var months []string
	for _, e := range *files {
		match := re.FindStringSubmatch(e.Name())

		if len(match) >= 2 {
			value := match[1] // The value "xx" is in the first capturing group
			if !is_in_array(&months, &value) {
				months = append(months, value)
			}
		}

	}

	return months, re
}

func make_dir(to *string, months *[]string) {
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

		// [?] TODO: Reescribir
		if len(match) >= 2 {
			value := match[1]
			old := *from + "\\" + e.Name()
			new := *to + "\\" + value + "\\" + e.Name()

			// [?] TODO: handle error. Tal vez no aplica
			os.Rename(old, new)
		}

	}

}

/*
[X] TODO: Reescribir
[X] TODO: Revisar los args por "\" al final de la cadena y quitarlos
*/

func check(str *string) {
	if strings.HasSuffix(*str, "\\") || strings.HasSuffix(*str, "/") {
		*str = (*str)[:len(*str)-1]
	}
}
