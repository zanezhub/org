// [-] TODO: Compatibilidad con Linux en "\" "/" (Cfilesetas)
package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	frm := flag.String("frm", "", "Folder that you want to read")
	to := flag.String("to", "", "Folder that you want to move to")
	flag.Parse()

	// [-] Reescribir
	if *frm == "" || *to == "" {
		fmt.Println("One of the args is nil")
		return
	}

	// Revisar si existe el folder
	// [?] Reescribir
	err := os.Chdir(*frm)
	if err != nil {
		fmt.Println(err)
		return
	}

	check(frm)
	check(to)

	// Leer contenidos
	entries, err := os.ReadDir(*frm)
	if err != nil {
		fmt.Println(err)
	}

	months, re := get_months(&entries, to, frm)
	make_dir(to, &months)
	move(&entries, to, frm, re)
}
