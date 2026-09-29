package seeder

import (
	"fmt"

	"github.com/Teepee/internal/database/entity"
	database "github.com/Teepee/internal/database/impl"
)

func SeedShops(db *database.Database, number int) error {
	if number <= 0 {
		return nil
	}

	shops := make([]entity.Shop, 0, number)

	for i := 0; i < number; i++ {
		shop := entity.Shop{
			Name:     fmt.Sprintf("Shop %d", i+1),
			Location: fmt.Sprintf("Location %d", i+1),
			ImgPath:  fmt.Sprintf("/images/shops/shop_%d.jpg", i+1),
		}

		shops = append(shops, shop)
	}

	// Insert shops in batches of 100
	if err := db.Db.CreateInBatches(&shops, 100).Error; err != nil {
		return fmt.Errorf("failed to seed shops: %w", err)
	}

	return nil
}
