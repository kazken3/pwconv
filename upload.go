package main

import (
	"fmt"
	"io"
	"os"

	"net/http"

	"github.com/labstack/echo"
)

func Upload() echo.HandlerFunc {
	return func(c echo.Context) error {
		ftype := c.FormValue("format")
		//-----------
		// Read file
		//-----------

		// Source
		file, err := c.FormFile("file")
		if err != nil {
			return err
		}
		src, err := file.Open()
		if err != nil {
			return err
		}
		defer src.Close()

		// Destination
		dst, err := os.Create("testoutput/" + file.Filename)
		if err != nil {
			return err
		}
		defer dst.Close()

		// Copy
		if _, err = io.Copy(dst, src); err != nil {
			return err
		}
		out := PandocExec(file.Filename, ftype)
		return c.HTML(http.StatusOK, fmt.Sprintf("<p>File %s uploaded and convert successfully <a href=\"dl/%s\">%s</a> .</p><a href=\"/\">top</a>", file.Filename, out, out))
	}
}
