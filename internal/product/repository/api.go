package products_repository

import (
	"context"
	"uuid"

	"github.com/Teepee/internal/database/entity"
)

type ProductRepository interface {
	GetItemListFromDatabase(ctx context.Context) ([]entity.Product, error)
	GetDetailItemFromDatabase(ctx context.Context, item_id uuid.UUID) (*entity.Product, error)
	AddProductItem(ctx context.Context, product *entity.Product) error
	SaveProductItemStock(ctx context.Context, product *entity.Product) error
	GetItemByShopId(ctx context.Context, shop_id uuid.UUID) ([]entity.Product, error)
}
