package db

import (
	"context"
	"fmt"
	"log"
	"spiropoulos94/emailchaser/invite/ent"
	"spiropoulos94/emailchaser/invite/ent/user"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

var database *ent.Client

func CreateFirstUser(email string) error {
	exists, err := database.User.Query().Where(user.Email(email)).Exist(context.Background())
	if err != nil {
		return err
	}

	if !exists {
		salesTeam, err := database.Team.Create().SetName("Sales").Save(context.Background())
		if err != nil {
			return err
		}

		_, err = database.User.Create().
			SetName("User A").
			SetEmail(email).
			SetPassword("password"). // Set an appropriate password here.
			AddTeams(salesTeam).
			Save(context.Background())
		if err != nil {
			return err
		}
	}

	return nil
}

func SetupDatabaseClient() {
	var err error
	database, err = ent.Open("postgres", "host=database port=5432 user=example_user dbname=example_db password=example_password sslmode=disable")
	if err != nil {
		log.Fatalf("failed opening connection to postgres: %v", err)
	}
	// defer database.Close()
	// Run the auto migration tool.
	if err := database.Schema.Create(&gin.Context{}); err != nil {
		defer database.Close()
		log.Fatalf("failed creating schema resources: %v", err)
	}

	err = CreateFirstUser("usera@emailchaser.com")
	if err != nil {
		fmt.Printf("failed to create user: %v", err)
	}

}

func GetDB() *ent.Client {
	return database
}
