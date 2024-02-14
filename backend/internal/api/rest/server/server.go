package server

import (
	"spiropoulos94/emailchaser/invite/internal/api/rest"
)

func Init() {
	r := rest.NewRouter()

	r.Run(":8080")
}
