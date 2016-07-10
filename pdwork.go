package main

import (
	"log"
	"os/exec"
)

func PandocExec(inputfile string) string {
	var in string = "./public/work/" + inputfile
	var outfile string = inputfile + ".odt"
	var out string = in + ".odt"
	cmd := exec.Command("pandoc", in, "-s", "-o", out)
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
	return outfile

}
