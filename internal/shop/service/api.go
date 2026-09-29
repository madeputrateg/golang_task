package shops_service

import (
	"context"

	"github.com/Teepee/internal/shop/dto"
)

type ShopService interface {
	CreateShopService(ctx context.Context, shopReq *dto.AddShopRequest) error
	GetShopItemDataService(ctx context.Context, shopRequest *dto.GetShopItemRequest) (*dto.ViewShopItemResponse, error)
	GetShopListService(ctx context.Context) (*dto.ShopDetailList, error)
}
