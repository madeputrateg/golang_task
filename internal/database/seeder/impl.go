package seeder

import (
	"fmt"

	database "github.com/Teepee/internal/database/impl"
)

func SeedDatabase(db *database.Database) error {
	err := SeedShops(db, 10)
	if err != nil {
		return fmt.Errorf("Failed to seed shops: %v", err)
	}
	err = SeedProdcuts(db, 100)
	if err != nil {
		return fmt.Errorf("Failed to seed product: %v", err)
	}
	return nil
}
