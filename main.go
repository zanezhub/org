// [-] TODO: Compatibilidad con Linux en "\" "/" (Cfilesetas)
package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	frm := flag.String("frm", "", "string")
	to := flag.String("to", "", "string")
	flag.Parse()

	// Revisar si existe el folder
	err := os.Chdir(*frm)
	if err != nil {
		fmt.Println(err)
	}

	check(frm, to)

	// Leer contenidos
	entries, err := os.ReadDir(*frm)
	if err != nil {
		fmt.Println(err)
	}

	months, re := grepMonths(&entries, to, frm)
	makeDir(to, &months)
	move(&entries, to, frm, re)
}
