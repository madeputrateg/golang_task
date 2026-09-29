package products_controller

import (
	"fmt"
	"net/http"

	"github.com/Teepee/internal/global"
	dto "github.com/Teepee/internal/product/dto"
	products_service "github.com/Teepee/internal/product/service"
	"github.com/Teepee/internal/utils"
	"github.com/labstack/echo/v5"
)

type ProductsController struct {
	service products_service.ProductsService
}

func InitializedProductsController(service products_service.ProductsService) *ProductsController {
	controller := ProductsController{
		service: service,
	}
	return &controller
}

func (controller *ProductsController) InitializedRoute(e *echo.Echo) {
	e.GET(global.RouteViewItemDetail, controller.ItemDetailController)
	e.GET(global.RouteViewItemList, controller.ItemListController)
	e.POST(global.RouteAddItemProduct, controller.AddItemController)
	e.PUT(global.RouteAddItemStock, controller.UpdateItemStockController)
	e.POST(global.RouteProductCheckout, controller.CheckoutProductItemController)
	e.PUT(global.RouteUpdateItemAvailabilty, controller.UpdateAvailabiltyProduct)
}

func (controller *ProductsController) ItemDetailController(ctx *echo.Context) error {
	item_id, err := echo.PathParam[string](ctx, "item_id")
	if err != nil {
		return echo.ErrBadRequest.Wrap(err)
	}
	payload, err := controller.service.GetViewItemDetail(ctx.Request().Context(), item_id)
	if err != nil {
		err = fmt.Errorf("error during service getting item detail: %v", err)
		response := utils.WrapPayload(nil, http.StatusInternalServerError, err)
		return ctx.JSON(http.StatusInternalServerError, response)
	}
	response := utils.WrapPayload(payload, http.StatusOK, nil)
	return ctx.JSON(http.StatusOK, response)
}

func (controller *ProductsController) ItemListController(ctx *echo.Context) error {
	productList, err := controller.service.GetItemListView(ctx.Request().Context(), 0, 0)
	if err != nil {
		err = fmt.Errorf("\nError happen during getting data from service : %v", err)
		response := utils.WrapPayload(nil, http.StatusInternalServerError, err)
		return ctx.JSON(http.StatusInternalServerError, response)
	}
	response := utils.WrapPayload(productList, http.StatusOK, nil)
	return ctx.JSON(http.StatusOK, response)
}

func (controller *ProductsController) AddItemController(ctx *echo.Context) error {
	request := new(dto.AddItemRequest)
	err := ctx.Bind(request)
	if err != nil {
		response := utils.WrapPayload(nil, http.StatusBadRequest, err)
		return ctx.JSON(http.StatusBadRequest, response)
	}
	err = controller.service.AddItemProduct(ctx.Request().Context(), request)
	if err != nil {
		response := utils.WrapPayload(nil, http.StatusInternalServerError, err)
		return ctx.JSON(http.StatusInternalServerError, response)
	}
	response := utils.WrapPayload(nil, http.StatusAccepted, err)
	return ctx.JSON(http.StatusAccepted, response)
}

func (controller *ProductsController) UpdateItemStockController(ctx *echo.Context) error {
	request := new(dto.AddItemStockRequest)
	err := ctx.Bind(request)
	if err != nil {
		response := utils.WrapPayload(nil, http.StatusBadRequest, err)
		return ctx.JSON(http.StatusBadRequest, response)
	}
	//UpdateItemService
	err = controller.service.AddStockProduct(ctx.Request().Context(), request)
	if err != nil {
		response := utils.WrapPayload(nil, http.StatusInternalServerError, err)
		return ctx.JSON(http.StatusInternalServerError, response)
	}
	response := utils.WrapPayload(nil, http.StatusAccepted, err)
	return ctx.JSON(http.StatusAccepted, response)
}

func (controller *ProductsController) CheckoutProductItemController(ctx *echo.Context) error {
	request := new(dto.CheckoutProductRequest)
	err := ctx.Bind(request)
	if err != nil {
		response := utils.WrapPayload(nil, http.StatusBadRequest, err)
		return ctx.JSON(http.StatusBadRequest, response)
	}

	// handle here
	err = controller.service.CheckoutProductItem(ctx.Request().Context(), request)
	if err != nil {
		response := utils.WrapPayload(nil, http.StatusInternalServerError, err)
		return ctx.JSON(http.StatusInternalServerError, response)
	}
	response := utils.WrapPayload(nil, http.StatusAccepted, nil)
	return ctx.JSON(http.StatusAccepted, response)
}

func (controller *ProductsController) UpdateAvailabiltyProduct(ctx *echo.Context) error {
	request := new(dto.UpdateAvailabiltyRequest)
	err := ctx.Bind(request)
	if err != nil {
		response := utils.WrapPayload(nil, http.StatusBadRequest, err)
		return ctx.JSON(http.StatusBadRequest, response)
	}
	err = controller.service.UpdateAvailabiltyProduct(ctx.Request().Context(), request)
	if err != nil {
		response := utils.WrapPayload(nil, http.StatusInternalServerError, err)
		return ctx.JSON(http.StatusInternalServerError, response)
	}
	response := utils.WrapPayload(nil, http.StatusAccepted, nil)
	return ctx.JSON(http.StatusAccepted, response)
}
