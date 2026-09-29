package web

import (
	"fmt"
	"net/http"
	"os"

	"github.com/Teepee/internal/utils"
	"github.com/labstack/echo/v5"
)

type WebController struct {
}

func (controller WebController) GetTemplateWeb(ctx *echo.Context) error {
	err := ctx.File("templates/index.html")
	dir, _ := os.Getwd()
	fmt.Println("Current directory:", dir)
	if err != nil {
		err = fmt.Errorf("loading page error : %v", err)
		response := utils.WrapPayload(nil, http.StatusInternalServerError, err)
		return ctx.JSON(http.StatusInternalServerError, response)
	}
	return ctx.JSON(http.StatusOK, nil)
}

func InitializedWebController() *WebController {
	controller := new(WebController)
	return controller
}

func (controller WebController) InitializedRoute(e *echo.Echo) {
	e.GET("/web", controller.GetTemplateWeb)
}
