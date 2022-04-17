package handlers

import (
	prisma "boilerplate/database"
	"boilerplate/models"
	"boilerplate/prisma/db"
	"log"
	"time"

	"golang.org/x/crypto/bcrypt"

	"github.com/gofiber/fiber/v2"
)

/* var (
	user  database.Account
	users []database.Account
) */

// UserGet returns a user
// @Summary Show all users
// @Description Show all users in the database
// @Tags User
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

	users, err := prisma.Client.User.FindMany().Exec(prisma.Ctx)
	if err != nil {
		return err
	}

	PrintJSON(users)

	return c.JSON(users)
}

// @Tags User
// @Param id path string true "Account ID"
// @Router /api/v1/users/{id} [get]
func UserFind(c *fiber.Ctx) error {

	user, err := prisma.Client.User.FindUnique(
		db.User.ID.Equals(c.Params("id")),
	).Exec(prisma.Ctx)
	if err != nil {
		return err
	}

	PrintJSON(user)

	return c.JSON(user)
}

// UserCreate registers a user
// @Tags User
//@Param        data  body      models.User  true  "Account Info"
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

	var payload models.User

	if err := c.BodyParser(&payload); err != nil {
		return err
	}

	PrintJSON(payload)

	password := []byte(payload.Password)
	hashedPassword, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	created, err := prisma.Client.User.CreateOne(
		db.User.Firstname.Set(payload.Firstname),
		db.User.Lastname.Set(payload.Lastname),
		db.User.Address.Set(payload.Address),
		db.User.Tel.Set(payload.Tel),
		db.User.Balance.Set(0),
		db.User.Email.Set(payload.Email),
		db.User.Password.Set(string(hashedPassword)),
		db.User.Bio.SetIfPresent(payload.Bio),
		db.User.PicFormat.SetOptional(payload.PicFormat),
	).Exec(prisma.Ctx)
	if err != nil {
		return err
	}

	PrintJSON(created)

	return c.JSON(created)
}

// @Tags User
// @Param id path string true "Account ID"
// @Router /api/v1/users/{id} [delete]
func UserDelete(c *fiber.Ctx) error {
	deleted, err := prisma.Client.User.FindUnique(db.User.ID.Equals(c.Params("id"))).Delete().Exec(prisma.Ctx)
	if err != nil {
		return err
	}

	return c.JSON(deleted)
}

// @Tags User
// @Param id path string true "Account ID"
//@Param        data  body      models.User  true  "Account Info"
// @Router /api/v1/users/{id} [patch]
func UserUpdate(c *fiber.Ctx) error {
	var payload models.User

	if err := c.BodyParser(&payload); err != nil {
		return err
	}

	updated, err := prisma.Client.User.FindUnique(db.User.ID.Equals(c.Params("id"))).Update(
		db.User.Firstname.Set(payload.Firstname),
		db.User.Lastname.Set(payload.Lastname),
		db.User.PicFormat.SetIfPresent(payload.PicFormat),
		db.User.Address.Set(payload.Address),
		db.User.Tel.Set(payload.Tel),
		db.User.Balance.Set(0),
		db.User.Email.Set(payload.Email),
		db.User.Password.Set(payload.Password),
		db.User.Bio.SetOptional(payload.Bio),
	).Exec(prisma.Ctx)
	if err != nil {
		return err
	}

	return c.JSON(updated)
}

// @Tags User
// @Router /api/v1/users [delete]
func DeleteAllUsers(c *fiber.Ctx) error {
	deleted, err := prisma.Client.User.FindMany().Delete().Exec(prisma.Ctx)
	if err != nil {
		return err
	}

	return c.JSON(deleted)
}
