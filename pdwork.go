package main

import (
	"log"
	"os/exec"
	"path"
	"strings"
)

var outformat = map[string]string{"odt": "odt", "rtf": "rtf", "json": "json", "html": "html", "html5": "html5", "md": "markdown"}

func PandocExec(inputfile, ftype string) string {
	var in string = "testoutput/" + inputfile
	i, j := strings.LastIndex(inputfile, "/")+1, strings.LastIndex(inputfile, path.Ext(inputfile))
	name := inputfile[i:j]

	var outfile string = name + "." + ftype
	var out string = "testoutput/" + name + "." + ftype

	cmd := exec.Command("pandoc", in, "-t", outformat[ftype], "-s", "-o", out)
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
	return outfile

}
