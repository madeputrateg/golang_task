package database

import (
	"fmt"
	"time"

	"github.com/Teepee/internal/config"
	"github.com/Teepee/internal/database/entity"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Database struct {
	Db *gorm.DB
}

func ConnectAndMigrateDatabase(config config.Config) (*Database, error) {

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		config.Host,
		config.User,
		config.Password,
		config.Dbname,
		config.Port,
	)

	var db *gorm.DB
	var err error

	maxRetry := 5
	delay := 2 * time.Second

	for attempt := 1; attempt <= maxRetry; attempt++ {

		fmt.Printf(
			"Connecting to database... attempt %d/%d\n",
			attempt,
			maxRetry,
		)

		db, err = gorm.Open(
			postgres.Open(dsn),
			&gorm.Config{},
		)

		if err == nil {

			fmt.Println(
				"Database connected successfully",
			)

			break
		}

		fmt.Println(
			"Database connection failed:",
			err,
		)

		if attempt < maxRetry {
			fmt.Printf(
				"Retrying in %v...\n",
				delay,
			)

			time.Sleep(delay)
		}

	}

	if err != nil {

		return nil, fmt.Errorf(
			"failed to connect database after %d attempts: %w",
			maxRetry,
			err,
		)

	}

	err = db.AutoMigrate(
		&entity.Shop{},
		&entity.Product{},
	)

	if err != nil {

		return nil, fmt.Errorf(
			"failed running migration: %w",
			err,
		)

	}

	DatabaseVariable := Database{
		Db: db,
	}

	return &DatabaseVariable, nil
}
