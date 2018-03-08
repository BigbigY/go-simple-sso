package main

import (
	"html/template"
	"io"
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/vanhtuan0409/go-simple-sso/ssoserver/datastore"
	"github.com/vanhtuan0409/go-simple-sso/ssoserver/handler"
	"github.com/vanhtuan0409/go-simple-sso/ssoserver/model"
)

type tpl struct {
	templates *template.Template
}

func newTpl(pattern string) *tpl {
	return &tpl{
		templates: template.Must(template.ParseGlob(pattern)),
	}
}

func (t *tpl) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func redirectMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		if cookie, err := c.Cookie("name"); err == nil && cookie.Value != "" {
			callback := c.Request().URL.Query().Get("callback")
			if callback == "" {
				callback = "http://web1.com:8081"
			}
			tokenAddedCallback := callback + "?name=" + cookie.Value
			return c.Redirect(http.StatusFound, tokenAddedCallback)
		}

		return next(c)
	}
}

func main() {
	e := echo.New()
	e.Use(middleware.Logger())

	// Init datastore and seed data
	ds := datastore.NewDatastore()
	ds.SaveUser(model.NewUser("member1@pav.com", "abc123", "Member 1"))
	ds.SaveUser(model.NewUser("member2@pav.com", "123abc", "Member 2"))

	// Set golang template
	t := newTpl("template/*.html")
	e.Renderer = t

	// Create handler
	h := handler.NewHandler(ds)

	// Routing
	e.GET("/", h.LoginView, redirectMiddleware)
	e.POST("/", h.LoginProcess, redirectMiddleware)
	e.GET("/logout", h.Logout)
	e.Start(":5000")
}