package products_service

import (
	"context"

	"github.com/Teepee/internal/product/dto"
)

type ProductsService interface {
	GetViewItemDetail(ctx context.Context, item_id string) (dto.ProductDetail, error)
	GetItemListView(ctx context.Context, pageStart int, pageEnd int) (dto.ProductSimpleList, error)
	AddItemProduct(ctx context.Context, product *dto.AddItemRequest) error
	AddStockProduct(ctx context.Context, stockUpdate *dto.AddItemStockRequest) error
	CheckoutProductItem(ctx context.Context, checkoutPayload *dto.CheckoutProductRequest) error
	UpdateAvailabiltyProduct(ctx context.Context, updatePayload *dto.UpdateAvailabiltyRequest) error
}
