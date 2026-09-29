package products_repository

import (
	"context"
	"fmt"
	"uuid"

	"github.com/Teepee/internal/database/entity"
	database "github.com/Teepee/internal/database/impl"

	"github.com/doug-martin/goqu/v9"
	_ "github.com/doug-martin/goqu/v9/dialect/postgres"
)

type ProductRepositoryImpl struct {
	Db *database.Database
}

func InitializedProductRepository(db *database.Database) *ProductRepositoryImpl {
	repository := ProductRepositoryImpl{
		Db: db,
	}
	return &repository
}

func (repo *ProductRepositoryImpl) GetDetailItemFromDatabase(ctx context.Context, item_id uuid.UUID) (*entity.Product, error) {
	var product *entity.Product
	dialect := goqu.Dialect("postgres")
	query, args, err := dialect.From("products").Where(goqu.Ex{"id": item_id.String()}).Limit(1).Prepared(true).ToSQL()
	if err != nil {
		err = fmt.Errorf("\n Error happen when creating sql product data:%v", err)
		return nil, err
	}
	err = repo.Db.Db.WithContext(ctx).Raw(query, args...).Scan(&product).Error
	if err != nil {
		err = fmt.Errorf("\n Error happen when fetching product data from database:%v", err)
		return nil, err
	}
	if product == nil {
		return nil, fmt.Errorf("product was not found")
	}
	return product, err
}

func (repo *ProductRepositoryImpl) GetItemListFromDatabase(ctx context.Context) ([]entity.Product, error) {
	var products []entity.Product
	dialect := goqu.Dialect("postgres")
	query, args, err := dialect.From("products").Prepared(true).ToSQL()
	if err != nil {
		err = fmt.Errorf("\n Error happen when creating sql product data:%v", err)
		return products, err
	}
	err = repo.Db.Db.WithContext(ctx).Raw(query, args...).Scan(&products).Error
	if err != nil {
		err = fmt.Errorf("\n Error happen when fetching product data from database:%v", err)
		return products, err
	}
	return products, nil
}

func (repo *ProductRepositoryImpl) AddProductItem(ctx context.Context, product *entity.Product) error {
	err := repo.Db.Db.WithContext(ctx).Create(product).Error
	// fmt.Print(product.ShopId.String())
	if err != nil {
		err = fmt.Errorf("Error happen when running query to add product %v", err)
		// log.Fatal(err)
		return err
	}
	return nil
}

func (repo *ProductRepositoryImpl) SaveProductItemStock(ctx context.Context, product *entity.Product) error {
	dialect := goqu.Dialect("postgres")
	query, args, err := dialect.
		Update("products").
		Set(goqu.Record{
			"stock":     product.Stock,
			"available": product.Available,
		}).
		Where(goqu.Ex{
			"id": product.Id.String(),
		}).
		Prepared(true).
		ToSQL()
	if err != nil {
		err = fmt.Errorf("err when creating sql query inside SaveProductItemStock : %v", err)
		return err
	}
	err = repo.Db.Db.WithContext(ctx).Exec(query, args...).Error
	if err != nil {
		err = fmt.Errorf("err when running sql query inside SaveProductItemStock: %v", err)
		return err
	}
	return nil
}

func (repo *ProductRepositoryImpl) GetItemByShopId(ctx context.Context, shop_id uuid.UUID) ([]entity.Product, error) {
	var products []entity.Product
	dialect := goqu.Dialect("postgres")
	query, args, err := dialect.From("products").Where(goqu.Ex{"shop_id": shop_id.String()}).Prepared(true).ToSQL()
	if err != nil {
		return products, fmt.Errorf("Error during creating sql query in repository : %v", err)
	}
	err = repo.Db.Db.Raw(query, args...).Scan(&products).Error
	if err != nil {
		return products, fmt.Errorf("Error during running sql query in repository : %v", err)
	}
	return products, err
}
