package main

import (
	"github.com/labstack/echo"
	"net/http"
)

type User struct {
	Name           string `json:"name"'`
	Surname        string `json:"surname"`
	Github_profile string `json:"github_profile"`
}

func handler(context echo.Context) error {
	me := &User{
		Name:           "Tolga",
		Surname:        "Cesur",
		Github_profile: "github.com/tolgacesur",
	}
	return context.JSON(http.StatusOK, me)
}

func main() {
	echoInstance := echo.New()

	echoInstance.GET("/", handler)

	echoInstance.Logger.Fatal(echoInstance.Start(":3000"))
}
