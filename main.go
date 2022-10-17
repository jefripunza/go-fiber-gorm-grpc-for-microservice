package main

import (
	"main-service/src/configs"
	"main-service/src/middlewares"
	"main-service/src/routers"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"

	run "main-service/src/app"
)

// @title GO Fiber Training
// @version 1.0
// @description lagi belajar asik golang sampe mampus :v
// @termsOfService http://swagger.io/terms/
// @contact.name Jefri Herdi Triyanto
// @contact.email jefriherditriyanto@gmail.com
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:8080
// @BasePath /
func main() {
	godotenv.Load(".env")

	// Define Fiber config.
	config := configs.FiberConfig()
	app := fiber.New(config)

	// import middleware (global) yang dibutuhkan
	middlewares.Cors(app)
	middlewares.Logger(app)

	// REGISTER ALL API ROUTERS
	api := app.Group(os.Getenv("ENDPOINT"))
	routers.Basic(api)
	routers.Main(api)

	// REQUIRE ROUTERS
	routers.Index(app)
	routers.Swagger(app)
	routers.NotFound(app)

	run.WebServer(app)
	run.DatabaseMigration() // activate in .env > DB_SYNC
	run.GrpcServer()

}
