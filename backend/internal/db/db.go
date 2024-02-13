package db

import (
	"log"
	"spiropoulos94/emailchaser/invite/ent"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

var database *ent.Client

func SetupDatabaseClient ()  {
    var err error
	database, err = ent.Open("postgres","host=database port=5432 user=example_user dbname=example_db password=example_password sslmode=disable")
    if err != nil {
        log.Fatalf("failed opening connection to postgres: %v", err)
    }
    // defer database.Close()
    // Run the auto migration tool.
    if err := database.Schema.Create(&gin.Context{}); err != nil {
        defer database.Close()
        log.Fatalf("failed creating schema resources: %v", err)
    }

}

func GetDB() *ent.Client {
	return database
}
