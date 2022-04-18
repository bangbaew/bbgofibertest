package main

import (
	"boilerplate/handlers"

	"flag"
	"log"

	_ "boilerplate/docs"

	swagger "github.com/arsmn/fiber-swagger/v2" // fiber-swagger middleware
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

var (
	port = flag.String("port", ":3000", "Port to listen on")
	prod = flag.Bool("prod", true, "Enable prefork in Production")
)

// @title Fiber Example API
// @version 1.0
// @description This is a sample swagger for Fiber
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email fiber@swagger.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// host localhost:8080
// BasePath /
func main() {

	//database.Gormtest()
	// Parse command-line flags
	flag.Parse()

	/* if err := run(); err != nil {
		panic(err)
	} */

	// Connected with database
	//database.Connect()

	// Create fiber app
	app := fiber.New(fiber.Config{
		Prefork: *prod, // go run app.go -prod
	})

	// Middleware
	app.Use(recover.New())
	app.Use(logger.New())

	app.Get("/swagger/*", swagger.HandlerDefault) // default

	// Create a /api/v1 endpoint
	v1 := app.Group("/api/v1")
	Route(v1)

	// Setup static files
	app.Static("/", "./static/public")

	// Handle not founds
	app.Use(handlers.NotFound)

	// Listen on port 3000
	log.Fatal(app.Listen(*port)) // go run app.go -port=:3000
}

/* func run() error {
	client := db.NewClient()
	if err := client.Prisma.Connect(); err != nil {
		return err
	}

	defer func() {
		if err := client.Prisma.Disconnect(); err != nil {
			panic(err)
		}
	}()

	ctx := context.Background()

	// create a post
	createdPost, err := client.Post.CreateOne(
		db.Post.Title.Set("Hi from Prisma!"),
		db.Post.Published.Set(true),
		db.Post.Desc.Set("Prisma is a database toolkit and makes databases easy."),
	).Exec(ctx)
	if err != nil {
		return err
	}

	result, _ := json.MarshalIndent(createdPost, "", "  ")
	fmt.Printf("created post: %s\n", result)

	// find a single post
	post, err := client.Post.FindUnique(
		db.Post.ID.Equals(createdPost.ID),
	).Exec(ctx)
	if err != nil {
		return err
	}

	result, _ = json.MarshalIndent(post, "", "  ")
	fmt.Printf("post: %s\n", result)

	// for optional/nullable values, you need to check the function and create two return values
	// `desc` is a string, and `ok` is a bool whether the record is null or not. If it's null,
	// `ok` is false, and `desc` will default to Go's default values; in this case an empty string (""). Otherwise,
	// `ok` is true and `desc` will be "my description".
	desc, ok := post.Desc()
	if !ok {
		return fmt.Errorf("post's description is null")
	}

	fmt.Printf("The posts's description is: %s\n", desc)

	return nil
} */

type CRUD interface {
	Create(*fiber.Ctx) error
	Find(*fiber.Ctx) error
	Update(*fiber.Ctx) error
	Delete(*fiber.Ctx) error

	FindAll(*fiber.Ctx) error
	DeleteAll(*fiber.Ctx) error
}

func Route(r fiber.Router) {
	// Bind handlers
	RouteCRUD(r, "/users", handlers.UserHandler{})
	RouteCRUD(r, "/artisans", handlers.ArtisanHandler{})
}

func RouteCRUD(r fiber.Router, groupPath string, handler CRUD) {
	group := r.Group(groupPath)
	{
		group.Post("/", handler.Create)
		group.Get("/:id", handler.Find)
		group.Patch("/:id", handler.Update)
		group.Delete("/:id", handler.Delete)

		group.Get("/", handler.FindAll)
		group.Delete("/", handler.DeleteAll)
	}
}
