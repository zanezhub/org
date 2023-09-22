// [-] TODO: Compatibilidad con Linux en "\" "/" (Carpetas)
package main

import (
	"flag"
	"fmt"
	"io/fs"
	"os"
	"regexp"
	"slices"
)

func makedirs(arp *[]fs.DirEntry, to *string, from *string) {
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

	os.Chdir(*to)
	for _, m := range months {
		err := os.Mkdir(m, os.ModeDir) // [X] TODO: Revisar si ya se crearon las carpetas y saltar esto
		if err != nil && !os.IsExist(err) {
			fmt.Println(err)
			return
		}
	}

	for _, e := range *arp {
		match := re.FindStringSubmatch(e.Name())

		if len(match) >= 2 {
			value := match[1]

			old := *from + "\\" + e.Name()
			new := *to + "\\" + value + "\\" + e.Name()

			os.Rename(old, new) // [?] TODO: handle error. Tal vez no aplica
		}

	}

}

// [?] TODO: Reescribir
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

	// [X] TODO: Revisar los args por "\" al final de la cadena y quitarlos
	check(frm, to)

	// Leer contenidos
	entries, err := os.ReadDir(*frm)
	if err != nil {
		fmt.Println(err)
	}

	makedirs(&entries, to, frm)

}
