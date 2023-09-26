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

func dir_exists(path *string) bool {
	//Return err?
	_, err := os.Stat(*path)

	if err != nil && os.IsNotExist(err) {
		fmt.Printf("Folder %s does not exist", *path)
		return false
	}

	return true
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

func make_dir(to *string, months *[]string) error {
	err := os.Chdir(*to)
	if err != nil {
		return err
	}

	for _, m := range *months {
		err := os.Mkdir(m, os.ModeDir)
		if err != nil && !os.IsExist(err) {
			return err
		}
	}

	return nil
}

func move(files *[]fs.DirEntry, to *string, from *string, re *regexp.Regexp) error {
	for _, dir := range *files {
		match := (*re).FindStringSubmatch(dir.Name())

		// [?] TODO: Reescribir
		if len(match) >= 2 {
			value := match[1]
			old := *from + "\\" + dir.Name()
			new := *to + "\\" + value + "\\" + dir.Name()

			err := os.Rename(old, new)
			if err != nil {
				return err
			}
		}

	}

	return nil
}

func clean_input(str *string) {
	if strings.HasSuffix(*str, "\\") || strings.HasSuffix(*str, "/") {
		// Eliminar último char
		*str = (*str)[:len(*str)-1]
	}

	if strings.HasPrefix(*str, ".\\") {
		*str = strings.TrimPrefix(*str, ".")
		current, _ := os.Getwd()

		*str = current + *str
	}

}
