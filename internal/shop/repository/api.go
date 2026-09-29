package shops_repository

import (
	"context"
	"uuid"

	"github.com/Teepee/internal/database/entity"
)

type ShopRepository interface {
	SaveShopEntity(ctx context.Context, shop *entity.Shop) error
	GetShopEntityById(ctx context.Context, shop_id uuid.UUID) (*entity.Shop, error)
	GetShopList(ctx context.Context) ([]entity.Shop, error)
}
