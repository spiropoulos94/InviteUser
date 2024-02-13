package main

import (
	"spiropoulos94/emailchaser/invite/internal/api/rest/server"
	"spiropoulos94/emailchaser/invite/internal/db"

	_ "github.com/lib/pq"
)

func main() {
	db.SetupDatabaseClient()
	server.Init()
}
