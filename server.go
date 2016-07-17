package main

import (
	"os"
	"os/user"

	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/standard"
	"github.com/labstack/echo/middleware"

	"github.com/GeertJohan/go.rice"
	"github.com/Sirupsen/logrus"
)

func init() {
	logrus.SetLevel(logrus.DebugLevel)
	logrus.SetFormatter(&logrus.JSONFormatter{})
}

func Exist(name string) bool {
	if _, err := os.Stat(name); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

func GetUserHomedir() string {
	user, err := user.Current()
	if err != nil {
		logrus.Fatal(err)
	}
	return user.HomeDir
}

func main() {

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.Static("public"))

	dirname := GetUserHomedir() + "/" + "testoutput"
	if !Exist(dirname) {
		if err := os.Mkdir(dirname, 0777); err != nil {
			logrus.Error("Directory error...")
			os.Exit(1)
		}
	}
	assetHandler := http.FileServer(rice.MustFindBox("public").HTTPBox())

	e.GET("/", standard.WrapHandler(assetHandler))
	e.Static("/dl", "testoutput")

	e.POST("/upload", Upload())
	logrus.Info("Listening...")

	e.Run(standard.New(":1323"))
	//log.Println("access to http://localhst:1323")
}
