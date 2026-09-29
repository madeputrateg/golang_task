package shops_repository

import (
	"context"
	"fmt"
	"uuid"

	"github.com/Teepee/internal/database/entity"
	database "github.com/Teepee/internal/database/impl"
	"github.com/doug-martin/goqu/v9"
)

type ShopRepositoryImpl struct {
	Db *database.Database
}

func InitializedShopRepository(db *database.Database) *ShopRepositoryImpl {
	repo := ShopRepositoryImpl{
		Db: db,
	}
	return &repo
}

func (repo *ShopRepositoryImpl) SaveShopEntity(ctx context.Context, shop *entity.Shop) error {
	err := repo.Db.Db.WithContext(ctx).Create(shop).Error
	if err != nil {
		return fmt.Errorf("error during saving data to database SaveShopEntity : %v", err)
	}
	return nil
}

func (repo *ShopRepositoryImpl) GetShopEntityById(ctx context.Context, shop_id uuid.UUID) (*entity.Shop, error) {
	var shopEntity *entity.Shop
	dialect := goqu.Dialect("postgres")
	query, args, err := dialect.From("shops").Where(goqu.Ex{"id": shop_id.String()}).Limit(1).Prepared(true).ToSQL()
	if err != nil {
		return nil, fmt.Errorf("Error happen during sql creation : %v", err)
	}
	err = repo.Db.Db.WithContext(ctx).Raw(query, args...).Scan(&shopEntity).Error

	if err != nil {
		return nil, fmt.Errorf("Error happen when running sql query : %v", err)
	}

	if shopEntity == nil {
		return nil, fmt.Errorf("Shop was not found")
	}

	return shopEntity, nil
}

func (repo *ShopRepositoryImpl) GetShopList(ctx context.Context) ([]entity.Shop, error) {
	var shopEntityList []entity.Shop
	dialect := goqu.Dialect("postgres")
	query, args, err := dialect.From("shops").Prepared(true).ToSQL()
	if err != nil {
		return shopEntityList, fmt.Errorf("Error during making query shop data : %v", err)
	}
	err = repo.Db.Db.Raw(query, args...).Scan(&shopEntityList).Error
	if err != nil {
		return shopEntityList, fmt.Errorf("Error during fetching shop data from database : %v", err)
	}
	return shopEntityList, nil
}
