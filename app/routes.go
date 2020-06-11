package app

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gopkg.in/go-playground/validator.v9"
)

type Validator struct {
	validator *validator.Validate
}

func (v *Validator) Validate(i interface{}) error {
	return v.validator.Struct(i)
}

type Binder struct{}

func (*Binder) Bind(i interface{}, c echo.Context) (err error) {
	// You may use default binder
	b := new(echo.DefaultBinder)
	if err = b.Bind(i, c); err != nil {
		return err
	}

	return c.Validate(i)
}

func init() {
	Router.Use(middleware.Logger())
	Router.Binder = &Binder{}
	Router.Validator = &Validator{validator: validator.New()}

	Router.GET("/", HomeController{}.index)
	Router.Resource("/user", UserController{})
}
