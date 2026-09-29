package shops_controller

import (
	"net/http"

	"github.com/Teepee/internal/global"
	"github.com/Teepee/internal/shop/dto"
	shop_service "github.com/Teepee/internal/shop/service"
	"github.com/Teepee/internal/utils"
	"github.com/labstack/echo/v5"
)

type ShopControllerImpl struct {
	service shop_service.ShopService
}

func InititalizedShopController(service shop_service.ShopService) *ShopControllerImpl {
	controller := ShopControllerImpl{
		service: service,
	}
	return &controller
}

func (controller ShopControllerImpl) InitializedRoute(e *echo.Echo) {
	e.POST(global.RouteAddShop, controller.SaveShopController)
	e.GET(global.RouteViewShopItemList, controller.GetShopItemController)
	e.GET(global.RouteViewShopList, controller.GetShopListController)
}

func (controller ShopControllerImpl) SaveShopController(ctx *echo.Context) error {
	request := new(dto.AddShopRequest)
	err := ctx.Bind(request)
	if err != nil {
		response := utils.WrapPayload(nil, http.StatusBadRequest, err)
		return ctx.JSON(http.StatusBadRequest, response)
	}
	err = controller.service.CreateShopService(ctx.Request().Context(), request)
	if err != nil {
		response := utils.WrapPayload(nil, http.StatusInternalServerError, err)
		return ctx.JSON(http.StatusInternalServerError, response)
	}
	response := utils.WrapPayload(nil, http.StatusAccepted, nil)
	return ctx.JSON(http.StatusAccepted, response)
}

func (controller ShopControllerImpl) GetShopItemController(ctx *echo.Context) error {
	request := new(dto.GetShopItemRequest)
	err := ctx.Bind(request)
	if err != nil {
		response := utils.WrapPayload(nil, http.StatusBadRequest, err)
		return ctx.JSON(http.StatusBadRequest, response)
	}
	responsePayload, err := controller.service.GetShopItemDataService(ctx.Request().Context(), request)
	if err != nil {
		response := utils.WrapPayload(nil, http.StatusInternalServerError, err)
		return ctx.JSON(http.StatusInternalServerError, response)
	}
	response := utils.WrapPayload(responsePayload, http.StatusAccepted, nil)
	return ctx.JSON(http.StatusAccepted, response)
}

func (controller ShopControllerImpl) GetShopListController(ctx *echo.Context) error {
	responsePayload, err := controller.service.GetShopListService(ctx.Request().Context())
	if err != nil {
		response := utils.WrapPayload(nil, http.StatusInternalServerError, err)
		return ctx.JSON(http.StatusInternalServerError, response)
	}
	response := utils.WrapPayload(responsePayload, http.StatusAccepted, nil)
	return ctx.JSON(http.StatusAccepted, response)
}
