package handlers

import (
	prisma "boilerplate/database"
	"boilerplate/models"
	"boilerplate/prisma/db"
	"encoding/json"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
)

/* var (
	user  database.Account
	users []database.Account
) */

// UserGet returns a user
// @Summary Show all users
// @Description Show all users in the database
// ID get-string-by-int
// Accept  json
// Produce  json
// Param id path int true "Account ID"
// @Router /api/v1/users [get]
func UserList(c *fiber.Ctx) error {
	//users := database.Get()
	/* db.Find(&users)
	fmt.Println(users) */
	log.Println(time.Now().Zone())
	log.Println(time.Now())
	log.Println(c)

	posts, err := prisma.Client.Post.FindMany().Exec(prisma.Ctx)
	if err != nil {
		return err
	}

	PrintJSON(posts)

	req := c.JSON(fiber.Map{
		"user": posts,
	})
	return req
}

// UserCreate registers a user
//@Param        data  body      models.Post  true  "Account Info"
// @Router /api/v1/users [post]
func UserCreate(c *fiber.Ctx) error {

	//database.Insert(user)

	/* newID, err := cuid.NewCrypto(rand.Reader)
	if err != nil {
		log.Fatal("error configuring the database: ", err)
	}

	user := &models.User{
		ID:   newID,
		Name: c.FormValue("user"),
	}

	if err := db.Create(&database.Account{ID: newID, Name: c.FormValue("user")}).Error; err != nil {
		return err
	} */

	var payload models.Post

	if err := c.BodyParser(&payload); err != nil {
		return err
	}

	PrintJSON(payload)

	createdPost, err := prisma.Client.Post.CreateOne(
		db.Post.Title.Set(payload.User),
		db.Post.Published.Set(payload.Published),
		db.Post.Desc.Set(payload.Desc),
	).Exec(prisma.Ctx)
	if err != nil {
		return err
	}

	PrintJSON(createdPost)

	return c.JSON(fiber.Map{"name": createdPost.Title, "published": createdPost.Published, "desc": createdPost.Desc})
}

// NotFound returns custom 404 page
func NotFound(c *fiber.Ctx) error {
	return c.Status(404).SendFile("./static/private/404.html")
}

func PrintJSON(data interface{}) {
	result, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("%s", result)
}
