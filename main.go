package main

import (
	"main-service/src/configs"
	"main-service/src/middlewares"
	"main-service/src/routers"
	"main-service/src/utils"

	"github.com/gofiber/fiber/v2"

	run "main-service/src/apps"
)

// @title GO Fiber Training
// @version 1.0
// @description lagi belajar asik golang sampe mampus :v
// @termsOfService http://swagger.io/terms/

// @contact.name Jefri Herdi Triyanto
// @contact.email jefriherditriyanto@gmail.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8082
// @BasePath /

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func main() {

	// get all value on environment
	utils.GetEnvironment()

	// Define Fiber config.
	config := configs.FiberConfig()
	app := fiber.New(config)

	// import middleware (global) yang dibutuhkan
	middlewares.Cors(app)
	middlewares.Logger(app)

	// REGISTER ALL API ROUTERS
	api := app.Group(configs.GetBaseEndpoint())
	routers.Basic(api) //> Connect to gRPC Service
	routers.Products(api)

	// REQUIRE ROUTERS
	routers.Index(app)
	routers.Swagger(app)
	routers.NotFound(app)

	//-> Execute all Apps
	run.WebServer(app)
	run.DatabaseMigration() // activate in .env > DB_SYNC
	run.GrpcServer()        // ending...

}
