package main

import (
	"os"
	"os/exec"
	"path"
	"strings"

	"github.com/Sirupsen/logrus"
)

var outformat = map[string]string{"odt": "odt", "rtf": "rtf", "json": "json", "html": "html", "html5": "html5", "md": "markdown"}

func PandocExec(inputfile, ftype string) string {
	var in string = GetUserHomedir() + "/" + "testoutput/" + inputfile
	i, j := strings.LastIndex(inputfile, "/")+1, strings.LastIndex(inputfile, path.Ext(inputfile))
	name := inputfile[i:j]

	var outfile string = name + "." + ftype
	var out string = GetUserHomedir() + "/" + "testoutput/" + name + "." + ftype
	logrus.Debug(in)
	logrus.Debug(out)
	logrus.Debug(os.Getenv("SNAP"))
	cmd := exec.Command("pandoc", in, "-t", outformat[ftype], "-s", "-o", out)

	logrus.Debug(cmd)
	err := cmd.Run()
	if err != nil {
		logrus.Fatal(err)
	}
	return outfile

}
