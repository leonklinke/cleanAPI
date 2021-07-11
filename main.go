package main

import (
	"cleanApi/infrastructure"
)

func main() {
	logger := infrastructure.NewLogger()
	infrastructure.LoadEnv(logger)
	infrastructure.ServeRoutes(logger)
}
