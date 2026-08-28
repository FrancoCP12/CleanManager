package main

import (
	"github.com/francocp12/CleanManager/internal/persistence/db"
	"github.com/francocp12/CleanManager/internal/persistence/repository"
)

func main() {

	db := db.Connect()

	ProductoRepository := repository.NewMySQLRepository(db)
	_ = ProductoRepository
}
