package app

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type UserController struct {
	ResourceController
	user User
}

/* func (u UserController) Model() interface{} {
	return u.user
} */

func (UserController) List(c echo.Context) error {
	var users []User
	DB.Find(&users)

	return c.JSON(http.StatusOK, users)
}

func (u UserController) Show(c echo.Context) error {

	id := c.Param("id")
	DB.Where("id = ?", id).First(&u.user)

	return c.JSON(http.StatusOK, u.user)
}

func (u UserController) Create(c echo.Context) error {

	if err := c.Bind(&u.user); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	DB.Create(&u.user)
	if u.user.ID == 0 {
		return echo.NewHTTPError(http.StatusBadRequest, "Couldn't create user")
	}
	return c.JSON(http.StatusOK, u.user)
}
