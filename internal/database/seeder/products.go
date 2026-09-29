package seeder

import (
	"fmt"
	"math/rand"
	"uuid"

	"github.com/Teepee/internal/database/entity"
	database "github.com/Teepee/internal/database/impl"
)

func SeedProdcuts(db *database.Database, number int) error {
	if number <= 0 {
		return nil
	}

	// Retrieve existing shop IDs
	var shopIDs []uuid.UUID

	if err := db.Db.Model(&entity.Shop{}).
		Pluck("id", &shopIDs).Error; err != nil {
		return fmt.Errorf("failed to retrieve shops: %w", err)
	}

	// Cannot seed products without shops
	if len(shopIDs) == 0 {
		return fmt.Errorf("cannot seed products: no shops found")
	}

	productNames := []string{
		"Laptop",
		"Smartphone",
		"Headphones",
		"Keyboard",
		"Mouse",
		"Monitor",
		"Webcam",
		"Speaker",
		"Tablet",
		"Smartwatch",
	}

	availablity := []string{
		"available",
		"out_of_stock",
		"not_available",
	}

	products := make([]entity.Product, 0, number)

	for i := 0; i < number; i++ {

		// Randomly select an existing shop
		shopID := shopIDs[rand.Intn(len(shopIDs))]

		// Random product name
		productName := productNames[rand.Intn(len(productNames))]

		// Random stock between 0 and 100
		stock := uint32(rand.Intn(101))

		// Random price between 10,000 and 1,000,000
		price := uint32(rand.Intn(990001) + 10000)

		idxavailable := uint32(rand.Intn(3))

		available := availablity[idxavailable]

		// if stock == 0 {
		// 	available = "unavailable"
		// }

		product := entity.Product{
			ShopId:    shopID,
			Name:      fmt.Sprintf("%s %d", productName, i+1),
			Stock:     stock,
			Price:     price,
			Available: available,
			ImgPath:   fmt.Sprintf("/images/products/product_%d.jpg", i+1),
		}

		products = append(products, product)
	}

	// Insert products in batches of 100
	if err := db.Db.CreateInBatches(&products, 100).Error; err != nil {
		return fmt.Errorf("failed to seed products: %w", err)
	}

	return nil
}
