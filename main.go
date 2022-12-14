package main

import (
	"microservice/src/configs"
	"microservice/src/consumer"
	"microservice/src/middlewares"
	"microservice/src/routers"
	"microservice/src/scheduler"
	"microservice/src/utils/api"
	"microservice/src/utils/environment"

	"github.com/gofiber/fiber/v2"

	run "microservice/src/apps"
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
	environment.Get()
	environment.RabbitMQ() // if use RabbitMQ

	// Define Fiber App.
	app := fiber.New(configs.FiberConfig())

	// import middleware (global) yang dibutuhkan
	middlewares.Cors(app)
	middlewares.Logger(app)

	//-> Require Routers
	routers.Register(app) // first !!!
	api.Index(app)        // if use
	api.Swagger(app)      // if use
	api.NotFound(app)

	//-> Use Consumer RabbitMQ
	consumer.Register() // if use
	//-> Use Task Scheduler
	scheduler.Register() // if use

	//-> Execute all Apps
	run.RabbitTryConnection()   // if use
	run.RedisTryConnection()    // if use
	run.DatabaseTryConnection() // if use // activate migration in .env > DB_SYNC
	run.GrpcServer()            // if use
	run.WebServer(app)          // ending...

}
