package main

import (
	"log"
	"os/exec"
	"path"
	"strings"
)

func PandocExec(inputfile string) string {
	var in string = "testoutput/" + inputfile
	i, j := strings.LastIndex(inputfile, "/")+1, strings.LastIndex(inputfile, path.Ext(inputfile))
	name := inputfile[i:j]

	var outfile string = name + ".odt"
	var out string = "testoutput/" + name + ".odt"

	cmd := exec.Command("pandoc", in, "-s", "-o", out)
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
	return outfile

}
