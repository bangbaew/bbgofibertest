package handlers

import (
	prisma "boilerplate/database"
	"boilerplate/models"
	"boilerplate/prisma/db"

	"github.com/gofiber/fiber/v2"
)

// @Summary Show all users
// @Description Show all users in the database
// @Tags Artisan
// Param id path int true "Account ID"
// @Router /api/v1/artisans [get]
func ArtisanList(c *fiber.Ctx) error {

	users, err := prisma.Client.Artisan.FindMany().Exec(prisma.Ctx)
	if err != nil {
		return err
	}

	PrintJSON(users)

	return c.JSON(users)
}

// @Tags Artisan
// @Param id path string true "Account ID"
// @Router /api/v1/artisans/{id} [get]
func ArtisanFind(c *fiber.Ctx) error {

	user, err := prisma.Client.Artisan.FindUnique(
		db.Artisan.ID.Equals(c.Params("id")),
	).Exec(prisma.Ctx)
	if err != nil {
		return err
	}

	PrintJSON(user)

	return c.JSON(user)
}

// @Tags Artisan
// UserCreate registers a user
//@Param        data  body      models.Artisan  true  "Account Info"
// @Router /api/v1/artisans [post]
func ArtisanCreate(c *fiber.Ctx) error {

	var payload models.Artisan

	if err := c.BodyParser(&payload); err != nil {
		return err
	}

	PrintJSON(payload)

	created, err := prisma.Client.Artisan.CreateOne(
		db.Artisan.Firstname.Set(payload.Firstname),
		db.Artisan.Lastname.Set(payload.Lastname),
		db.Artisan.Cluster.Link(db.Cluster.ID.Equals(payload.ClusterID)),
		db.Artisan.Tel.SetIfPresent(payload.Tel),
		db.Artisan.PicFormat.SetOptional(payload.PicFormat),
	).Exec(prisma.Ctx)
	if err != nil {
		return err
	}

	PrintJSON(created)

	return c.JSON(created)
}

// @Tags Artisan
// @Param id path string true "Account ID"
// @Router /api/v1/artisans/{id} [delete]
func ArtisanDelete(c *fiber.Ctx) error {
	deleted, err := prisma.Client.Artisan.FindUnique(db.Artisan.ID.Equals(c.Params("id"))).Delete().Exec(prisma.Ctx)
	if err != nil {
		return err
	}

	return c.JSON(deleted)
}

// @Tags Artisan
// @Param id path string true "Account ID"
//@Param        data  body      models.Artisan  true  "Account Info"
// @Router /api/v1/artisans/{id} [patch]
func ArtisanUpdate(c *fiber.Ctx) error {
	var payload models.Artisan

	if err := c.BodyParser(&payload); err != nil {
		return err
	}

	updated, err := prisma.Client.Artisan.FindUnique(db.Artisan.ID.Equals(c.Params("id"))).Update(
		db.Artisan.Firstname.Set(payload.Firstname),
		db.Artisan.Lastname.Set(payload.Lastname),
		db.Artisan.Cluster.Link(db.Cluster.ID.Equals(payload.ClusterID)),
		db.Artisan.Tel.SetIfPresent(payload.Tel),
		db.Artisan.PicFormat.SetOptional(payload.PicFormat),
	).Exec(prisma.Ctx)
	if err != nil {
		return err
	}

	return c.JSON(updated)
}

// @Tags Artisan
// @Router /api/v1/artisans [delete]
func DeleteAllArtisans(c *fiber.Ctx) error {
	deleted, err := prisma.Client.Artisan.FindMany().Delete().Exec(prisma.Ctx)
	if err != nil {
		return err
	}

	return c.JSON(deleted)
}
