package handler

import (
	"net/http"

	"github.com/labstack/echo"
)

type handler struct {
}

func NewHandler() *handler {
	return &handler{}
}

type homeViewModel struct {
	Name string
}

func (h *handler) Home(c echo.Context) error {
	data := homeViewModel{
		Name: "Tuan Vuong",
	}
	return c.Render(http.StatusOK, "home.html", data)
}
