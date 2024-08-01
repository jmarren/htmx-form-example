package routers

import (
	"fmt"

	"github.com/jmarren/htmx-form-example/internal/render"
	"github.com/labstack/echo/v4"
)

func GetSignInForm(c echo.Context) error {
	return render.RenderTemplate(c, "sign-in-form", nil)
}

func SignIn(c echo.Context) error {
	fmt.Println("hit POST/ sign-in")

	for i := 0; i < 1000000; i++ {
		fmt.Println(i)
	}

	c.Response().Header().Set("Hx-Push-Url", "/list/dashboard")
	return render.RenderTemplate(c, "sign-in-success", nil)
}
