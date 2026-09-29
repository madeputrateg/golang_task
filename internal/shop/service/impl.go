package shops_service

import (
	"context"
	"fmt"
	"uuid"

	"github.com/Teepee/internal/database/entity"
	dtoProd "github.com/Teepee/internal/product/dto"
	products_repository "github.com/Teepee/internal/product/repository"
	dto "github.com/Teepee/internal/shop/dto"
	shop_repository "github.com/Teepee/internal/shop/repository"
)

type ShopServiceImpl struct {
	repoShop    shop_repository.ShopRepository
	repoProduct products_repository.ProductRepository
}

func InitializedShopService(repoShop shop_repository.ShopRepository, repoProduct products_repository.ProductRepository) *ShopServiceImpl {
	service := ShopServiceImpl{
		repoShop:    repoShop,
		repoProduct: repoProduct,
	}
	return &service
}

func (service ShopServiceImpl) CreateShopService(ctx context.Context, shopReq *dto.AddShopRequest) error {
	shopEntity := entity.Shop{
		Name:     shopReq.Name,
		Location: shopReq.Location,
		ImgPath:  "temp",
	}
	err := service.repoShop.SaveShopEntity(ctx, &shopEntity)
	if err != nil {
		return fmt.Errorf("Error during saving data in repo : %v", err)
	}
	return nil
}

func (service ShopServiceImpl) GetShopItemDataService(ctx context.Context, shopRequest *dto.GetShopItemRequest) (*dto.ViewShopItemResponse, error) {
	var productList dtoProd.ProductSimpleList

	shop_uuid, err := uuid.Parse(shopRequest.ShopId)
	if err != nil {
		return nil, fmt.Errorf("Error during parsing shop id : %v", shop_uuid)
	}

	shopEntity, err := service.repoShop.GetShopEntityById(ctx, shop_uuid)
	if err != nil {
		return nil, fmt.Errorf("Error Getting Shop Entity from repo : %v", shop_uuid)
	}

	productEntitys, err := service.repoProduct.GetItemByShopId(ctx, shop_uuid)
	if err != nil {
		return nil, fmt.Errorf("Error Getting Item Entity from repo : %v", shop_uuid)
	}
	for _, productEntity := range productEntitys {
		product := dtoProd.ProductSimple{
			Id:        productEntity.Id.String(),
			Name:      productEntity.Name,
			Stock:     productEntity.Stock,
			Available: productEntity.Available,
		}
		productList = append(productList, product)
	}

	shopView := dto.ViewShopItemResponse{
		Name:     shopEntity.Name,
		Location: shopEntity.Location,
		ShopId:   shopRequest.ShopId,
		ImgPath:  "temp",
		Items:    productList,
	}
	return &shopView, nil
}

func (service ShopServiceImpl) GetShopListService(ctx context.Context) (*dto.ShopDetailList, error) {
	var shopList dto.ShopDetailList
	shopEntityList, err := service.repoShop.GetShopList(ctx)
	if err != nil {
		return nil, fmt.Errorf("Error inside shop repository : %v", err)
	}
	for _, shopEntity := range shopEntityList {
		shopData := dto.ShopDetail{
			Name:     shopEntity.Name,
			Location: shopEntity.Location,
			ShopId:   shopEntity.Id.String(),
			ImgPath:  "temp",
		}
		shopList = append(shopList, shopData)
	}
	return &shopList, nil
}
