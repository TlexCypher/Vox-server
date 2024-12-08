//go:build wireinject
// +build wireinject

package di

import (
	"Vox/db"
	"Vox/internal/infrastructure"

	"github.com/google/wire"
)

func InitUserRepository() infrastructure.UserRepositoryImpl {
	wire.Build(infrastructure.NewUserRepositoryImpl, db.NewDBConn)
	return infrastructure.UserRepositoryImpl{}
}

func InitInventoryRepository() infrastructure.InventoryRepositoryImpl {
	wire.Build(infrastructure.NewInventoryRepositoryImpl, db.NewDBConn)
	return infrastructure.InventoryRepositoryImpl{}
}

func InitCategoryRepository() infrastructure.CategoryRepositoryImpl {
	wire.Build(infrastructure.NewCategoryRepositoryImpl, db.NewDBConn)
	return infrastructure.CategoryRepositoryImpl{}
}
