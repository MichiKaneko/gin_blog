package main

import (
	"github.com/MichiKaneko/nekoblog/config"
	"github.com/MichiKaneko/nekoblog/db"
)

func main() {
	config.LoadEnv()

	db.DBConnect(config.DatabaseConfig())
	db.Migrate()

	r := InitRouter()
	r.Run(":8080")
}
