package products_service

import (
	"context"
	"fmt"
	"uuid"

	"github.com/Teepee/internal/database/entity"
	"github.com/Teepee/internal/global"
	"github.com/Teepee/internal/product/dto"
	products_repository "github.com/Teepee/internal/product/repository"
	"github.com/Teepee/internal/utils"
)

type ProductsServiceImpl struct {
	Repository products_repository.ProductRepository
}

func InitializedProductsService(repository products_repository.ProductRepository) *ProductsServiceImpl {
	service := ProductsServiceImpl{
		Repository: repository,
	}
	return &service
}

func (service *ProductsServiceImpl) GetViewItemDetail(ctx context.Context, item_id string) (dto.ProductDetail, error) {
	var product dto.ProductDetail
	itemUuid, err := uuid.Parse(item_id)
	if err != nil {
		err = fmt.Errorf("error during parsing item_id: %v", err)
		return product, err
	}
	productEntity, err := service.Repository.GetDetailItemFromDatabase(ctx, itemUuid)
	if err != nil {
		err = fmt.Errorf("error during getting product in repository: %v", err)
		return product, err
	}
	product = dto.ProductDetail{
		Id:        item_id,
		Name:      productEntity.Name,
		Price:     productEntity.Price,
		Stock:     productEntity.Stock,
		ShopId:    productEntity.ShopId.String(),
		Available: productEntity.Available,
	}
	return product, nil
}

func (service *ProductsServiceImpl) GetItemListView(ctx context.Context, pageStart int, pageEnd int) (dto.ProductSimpleList, error) {
	productEntityList, err := service.Repository.GetItemListFromDatabase(ctx)
	if err != nil {
		return nil, fmt.Errorf("\nError happen inside product repository %v", err)
	}
	var productSimpleListDto dto.ProductSimpleList
	for _, productEntity := range productEntityList {
		productSimple := dto.ProductSimple{
			Id:        productEntity.Id.String(),
			Name:      productEntity.Name,
			Stock:     productEntity.Stock,
			Available: productEntity.Available,
		}
		productSimpleListDto = append(productSimpleListDto, productSimple)
	}
	return productSimpleListDto, nil
}

func (service *ProductsServiceImpl) AddItemProduct(ctx context.Context, product *dto.AddItemRequest) error {
	shopId, err := uuid.Parse(product.ShopId)
	if err != nil {
		err = fmt.Errorf("Error inside AddItemProduct service because invalid parse %v", err)
		return err
	}
	productEntity := entity.Product{
		ShopId:    shopId,
		Name:      product.Name,
		Stock:     product.Stock,
		Price:     product.Price,
		Available: product.Available,
		ImgPath:   "temp",
	}
	err = service.Repository.AddProductItem(ctx, &productEntity)
	if err != nil {
		err = fmt.Errorf("Error inside repository product : %v", err)
		return err
	}
	return nil
}

func (service *ProductsServiceImpl) AddStockProduct(ctx context.Context, stockUpdate *dto.AddItemStockRequest) error {
	item_id, err := uuid.Parse((stockUpdate.ItemId))
	if err != nil {
		err = fmt.Errorf("Error inside AddStockProduct service because invalid parse %v", err)
		return err
	}
	product, err := service.Repository.GetDetailItemFromDatabase(ctx, item_id)
	if err != nil {
		err = fmt.Errorf("Error inside repository product : %v", err)
		return err
	}
	product.Stock = product.Stock + stockUpdate.Stock
	if product.Stock > 0 && product.Available == global.AvailableOut_of_stock {
		product.Available = global.AvailableAvailable
	}
	err = service.Repository.SaveProductItemStock(ctx, product)
	if err != nil {
		err = fmt.Errorf("Error inside repository product : %v", err)
		return err
	}
	return nil
}

func (service *ProductsServiceImpl) CheckoutProductItem(ctx context.Context, checkoutPayload *dto.CheckoutProductRequest) error {
	item_id, err := uuid.Parse((checkoutPayload.Id))
	if err != nil {
		err = fmt.Errorf("Error inside AddStockProduct service because invalid parse %v", err)
		return err
	}
	product, err := service.Repository.GetDetailItemFromDatabase(ctx, item_id)
	if err != nil {
		err = fmt.Errorf("Error inside repository product : %v", err)
		return err
	}
	if product.Available != "available" {
		err = fmt.Errorf("the product is not available")
		return err
	}
	if product.Stock >= checkoutPayload.Quantity {
		product.Stock = product.Stock - checkoutPayload.Quantity
		if product.Stock == 0 {
			product.Available = "out_of_stock"
		}
		service.Repository.SaveProductItemStock(ctx, product)
	} else {
		err = fmt.Errorf("the stock is not enough, there is only : %v but reqiured %v", product.Stock, checkoutPayload.Quantity)
		return err
	}
	return nil
}

func (service *ProductsServiceImpl) UpdateAvailabiltyProduct(ctx context.Context, updatePayload *dto.UpdateAvailabiltyRequest) error {
	if !utils.CheckAvailablity(updatePayload.Availabilty) {
		return fmt.Errorf("availabilty is not represented by attribute")
	}
	item_id, err := uuid.Parse((updatePayload.Id))
	if err != nil {
		err = fmt.Errorf("Error inside UpdateAvailabiltyProduct service because invalid parse %v", err)
		return err
	}
	product, err := service.Repository.GetDetailItemFromDatabase(ctx, item_id)
	if err != nil {
		err = fmt.Errorf("Error inside repository product : %v", err)
		return err
	}
	product.Available = updatePayload.Availabilty
	err = service.Repository.SaveProductItemStock(ctx, product)
	if err != nil {
		err = fmt.Errorf("Error inside repository product : %v", err)
		return err
	}
	return nil
}
